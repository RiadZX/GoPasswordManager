package main

import (
	"fmt"
	"strings"
)

func printMenu() {
	menutxt := getFileContent("./data/menu.txt")
	data := getJsonContent("./data/data.json")
	//print
	fmt.Println(menutxt)
	version := strings.Repeat(" ", 26) + data["version"].(string)
	fmt.Println(version)

}

func main() {
	printMenu()
}
