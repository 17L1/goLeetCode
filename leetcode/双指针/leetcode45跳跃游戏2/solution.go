/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-17 14:12:57
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-17 14:50:30
 */
package main

import "math"

func jump(nums []int) int {
	lenNums := len(nums)
	if lenNums <= 1 {
		return 0
	}
	dp := make([]int, lenNums)
	dp[0] = 0
	for i := 1; i < lenNums; i++ {
		dp[i] = math.MaxInt
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[lenNums-1]
}

func jump2(nums []int) int {
	lenNums := len(nums)
	if lenNums <= 1 {
		return 0
	}
	dp := make([]int, lenNums)
	dp[0] = 0
	for i, j := 1, 0; i < lenNums; i++ {
		if j+nums[j] < i {
			j++
		}
		dp[i] = dp[j] + 1
	}
	return dp[lenNums-1]
}

func jump3(nums []int) int {
	lenNums := len(nums)
	if lenNums < 2 {
		return 0
	}
	maxPosition := 0
	end := 0
	step := 0
	for i := 0; i < lenNums-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
