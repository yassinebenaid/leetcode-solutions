package easy

/*
*	Task: https://leetcode.com/problems/longest-common-prefix/
*
*	Write a function to find the longest common prefix string amongst an array of strings.
*	If there is no common prefix, return an empty string "".
*
*	Examples:
*	-------------------------------------------------
*	| Input						 | expected output	|
*   -------------------------------------------------
*	| ["flower","flow","flight"] |"fl"				|
*	-------------------------------------------------
*	| ["dog","racecar","car"]	 |""				|
*	-------------------------------------------------
 */
func findLongestCommonPrefix(strs []string) string {
	var prefix []byte
	var last byte

	for i := 0; ; i++ {
		for _, word := range strs {
			if len(word) <= i {
				return string(prefix)
			}

			if last == 0 {
				last = word[i]
			}

			if last != word[i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, last)
		last = 0
	}
}
