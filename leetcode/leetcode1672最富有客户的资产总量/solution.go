/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-14 12:15:53
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-14 12:28:53
 */
package main

func maximumWealth(accounts [][]int) int {
	ans := 0
	if len(accounts) == 0 {
		return ans
	}
	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[0]); j++ {
			temp += accounts[i][j]
		}
		if temp > ans {
			ans = temp
		}
	}
	return ans
}
