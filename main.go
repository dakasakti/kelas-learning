package main

import (
	"km-kelas-e/config"
	"km-kelas-e/database/migrate"
	"km-kelas-e/database/seeders"
	"km-kelas-e/routes"
)

func main() {
	//initiateDB
	config.InitDB()
	migrate.AutoMigrate()
	seeders.SetArticle()

	//initRoutes
	e := routes.New()

	//starting the server
	e.Start(":1234")
}
