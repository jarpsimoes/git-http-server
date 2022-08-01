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

	repo := utils.GetRepositoryConfigInstance()

	utils.CloneRepository(repo.GetRepo(), repo.GetBranch(), repo.GetTargetFolder(), true)
	
	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetClone()), handlers.CloneHandler)
	http.HandleFunc("/", handlers.StaticContentHandler)
	port := os.Getenv("HTTP_PORT")
	log.Printf("[STARTED] Listen port %s \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
