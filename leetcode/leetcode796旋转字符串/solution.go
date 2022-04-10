/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-07 10:23:03
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-07 15:04:52
 */
package main

import "strings"

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
