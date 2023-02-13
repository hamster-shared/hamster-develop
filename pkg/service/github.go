package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
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
	client := utils.NewGithubClient(g.ctx, token)
	_, res, _ := client.Repositories.Get(g.ctx, owner, projectName)
	if res.StatusCode == 404 {
		return true
	}
	return false
}

func (g *GithubService) CreateRepo(token, templateOwner, templateRepo, repoName, repoOwner string) (*github.Repository, *github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	var data github.TemplateRepoRequest
	//private := true
	data.Name = &repoName
	data.Owner = &repoOwner
	//data.Private = &private
	repo, res, err := client.Repositories.CreateFromTemplate(g.ctx, templateOwner, templateRepo, &data)
	if err != nil {
		log.Println("create github repository failed ", err.Error())
		return nil, res, err
	}
	return repo, res, nil
}

func (g *GithubService) GetUserInfo(token string) (*github.User, error) {
	client := utils.NewGithubClient(g.ctx, token)
	user, _, err := client.Users.Get(g.ctx, "")
	if err != nil {
		log.Println("get github user info failed ", err.Error())
		return nil, err
	}
	return user, nil
}

func (g *GithubService) UpdateRepo(token, owner, repoName, name string) (*github.Repository, *github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	var data github.Repository
	data.Name = &name
	repo, res, err := client.Repositories.Edit(g.ctx, owner, repoName, &data)
	if err != nil {
		log.Println("create github repository failed ", err.Error())
		return nil, res, err
	}
	return repo, res, nil
}

func (g *GithubService) DeleteRepo(token, owner, repoName string) (*github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	res, err := client.Repositories.Delete(g.ctx, owner, repoName)
	if err != nil {
		log.Println("delete github repository failed ", err.Error())
		return res, err
	}
	return res, nil
}

func (g *GithubService) AddFile(token, owner, repoName, content, fileName string) (*github.RepositoryContentResponse, *github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	var fileOptions github.RepositoryContentFileOptions
	path := fmt.Sprintf("contracts/%s.sol", fileName)
	message := "Initial commit"
	fileOptions.Message = &message
	fileOptions.Content = []byte(content)
	repoRes, res, err := client.Repositories.CreateFile(g.ctx, owner, repoName, path, &fileOptions)
	if err != nil {
		log.Println("add file failed: ", err.Error())
		return repoRes, res, err
	}
	return repoRes, res, nil
}

func (g *GithubService) GetCommitInfo(token, owner, repo, ref string) (string, *github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	hash, res, err := client.Repositories.GetCommitSHA1(g.ctx, owner, repo, ref, "")
	return hash, res, err
}
