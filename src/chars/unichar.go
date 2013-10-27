package main

import (
	"fmt"
	"strings"
)

func Decode(s string) {
	buffers := strings.TrimSuffix(strings.Split(s, "\r\n")[2], "\x1a")
		
}

func main() {
	fmt.Println(strings.Split("linux deepin is a", " ")[2])
}
