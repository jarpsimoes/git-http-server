package utils

import (
	"fmt"
	"log"
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

var baseRouteConfigInstance *BaseRouteConfig
var baseRepositoryConfigInstance *BaseRepositoryConfig

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
func GetRouteConfigInstance() *BaseRouteConfig {
	if baseRouteConfigInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if baseRouteConfigInstance == nil {
			log.Println("[BaseRouteConfig] Creating new instance ")

			// TODO Replace with Environment Variables
			baseRouteConfigInstance = &BaseRouteConfig{
				clonePath:   "_clone",
				webhookPath: "_hook",
				pullPath:    "_pull",
				versionPath: "_version",
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

			// TODO Replace with Environment Variables
			baseRepositoryConfigInstance = &BaseRepositoryConfig{
				repoUrl:      "https://github.com/jarpsimoes/git-http-server.git",
				branch:       "main",
				targetFolder: "target-git",
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
