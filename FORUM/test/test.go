package main

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func checkString(s string) error {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return errors.New("Empty string")
	}
	return nil
}

func main() {
	fmt.Println(checkEmail("qwe@wef.qwe"))
	fmt.Println(checkEmail("   "))
	fmt.Println(checkString(" wefwef "))
	fmt.Println(checkString("  "))
	fmt.Println(checkString(""))
}
