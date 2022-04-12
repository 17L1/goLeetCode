/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-12 12:54:06
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-12 13:17:47
 */
package main

func numberOfLines(widths []int, s string) []int {
	var res []int
	num := 0
	level := 1
	for _, k := range s {
		if num+widths[k-'a'] > 100 {
			level++
			num = widths[k-'a']
			continue
		}
		num += widths[k-'a']

	}
	res = append(res, level, num)

	return res
}
