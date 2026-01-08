package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, Go with buffered output hahhaha !\n")
	writer.Flush()
}
