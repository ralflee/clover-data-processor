package parser

import (
	"bufio"
	"clover-data-processor/app/constants"
	"clover-data-processor/app/model"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//GetSpecFilesFromPath get files from path
func GetSpecFilesFromPath(path string) []string {
	var files []string
	var validFile = regexp.MustCompile(`.*\.csv$`)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		if validFile.MatchString(fileName) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	return files
}

//GetDataFilesFromPath get files from path
func GetDataFilesFromPath(path string, specName string) []string {
	var files []string
	var validFile = regexp.MustCompile(specName + `_[0-9]{4}-[0-9]{2}-[0-9]{2}.*\.txt$`)
	var dateFormat = "2006-01-02"

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		//skip folders
		if info.IsDir() {
			return nil
		}

		fileName := info.Name()
		//skip file which the name is not valid
		if validFile.MatchString(fileName) {

			_, err := time.Parse(dateFormat, fileName[strings.LastIndex(fileName, "_")+1:strings.LastIndex(fileName, ".")])

			//skip file which the name format is not valid
			if err != nil {
				fmt.Print(err)
				return nil
			}
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	return files
}

//ConstructSpec construct spec from file path
func ConstructSpec(filePath string) (*model.Spec, error) {
	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Couldn't open the spec file")
	}

	r := csv.NewReader(csvfile)

	//skip the header line
	r.Read()

	var s model.Spec
	s.Name = filePath[strings.LastIndex(filePath, "/")+1 : strings.LastIndex(filePath, ".")]

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

//ConstructRecords construct record
func ConstructRecords(dataFilePath string, spec *model.Spec) ([]*model.Record, error) {
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
		//c := ""

		// for len(b) > 0 {
		// 	r, size := utf8.DecodeRune(b)
		// 	c += string(r)
		// 	b = b[size:len(b)]

		// }
		record := bytesToRecord(b, spec)
		records = append(records, record)
		//fmt.Println(row)
		//fmt.Println(len(row))
		//fmt.Print(row[0:10])
		//fmt.Print(",")
		//fmt.Print(row[11:12])
		//fmt.Print(",")
		//fmt.Print(row[13:15])

		// /start := 0
		//fmt.Println(spec.Columns)
		// for _, col := range spec.Columns {
		// 	fmt.Print(row[start : start+col.Width])
		// 	fmt.Print(",")
		// 	start += col.Width
		// }

	}

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

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
