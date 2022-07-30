package libs

import "fmt"

func Info(msg string) {
	fmt.Println("[ INFO ] " + msg)
}

func Warn(msg string) {
	fmt.Println("[ WARN ] " + msg)
}

func Error(msg string) {
	fmt.Println("[ ERROR ] " + msg)
}