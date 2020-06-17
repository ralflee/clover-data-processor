package main

import (
	"clover-data-processor/app/config"
	"clover-data-processor/app/db"
	"clover-data-processor/app/parser"
	"clover-data-processor/app/repository"
	"clover-data-processor/app/source"
	"clover-data-processor/app/wire"

	"log"

	"github.com/spf13/viper"
)

func main() {
	//init config
	config.InitAppConfig("./app/config")

	//init DB connection
	dbConnectionURL := viper.GetString("app.dbConnectionURL")
	db, err := db.InitDB(dbConnectionURL)
	if err != nil {
		log.Fatal(err)
	}

	//wire the data import service
	repo := repository.DataRepository{DB: db}
	fileDataParser := parser.FileDataParser{}
	dataSource := source.FileDataSource{}
	dataImportService := wire.InitDataImportService(&dataSource, &fileDataParser, &repo)

	//import data
	dataImportService.ImportData()
}
