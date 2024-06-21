package main

import (
	"errors"
	"fmt"
)

var divByZeroError = errors.New("division by zero is not allowed")

type divisionByZero struct {
	numA int
	numB int
	msg  string
}

func (d *divisionByZero) Error() string {
	return d.msg
}

func DoSomething() error {
	return errors.New("something went wrong")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &divisionByZero{a, b, "division by zero is not allowed"}
	}

	return a / b, nil
}

func doSomething() error {
	// some logic
	return fmt.Errorf("doSomething %w", errors.New("something went wrong"))
}

func doAnotherThing() error {
	// some logic
	if err := doSomething(); err != nil {
		return fmt.Errorf("doAnotherThing %w", err)
	}
	return fmt.Errorf("doAnotherThing %w", errors.New("an error occured"))
}

func main() {
	err := DoSomething()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Errorf("something went wrong: %w", err))

	a, b := 10, 0

	result, err := divide(a, b)

	if err != nil {
		var divErr *divisionByZero
		switch {
		case errors.As(err, &divErr):
			fmt.Println(divErr.Error())
		default:
			fmt.Println("unknown error")
		}
	}

	fmt.Println(result)

	err = doSomething()
	fmt.Println(err)

	err = doAnotherThing()
	fmt.Println(err)
}
