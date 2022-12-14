package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"strings"
)

// CommitData it's a class to wrap commit information
type CommitData struct {
	author  string
	hash    string
	message string
}

// ToString it's a method to show object data as string
func (cd CommitData) ToString() string {
	return fmt.Sprintf("commit=%s, author=%s, message=%s", cd.hash, cd.author, cd.message)
}

// NewCommitData it's a function to create CommitData
// object from git commit
func NewCommitData(commitObject *object.Commit) *CommitData {

	commit := new(CommitData)
	commit.author = commitObject.Author.Name
	commit.hash = commitObject.Hash.String()
	commit.message = commitObject.Message

	return commit
}

// BuildBranchPath it's path builder as target to clone repositories
// Should be defined as parameter the targetFolder name and branch
// Return pattern as _[target_folder]_[target_branch] without slash's
func BuildBranchPath(targetFolder string, branch string) string {
	baseRepositoryConfigInstance := GetRepositoryConfigInstance()

	if baseRepositoryConfigInstance.enabled {
		return fmt.Sprintf("_%s_%v", targetFolder, strings.ReplaceAll(branch, "/", "_"))
	} else {
		return fmt.Sprintf(targetFolder)
	}

}

// CheckContentExists it's function to check if exists content
// inside target folder
func CheckContentExists(target string, branch string) bool {
	if _, err := os.Stat(BuildBranchPath(target, branch)); !os.IsNotExist(err) {
		return true
	}
	return false
}

// CloneRepository it's repository clone operation.
// Should be provided as parameter, repository url, branch and target folder
// Must be selected if build path function will be used to generate folder name
// Return CommitData with the last commit
func CloneRepository(repoURL string, branch string, targetFolder string, buildPath bool) *CommitData {
	var targetFolderMultibranch string

	if buildPath {
		targetFolderMultibranch = BuildBranchPath(targetFolder, branch)
	} else {
		targetFolderMultibranch = targetFolder
	}

	if _, err := os.Stat(targetFolderMultibranch); !os.IsNotExist(err) {
		log.Println("Repository already exist. Will be pulled")
		return PullRepository(repoURL, targetFolderMultibranch, branch)
	}

	auth := GetBasicAuthenticationMethodInstance()
	var authResult http.BasicAuth

	if auth.BasicAuthAvailable() {
		authResult = auth.GetAuth()
	}

	r, err := git.PlainClone(targetFolderMultibranch, false, &git.CloneOptions{
		Auth:          &authResult,
		URL:           repoURL,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})

	if ErrorCheck(err) {
		return nil
	}

	return getCommit(r)

}
func getCommit(repository *git.Repository) *CommitData {

	ref, err := repository.Head()

	CriticalErrorCheck(err)

	commit, errorCommit := repository.CommitObject(ref.Hash())

	CriticalErrorCheck(errorCommit)

	commitData := NewCommitData(commit)
	return commitData

}

// CheckoutRepository it's checkout another branch on cloned repository
// Use CloneRepository function
// Return CommitData with the last commit
func CheckoutRepository(repoURL string, target string, branch string) *CommitData {

	if _, err := os.Stat(target); !os.IsNotExist(err) {
		os.RemoveAll(target)
	}

	return CloneRepository(repoURL, branch, target, false)
}

// PullRepository it's a git pull default operation
// Will be used to update repository content
// Return CommitData with the last commit
func PullRepository(repoURL string, target string, branch string) *CommitData {
	po, err := git.PlainOpen(target)

	h, errOpen := po.Head()

	ErrorCheck(errOpen)

	requestedBranch := fmt.Sprintf("refs/heads/%s", branch)
	presentBranch := fmt.Sprintf("%s", h.Name())

	if presentBranch != requestedBranch {
		log.Printf("Checkout: [%s]", branch)
		return CheckoutRepository(repoURL, target, branch)
	}
	if ErrorCheck(err) {
		return nil
	}

	w, errWorktree := po.Worktree()

	if ErrorCheck(errWorktree) {
		return nil
	}

	auth := GetBasicAuthenticationMethodInstance()
	var authResult http.BasicAuth

	if auth.BasicAuthAvailable() {
		authResult = auth.GetAuth()
	}

	errPull := w.Pull(&git.PullOptions{
		Auth:          &authResult,
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})
	if errPull != nil {
		ErrorCheck(errPull)
	}

	return getCommit(po)
}
