package easy

import "testing"

func TestFirstLetterToApearTwice(t *testing.T) {
	testCases := []struct {
		input          string
		expectedletter byte
	}{
		{"abccbaacz", 'c'},
		{"abcdd", 'd'},
		{"abc", 0},
	}

	for _, tc := range testCases {
		if result := firstLetterToApearTwice(tc.input); result != tc.expectedletter {
			t.Fatalf(`expected "%s" to return "%c", got "%c"`, tc.input, tc.expectedletter, result)
		}
	}
}
