package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"subber/models"
)

func GetLatestTag(ctx context.Context, repo string, token string) (string, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	req, err := http.NewRequestWithContext(ctx, "GET", link, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Go-Subber-App")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

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

	return release.LastSeenTag, nil
}

func CheckIfRepoExists(ctx context.Context, repo string, token string) (*http.Response, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	req, err := http.NewRequestWithContext(ctx, "HEAD", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Go-Subber-App")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client.Do(req)
}
