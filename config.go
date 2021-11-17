package main

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
)

type DBConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	sslMode    string
}


func (a *App) InitializeConfig(){
	config := DBConfig{}

	log.Println("In Consul config")
	vi := viper.New()
	err := vi.AddRemoteProvider("consul", "localhost:8500","local/beerdb/dsn")
	if err != nil {
		log.Fatal("Crap!!!!")
	}
	vi.SetConfigType("json")
	err = vi.ReadRemoteConfig()
	if err != nil {
		log.Fatal("Oh crap!!!  But closer")
	}

	config.dbUser = vi.GetString("dbUser")
	config.dbPassword = vi.GetString("dbPassword")
	config.sslMode = vi.GetString("sslMode")
	config.dbName = vi.GetString("dbName")
	config.dbHost = vi.GetString("dbHost")
	config.dbPort = vi.GetString("dbPort")
	a.DBConfig = config


}
