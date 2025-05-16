package _652_defuse_the_bomb

// https://leetcode.cn/problems/defuse-the-bomb/
func decrypt(code []int, k int) []int {
	n, res := len(code), make([]int, len(code))
	if k < 0 {
		for i := k; i < 0; i++ {
			res[0] += code[n+i]
		}

		for i := 1; i < n; i++ {
			res[i] = res[i-1] + code[i-1] - code[(n+k+i-1)%n]
		}
	} else if k == 0 {
		return res
	} else {
		for i := 0; i < k; i++ {
			res[0] += code[i+1]
		}

		for i := 1; i < n; i++ {
			res[i] = res[i-1] + code[(i+k)%n] - code[i]
		}

	}
	return res
}
