package main

import (
	"fmt"
	"os"

	"github.com/zub-bee/json-parser/parser"
)

func main() {
	fmt.Println("Hello, World!... Processing...")

	fileName := os.Args[len(os.Args)-1]

	// Open jsonFile
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(3)
	}
	defer file.Close()

	// Scan the file
	tokenList := parser.LexicalScan(file)

	if len(tokenList) == 0 {
		fmt.Println("err: empty file, invalid JSON")
		os.Exit(2)
	}

	// Parse the file
	result, err := parser.Parse(tokenList)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	// Print the result
	fmt.Println(result)
	os.Exit(0)
}
