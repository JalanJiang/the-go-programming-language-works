// exercise2.5：利用 x&(x-1) 可以清除最右边的非 0 位
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
		if x != x&(x-1) {
			// 最右边是 1
			result++
		}
		x = x >> 1
	}

	return result
}
