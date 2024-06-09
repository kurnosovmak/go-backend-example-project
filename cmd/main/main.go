package main

import (
	"example/cmd/routes"
	"fmt"
	"net/http"
)

func main() {

	routes.RegisterRoutes()
	fmt.Println("Starting server at port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
	}
}
