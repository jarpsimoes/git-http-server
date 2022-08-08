package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCloneRepository(t *testing.T) {
	token := os.Getenv("ACCESS_TOKEN")
	username := os.Getenv("ACCESS_USERNAME")

	type testsData struct {
		title        string
		repoUrl      string
		branch       string
		targetFolder string
	}

	var data [2]testsData

	data[0] = testsData{
		title:        "Test real repo without auth",
		repoUrl:      "https://github.com/jarpsimoes/ansible-configure-http-server",
		branch:       "main",
		targetFolder: "test1",
	}
	data[1] = testsData{
		title:        "Test real repo with PAT auth - Github",
		repoUrl:      "https://github.com/jarpsimoes/html_sample.git",
		branch:       "main",
		targetFolder: "test2",
	}

	for i, s := range data {

		if i == 1 {
			authInstance := GetBasicAuthenticationMethodInstance()
			authInstance.username = username
			authInstance.passwordToken = token

		}

		result := CloneRepository(s.repoUrl, s.branch, s.targetFolder, false)
		assert.NotNil(t, result)
	}

	for _, s := range data {
		os.RemoveAll(s.targetFolder)
	}

}
