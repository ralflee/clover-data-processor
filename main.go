package main

import (
	"fmt"

	"clover-data-processor/app/db"
	"clover-data-processor/app/model"
	"clover-data-processor/app/parser"
	"clover-data-processor/app/repository"

	"log"
)

func main() {
	//init DB connection
	dataSourceName := "postgres://test_user:12345678@localhost/test_db"
	db, err := db.InitDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	r := repository.DataRepository{DB: db}

	//get spec files
	specFiles := parser.GetSpecFilesFromPath("./specs")
	for _, file := range specFiles {

		fmt.Println(file)

	}

	specs := make([]*model.Spec, len(specFiles))
	for i, filePath := range specFiles {
		//parse spec files
		spec, err := parser.ConstructSpec(filePath)
		if err != nil {
			log.Print(err)
		}

		specs[i] = spec
	}

	for _, s := range specs {

		//create DB table
		if !r.CheckTableExists(s.Name) {
			err := r.CreateTable(s)
			if err != nil {
				log.Fatal(err)
				continue
			}
		}

		//get data files
		dataFiles := parser.GetDataFilesFromPath("./data", s.Name)
		fmt.Println(dataFiles)

		for _, file := range dataFiles {

			//parse Records
			records, err := parser.ConstructRecords(file, s)
			if err != nil {
				log.Fatal(err)
			}

			//insert records
			r.Insert(s, records)

		}

	}
}
