/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-16 14:19:13
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-16 14:40:35
 */
package main

func removeDuplicates(nums []int) int {
	lenNums := len(nums)
	if lenNums < 2 {
		return lenNums
	}
	index := 0
	// numDuplicate := 0
	for i := 0; i < lenNums; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	return index
}
