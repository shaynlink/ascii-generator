package main

import (
	"os"
	"fmt"

	"github.com/discord/fasthttp"
)

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if port == "" {
		port = "8080"
	}

	if host == "" {
		// You can use 0.0.0.0 if you want to access the server from outside
		host = "127.0.0.1"
	}

	fmt.Println("Server started on port " + port)

	if err := fasthttp.ListenAndServe(host + ":" + port, Handler); err != nil {
		panic(err)
	}
}


