package easy

/*
*	Task: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*
*	Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place
*	such that each unique element appears only once. The relative order of the elements should
*	be kept the same. Then return the number of unique elements in nums.
*
*	Examples:
*	---------------------------------------------
*	| Input			 		| expected output	|
*   ---------------------------------------------
*	| [1,1,2]   	 		| 2 				|
*	---------------------------------------------
*	| [0,0,1,1,1,2,2,3,3,4] | 5					|
*	---------------------------------------------
 */

func removeDuplicatesFromSortedArray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var last_unique_index int
	var num_unique = 1

	for _, num := range nums {
		if num != nums[last_unique_index] {
			last_unique_index++
			num_unique++
			nums[last_unique_index] = num
		}
	}

	return num_unique
}
