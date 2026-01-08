package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What is yor name?")

	if scanner.Scan() {
		fmt.Printf("Hello, %s!\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input", err)
	}

}


