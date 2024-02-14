package easy

/*
*	Task: https://leetcode.com/problems/roman-to-integer/
*
*	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
*
*	For example, 2 is written as II in Roman numeral, just two ones added together.
*	12 is written as XII, which is simply X + II.
*	The number 27 is written as XXVII, which is XX + V + II.
*
*	Roman numerals are usually written largest to smallest from left to right.
*   However, the numeral for four is not IIII. Instead, the number four is written as IV.
*	Because the one is before the five we subtract it making four.
*	The same principle applies to the number nine, which is written as IX.
*	There are six instances where subtraction is used:
*
*	- I can be placed before V (5) and X (10) to make 4 and 9.
*	- X can be placed before L (50) and C (100) to make 40 and 90.
*	- C can be placed before D (500) and M (1000) to make 400 and 900.
*
*	Given a roman numeral, convert it to an integer.
*
*	Examples:
*	-------------------------------------
*	| Input			| expected output	|
*   -------------------------------------
*	| "III"			| 3 				|
*	-------------------------------------
*	| "LVIII"		| 58				|
*	-------------------------------------
*	| "MCMXCIV"		| 1994				|
*	-------------------------------------
*
 */

func romanToInt(s string) int {

	var roman = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	var should_skip bool

	for i, c := range s {
		if should_skip {
			should_skip = false
			continue
		}

		if c == 'I' && i < len(s)-1 && (s[i+1] == 'V' || s[i+1] == 'X') {
			result += roman[rune(s[i+1])] - roman[c]
			should_skip = true
			continue
		}

		if c == 'X' && i < len(s)-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
			result += roman[rune(s[i+1])] - roman[c]
			should_skip = true
			continue
		}

		if c == 'C' && i < len(s)-1 && (s[i+1] == 'D' || s[i+1] == 'M') {
			result += roman[rune(s[i+1])] - roman[c]
			should_skip = true
			continue
		}

		result += roman[c]
	}

	return result
}
