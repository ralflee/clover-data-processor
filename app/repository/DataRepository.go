package repository

import (
	"clover-data-processor/app/model"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"clover-data-processor/app/constants"
)

type DataRepository struct {
	DB *sql.DB
}

func (r DataRepository) CheckTableExists(tableName string) bool {
	//var check string

	//r.DB.QueryRow("SELECT 1 from '"+tableName+"'", tableName).Scan(check)
	fmt.Println("aaa")
	_, err := r.DB.Query("select 1 from " + tableName)
	if err != nil {
		return false
	}

	return true
}

//CreateTable Create table
func (r DataRepository) CreateTable(spec *model.Spec) error {
	sql := "create table " + spec.Name + " ("

	cols := make([]string, len(spec.Columns))
	for i, col := range spec.Columns {
		var colType string
		switch col.Type {
		case constants.ColumnTypeText:
			colType = "text"
		case constants.ColumnTypeBoolean:
			colType = "boolean"
		case constants.ColumnTypeInteger:
			colType = "int8"
		}

		cols[i] = col.Name + " " + colType
	}

	sql += strings.Join(cols, ",")
	sql += ")"

	_, err := r.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("db table %s created", spec.Name)
	return nil
}

//Insert insert record into database
func (r DataRepository) Insert(spec *model.Spec, records []*model.Record) error {

	sql := "insert into " + spec.Name + " ("

	cols := make([]string, len(spec.Columns))
	for i, col := range spec.Columns {
		cols[i] = col.Name
	}

	sql += strings.Join(cols, ",")
	sql += ") values "

	values := make([]string, len(records))

	for i, r := range records {
		valueArr := make([]string, len(r.Columns))
		for j, c := range spec.Columns {
			if c.Type == constants.ColumnTypeText {
				valueArr[j] = fmt.Sprintf("'%v'", r.Columns[j])
			} else {
				valueArr[j] = fmt.Sprintf("%v", r.Columns[j])
			}

		}

		values[i] = "(" + strings.Join(valueArr, ",") + ")"
	}

	sql += strings.Join(values, ",")

	fmt.Println(sql)
	_, err := r.DB.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
