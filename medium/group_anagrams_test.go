package medium

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	testCases := []struct {
		input    []string
		expected [][]string
	}{
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}

	for _, tc := range testCases {
		// the result may be in any order
		// we need to sort the result to match our test case
		result := groupAnagrams(tc.input)
		sort.Slice(result, func(i, j int) bool {
			return len(result[i]) > len(result[j])
		})

		t.Log(result)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Fatalf(`expected "%s" to return "%v", got "%v"`, tc.input, tc.expected, result)
		}
	}
}

func BenchmarkGroupAnagram1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagrams([]string{"eat",
			"tea", "tan", "ate", "nat",
			"bat", "eat", "tea", "tan",
			"ate", "nat", "bat", "eat",
			"tea", "tan", "ate", "nat",
			"bat", "eat", "tea", "tan",
			"ate", "nat", "bat", "tea",
			"tan", "ate", "nat", "bat",
			"eat", "tea", "tan", "ate",
			"nat", "bat", "eat", "tea",
			"tan", "ate", "nat", "bat",
			"eat", "tea", "tan", "ate",
			"nat", "bat", "tea", "tan",
			"ate", "nat", "bat", "eat",
			"tea", "tan", "ate", "nat",
			"bat", "eat", "tea", "tan",
			"ate", "nat", "bat", "eat",
			"tea", "tan", "ate", "nat",
			"bat", "tea", "tan", "ate",
			"nat", "bat", "eat", "tea",
			"tan", "ate", "nat", "bat",
			"eat", "tea", "tan", "ate",
			"nat", "bat", "eat", "tea",
			"tan", "ate", "nat", "bat",
		})
	}
}
