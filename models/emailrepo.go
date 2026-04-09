package models

type EmailRepo struct {
	Email string `json:"email"` // Email address
	Repo  string `json:"repo"`  // GitHub repository in owner/repo format
}
