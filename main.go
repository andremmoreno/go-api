package main

import "github.com/andremmoreno/go-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
