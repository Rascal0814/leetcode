package _38_product_of_array_except_self

const _ = `
238. 除自身以外数组的乘积

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。`

func productExceptSelf(nums []int) []int {
	var preProduct = make(map[int]int)
	preProduct[0] = 1
	var sufProduct = make(map[int]int)
	sufProduct[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		preProduct[i] = nums[i-1] * preProduct[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		sufProduct[i] = nums[i+1] * sufProduct[i+1]
	}

	var res []int

	for i := 0; i < len(nums); i++ {
		res = append(res, preProduct[i]*sufProduct[i])
	}

	return res
}

// o1空间复杂度
func productExceptSelf2(nums []int) []int {
	var res = make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i-1] * res[i-1]
	}

	var x = 1
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] *= x * nums[i+1]
		x *= nums[i+1]
	}

	return res
}
