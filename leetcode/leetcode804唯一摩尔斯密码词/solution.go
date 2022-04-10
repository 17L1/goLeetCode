/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-10 09:06:14
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-10 09:19:41
 */
package main

func uniqueMorseRepresentations(words []string) int {
	password := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
		"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	nums := make(map[string]int, 0)

	lenWords := len(words)
	if lenWords < 1 {
		return 0
	}
	for i := 0; i < lenWords; i++ {
		var tmp string
		for _, v := range words[i] {
			tmp += password[v-'a']
		}
		if _, ok := nums[tmp]; ok {
			nums[tmp]++
			continue
		}
		nums[tmp] = 1
	}
	return len(nums)
}
