package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkExt(filePath string) bool {

	splitted := strings.Split(filePath, ".")
	return splitted[len(splitted) -1] == "bf"
}

func getBfContent(filePath string) string {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No Brainfuck file provided")
		os.Exit(1)
	}

	programPath := os.Args[1]

	if (!checkExt(programPath)) {
		fmt.Println("Please provide a valid Brainfuck file")
		os.Exit(1)
	}

	fileCont := getBfContent(programPath)
	compiled := compile(fileCont)

	fmt.Println(compiled)

}