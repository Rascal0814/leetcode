package _07_range_sum_query_mutable

const _ = `
给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组 nums 下标对应的值
另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
实现 NumArray 类：

NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）


示例 1：

输入：
	["NumArray", "sumRange", "update", "sumRange"]
	[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
输出：
	[null, 9, null, 8]

解释：
	NumArray numArray = new NumArray([1, 3, 5]);
	numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
	numArray.update(1, 2);   // nums = [1,2,5]
	numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
`

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (this *NumArray) Update(index int, val int) {

	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
