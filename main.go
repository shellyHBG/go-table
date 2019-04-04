package main

import (
	"fmt"
	"io/ioutil"

	"./table"
)

func main() {
	fmt.Println("Hi there!")

	// 1. load txt file
	// https://golang.org/pkg/io/ioutil/
	// func ReadFile(filename string) ([]byte, error)
	// ReadFile reads the file named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF.
	// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
	data_bytes, err := ioutil.ReadFile("LocalContent.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data_bytes)

	// force type converting from byte[] to string
	data := string(data_bytes)
	//fmt.Println(data)

	// 2. parse data
	table, err := table.Parse("LocalContent", data)
	if err != nil {
		fmt.Println(err)
	}

	// 3. test
	fmt.Printf("GetAll > %q\n", table.GetAll("1000"))
	//fmt.Printf("GetAll > %q\n", table.GetAll(1000))  // error: cannot use 1000 (type int) as type string in argument to table.GetAll

	fmt.Printf("Get > %v\n", table.Get("1000", "EN"))
}
