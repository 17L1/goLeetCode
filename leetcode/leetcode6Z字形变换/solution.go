/*
 * @Description:Z字形变换，字符串,时间空间复杂度比较低
 * @Author: lilongguang
 * @Date: 2022-04-12 13:19:56
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-12 13:52:10
 */
package main

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i, flag := 0, -1
	for _, v := range s {
		res[i] += string(v)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	var ans string
	for _, v := range res {
		ans += v
	}
	return ans
}
