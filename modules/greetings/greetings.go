package greetings

import (
	"errors"
	"fmt"
	"math/rand"

	"rsc.io/quote/v3"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randomGreeting(), name), nil
}

// Hellos retuns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		mesage, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = mesage
	}

	return messages, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func randomGreeting() string {
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return greetings[rand.Intn(len(greetings))]
}

func HelloWorld() string {
	return quote.HelloV3()
}
