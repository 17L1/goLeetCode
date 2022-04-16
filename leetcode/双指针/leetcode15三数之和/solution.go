/*
 * @Description:双指针
 * @Author: lilongguang
 * @Date: 2022-04-14 13:14:38
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-15 18:27:12
 */
package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	if lenNums < 3 {
		return make([][]int, 0)
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < lenNums-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, lenNums-1
		for left < right {
			for left > i+1 && left < lenNums && nums[left] == nums[left-1] {
				left++
				continue
			}
			if left >= right {
				break
			}
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
				continue
			}
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
				continue
			}
			res = append(res, []int{nums[i], nums[left], nums[right]})
			left++
			right--
		}
	}
	return res
}
