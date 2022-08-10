package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetRouteConfigInstance(t *testing.T) {

	os.Setenv("PATH_CLONE", "_clone1")
	os.Setenv("PATH_WEBHOOK", "_hook1")
	os.Setenv("PATH_PULL", "_pull1")
	os.Setenv("PATH_VERSION", "_version1")

	routeConfigInstance := GetRouteConfigInstance()

	assert.Equal(t, "_clone1", routeConfigInstance.GetClone(), "Check clone path")
	assert.Equal(t, "_hook1", routeConfigInstance.GetWebHook(), "Check Webhook path")
	assert.Equal(t, "_pull1", routeConfigInstance.GetPull(), "Check pull path")
	assert.Equal(t, "_version1", routeConfigInstance.GetVersion(), "Check version path")

	routeConfigInstance1 := GetRouteConfigInstance()

	assert.Equal(t, "_clone1", routeConfigInstance1.GetClone(), "Check clone path")
	assert.Equal(t, "_hook1", routeConfigInstance1.GetWebHook(), "Check Webhook path")
	assert.Equal(t, "_pull1", routeConfigInstance1.GetPull(), "Check pull path")
	assert.Equal(t, "_version1", routeConfigInstance1.GetVersion(), "Check version path")
}
func TestGetRepositoryConfigInstance(t *testing.T) {
	os.Setenv("REPO_URL", "https://test.com/repo.git")
	os.Setenv("REPO_BRANCH", "main")
	os.Setenv("REPO_TARGET_FOLDER", "test1")

	repositoryConfigInstance := GetRepositoryConfigInstance()

	assert.Equal(t, "https://test.com/repo.git", repositoryConfigInstance.GetRepo(), "Check repo url")
	assert.Equal(t, "main", repositoryConfigInstance.GetBranch(), "Check repo url")
	assert.Equal(t, "test1", repositoryConfigInstance.GetTargetFolder(), "Check repo url")

	repositoryConfigInstance1 := GetRepositoryConfigInstance()

	assert.Equal(t, "https://test.com/repo.git", repositoryConfigInstance1.GetRepo(), "Check repo url")
	assert.Equal(t, "main", repositoryConfigInstance1.GetBranch(), "Check repo url")
	assert.Equal(t, "test1", repositoryConfigInstance1.GetTargetFolder(), "Check repo url")
}
func TestGetBasicAuthenticationMethodInstance(t *testing.T) {
	os.Setenv("REPO_USERNAME", os.Getenv("ACCESS_USERNAME"))
	os.Setenv("REPO_PASSWORD", os.Getenv("ACCESS_TOKEN"))

	basicAuthInstance := GetBasicAuthenticationMethodInstance()

	assert.Equal(t, "test1", basicAuthInstance.GetAuth().Username)
	assert.Equal(t, "testPwd", basicAuthInstance.GetAuth().Password)

	basicAuthInstance1 := GetBasicAuthenticationMethodInstance()

	assert.Equal(t, "test1", basicAuthInstance1.GetAuth().Username)
	assert.Equal(t, "testPwd", basicAuthInstance1.GetAuth().Password)
}
