// 循环写法替代 popcount
// TODO：和原写法性能对比
package exercise23

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(0*8))])
	}
	return result
}
