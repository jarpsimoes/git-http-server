package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type testsData struct {
	title        string
	repoUrl      string
	branch       string
	targetFolder string
}

func TestCloneRepository(t *testing.T) {
	token := os.Getenv("ACCESS_TOKEN")
	username := os.Getenv("ACCESS_USERNAME")

	var data [2]testsData

	data[0] = getPublicRepo()
	data[1] = getRepoAuth()

	for i, s := range data {

		if i == 1 {
			authInstance := GetBasicAuthenticationMethodInstance()
			authInstance.username = username
			authInstance.passwordToken = token

		}

		result := CloneRepository(s.repoUrl, s.branch, s.targetFolder, false)
		assert.NotNil(t, result)

		assert.NotEmpty(t, result.hash)
		assert.NotEmpty(t, result.author)

	}

	for _, s := range data {
		os.RemoveAll(s.targetFolder)
	}

}

func TestCheckoutRepository(t *testing.T) {
	token := os.Getenv("ACCESS_TOKEN")
	username := os.Getenv("ACCESS_USERNAME")

	data := getRepoAuth()

	authInstance := GetBasicAuthenticationMethodInstance()

	authInstance.username = username
	authInstance.passwordToken = token

	result := CheckoutRepository(data.repoUrl, data.targetFolder, data.branch)
	assert.NotNil(t, result)

	assert.NotEmpty(t, result.hash)
	assert.NotEmpty(t, result.author)
	os.RemoveAll(data.targetFolder)
}
func TestPullRepository(t *testing.T) {

	data := getRepoAuth()

	token := os.Getenv("ACCESS_TOKEN")
	username := os.Getenv("ACCESS_USERNAME")

	authInstance := GetBasicAuthenticationMethodInstance()

	authInstance.username = username
	authInstance.passwordToken = token

	cloneResult := CloneRepository(data.repoUrl, data.branch, data.targetFolder, false)
	assert.NotNil(t, cloneResult)

	pullResult := PullRepository(data.repoUrl, data.targetFolder, data.branch)
	assert.NotNil(t, pullResult)

	os.RemoveAll(data.targetFolder)

}
func getPublicRepo() testsData {
	return testsData{
		title:        "Test real repo without auth",
		repoUrl:      "https://github.com/jarpsimoes/ansible-configure-http-server",
		branch:       "main",
		targetFolder: "test1",
	}
}
func getRepoAuth() testsData {
	return testsData{
		title:        "Test real repo with PAT auth - Github",
		repoUrl:      "https://github.com/jarpsimoes/html_sample.git",
		branch:       "main",
		targetFolder: "test2",
	}
}
