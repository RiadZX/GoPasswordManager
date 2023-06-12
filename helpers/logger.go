package helpers

import "fmt"

type Logger struct {
	Level  int
	Prefix string
}

func (l Logger) Log(message string) {
	fmt.Println(l.Prefix, message)
}

func (l Logger) LogList(message []string) {
	fmt.Println(l.Prefix, "@options:")

	for i := 0; i < len(message); i++ {
		fmt.Print(l.Prefix, "[", i+1, "] ", message[i], "\n")
	}

}
