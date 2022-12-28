package service

import (
	"context"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/a-line/pkg/utils"
	"log"
)

type IGithubService interface {
	CheckName(owner, projectName string) (bool, error)
	CreateRepo(templateOwner, templateRepo, repoName, repoOwner string) error
}

type GithubService struct {
	ctx context.Context
}

func NewGithubService() *GithubService {
	return &GithubService{
		ctx: context.Background(),
	}
}

func (g *GithubService) CheckName(token, owner, projectName string) bool {
	tokenData := "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"
	client := utils.NewGithubClient(g.ctx, tokenData)
	_, res, _ := client.Repositories.Get(g.ctx, owner, projectName)
	if res.StatusCode == 404 {
		return true
	}
	return false
}

func (g *GithubService) CreateRepo(token, templateOwner, templateRepo, repoName, repoOwner string) (*github.Repository, error) {
	tokenData := "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"
	client := utils.NewGithubClient(g.ctx, tokenData)
	var data github.TemplateRepoRequest
	data.Name = &repoName
	data.Owner = &repoOwner
	repo, _, err := client.Repositories.CreateFromTemplate(g.ctx, templateOwner, templateRepo, &data)
	if err != nil {
		log.Println("create github repository failed ", err.Error())
		return nil, err
	}
	return repo, nil
}
