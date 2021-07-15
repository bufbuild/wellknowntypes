package githubutil

import (
	"context"
	"net/http"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

const (
	githubPerPage     = 50
	githubTokenEnvKey = "GITHUB_TOKEN"
)

// NewHTTPClient returns a new *http.Client that has an oauth2 token from
// the environment variable GITHUB_TOKEN if present.
func NewHTTPClient(ctx context.Context) *http.Client {
	if githubToken := os.Getenv(githubTokenEnvKey); githubToken != "" {
		return oauth2.NewClient(
			ctx,
			oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: githubToken,
				},
			),
		)
	}
	return &http.Client{}
}

// NewGitbubClient returns a new *github.Client.
func NewGithubClient(httpClient *http.Client) *github.Client {
	return github.NewClient(httpClient)
}

// AllReleaseTagNames gets all release tag names for the repository.
func AllReleaseTagNames(
	ctx context.Context,
	githubClient *github.Client,
	owner string,
	repository string,
) ([]string, error) {
	var tagNames []string
	nextPage := 0
	for {
		repositoryReleases, response, err := githubClient.Repositories.ListReleases(
			ctx,
			owner,
			repository,
			&github.ListOptions{
				Page:    nextPage,
				PerPage: githubPerPage,
			},
		)
		if err != nil {
			return nil, err
		}
		for _, repositoryRelease := range repositoryReleases {
			if tagName := repositoryRelease.TagName; tagName != nil {
				tagNames = append(tagNames, *tagName)
			}
		}
		nextPage = response.NextPage
		if nextPage == 0 {
			return tagNames, nil
		}
	}
}
