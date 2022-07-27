package main

import (
	"fmt"
	"jarpsystems.net/server/handlers"
	"jarpsystems.net/server/utils"
	"log"
	"net/http"
)

func main() {
	routeConfig := utils.GetRouteConfigInstance()
	http.HandleFunc(fmt.Sprintf("/%s/", routeConfig.GetClone()), handlers.CloneHandler)

	// TODO - Make a port as environment variable
	log.Fatal(http.ListenAndServe(":8081", nil))
}
