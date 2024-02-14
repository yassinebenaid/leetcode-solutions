package easy

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 2, 3, 3, 3, 5, 5, 5, 5, 5}, []int{1, 2, 3, 5}},
	}

	for _, tc := range testCases {
		k := removeDuplicatesFromSortedArray(tc.input)

		if k != len(tc.expected) {
			t.Fatalf("expected returned int to be '%d', got '%d'", len(tc.expected), k)
		}

		for i := 0; i < k; i++ {
			if tc.expected[i] != tc.input[i] {
				t.Fatalf("failed to assert that tc.expected[i] == tc.input[i]")
			}
		}
	}
}
