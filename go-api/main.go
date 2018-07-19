package main

import (
	"fmt"
	"github.com/ccdle12/Blocksage/go-api/dependencyinjector"
	"log"
	"net/http"
)

func main() {
	injector := dependencyinjector.DependencyInjector{}
	API := injector.InjectAPI()

	API.Router.HandleFunc("/api/network-info", API.NetworkInfo).Methods("GET")

	fmt.Println("Serving on Port 8548...")
	log.Fatal(http.ListenAndServe(":8548", API.Router))
}
