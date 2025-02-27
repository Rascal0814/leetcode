package _0_remove_duplicates_from_sorted_array_ii

const _ = `
80. 删除有序数组中的重复项 II


给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
`

func removeDuplicates(nums []int) int {
	var exitCount = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c, ok := exitCount[nums[i]]
		if c == 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}

		if !ok || c < 2 {
			exitCount[nums[i]]++
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	var fast, slow = 2, 2
	if len(nums) <= 2 {
		return len(nums)
	}

	for fast < len(nums) {
		// 已经检查完的前两个和当前待检查数据进行比较
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return len(nums[:slow])
}
