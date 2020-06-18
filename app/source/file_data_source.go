package source

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//FileDataSource File data source
type FileDataSource struct {
}

//GetSpecPath get files from path
func (s *FileDataSource) GetSpecPath(basePath string) []string {
	var files []string
	var validFile = regexp.MustCompile(`.*\.csv$`)
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {

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
		log.Err(err).Send()
		return []string{}
	}

	logArray := zerolog.Arr()
	for _, f := range files {
		logArray.Str(f)
	}
	log.Info().Array("specFiles", logArray).Msg("Get spec files success")

	return files
}

//GetDataPath get files from path
func (s *FileDataSource) GetDataPath(basePath string, specName string) []string {
	var files []string
	var validFile = regexp.MustCompile(specName + `_[0-9]{4}-[0-9]{2}-[0-9]{2}.*\.txt$`)
	var dateFormat = "2006-01-02"

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {

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
				log.Err(err).Send()
				return nil
			}
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Err(err).Send()
		return []string{}
	}

	logArray := zerolog.Arr()
	for _, f := range files {
		logArray.Str(f)
	}
	log.Info().Array("specFiles", logArray).Msg("Get data files success")

	return files
}
