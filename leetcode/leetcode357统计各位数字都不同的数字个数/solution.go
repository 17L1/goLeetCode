/*
 * @Description:排列组合方法
 * @Author: lilongguang
 * @Date: 2022-04-11 11:26:50
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-11 22:07:10
 */
package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res := 10

	for i, last := 2, 9; i <= n; i++ {
		cur := last * (10 - i + 1)
		res += cur
		last = cur
	}
	return res

}
