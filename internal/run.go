package internal

import (
	"fmt"

	files "github.com/wellingtonbrunodev/quake-log/pkg"
)

func Run() {
	content := files.ReadFile("./pkg/input_files/qgames.log")
	fmt.Println(content)
}
