package main

import (
	"fmt"

	"github.com/jackfr0st13/pass-check/pass"
)

func main() {
	pd := pass.NewPasswordData(pass.WithPassword("mytestnimda"), pass.WithUsername("admin"))
	validator := pass.NewPassChecker(pass.LengthRule(10, 20), pass.UsernameRule(true, true))
	result, err := validator.Check(*pd)
	if err != nil {
		fmt.Println(err)
	}
	if !result.IsValid() {
		messages := validator.GetMessages(result)
		for _, message := range messages {
			fmt.Println(message)
		}
	} else {
		fmt.Println("Password Check successful")
	}

}
