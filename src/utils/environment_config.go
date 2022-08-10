package utils

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

// BaseRouteConfig it's a struct to represent routes configuration
// Singleton
type BaseRouteConfig struct {
	clonePath   string
	webhookPath string
	pullPath    string
	versionPath string
}

// BaseRepositoryConfig it's a struct to represent repository configuration
// Singleton
type BaseRepositoryConfig struct {
	repoUrl      string
	branch       string
	targetFolder string
}

// BasicAuthenticationMethod it's a struct to represent authorization configuration
// Singleton
type BasicAuthenticationMethod struct {
	username      string
	passwordToken string
}

var baseRouteConfigInstance *BaseRouteConfig
var baseRepositoryConfigInstance *BaseRepositoryConfig
var basicAuthenticationMethod *BasicAuthenticationMethod

// Show (BaseRouteConfig) it's a function print strut content as string
func (brc BaseRouteConfig) Show() string {
	return fmt.Sprintf("clone=%v, webhook=%v, pull=%v, version=%v",
		brc.clonePath, brc.webhookPath, brc.pullPath, brc.versionPath)
}

// GetClone (BaseRouteConfig) it's a method to get Clone path configured
func (brc BaseRouteConfig) GetClone() string {
	return brc.clonePath
}

// GetWebHook (BaseRouteConfig) it's a method to get Webhook path configured
func (brc BaseRouteConfig) GetWebHook() string {
	return brc.webhookPath
}

// GetPull (BaseRouteConfig) it's a method to get Pull path configured
func (brc BaseRouteConfig) GetPull() string {
	return brc.pullPath
}

// GetVersion (BaseRouteConfig) it's a method to get Version path configured
func (brc BaseRouteConfig) GetVersion() string {
	return brc.versionPath
}

// Show (BaseRepositoryConfig) it's a function print strut content as string
func (brc BaseRepositoryConfig) Show() string {
	return fmt.Sprintf("repository=%v, branch=%v", brc.repoUrl, brc.branch)
}

// GetRepo (BaseRouteConfig) it's a method to get configured repository
func (brc BaseRepositoryConfig) GetRepo() string {
	return brc.repoUrl
}

// GetBranch (BaseRepositoryConfig) it's a method to get configured default branch
func (brc BaseRepositoryConfig) GetBranch() string {
	return brc.branch
}

// GetTargetFolder (BaseRepositoryConfig) it's a method to get configured target folder
func (brc BaseRepositoryConfig) GetTargetFolder() string {
	return brc.targetFolder
}

// Show (BasicAuthenticationMethod) it's a function print strut content as string
func (bam BasicAuthenticationMethod) Show() string {
	return fmt.Sprintf("username=%v, password=******", bam.username)
}

// BasicAuthAvailable (BasicAuthenticationMethod) return true if Basic Authentication config
// is available
func (bam BasicAuthenticationMethod) BasicAuthAvailable() bool {
	return bam.username != "" && bam.passwordToken != ""
}

// GetAuth (BasicAuthenticationMethod) return basic auth data
func (bam BasicAuthenticationMethod) GetAuth() http.BasicAuth {
	return http.BasicAuth{
		Username: bam.username,
		Password: bam.passwordToken,
	}
}

// GetRouteConfigInstance return singleton instance
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

// GetRepositoryConfigInstance return singleton instance
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

// GetBasicAuthenticationMethodInstance return singleton instance
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
