package easy

import "testing"

func TestFirstUnqieCharacter(t *testing.T) {
	testCases := []struct {
		input         string
		expectedIndex int
	}{
		{"test", 1},
		{"aaee", -1},
		{"foobar", 0},
		{"foofb", 4},
	}

	for _, tc := range testCases {
		if i := firstUniqueCharacter(tc.input); i != tc.expectedIndex {
			t.Fatalf(`expected "%s" to return "%d", got "%d"`, tc.input, tc.expectedIndex, i)
		}
	}
}
