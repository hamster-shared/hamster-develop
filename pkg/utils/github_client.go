package utils

import (
	"context"
	"errors"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/aline-engine/logger"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strconv"
)

func NewGithubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func NewGithubClientWithPrivateKey() (*github.Client, error) {
	appIdString, exist := os.LookupEnv("GITHUB_APP_ID")
	if !exist {
		logger.Errorf("please contact the administrator to configure 'GITHUB_APP_ID'")
		return nil, errors.New("please contact the administrator to configure 'GITHUB_APP_ID'")
	}
	appId, err := strconv.Atoi(appIdString)
	if err != nil {
		logger.Errorf("app id format failed:%s", err)
		return nil, err
	}
	appPemPath, exist := os.LookupEnv("GITHUB_APP_PEM")
	if !exist {
		logger.Errorf("please contact the administrator to configure 'GITHUB_APP_PEM'")
		return nil, errors.New("please contact the administrator to configure 'GITHUB_APP_PEM'")
	}
	atr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, int64(appId), appPemPath)
	if err != nil {
		logger.Errorf("get github client by private key failed:%s", err)
		return nil, err
	}
	client := github.NewClient(&http.Client{Transport: atr})
	return client, nil
}

func NewGithubClientWithEmpty() *github.Client {
	return github.NewClient(nil)
}
