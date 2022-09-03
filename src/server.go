package main

import (
	"fmt"
	"github.com/jarpsimoes/git-http-server/handlers"
	"github.com/jarpsimoes/git-http-server/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	routeConfig := utils.GetRouteConfigInstance()
	healthCheckControl := utils.GetHealthCheckControlInstance()

	repo := utils.GetRepositoryConfigInstance()
	customPaths := utils.GetCustomPathsInstance()

	utils.CloneRepository(repo.GetRepo(), repo.GetBranch(), repo.GetTargetFolder(), true)

	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetClone()), handlers.CloneHandler)
	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetPull()), handlers.PullHandler)
	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetHealthCheck()), handlers.HealthCheckHandler)
	if len(*customPaths) > 0 {

		for _, customPath := range *customPaths {
			http.HandleFunc(fmt.Sprintf("/%s/", customPath.GetPath()), handlers.ServeReverseProxy)
			log.Printf("[CUSTOM_PATH] /%s with target: %s", customPath.GetPath(), customPath.GetTarget())
		}

	}
	http.HandleFunc("/", handlers.StaticContentHandler)

	port := os.Getenv("HTTP_PORT")

	log.Printf("[STARTED] Listen port %s \n", port)
	healthCheckControl.UpdateState(true)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
