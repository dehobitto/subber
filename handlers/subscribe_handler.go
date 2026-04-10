package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"subber/models"
	"subber/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Subscribe(c *gin.Context) {
	var newOwnerRepo models.Subscription

	if err := c.ShouldBindJSON(&newOwnerRepo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if !utils.IsValidRepo(newOwnerRepo.Repo) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repository"})
		return
	}

	resp, err := checkIfRepoExists(c.Request.Context(), newOwnerRepo.Repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach GitHub API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "GitHub API rate limit exceeded. Try again later."})
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "This repository does not exist"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "External API error"})
		return
	}

	tag, err := getLatestTag(c.Request.Context(), newOwnerRepo.Repo)
	if err != nil {
		log.Println("Could not fetch tag:", err)
	}

	newOwnerRepo.LastSeenTag = tag
	newOwnerRepo.Token = uuid.New().String()
	newOwnerRepo.Confirmed = false

	err = h.Repo.SaveSubscription(c.Request.Context(), newOwnerRepo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database saving failed"})
		return
	}

	err = h.Repo.ConfirmSubscriptionByToken(c.Request.Context(), newOwnerRepo.Token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Subscription successfull!"})
}

func checkIfRepoExists(ctx context.Context, repo string) (*http.Response, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	req, err := http.NewRequestWithContext(ctx, "HEAD", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Go-Subber-App")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client.Do(req)
}

func getLatestTag(ctx context.Context, repo string) (string, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	req, err := http.NewRequestWithContext(ctx, "GET", link, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Go-Subber-App")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("github error: %d", resp.StatusCode)
	}

	var release models.GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	return release.TagName, nil
}
