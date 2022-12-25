package main

import (
	"01._Structures/internal/models"
	"fmt"
)

func main() {
	errT := func(err error) string {
		if err != nil {
			return fmt.Sprintf("Error occurred: %v\n", err)
		}
		return "Phone was successfully changed!\n"
	}
	igor := models.User{
		Email:    "igor@gmail.com",
		Password: "qwerty009",
		Age:      30,
		Phone:    "+78005553535",
	}
	test := []string{"", "+9872621330", "+88987221330", "988987221330", "+79002281337"}
	for _, s := range test {
		err := igor.ChangePhone(s)
		fmt.Print(errT(err))
		igor.PrintInformation()
	}
}
