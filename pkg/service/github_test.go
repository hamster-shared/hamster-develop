package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/stretchr/testify/assert"
	"io"
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
