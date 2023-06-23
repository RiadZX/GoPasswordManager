package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var PREFIX = ""

func Log(message string) {
	color.Set(color.FgHiCyan)
	fmt.Println(message)
	color.Unset()
}

func LogoPrint(logo string) {
	stringlist := strings.Split(logo, "\n")
	color.Set(color.FgHiMagenta)
	fmt.Println(stringlist[0])
	color.Unset()

	color.Set(color.FgHiBlue)
	fmt.Println(stringlist[1])
	color.Unset()

	color.Set(color.FgHiGreen)
	fmt.Println(stringlist[2])
	color.Unset()

	color.Set(color.FgHiRed)
	fmt.Println(stringlist[3])
	color.Unset()

	color.Set(color.FgHiCyan)
	fmt.Println(stringlist[4])
	color.Unset()

	color.Set(color.FgHiYellow)
	fmt.Println(stringlist[5])
	color.Unset()

	color.Set(color.FgHiWhite)
	fmt.Println(stringlist[6])
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
