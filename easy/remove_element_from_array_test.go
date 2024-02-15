package easy

import "testing"

func TestRemoveElementFromArray(t *testing.T) {
	testCases := []struct {
		input    []int
		element  int
		expected []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1}, 0, []int{1}},
		{[]int{1, 2}, 2, []int{1}},
		{[]int{1, 2, 2, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9}, 2, []int{1, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tc := range testCases {
		k := removeElementFromArray(tc.input, tc.element)

		if k != len(tc.expected) {
			t.Fatalf("expected returned int to be '%d', got '%d'", len(tc.expected), k)
		}

		for i := 0; i < k; i++ {
			if tc.expected[i] != tc.input[i] {
				t.Fatalf("failed to assert that tc.expected[i] == tc.input[i] , %v is not %v", tc.expected, tc.input[:k])
			}
		}
	}
}
