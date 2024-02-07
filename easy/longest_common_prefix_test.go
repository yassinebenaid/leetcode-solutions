package easy

import "testing"

func TestFindCommenPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight", "fl"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"flowering", "flowering", "flowering", "flowering"}, "flowering"},
		{[]string{}, ""},
	}

	for _, tc := range testCases {
		if result := findLongestCommonPrefix(tc.input); result != tc.expected {
			t.Fatalf(`expected "%s" to return "%v", got "%v"`, tc.input, tc.expected, result)
		}
	}
}
