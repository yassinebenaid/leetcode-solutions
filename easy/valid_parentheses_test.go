package easy

import "testing"

func TestIsValidParentheses(t *testing.T) {
	testCases := []struct {
		input   string
		isValid bool
	}{
		{"()", true},
		{"(]", false},
		{"()[]{}", true},
		{"()[{}", false},
		{"{([])}", true},
		{"([])}", false},
	}

	for _, tc := range testCases {
		if result := isValidParentheses(tc.input); result != tc.isValid {
			t.Fatalf(`expected "%s" to return "%t", got "%t"`, tc.input, tc.isValid, result)
		}
	}
}
