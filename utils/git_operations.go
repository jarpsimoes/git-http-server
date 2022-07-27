package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

type CommitData struct {
	author  string
	hash    string
	message string
}

func (cd CommitData) toString() string {
	return fmt.Sprintf("commit=%s, author=%s, message=%s", cd.hash, cd.author, cd.message)
}
func NewCommitData(commitObject *object.Commit) *CommitData {

	commit := new(CommitData)
	commit.author = commitObject.Author.Name
	commit.hash = commitObject.Hash.String()
	commit.message = commitObject.Message

	return commit
}
func CloneRepository(repoUrl string, branch string, targetFolder string) {
	log.Println(branch)
	r, err := git.PlainClone(targetFolder, false, &git.CloneOptions{
		URL:           repoUrl,
		ReferenceName: "refs/heads/feature/test1",
	})

	if err != nil {
		log.Printf("ERROR: %v", err.Error())
	} else {

		getCommit(r)
	}

}
func getCommit(repository *git.Repository) {

	ref, err := repository.Head()

	if err != nil {
		log.Fatal(err)
	} else {
		commit, errorCommit := repository.CommitObject(ref.Hash())

		if errorCommit != nil {
			log.Fatal(errorCommit)
		} else {
			commitData := NewCommitData(commit)
			fmt.Println(commitData.toString())
		}
	}

}
func PullRepository() {

}
