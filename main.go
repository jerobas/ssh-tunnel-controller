package main

import (
	"fmt"
	"log"
	"net/http"

	"jerobas.com/yepee/repos"
	"jerobas.com/yepee/routes"
)

func main() {
	repos.LoadRoutesVariable()

	http.HandleFunc("/", routes.GetRoute)
	http.HandleFunc("/create", routes.PostRoute)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
