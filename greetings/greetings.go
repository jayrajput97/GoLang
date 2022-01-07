package greetings

import (
	"errors"
	"fmt"
)

//Hello returns a greeting for the named person
func Hello(name string) (string,error) {

	//If no name is mentioned, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	//If the name is mentioned, return a greeting for the person
	message := fmt.Sprintf("Hello! %v. Welcome and Congratulations of your first module.", name)
	return message, nil
}