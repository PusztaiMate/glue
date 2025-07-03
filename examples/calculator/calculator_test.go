package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20
	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; want %d", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	result := Divide(10, 0)
	expected := 0
	if result != expected {
		t.Errorf("Divide(10, 0) = %d; want %d", result, expected)
	}
}
