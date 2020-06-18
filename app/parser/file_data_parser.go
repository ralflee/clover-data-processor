package parser

import (
	"bufio"
	"clover-data-processor/app/constants"
	"clover-data-processor/app/model"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
)

//FileDataParser File data parser
type FileDataParser struct {
}

//ConstructSpec Construct spec
func (p *FileDataParser) ConstructSpec(specFilePath string) (*model.Spec, error) {
	log.Info().Str("filePath", specFilePath).Msg("Parsing spec file")

	csvfile, err := os.Open(specFilePath)
	if err != nil {
		log.Error().Str("filePath", specFilePath).Msg("Couldn't open the spec file")
		return nil, err
	}

	r := csv.NewReader(csvfile)

	//skip the header line
	r.Read()

	var s model.Spec
	s.Name = specFilePath[strings.LastIndex(specFilePath, "/")+1 : strings.LastIndex(specFilePath, ".")]

	//iterate through the records
	for {
		// Read each record from csv
		colDef, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Err(err).Send()
			return nil, err
		}

		//if len is not 3, skip the line
		if len(colDef) != 3 {
			log.Error().Msg("column definition not correct")
			continue
		}

		//pasre width
		width, err := strconv.Atoi(colDef[1])
		if err != nil {
			log.Err(err).Send()
			continue
		}
		if width < 1 {
			log.Err(err).Msg("width < 1")
			continue
		}

		s.Columns = append(s.Columns, &model.Column{Name: colDef[0], Width: width, Type: colDef[2]})
	}

	log.Info().Str("filePath", specFilePath).Msg("Construct spec success")

	return &s, nil
}

//ConstructRecords construct records
func (p *FileDataParser) ConstructRecords(dataFilePath string, spec *model.Spec) ([]*model.Record, error) {
	log.Info().Str("filePath", dataFilePath).Msg("Parsing data file")

	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Error().Str("filePath", dataFilePath).Msg("Couldn't open the data file")
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

	log.Info().Str("filePath", dataFilePath).Msg("Construct data success")

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
