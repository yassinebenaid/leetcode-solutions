package easy

/*
*	Task: https://leetcode.com/problems/valid-parentheses/
*
*	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
*   determine if the input string is valid.
*
*	An input string is valid if:
*
*	 1. Open brackets must be closed by the same type of brackets.
*	 2. Open brackets must be closed in the correct order.
*	 3. Every close bracket has a corresponding open bracket of the same type.
*
*
*	Examples:
*	-------------------------------------
*	| Input			| expected output	|
*   -------------------------------------
*	| "()"			| true				|
*	-------------------------------------
*	| "()[]{}"		| true				|
*	-------------------------------------
*	| "(]"			| false				|
*	-------------------------------------
*	| "{([])}"		| true				|
*	-------------------------------------
*
 */

func isValidParentheses(str string) bool {

	var parentheses = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []rune

	for _, r := range str {
		// if the current character is an openning parentheses
		// such as : (, [ or { we push it to the stack
		if _, is_openning := parentheses[r]; is_openning {
			stack = append(stack, r)
			continue
		}

		// if we reach here, it means the current character is a closing parentheses
		// we check if we have any opened parentheses
		if len(stack) == 0 {
			return false
		}

		// we need to check if the parentheses is of the same type
		// as the last opened one
		last_openning := stack[len(stack)-1]
		if parentheses[last_openning] != r {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
