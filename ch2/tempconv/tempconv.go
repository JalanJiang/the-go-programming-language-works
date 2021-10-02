package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvins float64 // 热力学温度

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 摄氏转华氏
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 华氏转摄氏
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 开尔文温度转摄氏度
func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
