package main

import (
	"fmt"

	"jerobas.com/ssh-tunnel-controller/repos"
)

func main() {
	// http.HandleFunc("/", routes.GetRoute)
	// http.HandleFunc("/create", routes.PostRoute)

	// fmt.Println("Server running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println(repos.GetTunnels())
}
