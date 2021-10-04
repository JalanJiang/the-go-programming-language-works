// exercise2.4：移位，每次判断最右边的位来计算 1 的个数
package exercise24

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result int

	for i := 0; i < 64; i++ {
		// x 向右移动 1 位
		x = x >> i
		// 和 1 进行与操作保留最右边 1 位
		x = x & 1
		// 结果加到 result 中
		result += int(x)
	}

	return result
}
