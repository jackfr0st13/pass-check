# PassCheck: Your next password validator for Go
A useful password rules validator for Go

## Installation

~~~sh
go get -u github.com/jackfr0st13/pass-check
~~~

## Usage

~~~go
import (
    pass "github.com/jackfr0st13/pass-check/pass"
)

pd := pass.NewPasswordData(pass.WithPassword("mytestnimda"), pass.WithUsername("admin"))
	validator := pass.NewPassChecker(pass.LengthRule(10, 20), pass.UsernameRule(true, true))
	result, err := validator.Check(*pd)
	if err != nil {
		panic(err)
	}
	if !result.IsValid() {
		errorMessages := validator.GetMessages(result)
		for _, errorMsg := range errorMessages {
			fmt.Println(errorMsg)
		}
	} else {
		fmt.Println("Password Check successful")
	}
~~~
---

## API Overview

The API consists of 3 main structures:
* Rule: A validation clause the password has to pass in order for it to be deemed valid. 
* PassChecker: A password checker which when provided with a list of Rules, validates the password against them. 

## Rule Overview


Rules are building blocks for password validation and **pass-check** allows for the user to specify two types of rules :

* **Positive Rule:** Passes the password that matches the rule.
* **Negative Rule:** Rejects the password that matches the rule.

### Positive Matching Rules

* **AllowedRegex Rule:** : requires a password to conform to the given pattern 
* **Length Rule** : requires a password to be within the given length range.

### Negative Matching Rules

* **Username Rule** : rejects passwords that contain the username of the user providing the password
* **Whitespace Rule** : rejects passwords that contain any whitespace
* **Number Range Rule** : rejects passwords that contain any number within the given range (inclusive of limits).
* **Illegal Regex Rule** : rejects passwords that conform to a regular expression










   

