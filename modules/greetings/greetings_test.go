package greetings

import "testing"

func TestHelloValidName(t *testing.T) {
	// Arrange
	name := "Bob"

	// Act
	actual, err := Hello(name)

	// Assert

	if err != nil {
		t.Errorf("Hello(%v) returned an error: %v", name, err)

	}

	if actual == "" {
		t.Errorf("Hello(%v) = '', expected %v", name, actual)
	}
}

func TestHelloEmptyName(t *testing.T) {
	// Arrange
	name := ""
	expected := "empty name"

	// Act
	actual, err := Hello(name)

	// Assert
	if err == nil {
		t.Errorf("Hello(%v) did not return an error", name)
	}

	if actual != "" {
		t.Errorf("Hello(%v) = %v, expected %v", name, actual, "")
	}

	if err.Error() != expected {
		t.Errorf("Hello(%v) error = %v, expected %v", name, err, expected)
	}

}

func TestHelloWorld(t *testing.T) {
	// Arrange
	expected := "Hello, world."

	// Act
	actual := HelloWorld()

	// Assert
	if actual != expected {
		t.Errorf("HelloWorld() = %v, expected %v", actual, expected)
	}
}
