package helpers

import (
	"fmt"
	"github.com/fatih/color"
)

func GetIntInput() int {
	color.Set(color.FgHiCyan)
	fmt.Print(">> ")
	color.Unset()
	var o int
	// Taking input from user
	fmt.Scanln(&o)
	return o
}

func GetStrInput(message string) string {
	color.Set(color.FgHiCyan)
	fmt.Print(">> ", message)
	color.Unset()
	var o string
	// Taking input from user
	fmt.Scanln(&o)

	if o == "" {
		return " "
	}
	return o
}
