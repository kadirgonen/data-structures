package main

import (
	"fmt"
	"strings"
)

func main() {
	result := reverse("Hello World")
	fmt.Println(result)

}

func reverse(input string) string {

	var slice []string
	for i := len(input) - 1; i >= 0; i-- {
		slice = append(slice, string(input[i]))
	}
	return strings.Join(slice, "")
}
