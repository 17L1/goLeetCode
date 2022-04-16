/*
 * @Description:双指针
 * @Author: lilongguang
 * @Date: 2022-04-15 18:46:49
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-15 19:31:51
 */
package main

import (
	"fmt"
	"sort"
)

func fourSum(num []int, target int) [][]int {
	// numsMap := make(map[int]int, 0)
	lenNums := len(num)
	ans := make([][]int, 0)
	sort.Ints(num)
	if lenNums < 4 {
		return [][]int{}
	}
	// for _, v := range num {
	// 	numsMap[v] = target - v
	// }
	for i := 0; i < lenNums-3; i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		temp := target - num[i]
		for j := i + 1; j < lenNums-2; j++ {
			if j > i+1 && num[j] == num[j-1] {
				continue
			}
			left, right := j+1, lenNums-1
			for left < right {
				if left > j+1 && left < lenNums && num[left] == num[left-1] {
					left++
					continue
				}
				if left >= right {
					break
				}
				if num[j]+num[left]+num[right] > temp {
					right--
					continue
				}
				if num[j]+num[left]+num[right] < temp {
					left++
					continue
				}
				ans = append(ans, []int{num[i], num[j], num[left], num[right]})
				right--
				left++
			}
		}
	}
	return ans
}

func main() {
	num := []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(num, 8))
}
