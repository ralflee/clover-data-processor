package wire

import "clover-data-processor/app/service"

//InitDataImportService Initialize data import service
func InitDataImportService(dataSource service.RawDataSource, parser service.DataParser, repository service.DataRepository) *service.DataImportService {
	return &service.DataImportService{DataSource: dataSource, Parser: parser, Repository: repository}
}
