package main

import (
	"v1/db"
	"v1/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}