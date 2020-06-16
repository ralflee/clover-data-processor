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
	dataSourceName := "postgres://test_user:12345678@localhost/test_db"
	db, err := db.InitDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	r := repository.DataRepository{DB: db}

	specFiles := parser.GetSpecFilesFromPath("./specs")
	for _, file := range specFiles {

		fmt.Println(file)

	}

	// dataFiles := parser.GetFilesFromPath("./data")
	// for _, file := range dataFiles {

	// 	fmt.Println(file)

	// }

	//print file content
	// for _, file := range files {
	// 	dat, err := ioutil.ReadFile(file)

	// 	if err != nil {
	// 		fmt.Println(err)

	// 	}
	// 	fmt.Print(string(dat))
	// }

	//print flie content line by line
	// for _, f := range files {
	// 	file, err := os.Open(f)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer file.Close()

	// 	scanner := bufio.NewScanner(file)
	// 	for scanner.Scan() {
	// 		fmt.Println(scanner.Text())
	// 	}

	// 	if err := scanner.Err(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	specs := make([]*model.Spec, len(specFiles))
	for i, file := range specFiles {
		spec, err := parser.ConstructSpec(file)
		if err != nil {
			log.Print(err)
		}

		specs[i] = spec
		//fmt.Println(specs[0].Name)
		//fmt.Println(specs[0].Columns[0].Name)
		//fmt.Println(specs[0].Columns[0].Width)
		//fmt.Println(specs[0].Columns[0].Type)

		// csvfile, err := os.Open(file)
		// if err != nil {
		// 	log.Fatalln("Couldn't open the csv file", err)
		// }

		// r := csv.NewReader(csvfile)

		// //skip the header line
		// r.Read()

		// // Iterate through the records
		// for {
		// 	// Read each record from csv
		// 	record, err := r.Read()
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	fmt.Println(len(record))

		// 	fmt.Printf("%s %s %s\n", record[0], record[1], record[2])
		// }
	}

	//create db table
	for _, s := range specs {
		if !r.CheckTableExists(s.Name) {
			r.CreateTable(s)
		}

	}

	for _, s := range specs {
		dataFiles := parser.GetDataFilesFromPath("./data", s.Name)
		fmt.Println(dataFiles)

		for _, file := range dataFiles {
			record, err := parser.ConstructRecord(file, s)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(record)
		}

	}

}
