/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-10 09:21:33
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-10 09:42:49
 */
package main

func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS < 1 {
		return s
	}
	var res string
	for i := 0; i < lenS; i++ {
		//如果回文串长度为奇数
		l := i - 1
		r := i + 1
		tmp := getString(s, l, r)
		if len(tmp) > len(res) {
			res = tmp
		}

		//如果回文串长度为偶数
		l = i - 1
		r = i
		tmp = getString(s, l, r)
		if len(tmp) > len(res) {
			res = tmp
		}
	}
	return res
}

func getString(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l : r+1]
}
