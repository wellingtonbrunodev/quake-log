package files

import "testing"

func TestReadFiles(t *testing.T) {
	content := ReadFile("./input_files/test.log")
	var expected = "Hello, World!"
	if content != expected {
		t.Fatalf(`Message should be = %s, but is %s`, expected, content)
	}
}
