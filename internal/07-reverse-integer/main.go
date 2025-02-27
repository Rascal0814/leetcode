package _7_reverse_integer

import "math"

func reverse(x int) int {
	var res int
	for x != 0 {
		//越界 math.Maxint32 < 10 * res + temp < math.MinInt32
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		// 取余
		temp := x % 10
		// 将余数放到res上
		res = res*10 + temp
		// 将x取整
		x = x / 10
	}

	return res
}
