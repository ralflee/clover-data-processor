# Clover Data Processor

## Prerequisite

1. golang 1.14
2. setup go root and go path
3. postgreSQL 11+

In Linux, set up go root and go path

```bash
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
```

## Quick start

```bash
go run main.go
```

## Introduction

### Why choose golang

Golang is a static typed general purpose programming lanuague. It is famous for it simple syntax, built-in concurrency, fast compile time and light in memory.

It is to showcase that golang can be a efficient *script* language to do simple and routine tasks, just like the nature of this assignment

### Design

After requirement analysis, the program has 6 use cases.

1. Get all spec file path from "./spec"
2. Get all data file path from "./data"
3. Parse spec file to become a spec struct (in Java terms, Spec Objects)
4. Parse data file to become a data struct (in Java terms, Data Objects)
5. Create DB Table
6. Insert data into DB Table

The data parsing program has 4 components. Each components has its only reponsibilities. With this design we can plugin and play and further enhance the program if we wish.
E.g. The spec and data can be from ftp rather than local storage. Or, the DB can be a NoSQL DB for analytic usage.

- Service
- Data Source
- Parser
- Repository

## Service

InitDataImportService is the core of this program, it contains the highest level of the program logic. In short it do followings:

1. Read spec and data files
2. Parse spec and data files to get spec and data structs
3. Create DB tables accordingly
4. Insert data into the corresponding DB table

It depends on a data source, a parser and a repository to execute.

Remarks: refer to data_import_service.go

## Data Source

Data source here don't mean a DB connection data source. It means the source of the data and spec.

It handles the following use cases:

- Get all spec file path from "./spec"
- Get all data file path from "./data"

Remarks: refer to file_data_source.go

## Parser

It parse data source and transform it into model.Spec or model.Record.

It handles the following use cases:

- Parse spec file to become a spec struct (in Java terms, Spec Objects)
- Parse data file to become a data struct (in Java terms, Data Objects)

Remarks: refer to file_data_parser.go

## Repository

The nature of data processer software is quite different from a typical CRUD. There is no known table structure util a spec file is read by the program. So, this program didn't use ORM framework / library. Rather, it use built-in sql library to run sql statements.

It handles the following use cases:

- Create DB Table
- Insert data into DB Table

Remarks: refer to data_repository.go

## Configuration

The configuration is put in config/config.yml

## Abnormal data handling

Spec:

- if each row of data has more than or less than 3 values, the row will be skipped

## DB table

In this execise, the db table creation depends only on the spec file. In a real world application, we may want to added following columns automatically.

- primary key
- creation time

## Test

The offical [The cover story](https://blog.golang.org/cover). A recommended reading.

Remarks: In golang, test files' name has suffix *_test.go*. In convention, test files are put in same directory of the target go file

### Unit Test

```bash
go test ./...  
```

### Unit test with coverage report

The current code coverage is around 50%.  

The code coverage report can be display through following commands:

```bash
go test -v -coverpkg=./... -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Mocking - Mockery

Mockery is used to generate mock struts

Install Mockery through homebrew

```bash
brew install vektra/tap/mockery
brew upgrade mockery
```

Generate the mocks by calling the mockery command.  

E.g. Generate mock repositories:

```bash
mockery -name=DataRepository -dir=./app/service -output=./app/service/mocks -filename=mock_data_repository.go
```

## Improvements

Potential improvements can be done it the program

- Table name checking, check whether special characters which is allowed for a file name but not allowed for a DB table name
- Unit test should cover error cases.
- Unit test should cover repository by mocking a database. Mocking a database is possible in golang [Sql driver mock for Golang](https://github.com/DATA-DOG/go-sqlmock)
