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
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567891111"))
}
