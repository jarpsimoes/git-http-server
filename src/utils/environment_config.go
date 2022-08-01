package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type BaseRouteConfig struct {
	clonePath   string
	webhookPath string
	pullPath    string
	versionPath string
}

type BaseRepositoryConfig struct {
	repoUrl      string
	branch       string
	targetFolder string
}
type BasicAuthenticationMethod struct {
	username      string
	passwordToken string
}

var baseRouteConfigInstance *BaseRouteConfig
var baseRepositoryConfigInstance *BaseRepositoryConfig
var basicAuthenticationMethod *BasicAuthenticationMethod

func (brc BaseRouteConfig) Show() string {

	return fmt.Sprintf("clone=%v, webhook=%v, pull=%v, version=%v",
		brc.clonePath, brc.webhookPath, brc.pullPath, brc.versionPath)

}
func (brc BaseRouteConfig) GetClone() string {
	return brc.clonePath
}
func (brc BaseRouteConfig) GetWebHook() string {
	return brc.webhookPath
}
func (brc BaseRouteConfig) GetPull() string {
	return brc.pullPath
}
func (brc BaseRouteConfig) GetVersion() string {
	return brc.versionPath
}
func (brc BaseRepositoryConfig) Show() string {
	return fmt.Sprintf("repository=%v, branch=%v", brc.repoUrl, brc.branch)
}
func (brc BaseRepositoryConfig) GetRepo() string {
	return brc.repoUrl
}
func (brc BaseRepositoryConfig) GetBranch() string {
	return brc.branch
}
func (brc BaseRepositoryConfig) GetTargetFolder() string {
	return brc.targetFolder
}
func (bam BasicAuthenticationMethod) Show() string {
	return fmt.Sprintf("username=%v, password=******", bam.username)
}
func (bam BasicAuthenticationMethod) GetAuth() http.BasicAuth {
	return http.BasicAuth{
		Username: bam.username,
		Password: bam.passwordToken,
	}
}
func GetRouteConfigInstance() *BaseRouteConfig {
	if baseRouteConfigInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if baseRouteConfigInstance == nil {
			log.Println("[BaseRouteConfig] Creating new instance ")

			// TODO Replace with Environment Variables
			baseRouteConfigInstance = &BaseRouteConfig{
				clonePath:   os.Getenv("PATH_CLONE"),
				webhookPath: os.Getenv("PATH_WEBHOOK"),
				pullPath:    os.Getenv("PATH_PULL"),
				versionPath: os.Getenv("PATH_VERSION"),
			}
		} else {
			log.Println("[BaseRouteConfig] Instance already created")
		}
	} else {
		log.Println("[BaseRouteConfig] Instance already created")
	}
	log.Println(BaseRouteConfig.Show(*baseRouteConfigInstance))
	return baseRouteConfigInstance
}
func GetRepositoryConfigInstance() *BaseRepositoryConfig {
	if baseRepositoryConfigInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if baseRepositoryConfigInstance == nil {
			log.Println("[BaseRepositoryConfigInstance] Creating new instance")

			baseRepositoryConfigInstance = &BaseRepositoryConfig{
				repoUrl:      os.Getenv("REPO_URL"),
				branch:       os.Getenv("REPO_BRANCH"),
				targetFolder: os.Getenv("REPO_TARGET_FOLDER"),
			}

		} else {
			log.Println("[BaseRepositoryConfigInstance] Instance already created")
		}

	} else {
		log.Println("[BaseRepositoryConfigInstance] Instance already created")
	}
	log.Println(BaseRepositoryConfig.Show(*baseRepositoryConfigInstance))
	return baseRepositoryConfigInstance
}
func GetBasicAuthenticationMethodInstance() *BasicAuthenticationMethod {
	if basicAuthenticationMethod == nil {
		lock.Lock()
		defer lock.Unlock()

		if basicAuthenticationMethod == nil {
			log.Println("[BasicAuthenticationMethod] Creating new instance")

			basicAuthenticationMethod = &BasicAuthenticationMethod{
				username:      os.Getenv("REPO_USERNAME"),
				passwordToken: os.Getenv("REPO_PASSWORD"),
			}
		} else {
			log.Println("[BasicAuthenticationMethod] Instance already created")
		}
	} else {
		log.Println("[BasicAuthenticationMethod] Instance already created")
	}

	log.Println(BasicAuthenticationMethod.Show(*basicAuthenticationMethod))
	return basicAuthenticationMethod
}
