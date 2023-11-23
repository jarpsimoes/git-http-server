package handlers

import (
	"fmt"
	"github.com/jarpsimoes/git-http-server/utils"
	"html"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// StaticContentHandler it's a provider static content cloned from repository
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

	if repoConfig.GetRootFolder() != "" {

		basePath = fmt.Sprintf("%s/%s", repoConfig.GetRootFolder(), basePath)

		if strings.Contains(basePath, "//") {
			basePath = strings.ReplaceAll(basePath, "//", "/")
		}
	}

	log.Printf("[Request] Path %s", basePath)

	fs := http.FileServer(http.Dir(basePath))
	fs.ServeHTTP(w, r)
}

// CloneHandler it's a handler to clone source code from configured repository
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

// PullHandler it's a handler to update repository from configured repository
func PullHandler(w http.ResponseWriter, r *http.Request) {
	repoConfig := utils.GetRepositoryConfigInstance()
	commit := utils.PullRepository(repoConfig.GetRepo(), utils.BuildBranchPath(repoConfig.GetTargetFolder(), repoConfig.GetBranch()), repoConfig.GetBranch())

	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintf(w, "Branch %s pulled successfull \n", repoConfig.GetBranch())
	fmt.Fprintf(w, "Last commit [%s]", commit.ToString())
}

// FeatureNotEnabled it's a handler to response feature not enabled
func FeatureNotEnabled(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotAcceptable)

	fmt.Fprintf(w, "Operation %s not supported with present configuration \n", r.URL)

}

// HealthCheckHandler it's a handler to return server status
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthCheck := utils.GetHealthCheckControlInstance()

	w.Header().Add("Content-Type", "application/json")

	if healthCheck.IsHealthy() {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusExpectationFailed)
	}

	fmt.Fprintf(w, healthCheck.JSONHealthCheck())
}

// ServeReverseProxy reverse proxy handler
// Handle custom paths request
func ServeReverseProxy(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	exists, customPath := utils.FindPath(path)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	targetUrl, _ := url.Parse(customPath.GetTarget())

	if customPath.IsRewrite() {
		r.URL.Path = strings.ReplaceAll(path, fmt.Sprintf("/%s", customPath.GetPath()), "")
	}

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.Director = func(req *http.Request) {
		req.Header = r.Header
		req.Host = targetUrl.Host
		req.URL.Scheme = targetUrl.Scheme
		req.URL.Host = targetUrl.Host
		req.URL.Path = r.URL.Path

		log.Printf("Proxy to: %s \n", req.URL.String())
	}

	proxy.ServeHTTP(w, r)
}
