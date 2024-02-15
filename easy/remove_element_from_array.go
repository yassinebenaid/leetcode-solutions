package easy

/*
*	Task: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*
*	Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
*	The order of the elements may be changed. Then return the number of elements in nums which are
*	not equal to val.
*
*	Examples:
*	-----------------------------------------------------------------
*	| Input			 	    | elemant			| expected output	|
*   -----------------------------------------------------------------
*	| [3,2,2,3]   	 		| 2 				|	2				|
*	-----------------------------------------------------------------
*	| [0,1,5,5,3,0,4,2,9,8]	| 5					|	7				|
*	-----------------------------------------------------------------
 */

func removeElementFromArray(nums []int, e int) int {
	var unique_items int
	var last_index int

	for _, num := range nums {
		if num != e {
			nums[last_index] = num
			unique_items++
			last_index++
		}
	}

	return unique_items
}
