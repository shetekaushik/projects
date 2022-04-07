package main

import (
	"projects/project1/route"
)

func main() {
	server := route.InitServer()
	server.Logger.Fatal(server.Start(":8000"))
}
