package parser_test

import (
	"clover-data-processor/app/constants"
	"clover-data-processor/app/model"
	"clover-data-processor/app/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileDataParser(t *testing.T) {
	parser := parser.FileDataParser{}

	t.Run("ConstructSpecSuccess", func(t *testing.T) {
		spec, err := parser.ConstructSpec("../../test/specs/testformat1.csv")
		assert.Nil(t, err)
		assert.Equal(t, 3, len(spec.Columns))

		//test parse TEXT
		assert.Equal(t, "name", spec.Columns[0].Name)
		assert.Equal(t, 10, spec.Columns[0].Width)
		assert.Equal(t, constants.ColumnTypeText, spec.Columns[0].Type)

		//test parse BOOLEAN
		assert.Equal(t, "valid", spec.Columns[1].Name)
		assert.Equal(t, 1, spec.Columns[1].Width)
		assert.Equal(t, constants.ColumnTypeBoolean, spec.Columns[1].Type)

		//test parse INTEGER
		assert.Equal(t, "count", spec.Columns[2].Name)
		assert.Equal(t, 3, spec.Columns[2].Width)
		assert.Equal(t, constants.ColumnTypeInteger, spec.Columns[2].Type)
	})

	t.Run("ConstructRecords", func(t *testing.T) {
		spec := model.Spec{}

		spec.Name = "testformat1"
		spec.Columns = make([]*model.Column, 3)
		spec.Columns[0] = &model.Column{Name: "name", Width: 10, Type: constants.ColumnTypeText}
		spec.Columns[1] = &model.Column{Name: "valid", Width: 1, Type: constants.ColumnTypeBoolean}
		spec.Columns[2] = &model.Column{Name: "count", Width: 3, Type: constants.ColumnTypeInteger}
		// spec.Columns = [3]*model.Column{
		// 	&model.Column{Name: "name", Width: 10, Type: constants.ColumnTypeText},
		// 	&model.Column{Name: "valid", Width: 1, Type: constants.ColumnTypeBoolean},
		// 	&model.Column{Name: "count", Width: 3, Type: constants.ColumnTypeInteger},
		// }

		records, err := parser.ConstructRecords("../../test/data/testformat1_2015-06-28.txt", &spec)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(records))

		assert.Equal(t, "Foonyor", records[0].Columns[0])
		assert.Equal(t, true, records[0].Columns[1])
		assert.Equal(t, 1, records[0].Columns[2])
	})
}
