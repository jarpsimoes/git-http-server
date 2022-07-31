package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"os"
	"strings"
)

type CommitData struct {
	author  string
	hash    string
	message string
}

func (cd CommitData) ToString() string {
	return fmt.Sprintf("commit=%s, author=%s, message=%s", cd.hash, cd.author, cd.message)
}
func NewCommitData(commitObject *object.Commit) *CommitData {

	commit := new(CommitData)
	commit.author = commitObject.Author.Name
	commit.hash = commitObject.Hash.String()
	commit.message = commitObject.Message

	return commit
}
func CloneRepository(repoUrl string, branch string, targetFolder string, buildPath bool) *CommitData {
	var targetFolderMultibranch string

	if buildPath {
		targetFolderMultibranch = fmt.Sprintf("_%s_%v", targetFolder, strings.ReplaceAll(branch, "/", "_"))
	} else {
		targetFolderMultibranch = targetFolder
	}
	
	if _, err := os.Stat(targetFolderMultibranch); !os.IsNotExist(err) {
		log.Println("Repository already exist. Will be pulled")
		return PullRepository(repoUrl, targetFolderMultibranch, branch)
	}

	r, err := git.PlainClone(targetFolderMultibranch, false, &git.CloneOptions{
		URL:           repoUrl,
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
func CheckoutRepository(repoUrl string, target string, branch string) *CommitData {

	if _, err := os.Stat(target); !os.IsNotExist(err) {
		os.RemoveAll(target)
	}

	return CloneRepository(repoUrl, branch, target, false)
}
func PullRepository(repoUrl string, target string, branch string) *CommitData {
	po, err := git.PlainOpen(target)

	h, errOpen := po.Head()

	ErrorCheck(errOpen)

	if fmt.Sprintf("/refs/heads/%s", h.Name()) != branch {
		log.Printf("Checkout: [%s]", branch)
		return CheckoutRepository(repoUrl, target, branch)
	}
	if ErrorCheck(err) {
		return nil
	}

	w, errWorktree := po.Worktree()

	if ErrorCheck(errWorktree) {
		return nil
	}

	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})

	return getCommit(po)
}
