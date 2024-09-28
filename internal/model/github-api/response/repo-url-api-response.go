package response

import "time"

type GitHubRepoUrlApiResponse []struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Private     bool      `json:"private"`
	HTMLURL     string    `json:"html_url"`
	Description string    `json:"description"`
	Fork        bool      `json:"fork"`
	CreatedAt   time.Time `json:"created_at"`
}
