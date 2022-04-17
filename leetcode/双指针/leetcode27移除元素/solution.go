/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-17 14:02:07
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-17 14:11:42
 */
package main

func removeElement(nums []int, val int) int {
	lenNums := len(nums)
	index := 0
	if lenNums < 1 {
		return 0
	}
	for i := 0; i < lenNums; i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
