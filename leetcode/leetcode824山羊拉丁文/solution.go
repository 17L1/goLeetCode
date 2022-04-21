/*
 * @Description:
 * @Author: 17L1
 * @Date: 2022-04-21 11:44:29
 * @LastEditors: 17L1
 * @LastEditTime: 2022-04-21 12:31:41
 */
package main

import "strings"

var vowels = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}

func toGoatLatin(sentence string) string {
	ans := &strings.Builder{}
	for i, cnt, n := 0, 1, len(sentence); i < n; i++ {
		if cnt > 1 {
			ans.WriteByte(' ')
		}
		start := i
		for i++; i < n && sentence[i] != ' '; i++ {
		}
		cnt++
		if _, ok := vowels[sentence[start]]; ok {
			ans.WriteString(sentence[start:i])
		} else {
			ans.WriteString(sentence[start+1 : i])
			ans.WriteByte(sentence[start])
		}
		ans.WriteByte('m')
		ans.WriteString(strings.Repeat("a", cnt))
	}
	return ans.String()
}
