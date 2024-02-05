package easy

/*
*	Task: https://leetcode.com/problems/first-unique-character-in-a-string/
*
*	Given a string s, find the first non-repeating character in it and return its index.
*   If it does not exist, return -1.
*
*	Examples:
*	-------------------------------------
*	| Input			 | expected output	|
*   -------------------------------------
*	| "leetcode"	 | 0 				|
*	-------------------------------------
*	| "loveleetcode" | 2				|
*	-------------------------------------
*	| "aabb"		 | -1   			|
*	-------------------------------------
 */

func firstUniqueCharacter(str string) int {
outer:
	for i, c := range str {
		for j, c2 := range str {
			if i != j && c == c2 {
				continue outer
			}
		}

		return i
	}

	return -1
}
