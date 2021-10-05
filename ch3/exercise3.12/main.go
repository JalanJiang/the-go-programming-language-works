// 判断两个字符串是否同文异构
package main

import "fmt"

func isSame(a string, b string) bool {
	// 长度不同不符合要求
	if len(a) != len(b) {
		return false
	}
	table := make(map[rune]int)
	for _, c := range a {
		table[c]++
	}
	for _, c := range b {
		if table[c] == 0 {
			return false
		}
		table[c]--
	}
	return true
}

func main() {
	fmt.Println(isSame("abcd", "dbca"))
	fmt.Println(isSame("abcdeee", "dbcaeea"))
	fmt.Println(isSame("abcdeee", "dbecaee"))
}
