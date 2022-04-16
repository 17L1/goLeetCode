/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-16 13:27:43
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-16 14:00:36
 */
package main

import (
	"fmt"
	"math"
)

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	//或取回文数的前半部分
	upper := int(math.Pow10(n)) - 1
	for left := upper; ; left-- {
		p := left
		//构造回文数的后半部分
		for i := left; i > 0; i /= 10 {
			p = p*10 + i%10
		}
		//枚举获取因子
		for x := upper; x*x >= p; x-- {
			if p%x == 0 {
				return p % 1337
			}
		}
	}
}

func main() {
	n := 2
	fmt.Println(largestPalindrome(n))
}
