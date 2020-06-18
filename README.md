# Clover Data Processor

## Prerequisite

go 1.14 installed
setup go root
setup go patha
running on linux

## Quick start

go run main.go

## Introduction

Why choose golang

How program due with invalid data:
skip

The nature of data processer software is quite different from a typical CRUD 

Abstraction:
- service
- data source
- parser
- repository

## Service

It own the data import logic:

1. From a spec base path, get all spec path
2. From a data base path, get all data path
3. From a spec path, get spec struct(in java terms, get all spec objects)
4. From a data path, get data struct(in java terms, get all data objects)
5. For a spec struct, create a DB table
6. For a data struct, insert a DB record

## Abnormal data handling

spec 
- any abnormal behavior, skip the spec
- length of column name
- duplicated column names
- more logging

need to handle
table scheam didn't match spec

## DB table

- primary key
- creation time

## Configuration

- db connection url

## Test

The offical [The cover story](https://blog.golang.org/cover). A recommended reading.

Remarks: In golang, test files' name has suffix *_test.go*. In convention, test files are put in same directory of the target go file

### Unit Test

```bash
go test ./...  
```

### Unit test with coverage report

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

## TODO

table name checking
structure log

## Improvements

- DB table patching
- NoSQL?
- handle spec file name ordering