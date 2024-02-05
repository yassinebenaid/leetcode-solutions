package easy

/*
*	Task: https://leetcode.com/problems/first-letter-to-appear-twice/
*
*	Given a string s consisting of lowercase English letters, return the first letter to appear twice.
*
*   Note:
*  	 - A letter a appears twice before another letter b if the second occurrence of a is
*      before the second occurrence of b.
*	 - str will contain at least one letter that appears twice.
*
*	Examples:
*	-------------------------------------
*	| Input			 | expected output	|
*   -------------------------------------
*	| "abccbaacz"	 | c 				|
*	-------------------------------------
*	| "abcdd"        | d				|
*	-------------------------------------
 */

func firstLetterToApearTwice(str string) byte {
	var set []byte

	have_seen := func(c byte) bool {
		for _, i := range set {
			if i == c {
				return true
			}
		}
		return false
	}

	for _, c := range str {
		if have_seen(byte(c)) {
			return byte(c)
		} else {
			set = append(set, byte(c))
		}
	}

	return 0
}
