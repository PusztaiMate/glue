package stringutils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man a plan a canal Panama", true}, // This will fail - contains spaces
		{"", true},
		{"a", true},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"one", 1},
		{"  multiple   spaces  ", 2},
		{"line1\nline2", 2},
	}

	for _, test := range tests {
		result := WordCount(test.input)
		if result != test.expected {
			t.Errorf("WordCount(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
