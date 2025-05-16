package _052_grumpy_bookstore_owner

// https://leetcode.cn/problems/grumpy-bookstore-owner/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var res, total int

	// 计算所有不生气顾客
	for i, v := range customers {
		if grumpy[i] == 0 {
			total += v
		}
	}

	// 计算生气中最大顾客数
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 1 {
			total += customers[i]
		}
		if i < minutes {
			res = total
			continue
		}
		if grumpy[i-minutes] == 1 {
			total -= customers[i-minutes]
		}
		res = max(total, res)
	}
	return res
}
