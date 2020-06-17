package parser

import (
	"bufio"
	"clover-data-processor/app/constants"
	"clover-data-processor/app/model"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

//FileDataParser File data parser
type FileDataParser struct {
}

//ConstructSpec Construct spec
func (p *FileDataParser) ConstructSpec(specFilePath string) (*model.Spec, error) {
	csvfile, err := os.Open(specFilePath)
	if err != nil {
		return nil, errors.New("Couldn't open the spec file")
	}

	r := csv.NewReader(csvfile)

	//skip the header line
	r.Read()

	var s model.Spec
	s.Name = specFilePath[strings.LastIndex(specFilePath, "/")+1 : strings.LastIndex(specFilePath, ".")]

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//TODO, if len is not 3, skip the line
		width, err := strconv.Atoi(record[1])
		if err != nil {

			continue
		}
		if width < 1 {
			log.Fatal("width < 1")
		}

		s.Columns = append(s.Columns, &model.Column{Name: record[0], Width: width, Type: record[2]})
	}

	return &s, nil
}

//ConstructRecords construct records
func (p *FileDataParser) ConstructRecords(dataFilePath string, spec *model.Spec) ([]*model.Record, error) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	var records []*model.Record
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := scanner.Bytes()

		record := bytesToRecord(b, spec)
		records = append(records, record)
	}

	return records, nil
}

func bytesToRecord(bytes []byte, spec *model.Spec) *model.Record {
	var record model.Record
	record.Columns = make([]interface{}, len(spec.Columns))
	for i, col := range spec.Columns {
		c := ""
		for j := 0; j < col.Width; j++ {
			r, size := utf8.DecodeRune(bytes)
			c += string(r)
			bytes = bytes[size:]
		}

		if col.Type == constants.ColumnTypeBoolean {
			record.Columns[i] = c == "1"
		} else if col.Type == constants.ColumnTypeInteger {
			v, err := strconv.Atoi(strings.TrimSpace(c))
			if err != nil {
				continue
			}

			record.Columns[i] = v
		} else {
			//trim white spaces
			record.Columns[i] = strings.TrimSpace(c)
		}

	}

	return &record
}
