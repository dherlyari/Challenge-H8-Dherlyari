package main

import (
	"Challenge-H8-Dherlyari/cli"
	"Challenge-H8-Dherlyari/config"
)

func main() {
	config.DBConnect()
	cli.MainMenu()
}
