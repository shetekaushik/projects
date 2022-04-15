package main

import (
	"projects/project2/route"
)

func main() {
	server := route.InitServer()
	server.Logger.Fatal(server.Start(":9000"))
}
