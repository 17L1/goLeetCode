/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-06 14:55:50
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-06 15:07:04
 */

package main

func search(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums < 1 {
		return -1
	}
	left, right := 0, lenNums-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return -1
}
