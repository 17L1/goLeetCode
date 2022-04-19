/*
 * @Description:
 * @Author: 17L1
 * @Date: 2022-04-19 12:42:35
 * @LastEditors: 17L1
 * @LastEditTime: 2022-04-19 13:32:44
 */
package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	lenS := len(s)
	if lenS < 1 {
		return []int{}
	}
	ans := make([]int, lenS)
	indexC := make([]int, 0, lenS)
	for i := 0; i < lenS; i++ {
		if s[i] == c {
			indexC = append(indexC, i)
		}
	}
	for i := 0; i < lenS; i++ {
		temp := math.MaxInt
		for j := 0; j < len(indexC); j++ {
			if indexC[j]-i == 0 {
				temp = 0
				break
			}
			if int(math.Abs(float64(indexC[j]-i))) < temp {
				temp = int(math.Abs(float64(indexC[j] - i)))
			}
		}
		ans[i] = temp
	}
	return ans
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
