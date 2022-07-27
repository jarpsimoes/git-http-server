package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"os"
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
func CloneRepository(repoUrl string, branch string, targetFolder string) *CommitData {

	if _, err := os.Stat(targetFolder); !os.IsNotExist(err) {
		log.Println("Repository already exist. Will be pulled")
		return PullRepository(targetFolder, branch)
	}

	r, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL:           repoUrl,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})

	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		return nil
	} else {

		return getCommit(r)
	}

}
func getCommit(repository *git.Repository) *CommitData {

	ref, err := repository.Head()

	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		commit, errorCommit := repository.CommitObject(ref.Hash())

		if errorCommit != nil {
			log.Fatal(errorCommit)
			return nil
		} else {
			commitData := NewCommitData(commit)
			fmt.Println(commitData.ToString())
			return commitData
		}
	}

}
func PullRepository(target string, branch string) *CommitData {
	po, err := git.PlainOpen(target)

	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil
	}

	w, errW := po.Worktree()

	if errW != nil {
		log.Printf("ERROR: %s", errW.Error())
		return nil
	}

	w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})

	return getCommit(po)
}
