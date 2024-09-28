package usecase

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-github/internal/model/domain"
	"github.com/pgsilva/go-github/internal/model/github-api/response"

	"github.com/pgsilva/go-github/pkg/config"
)

func SearchUsers(query string) ([]domain.GoHubSearchResponse, error) {
	slog.Info("Searching users with query: " + query)

	apiResult, err := callApi(query)
	if err != nil {
		slog.Error("Error retrieve GitHub API response", "err", err)
		return []domain.GoHubSearchResponse{}, err
	}

	listResponse := []domain.GoHubSearchResponse{}
	for _, item := range apiResult.Items {
		followersQtd, err := getQuantity(item.FollowersURL)
		if err != nil {
			slog.Error("Error getting followers quantity", "err", err)
			return []domain.GoHubSearchResponse{}, err
		}

		followingQtd, err := getQuantity(item.FollowingURL)
		if err != nil {
			slog.Error("Error getting following quantity", "err", err)
			return []domain.GoHubSearchResponse{}, err
		}

		repos, err := getRepos(item.ReposURL)
		if err != nil {
			slog.Error("Error getting repos quantity", "err", err)
			return []domain.GoHubSearchResponse{}, err
		}

		listResponse = append(listResponse, domain.GoHubSearchResponse{
			Username:  item.Login,
			ID:        item.ID,
			AvatarURL: item.AvatarURL,
			URL:       item.HTMLURL,
			Followers: strconv.Itoa(followersQtd),
			Following: strconv.Itoa(followingQtd),
			Repos:     repos,
		})
	}

	return listResponse, nil
}

func callApi(query string) (response.GitHubSearchUserApiResponse, error) {
	request, err := makeSearchRequest(query)
	if err != nil {
		slog.Error("Error making request", "err", err)
		return response.GitHubSearchUserApiResponse{}, err
	}

	client := config.GetHttpClient()
	resp, err := client.Do(request)
	if err != nil {
		slog.Error("Error calling GitHub API", "err", err)
		return response.GitHubSearchUserApiResponse{}, err
	}

	search := response.GitHubSearchUserApiResponse{}
	if err := json.Unmarshal(resp, &search); err != nil {
		return response.GitHubSearchUserApiResponse{}, err
	}

	return search, nil
}

func makeSearchRequest(query string) (*http.Request, error) {
	slog.Info("Making request to GitHub API with query: " + query)

	apiPath, err := url.Parse(config.GitHubApiUrl + "/search/users")
	if err != nil {
		slog.Error("Error building API path", "err", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("q", query)
	params.Add("sort", "updated")
	params.Add("order", "asc")
	apiPath.RawQuery = params.Encode()

	slog.Info("Request Data", "url", apiPath, "data", params)

	req, err := http.NewRequest(fiber.MethodGet, apiPath.String(), nil)
	if err != nil {
		slog.Error("Error building request", "err", err)
		return nil, err
	}

	return req, nil
}

func getQuantity(path string) (int, error) {
	slog.Info("Getting quantity from URL: " + path)

	client := config.GetHttpClient()
	apiPath, err := url.Parse(path)
	if err != nil {
		slog.Error("Error building API path", "err", err)
		return 0, err
	}

	req, err := http.NewRequest(fiber.MethodGet, apiPath.String(), nil)
	if err != nil {
		slog.Error("Error building request", "err", err)
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error calling GitHub API", "err", err)
		return 0, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		slog.Error("Error decoding JSON", "err", err)
		return 0, err
	}

	return len(result), nil

}

func getRepos(path string) ([]domain.GoHubSearchRepositorieItem, error) {
	slog.Info("Getting repos from URL: " + path)

	client := config.GetHttpClient()
	apiPath, err := url.Parse(path)
	if err != nil {
		slog.Error("Error building API path", "err", err)
		return []domain.GoHubSearchRepositorieItem{}, err
	}

	req, err := http.NewRequest(fiber.MethodGet, apiPath.String(), nil)
	if err != nil {
		slog.Error("Error building request", "err", err)
		return []domain.GoHubSearchRepositorieItem{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error calling GitHub API", "err", err)
		return []domain.GoHubSearchRepositorieItem{}, err
	}

	repos := response.GitHubRepoUrlApiResponse{}
	if err := json.Unmarshal(resp, &repos); err != nil {
		return []domain.GoHubSearchRepositorieItem{}, err
	}

	listRepos := []domain.GoHubSearchRepositorieItem{}
	for _, item := range repos {
		listRepos = append(listRepos, domain.GoHubSearchRepositorieItem{
			ID:          item.ID,
			Name:        item.Name,
			FullName:    item.FullName,
			Private:     item.Private,
			URL:         item.HTMLURL,
			Description: item.Description,
			Fork:        item.Fork,
			CreatedAt:   item.CreatedAt,
		})
	}

	return listRepos, nil
}
