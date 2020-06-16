package parser

import (
	"bufio"
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
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		files = append(files, path)

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

//ConstructRecord construct record
func ConstructRecord(dataFilePath string, spec *model.Spec) (*model.Record, error) {
	file, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := scanner.Bytes()
		c := ""

		for len(b) > 0 {
			r, size := utf8.DecodeRune(b)
			c += string(r)
			b = b[size:len(b)]

		}
		fmt.Print(c)
		fmt.Println("")
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
