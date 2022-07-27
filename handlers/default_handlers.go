package handlers

import (
	"fmt"
	"html"
	"jarpsystems.net/server/utils"
	"net/http"
)

func CloneHandler(w http.ResponseWriter, r *http.Request) {
	routeConfig := utils.GetRouteConfigInstance()
	repoConfig := utils.GetRepositoryConfigInstance()

	path := html.EscapeString(r.URL.Path)
	cloneMain := path[len(fmt.Sprintf("/%v/", routeConfig.GetClone())):len(path)]

	fmt.Fprintf(w, "Source cloned [branch: %v]", cloneMain)

	if cloneMain == "" {
		cloneMain = "main"
	}
	utils.CloneRepository(repoConfig.GetRepo(), cloneMain, repoConfig.GetTargetFolder())
}
