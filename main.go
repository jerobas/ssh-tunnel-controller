package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jerobas/territo/config"
	"github.com/jerobas/territo/handlers"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/", handlers.MainRoute)

	fmt.Printf("Server running at http://localhost:%d\n", config.GetConfig().Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.GetConfig().Port), nil))
}
