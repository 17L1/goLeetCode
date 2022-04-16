/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-15 17:00:34
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-15 18:45:22
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	lenNums := len(nums)
	res := math.MaxInt
	ans := 0
	sort.Ints(nums)
	for i := 0; i < lenNums-2; i++ {
		left, right := i+1, lenNums-1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if int(math.Abs(float64(temp-target))) < res {
				res = int(math.Abs(float64(temp - target)))
				ans = temp
			}
			if temp-target > 0 {
				right--
				continue
			}
			if temp-target < 0 {
				left++
				continue
			}
			return ans
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
