package lengthunit

import "fmt"

type Centimeter float64
type Metre float64

// 厘米转米
func CToM(c Centimeter) Metre {
	return Metre(c / 100)
}

// 米转厘米
func MToC(m Metre) Centimeter {
	return Centimeter(m * 100)
}

// 实现 String()
func (c Centimeter) String() string {
	return fmt.Sprintf("%gcm", c)
}

func (m Metre) String() string {
	return fmt.Sprintf("%gm", m)
}
