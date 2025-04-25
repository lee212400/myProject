package linttest

import "testing"

func TestGoCyclo(t *testing.T) {
	tests := []struct {
		name string
		age  int
	}{
		{"age 1", 1},
		{"age 2", 2},
		{"age 3", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoCyclo(tt.age)

		})
	}

}

func TestUnparam(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"match", "myTest", 1},
		{"non-match", "otherKey", 0},
		{"empty str", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unparam(tt.input)
			if result != tt.expected {
				t.Errorf("Unparam(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}

}
