package main

import (
	"get.cutie.cafe/rainy/server"
	"get.cutie.cafe/rainy/upload/filesystem"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	filesystem.New("hello")

	server.New().Listen()
}
