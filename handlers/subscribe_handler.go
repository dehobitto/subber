package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"subber/models"
)

func Subscribe(c *gin.Context) {
	var newOwnerRepo models.EmailRepo

	if err := c.ShouldBindJSON(&newOwnerRepo); err != nil {
		// TODO
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if !isValid(newOwnerRepo.Repo) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repository"})
		return
	}

	resp, err := checkIfRepoExists(newOwnerRepo.Repo)

	if err != nil {
		// do smth
	}

	log.Println(resp)

}

func checkIfRepoExists(repo string) (*http.Response, error) {
	link := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)

	req.Header.Set("User-Agent", "Go-Subber-App")

	return client.Do(req)
}

// another file TODO
var re = regexp.MustCompile(`^[^/]+/[^/]+$`)

func isValid(repo string) bool {
	return re.MatchString(repo)
}
