package response

type GitHubSearchUserApiResponse struct {
	TotalCount        int              `json:"total_count"`
	IncompleteResults bool             `json:"incomplete_results"`
	Items             []GitHubUserItem `json:"items"`
}

type GitHubUserItem struct {
	Login        string `json:"login"`
	ID           int    `json:"id"`
	AvatarURL    string `json:"avatar_url"`
	URL          string `json:"url"`
	HTMLURL      string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	ReposURL     string `json:"repos_url"`
	Type         string `json:"type"`
}
