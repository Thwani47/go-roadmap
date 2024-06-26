package hello

import "testing"

func TestHello(t *testing.T) {
	// Arrange
	expected := "Hello, world."

	// Act
	actual := Hello()

	// Assert
	if actual != expected {
		t.Errorf("Hello() = %q, expected %q", actual, expected)
	}
}

func TestProverb(t *testing.T) {
	// Arrange
	expected := "Concurrency is not parallelism."

	// Act
	actual := Proverb()

	// Assert
	if actual != expected {
		t.Errorf("Proverb() = %q, expected %q", actual, expected)
	}
}
