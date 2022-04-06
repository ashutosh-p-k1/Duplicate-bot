package main

import (
	"auto-approval/pckg"
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/go-github/v43/github"
)

var (
	client *github.Client
	err    error
)

func main() {
	//TODO Authentication For the Github Bot
	client, err = pckg.Initialize()
	if err != nil {
		log.Fatal(err.Error())

	}
	// gin.New returns a Blank engine .
	r := gin.New()

	//TODO - Setting Up the server .

	server := http.Server{
		Addr:    os.Getenv("ADDR"),
		Handler: r,
	}
	r.POST("/api/v1/", Handler)
	log.Println("Starting Server On : ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Successfully Started The Server")
	}
}

func Handler(ctx *gin.Context) {
	event := ctx.GetHeader("X-Github-Event")
	fmt.Println("The Event is :=>", event)
	payload, _ := ioutil.ReadAll(ctx.Request.Body)
	if VerifySignature(payload, ctx.GetHeader("X-Hub-Signature-256")) && len(payload) > 0 {
		fmt.Println("Github_Signature_Verified")
		switch event {
		case "pull_request":
			pullEvent(payload)
		default:
			writeTofile(payload)

		}
	}
}

/*TODO - Verify that ,is Github is Really the Service that Genrated the event ?*/
func VerifySignature(payload []byte, signature string) bool {
	key := hmac.New(sha256.New, []byte(os.Getenv("WEBHOOK_SECRET")))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))
	log.Printf("computed signature: %s", computedSignature)

	return computedSignature == signature
}

/*If Case is Pull Request .*/
func pullEvent(payload []byte) { // Takes Payload as input .
	writeTofile(payload) // write the Payloads into a Json File
	pr := pckg.PullRequest{}

	if err := json.Unmarshal(payload, &pr); err != nil { // Unmarshalling the Payload w.r.t  Pull request struct
		log.Println(err.Error())
		return
	}
	fmt.Println("Pr.Action :=>", pr.Action)
	if pr.Action == "opened" || pr.Action == "reopened" {
		files := pr.ChangedFilesFromPullRequest(client) // To get the modifified files from the Pull request .

		for _, f := range files {
			if filepath.Ext(f.GetFilename()) == ".yaml" && f.GetStatus() == "modified" {
				value, n := isValid(f)
				fmt.Println("The value is    :", value)
				if !value {
					body := fmt.Sprintf("Hey @%s ,Your policy does not validate with the basepolicy",
						pr.PullRequestData.User.Login)

					comment := &github.PullRequestComment{
						Body:     github.String(body),
						Path:     f.Filename,
						CommitID: github.String(pr.PullRequestData.Head.Sha),
						Position: github.Int(n),
					}
					pr.CommentOnPullRequest(client, comment)
					pr.ClosePullRequest(client)
					return
				}
				if value {
					fmt.Println("IN MERGING :=>")
					pr.Merging(client)
					return
				}
			}
		}

	}
}
func isValid(file *github.CommitFile) (bool, int) {

	buff := bytes.NewBuffer([]byte(file.GetPatch()))
	scanner := bufio.NewScanner(buff)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count++
		if strings.HasPrefix(line, "-") {
			return false, count
		}
	}
	return true, count
}
func writeTofile(data []byte) {
	name := fmt.Sprintf("temp/logs/%d-action.json", time.Now().UnixNano())
	w, err := os.Create(name)
	if err != nil {
		log.Println(err.Error())
	}
	defer w.Close()

	if _, err := w.Write(data); err != nil {
		log.Println(err.Error())
	}
}
