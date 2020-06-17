package service

import (
	"clover-data-processor/app/model"
	"fmt"
	"log"
)

//RawDataSource Raw data source interface
type RawDataSource interface {
	GetSpecPath(basePath string) []string
	GetDataPath(basePath string, specName string) []string
}

//DataParser Data parser interface
type DataParser interface {
	ConstructSpec(specFilePath string) (*model.Spec, error)
	ConstructRecords(dataFilePath string, spec *model.Spec) ([]*model.Record, error)
}

//DataRepository Data repository interface
type DataRepository interface {
	CheckTableExists(tableName string) bool
	CreateTable(spec *model.Spec) error
	Insert(spec *model.Spec, records []*model.Record) error
}

//DataImportService Data import service
type DataImportService struct {
	DataSource RawDataSource
	Parser     DataParser
	Repository DataRepository
}

//ImportData Import data
func (s *DataImportService) ImportData() error {
	//get spec files
	specFiles := s.DataSource.GetSpecPath("./specs")
	for _, file := range specFiles {

		fmt.Println(file)

	}

	specs := make([]*model.Spec, len(specFiles))
	for i, filePath := range specFiles {
		//parse spec files
		spec, err := s.Parser.ConstructSpec(filePath)
		if err != nil {
			log.Print(err)
		}

		specs[i] = spec
	}

	for _, spec := range specs {

		//create DB table
		if !s.Repository.CheckTableExists(spec.Name) {
			err := s.Repository.CreateTable(spec)
			if err != nil {
				log.Fatal(err)
				continue
			}
		}

		//get data files
		dataFiles := s.DataSource.GetDataPath("./data", spec.Name)
		fmt.Println(dataFiles)

		for _, file := range dataFiles {

			//parse Records
			records, err := s.Parser.ConstructRecords(file, spec)
			if err != nil {
				log.Fatal(err)
			}

			//insert records
			s.Repository.Insert(spec, records)
		}
	}

	return nil
}
