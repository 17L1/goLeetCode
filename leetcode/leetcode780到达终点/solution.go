/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-09 15:39:47
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-09 16:09:27
 */
package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for sx < tx && sy < ty {
		if tx < ty {
			ty %= tx
		} else {
			tx %= ty
		}
	}
	if tx < sx || ty < sy {
		return false
	}
	if sx == tx {
		return (ty-sy)%tx == 0
	} else {
		return (tx-sx)%ty == 0
	}
}
