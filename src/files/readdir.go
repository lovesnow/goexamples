package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ProcessDir(path string) {
	file_list, err := ioutil.ReadDir(path)
	if err == nil {
		for _, f := range file_list {
			if f.IsDir() {
				fmt.Println(f.Name(), " is dir")
								
			} else {
				fmt.Println(f.Name(), " is file")
			}
		}
	}
	
}


func main () {
	
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], " Directory")
		return 
		
	} else {
		ProcessDir(os.Args[1])
	}
}
