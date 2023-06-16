package main

import (
	"GoPasswordManager/helpers"
	"GoPasswordManager/passwords"
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
	menutxt := helpers.GetFileContent("./data/menu.txt")
	data := helpers.GetJsonContent("./data/data.json")
	//print
	fmt.Println(menutxt)
	version := strings.Repeat(" ", 26) + data["version"].(string)
	fmt.Println(version)

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
	logger := helpers.Logger{Level: 1, Prefix: "[GoPassword] | "}

	for true {
		printLogo()
		logger.LogList(OPTIONS)
		option := helpers.GetIntInput()
		switch option {
		case 1:
			entry := passwords.AddEntry()
			helpers.SavePassword(entry)
		case 2:
			logger.Log("Chosen 2")
		case 3:
			logger.Log("Chosen 3")
		case 4:
			logger.Log("Chosen 4")
		}
	}
}
