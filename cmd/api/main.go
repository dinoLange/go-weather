package main

import (
	"fmt"
	"go-weather/internal/server"
)

func main() {

	server := server.NewServer()
	fmt.Println("Server started!")
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
