package main

import (
	"leong-01-goapp/config"
	"leong-01-goapp/controller"
	mysql "leong-01-goapp/db/mysql"
)

func main() {
	config.Init()
	mysql.Init()
	controller.Init()
}
