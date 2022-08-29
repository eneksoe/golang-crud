package controllers

import (
	"fmt"
	"testing"
)

func Test_valid(t *testing.T) {
	for _, email := range []string{
		"good@exmaple.com",
		"bad-example",
	} {
		fmt.Printf("%18s emailValidator: %t\n", email, emailValidator(email))
	}
}
