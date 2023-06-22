package helpers

import (
	"fmt"
	"github.com/fatih/color"
)

func GetIntInput(message string) int {
	color.Set(color.FgHiCyan)
	fmt.Print(">>", message, ": ")
	color.Unset()
	var o int
	// Taking input from user
	fmt.Scanln(&o)
	return o
}

func GetStrInput(message string) string {
	color.Set(color.FgHiCyan)
	fmt.Print(">>", message, ": ")
	color.Unset()
	var o string
	// Taking input from user
	fmt.Scanln(&o)

	if o == "" {
		return " "
	}
	return o
}

func GetPasswordInput() string {
	resp := GetStrInput("Do you want to generate a new password? (y/n): ")
	if resp == "y" || resp == "Y" || resp == "yes" || resp == "YES" || resp == "1" {
		pw := GeneratePassword()
		Warn("Password generated:" + pw)
		return pw

	} else {
		pw := GetStrInput("Enter new password")
		return pw
	}
}
