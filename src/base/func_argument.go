package main

import (
	"fmt"
)


func Expand(s string, mapping func(string) string) string {
	buf := make([]byte, 0, len(s))
	
	for i := 0; i < len(s); i++ {
		buf = append(buf, s[i])
		fmt.Println(string(s[i]))
		// fmt.Println(strconv.Itoa(int(s[i])))
		print(mapping(s[i:i+2]))
		
		fmt.Println()
	}
	return string(buf)
}


func Upper(s string) string {
	return s
}

func main() {
	Expand("linuxdeepin", Upper) 
}
