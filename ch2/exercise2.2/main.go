package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JalanJiang/the-go-programming-language-works/ch2/exercise2.2/lengthunit"
)

func main() {
	// 读取命令行参数
	for _, arg := range os.Args[1:] {
		// 转为 64 位小数
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := lengthunit.Centimeter(t)
		m := lengthunit.Metre(t)
		fmt.Printf("%s = %s, %s = %s\n", c, lengthunit.CToM(c), m, lengthunit.MToC(m))
	}
}
