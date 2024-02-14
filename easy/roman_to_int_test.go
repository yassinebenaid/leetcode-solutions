package easy

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"I", 1},
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		if result := romanToInt(tc.input); result != tc.expected {
			t.Fatalf(`expected "%s" to return "%d", got "%d"`, tc.input, tc.expected, result)
		}
	}
}
