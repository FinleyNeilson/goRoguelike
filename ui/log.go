package ui

import "fmt"

func LogMessage(str string) {
	logToConsole(str)
}

func logToConsole(str string) {
	fmt.Println(str)
}
