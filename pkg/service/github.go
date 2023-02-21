package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/pkg/errors"
	"github.com/wujiangweiphp/go-curl"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func (g *GithubService) CreateRepository(token, repoName string) (*github.Repository, *github.Response, error) {
	client := utils.NewGithubClient(g.ctx, token)
	var data github.Repository
	data.Name = &repoName
	return client.Repositories.Create(g.ctx, "", &data)
}

func (g *GithubService) CommitAndPush(token, repoUrl, owner, email, templateUrl, templateName string) error {
	cloneDir := filepath.Join(utils.DefaultRepoDir(), owner)
	_, err := os.Stat(cloneDir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(cloneDir, os.ModePerm)
		if err != nil {
			log.Println("create workdir failed", err.Error())
			return err
		}
	}
	gitClone := exec.Command("git", "clone", templateUrl)
	gitClone.Dir = cloneDir
	err = gitClone.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git clone failed", err.Error())
		return err
	}
	workdir := filepath.Join(utils.DefaultRepoDir(), owner, templateName)
	deleteGit := exec.Command("rm", "-rf", ".git")
	deleteGit.Dir = workdir
	err = deleteGit.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("delete .git failed", err.Error())
		return err
	}
	gitInit := exec.Command("git", "init", "-b", "main")
	gitInit.Dir = workdir
	err = gitInit.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git init failed", err.Error())
		return err
	}
	configName := exec.Command("git", "config", "user.name", owner)
	configName.Dir = workdir
	err = configName.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("config git user name failed", err.Error())
		return err
	}
	configEmail := exec.Command("git", "config", "user.email", email)
	configEmail.Dir = workdir
	err = configEmail.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("config git user email failed", err.Error())
		return err
	}
	first := strings.Index(repoUrl, "/")
	index := first + 2
	originUrl := repoUrl[:index] + fmt.Sprintf("%s@", token) + repoUrl[index:]
	addOrigin := exec.Command("git", "remote", "add", "origin", originUrl)
	addOrigin.Dir = workdir
	err = addOrigin.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git add origin failed", err.Error())
		return err
	}
	fileAdd := exec.Command("git", "add", ".")
	fileAdd.Dir = workdir
	err = fileAdd.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git file add failed", err.Error())
		return err
	}
	gitCommit := exec.Command("git", "commit", "-m", "Initial commit")
	gitCommit.Dir = workdir
	err = gitCommit.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git commit failed", err.Error())
		return err
	}
	gitPush := exec.Command("git", "push", "origin", "main")
	gitPush.Dir = workdir
	err = gitPush.Run()
	if err != nil {
		deleteOwnerDir(owner)
		log.Println("git push failed", err.Error())
		return err
	}
	return nil
}

func (g *GithubService) GetUserEmail(token string) (string, error) {
	var data []vo.GithubEmail
	auth := fmt.Sprintf("Bearer %s", token)
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        auth,
		"X-GitHub-Api-Version": "2022-11-28",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl("https://api.github.com/user/emails").
		SetHeaders(headers).
		Get()
	if err != nil {
		log.Println("github email request failed: ", err.Error())
		return "", err
	} else {
		if resp.IsOk() {
			json.Unmarshal([]byte(resp.Body), &data)
		} else {
			log.Printf("%v\n", resp.Raw)
			return "", errors.New("github email request failed")
		}
	}
	if len(data) > 0 {
		return data[0].Email, nil
	}
	return "", nil
}

func deleteOwnerDir(owner string) {
	deleteCmd := exec.Command("rm", "-rf", owner)
	deleteCmd.Dir = utils.DefaultRepoDir()
	deleteCmd.Start()
}
