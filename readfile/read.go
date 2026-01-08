package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var file *os.File
	file, err := os.Open("readfile/haha.txt")
	if err != nil {
		fmt.Println("Error opening file!:", err)
		return
	}
	defer file.Close()
	/*In Go, the defer keyword is used to delay the execution
	of a function until the surrounding function returns. This is commonly
	used for cleanup tasks, such as closing files or releasing resources,
	ensuring that these actions are performed regardless of how the function exits*/

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')//reads a line at a time from the file until it hits a newline.
		if err != nil {
			break

		}
		fmt.Print(line)
	}
}
