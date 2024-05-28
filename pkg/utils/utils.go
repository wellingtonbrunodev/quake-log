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

// ReadFile reads the content from the given file.
// It returns the content splitted by lines
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

// WriteFile writes the content to the given filePath
func WriteFile(content, filePath string) {
	data := []byte(content)
	err := os.WriteFile(filePath, data, 0666)
	check(err)
}
