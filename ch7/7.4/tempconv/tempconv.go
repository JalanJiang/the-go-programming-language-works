package tempconv

import (
	"flag"
	"fmt"

	"github.com/JalanJiang/the-go-programming-language-works/ch2/tempconv"
)

// 满足 flag.Value 的接口
// type Value interface {
// 	String() string
// 	Set(string) error
// }

type celsiusFlag struct{ tempconv.Celsius }

// Set 函数
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "­°C":
		f.Celsius = tempconv.Celsius(value)
	case "F", "­°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
	case "K":
		f.Celsius = tempconv.KToC(tempconv.Kelvins(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	// 【重点】f 需要满足 Var.Value 接口
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
