package library

import (
	"fmt"
)

func Message() {
	fmt.Println("Hello World")
}

func Introduce(s string) string {
	result := fmt.Sprintf("Hello %v", s)
	return result
}
