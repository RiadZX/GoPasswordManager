package helpers

import (
	"fmt"
	"github.com/fatih/color"
)

var PREFIX = "[🔒]"

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

func Danger(message string) {
	color.Set(color.FgHiRed)
	fmt.Println(message)
	color.Unset()
}

func LogList(message []string) {
	for i := 0; i < len(message); i++ {
		fmt.Print(PREFIX, "[", i+1, "] ", message[i], "\n")
	}

}
