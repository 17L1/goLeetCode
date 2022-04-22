/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-03-23 14:14:32
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-06 14:23:10
 */
package main

import "fmt"

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, left, right := 1, 0, 0
	var ch [256]int
	for right < len(s) {
		ch[s[right]]++
		for ch[s[right]] > 1 {
			ch[s[left]]--
			left++
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
