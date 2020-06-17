package main

import (
	"clover-data-processor/app/db"
	"clover-data-processor/app/parser"
	"clover-data-processor/app/repository"
	"clover-data-processor/app/source"
	"clover-data-processor/app/wire"

	"log"
)

func main() {
	//init DB connection
	dataSourceName := "postgres://test_user:12345678@localhost/test_db"
	db, err := db.InitDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.DataRepository{DB: db}
	fileDataParser := parser.FileDataParser{}
	dataSource := source.FileDataSource{}

	dataImportService := wire.InitDataImportService(&dataSource, &fileDataParser, &repo)
	dataImportService.ImportData()
}
