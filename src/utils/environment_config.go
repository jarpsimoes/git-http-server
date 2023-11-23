package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

// BaseRouteConfig it's a struct to represent routes configuration
// Singleton
type BaseRouteConfig struct {
	clonePath          string
	webhookPath        string
	pullPath           string
	versionPath        string
	healthCheckControl string
}

// BaseRepositoryConfig it's a struct to represent repository configuration
// Singleton
type BaseRepositoryConfig struct {
	enabled      bool
	repoURL      string
	branch       string
	targetFolder string
	rootFolder   string
}

// BasicAuthenticationMethod it's a struct to represent authorization configuration
// Singleton
type BasicAuthenticationMethod struct {
	username      string
	passwordToken string
}

// HealthCheckControl it's a struct to represent health check result
type HealthCheckControl struct {
	Status           bool
	Port             string
	StartTime        string
	StatusUpdateTime string
}

var baseRouteConfigInstance *BaseRouteConfig
var baseRepositoryConfigInstance *BaseRepositoryConfig
var basicAuthenticationMethod *BasicAuthenticationMethod
var healthCheckControl *HealthCheckControl
var pathSecurityCheck *PathSecurityCheck
var basePath, _ = os.Getwd()

// UpdateState [HealthCheckControl] it's a function to update Status
func (hcc *HealthCheckControl) UpdateState(status bool) {
	hcc.Status = status
	hcc.StatusUpdateTime = time.Now().String()

	healthCheckControl = hcc
}

// JSONHealthCheck [HealthCheckControl] it's a function to get heath check as a json string
func (hcc HealthCheckControl) JSONHealthCheck() string {
	jsonContent, _ := json.Marshal(hcc)

	return string(jsonContent)
}

// IsHealthy [HealthCheckControl] it's a function to check if server is healthy
func (hcc HealthCheckControl) IsHealthy() bool {
	return hcc.Status
}

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

// GetHealthCheck (BaseRouteConfig) it's a method to get Version path configured
func (brc BaseRouteConfig) GetHealthCheck() string {
	return brc.healthCheckControl
}

// Show (BaseRepositoryConfig) it's a function print strut content as string
func (brc BaseRepositoryConfig) Show() string {
	return fmt.Sprintf("repository=%v, branch=%v", brc.repoURL, brc.branch)
}

// GetRepo (BaseRouteConfig) it's a method to get configured repository
func (brc BaseRepositoryConfig) GetRepo() string {
	return brc.repoURL
}

// GetBranch (BaseRepositoryConfig) it's a method to get configured default branch
func (brc BaseRepositoryConfig) GetBranch() string {
	return brc.branch
}

// GetTargetFolder (BaseRepositoryConfig) it's a method to get configured target folder
func (brc BaseRepositoryConfig) GetTargetFolder() string {
	return brc.targetFolder
}

// GetRootFolder (baseRepositoryConfig) it's a method to get root folder
func (brc BaseRepositoryConfig) GetRootFolder() string {
	return brc.rootFolder
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

			baseRouteConfigInstance = &BaseRouteConfig{
				clonePath:          os.Getenv("PATH_CLONE"),
				webhookPath:        os.Getenv("PATH_WEBHOOK"),
				pullPath:           os.Getenv("PATH_PULL"),
				versionPath:        os.Getenv("PATH_VERSION"),
				healthCheckControl: os.Getenv("PATH_HEALTH"),
			}

		}
	}
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
				enabled:      os.Getenv("REPO_URL") != "",
				repoURL:      os.Getenv("REPO_URL"),
				branch:       os.Getenv("REPO_BRANCH"),
				targetFolder: os.Getenv("REPO_TARGET_FOLDER"),
				rootFolder:   basePath,
			}

			if pathSecurityCheck.IsValidPath(baseRepositoryConfigInstance.targetFolder) {
				log.Printf("[BaseRepositoryConfigInstance] Target folder %v is valid", baseRepositoryConfigInstance.targetFolder)
			} else {
				log.Printf("[BaseRepositoryConfigInstance] Target folder %v is invalid", baseRepositoryConfigInstance.targetFolder)
				log.Printf("[BaseRepositoryConfigInstance] Will be changed to target")
				baseRepositoryConfigInstance.targetFolder = "target_folder"
			}
		}
	}

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
		}
	}

	return basicAuthenticationMethod
}

// GetHealthCheckControlInstance return singleton instance
func GetHealthCheckControlInstance() *HealthCheckControl {
	if healthCheckControl == nil {
		lock.Lock()
		defer lock.Unlock()
	}

	if healthCheckControl == nil {
		log.Println("[HealthCheckControl] Creating new instance")

		currentTime := time.Now()

		healthCheckControl = &HealthCheckControl{
			Status:           false,
			Port:             os.Getenv("HTTP_PORT"),
			StartTime:        currentTime.String(),
			StatusUpdateTime: currentTime.String(),
		}
	}

	return healthCheckControl
}
