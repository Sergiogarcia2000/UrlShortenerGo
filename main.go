package main

import (
	"UrlShortenerGoLang/api"
	"UrlShortenerGoLang/storage"
	"fmt"
)

func main() {

	store := storage.NewMemoryUrlStore()

	server := api.NewServer(":8080", store)
	fmt.Println("Server running on port 8080")
	server.Start()
}
