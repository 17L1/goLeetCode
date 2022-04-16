/*
 * @Description:双指针，暴力方法不行
 * @Author: lilongguang
 * @Date: 2022-04-14 12:54:48
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-14 13:13:43
 */
package main

func maxArea(height []int) int {
	lenArray := len(height)
	if lenArray <= 1 {
		return 0
	}
	res := 0
	left, right := 0, lenArray-1
	for left < right {
		temp := 0
		if height[left] < height[right] {
			temp = height[left] * (right - left)
			if temp > res {
				res = temp
			}
			left++
			continue
		}
		temp = height[right] * (right - left)
		if temp > res {
			res = temp
		}
		right--
	}
	return res

}
