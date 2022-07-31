package handlers

import (
	"fmt"
	"github.com/jarpsimoes/git-http-server/utils"
	"html"
	"log"
	"net/http"
)

func StaticContentHandler(w http.ResponseWriter, r *http.Request) {
	repoConfig := utils.GetRepositoryConfigInstance()
	keys, ok := r.URL.Query()["_branch"]

	branch := repoConfig.GetBranch()

	if !ok || len(keys[0]) < 1 {
		log.Printf("Default branch %s \n", branch)
	} else {
		branch = keys[0]
	}

	basePath := utils.BuildBranchPath(repoConfig.GetTargetFolder(), branch)
	log.Printf("[Request] Path %s", basePath)

	fs := http.FileServer(http.Dir(basePath))
	fs.ServeHTTP(w, r)
}
func CloneHandler(w http.ResponseWriter, r *http.Request) {
	routeConfig := utils.GetRouteConfigInstance()
	repoConfig := utils.GetRepositoryConfigInstance()

	path := html.EscapeString(r.URL.Path)
	cloneMain := path[len(fmt.Sprintf("/%v/", routeConfig.GetClone())):len(path)]

	if cloneMain == "" {
		cloneMain = "main"
	}
	commit := utils.CloneRepository(repoConfig.GetRepo(), cloneMain, repoConfig.GetTargetFolder(), true)

	if commit == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Branch %v not found \n", cloneMain)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Source cloned [branch: %v] \n", cloneMain)
	fmt.Fprintf(w, "Last Commit [%s]", commit.ToString())
}
