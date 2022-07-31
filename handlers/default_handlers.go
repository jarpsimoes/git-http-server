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
