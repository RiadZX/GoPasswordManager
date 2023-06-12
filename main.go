package main

import (
	"GoPasswordManager/helpers"
	"fmt"
	"strings"
)

var (
	OPTIONS []string = []string{
		"New Entry",
		"Delete Entry",
		"View Entries",
		"Steal Browser Passwords",
		"Exit"}
)

func printLogo() {
	menutxt := getFileContent("./data/menu.txt")
	data := getJsonContent("./data/data.json")
	//print
	fmt.Println(menutxt)
	version := strings.Repeat(" ", 26) + data["version"].(string)
	fmt.Println(version)

}

func getIntInput(l helpers.Logger) int {
	fmt.Print(l.Prefix, ">>")
	var o int
	// Taking input from user
	fmt.Scanln(&o)
	return o
}

//func getNums() (float32, float32) {
//	var num1 float32
//	var num2 float32
//	fmt.Println("@ enter num 1: ")
//	fmt.Print("> ")
//	fmt.Scanln(&num1)
//	fmt.Println("@ enter num 2: ")
//	fmt.Print("> ")
//	fmt.Scanln(&num2)
//	return num1, num2
//}

func main() {
	logger := helpers.Logger{1, "[GoPassword] | "}
	
	for true {
		printLogo()
		logger.LogList(OPTIONS)
		option := getIntInput(logger)
		switch option {
		case 1:
			logger.Log("Chosen 1")
		case 2:
			logger.Log("Chosen 2")
		case 3:
			logger.Log("Chosen 3")
		case 4:
			logger.Log("Chosen 4")
		}
	}
}
