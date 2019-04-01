# Table Parser by golang

## Introduction
This table parser combines two functionality. One is reading file, the other is processing the file data and generate a kind of data structure `type Table struct` including its method.
  
This project create a local package named `Table`. `Table` package provide a function to parse raw string and return a `Table struct` and an error:
```
func Parse(name, data string) (result Table, err error) 
```
```
type Table struct {
	name    string // table name
	content map[string]map[string]string
}
```
  
User can also fetch the table content through this 2 function:
```
func (t *Table) Get(key, subKey string) string
```
and
```
func (t *Table) GetAll(key string) map[string]string
```

## Requirements
Following are standard library (external package supported by go) used in this project:
* fmt
* io/ioutil
* errors
* strings

## Installation
No, just use it :D

## Usage
Compile and run this code by typing the command line:
```
> go run 
```
Build an executable file using:
```
> go build 
```

## Authors
By shellyHBG
