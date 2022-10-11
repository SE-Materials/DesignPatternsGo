package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"
	sb := strings.Builder{}

	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Println(sb.String())

	words := []string{"Hello", "World!"}
	sb.Reset()

	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")

	fmt.Println(sb.String())
}
