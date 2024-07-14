package main

import (
	"football-simulation/cmd/api"
	"football-simulation/config"
	"football-simulation/database"
	"log"
)

func main() {

	dbConfig := database.DBConfig{
		User:     config.Envs.User,
		Password: config.Envs.Password,
		DBName:   config.Envs.DBName,
		Host:     config.Envs.Host,
		DBPort:   config.Envs.DBPort,
		SSLMode:  config.Envs.SSLMode,
	}

	db, err := database.NewPostgreSQLStorage(dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
