package helpers

import (
	"fmt"
)

func GetIntInput() int {
	fmt.Print(">>")
	var o int
	// Taking input from user
	fmt.Scanln(&o)
	return o
}

func GetStrInput() string {
	fmt.Print(">>")
	var o string
	// Taking input from user
	fmt.Scanln(&o)
	return o
}
