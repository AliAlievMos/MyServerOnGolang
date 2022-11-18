package main

import (
	_ "fmt"

	"MyServer/server"
)

func main() {

	go func() {
		server.Server()
	}()

}
