// 处理浮点数和正负数
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	// 数字每 3 位插入逗号
	// 处理正负号
	symbol := s[:1]
	if symbol == "+" || symbol == "-" {
		// 摘除
		s = s[1:]
	} else {
		symbol = ""
	}
	// 判断是否有小数
	// 找不到 slash = -1
	last := ""
	if slash := strings.LastIndex(s, "."); slash >= 0 {
		// 去掉小数部分
		last = s[slash:]
		s = s[:slash]
	}

	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	tmp := buf.String()
	// 最终拼接
	return symbol + tmp + last
}

func main() {
	fmt.Println(comma("1234567891111"))
	fmt.Println(comma("-1234567891111"))
	fmt.Println(comma("1234567891111.2234"))
	fmt.Println(comma("+1234567891111.2234"))
}
