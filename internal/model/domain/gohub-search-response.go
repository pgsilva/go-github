package domain

import "time"

type GoHubSearchResponse struct {
	Username  string                       `json:"username"`
	ID        int                          `json:"id"`
	AvatarURL string                       `json:"avatar_url"`
	URL       string                       `json:"url"`
	Followers string                       `json:"followers"`
	Following string                       `json:"following"`
	Repos     []GoHubSearchRepositorieItem `json:"repos"`
}

type GoHubSearchRepositorieItem struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	Private     bool      `json:"private"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Fork        bool      `json:"fork"`
	CreatedAt   time.Time `json:"created_at"`
}
