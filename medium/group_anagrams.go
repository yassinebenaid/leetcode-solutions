package medium

import (
	"sort"
)

/*
*	Task: https://leetcode.com/problems/group-anagrams/
*
*	Given an array of strings strs, group the anagrams together. You can return the answer in any order.
*
*	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
*	typically using all the original letters exactly once.

*
*	Examples:
*	-------------------------------------------------------------------------------------------------
*	| Input											| expected output								|
*   -------------------------------------------------------------------------------------------------
*	| ["eat","tea","tan","ate","nat","bat"]			| [["bat"],["nat","tan"],["ate","eat","tea"]]	|
*	-------------------------------------------------------------------------------------------------
*	| [""]											| [[""]]										|
*	-------------------------------------------------------------------------------------------------
*	| ["a"]											| [["a"]]										|
*	-------------------------------------------------------------------------------------------------
 */

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	// two words are anagram if they have the same length, the same characters
	// with the exact same occurences. but they may be differnt in the order of the characters.
	// This means if we sort the characters , both words should be equal.
	for _, word := range strs {
		bytes := []byte(word)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		sorted_word := string(bytes)
		groups[sorted_word] = append(groups[sorted_word], word)
	}

	groups_slice := make([][]string, 0)
	for key := range groups {
		groups_slice = append(groups_slice, groups[key])
	}

	return groups_slice
}
