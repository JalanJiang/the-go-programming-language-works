package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	// 数字每 3 位插入逗号
	var n = len(s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if (len(s)-i)%3 == 1 && i != len(s)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567891111"))
}
