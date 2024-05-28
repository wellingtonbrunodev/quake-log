package utils

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) []string {

	readFile, err := os.Open(filepath)
	check(err)

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
