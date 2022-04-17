/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-17 13:24:58
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-17 13:59:20
 */
package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	var ans string
	ansCount := 0
	wordMap := make(map[string]int, 0)
	bannedMap := make(map[string]bool, 0)
	for _, v := range banned {
		bannedMap[v] = true
	}
	var temp string
	for i, v := range paragraph {
		if v-'a' < 0 || v-'z' > 0 {
			if temp == "" {
				continue
			}
			if _, ok := wordMap[temp]; !ok {
				wordMap[temp] = 1
			} else {
				wordMap[temp]++
			}
			_, ok := bannedMap[temp]
			if wordMap[temp] > ansCount && !ok {
				ans = temp
				ansCount = wordMap[temp]
			}
			temp = ""
			continue
		}
		temp += string(v)
		if i == len(paragraph)-1 {
			if temp == "" {
				continue
			}
			if _, ok := wordMap[temp]; !ok {
				wordMap[temp] = 1
			} else {
				wordMap[temp]++
			}
			_, ok := bannedMap[temp]
			if wordMap[temp] > ansCount && !ok {
				ans = temp
				ansCount = wordMap[temp]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
