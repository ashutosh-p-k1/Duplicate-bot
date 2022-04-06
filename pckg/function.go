package pckg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v43/github"
)

/*Function Definition to Tackle Authentication*/
func Initialize() (*github.Client, error) {
	appid, _ := strconv.Atoi(os.Getenv("APP_ID"))
	installid, _ := strconv.Atoi(os.Getenv("INSTALL_ID"))

	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, int64(appid), int64(installid), os.Getenv("KEYPATH"))

	if err != nil {
		log.Fatal(err)
	}
	return github.NewClient(&http.Client{Transport: itr}), nil
}

/*Return all files From The pull request Which are Changed .*/
func (pr *PullRequest) ChangedFilesFromPullRequest(client *github.Client) []*github.CommitFile {

	files, _, err := client.PullRequests.ListFiles(context.Background(), pr.PullRequestData.Head.User.Login,
		pr.PullRequestData.Head.Repo.Name, int(pr.Number), &github.ListOptions{})

	if err != nil {
		log.Println(err.Error())
		return []*github.CommitFile{}
	}
	return files
}

/*Merge the pull request*/

/*func (pr *PullRequest) MergePullRequest(client *github.Client) (*github.PullRequestMergeResult, error) {

	log.Printf("Merging pull request %d on %s", int(pr.Number), pr.PullRequestData.Base.Repo.Name)

	result, _, err := client.PullRequests.Merge(context.Background(), pr.PullRequestData.Base.User.Login, pr.PullRequestData.Base.Repo.Name, int(pr.Number),

		fmt.Sprintf("Merging pull request %d", pr.Number),
		&github.PullRequestOptions{})

	return result, err

}
*/
func (pr *PullRequest) Merging(client *github.Client) (*github.PullRequestMergeResult, *github.Response, error) {
	result, response, err := client.PullRequests.Merge(context.Background(), pr.PullRequestData.Base.User.Login, pr.PullRequestData.Base.Repo.Name, int(pr.Number),
		fmt.Sprintf("Merging pull request %d", pr.Number), &github.PullRequestOptions{})

	return result, response, err

}

/*closes the pull request*/

func (pr *PullRequest) ClosePullRequest(client *github.Client) {

	log.Printf("Closing pull request %d on %s",
		int(pr.Number),
		pr.PullRequestData.Base.Repo.Name)

	result, res, err := client.PullRequests.Edit(context.Background(), pr.PullRequestData.Base.User.Login,
		pr.PullRequestData.Base.Repo.Name, int(pr.Number), &github.PullRequest{
			State: github.String("closed"),
		})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res.StatusCode)
	log.Println(*result.State)
}

// Creating Comment on the files ..
func (pr PullRequest) CommentOnPullRequest(client *github.Client, comment *github.PullRequestComment) {

	_, res, err := client.PullRequests.CreateComment(context.Background(),
		pr.PullRequestData.Base.User.Login,
		pr.PullRequestData.Base.Repo.Name, int(pr.Number),
		comment)

	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res.Status)

}
