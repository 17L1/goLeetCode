/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-18 11:01:34
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-18 12:42:00
 */
package main

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return ans
}
