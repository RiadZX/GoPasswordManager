package helpers

import (
	"fmt"
	"github.com/fatih/color"
)

var PREFIX = ""

func Log(message string) {
	color.Set(color.FgHiCyan)
	fmt.Println(message)
	color.Unset()
}

func Warn(message string) {
	color.Set(color.FgHiYellow)
	fmt.Println(message)
	color.Unset()
}

func Success(message string) {
	color.Set(color.FgHiGreen)
	fmt.Println(message)
	color.Unset()
}

func LogEntry(email string, username string, password string, website string, category string, notes string) {
	color.Set(color.FgHiMagenta)
	fmt.Println("---------Entry---------")
	fmt.Println("> Email: " + email)
	fmt.Println("> Username: " + username)
	fmt.Println("> Password: " + password)
	fmt.Println("> Website: " + website)
	fmt.Println("> Category: " + category)
	fmt.Println("> Notes: " + notes)
	fmt.Println("-----------------------")
	color.Unset()

}

func Danger(message string) {
	color.Set(color.FgHiRed)
	fmt.Println(message)
	color.Unset()
}

func LogList(message []string) {
	fmt.Println("---------Menu---------")
	for i := 0; i < len(message); i++ {
		fmt.Print(PREFIX, i+1, ".", message[i], "\n")
	}
	fmt.Println("-----------------------")

}
