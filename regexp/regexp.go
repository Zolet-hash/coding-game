package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "^([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)$"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	fmt.Println(re.MatchString("example@examle.com"))
}

func isValidEmail(email string) bool {
	re, err := regexp.Compile("^([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)$")
	if err != nil {
		return false
	}
	return re.MatchString(email)
}

var emailRegexp = regexp.MustCompile("^([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)$")

func isValidEmail2(email string) bool {
	return emailRegexp.MatchString(email)
}
