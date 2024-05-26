package files

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	check(err)
	return string(content)
}
