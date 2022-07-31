package main

import (
	"fmt"
	"github.com/jarpsimoes/git-http-server/handlers"
	"github.com/jarpsimoes/git-http-server/utils"
	"log"
	"net/http"
)

func main() {
	routeConfig := utils.GetRouteConfigInstance()
	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetClone()), handlers.CloneHandler)
	http.HandleFunc("/", handlers.StaticContentHandler)
	// TODO - Make a port as environment variable
	log.Fatal(http.ListenAndServe(":8081", nil))
}
