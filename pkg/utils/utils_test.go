package utils

import "testing"

func TestReadFiles(t *testing.T) {
	content := ReadFile("./../input_files/test.log")
	var expected = []string{"Hello, World!"}
	if content[0] != expected[0] {
		t.Fatalf(`Message should be = %s, but is %s`, expected, content)
	}
}

func TestWriteFile(t *testing.T) {
	content := "Hello World!"
	filePath := "./../output_files/output_test.json"
	WriteFile(content, filePath)
	savedContent := ReadFile(filePath)

	if content != savedContent[0] {
		t.Fatalf(`Message should be = %s, but is %s`, content, savedContent)
	}
}
