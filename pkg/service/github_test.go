package service

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestGithubCreateRepo(t *testing.T) {
	ctx := context.Background()
	token := ""
	client := utils.NewGithubClient(ctx, token)

	owner := "mohaijiang"
	repoName := "my-starknet"

	_, resp, _ := client.Repositories.Get(ctx, owner, repoName)
	if resp.StatusCode == 200 {
		deleteRepo(ctx, client, owner, repoName)
	}

	var data github.TemplateRepoRequest
	//private := true
	data.Name = &repoName
	data.Owner = &owner
	//data.Private = &private
	repo, res, _ := client.Repositories.CreateFromTemplate(ctx, consts.TemplateOwner, consts.TemplateRepoName, &data)
	fmt.Println(res.StatusCode)

	fmt.Println(repo.GetDefaultBranch())

	repo, resp, _ = client.Repositories.Get(ctx, owner, repoName)

	for {
		commitSha1, resp, err := client.Repositories.GetCommitSHA1(ctx, owner, repoName, repo.GetDefaultBranch(), "")
		fmt.Println("get commit err : ", err)
		fmt.Println("get commit resp: ", resp.Status)
		fmt.Println("get commit sha1: ", commitSha1)
		if resp.StatusCode == 200 {
			break
		}
	}

	var fileOptions github.RepositoryContentFileOptions
	path := fmt.Sprintf("contracts/%s.md", "ERC20")
	content := "#ERC20 "
	message := "Second commit"
	fileOptions.Branch = repo.DefaultBranch
	fileOptions.Message = &message
	fileOptions.Content = []byte(content)
	contentResp, resp, err := client.Repositories.CreateFile(ctx, owner, repoName, path, &fileOptions)

	assert.NoError(t, err)
	fmt.Println(resp.Status)

	fmt.Println(contentResp.GetHTMLURL())

}

func deleteRepo(ctx context.Context, client *github.Client, owner, repo string) {

	res, err := client.Repositories.Delete(ctx, owner, repo)
	fmt.Println("delete repo Response code :", res.StatusCode)
	if res.StatusCode != 204 {
		fmt.Println(err.Error())
		data, _ := io.ReadAll(res.Body)
		fmt.Println(string(data))
		return
	}
}

func TestDeleteRepo(t *testing.T) {
	ctx := context.Background()
	token := "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"
	client := utils.NewGithubClient(ctx, token)

	option := &github.RepositoryListOptions{
		Sort:      "created",
		Direction: "desc",
	}
	repositorys, _, err := client.Repositories.List(ctx, "mohaijiang", option)
	if err != nil {
		panic(err)
		return
	}

	for _, r := range repositorys {
		fmt.Println(r.GetName())
		if strings.HasPrefix(r.GetName(), "my-") {
			deleteRepo(ctx, client, "mohaijiang", r.GetName())
		}
	}

	//deleteRepo(ctx, client, owner, repoName)
}

func TestGetGitHubAppInfo(t *testing.T) {
	appId := int64(294815)
	//installationID := int64(34695244)
	pemPath := "/Users/abing/Desktop/hamster-test.2023-10-09.private-key.pem"
	////tr := http.DefaultTransport
	////_, err := ghinstallation.NewKeyFromFile(tr, appId, installationID, pemPath)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//ctx := context.Background()
	////_, err = itr.Token(ctx)
	////fmt.Println("Fetching token:", err)
	////client := github.NewClient(&http.Client{Transport: itr})
	////installation, _, err := client.Apps.GetInstallation(ctx, installationID)
	////if err != nil {
	////	log.Fatal(err)
	////}
	////fmt.Println(installation)
	//
	atr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, appId, pemPath)
	if err != nil {
		panic(err)
	}
	client := github.NewClient(&http.Client{Transport: atr})
	//getInstallation, _, err := client.Apps.GetInstallation(ctx, installationID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(fmt.Sprintf("用户是 %s, installationID是 %d, 选择的仓库权限是 %s \n", *getInstallation.Account.Login, *getInstallation.ID, *getInstallation.RepositorySelection))
	//
	//opt := &github.ListOptions{}
	//installations, _, err := client.Apps.ListInstallations(ctx, opt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for i, installation := range installations {
	//	fmt.Println(i, "安装者是--> ", installation)
	//}
	ctx := context.Background()
	//opt := &github.ListOptions{}
	githubClient := utils.NewGithubClient(ctx, "ghu_CTh3pu2XMAQNF0Rm3rfWUuG3pCm8PF0WlA81")
	userInstallations, _, err := githubClient.Apps.ListUserInstallations(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for i, installation := range userInstallations {
		fmt.Println(i, "用户安装的app有--> ", installation)
		fmt.Println(fmt.Sprintf("用户是 %s, installationID是 %d, appId是%d,  选择的仓库权限是 %s \n", *installation.Account.Login, *installation.ID, *installation.AppID, *installation.RepositorySelection))
	}

	installation, _, err := client.Apps.FindUserInstallation(ctx, "abing258")
	if err != nil {
		log.Fatal(err)
	}
	//for i, installation := range userInstallations {
	//	fmt.Println(i, "用户安装的app有--> ", installation)
	//	fmt.Println(fmt.Sprintf("用户是 %s, installationID是 %d, appId是%d,  选择的仓库权限是 %s \n", *installation.Account.Login, *installation.ID, *installation.AppID, *installation.RepositorySelection))
	//}
	fmt.Println(installation)
}

func TestGetRepo(t *testing.T) {
	token := "ghu_MZi4pD28DHbfrYfKHBuMooIYNgtJMk2zJpeW"
	owner := "abing258"
	ctx := context.Background()
	client := utils.NewGithubClient(ctx, token)
	repo, _, err := client.Repositories.Get(ctx, owner, "computeshare-server")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("仓库的信息是%v\n", repo)
	commits, _, err := client.Repositories.ListCommits(ctx, "abing258", "computeshare-server", &github.CommitsListOptions{SHA: "main"})
	if err != nil {
		log.Fatal(err)
	}
	for _, commit := range commits {
		//fmt.Printf("仓库的commit信息是%v\n", commits)
		message := *commit.Commit.Message
		time := *commit.Commit.Author.Date
		sha := *commit.SHA
		fmt.Println(fmt.Sprintf("SHA信息 %s, 提交时间是 %s, Message是 %s \n", sha, time, message))
	}
}
