package main

import (
	"fmt"
	"os"
)

func main() {
	// demo
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
	// 1.1：输出命令的名字
	fmt.Println(os.Args[0])
	// 1.2：输出参数和索引值
	for index, value := range os.Args[1:] {
		echoString := fmt.Sprintf("index=%d, value=%s", index, value)
		fmt.Println(echoString)
	}
	// TODO 1.3：strings.Join 效率测试，等到 1.6 节完成
}
