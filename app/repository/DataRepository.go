package repository

import (
	"clover-data-processor/app/model"
	"database/sql"
	"fmt"

	//"log"
	"clover-data-processor/app/constants"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

//DataRepository data repository
type DataRepository struct {
	DB *sql.DB
}

//CheckTableExists Check table exists
func (r *DataRepository) CheckTableExists(tableName string) bool {
	sql := "select 1 from " + tableName

	if viper.GetBool("app.logSQL") {
		log.Debug().Str("sql", sql).Send()
	}

	_, err := r.DB.Query(sql)
	if err != nil {
		return false
	}

	return true
}

//CreateTable Create table
func (r *DataRepository) CreateTable(spec *model.Spec) error {
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

	if viper.GetBool("app.logSQL") {
		log.Debug().Str("sql", sql).Send()
	}

	_, err := r.DB.Exec(sql)
	if err != nil {
		log.Err(err).Send()
		return err
	}

	log.Printf("db table %s created", spec.Name)
	return nil
}

//Insert insert record into database
func (r *DataRepository) Insert(spec *model.Spec, records []*model.Record) error {

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

	if viper.GetBool("app.logSQL") {
		log.Debug().Str("sql", sql).Send()
	}

	_, err := r.DB.Exec(sql)
	if err != nil {
		log.Err(err).Send()
		return err
	}

	return nil
}
