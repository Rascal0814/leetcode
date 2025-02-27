package main

import "fmt"

func main() {
	//fmt.Println(removeDuplicates([]int{1, 2, 3, 4, 4, 4, 5, 7}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates3([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

}

func removeDuplicates2(nums []int) int {
	newArr := make([]int, 0)
	n := 0
	if len(nums) <= 1 {
		return len(nums)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		newArr = append(newArr, nums[i])
		n++
	}

	if len(nums) > 2 && nums[len(nums)-1] != nums[len(nums)-2] {
		newArr = append(newArr, nums[len(nums)-1])
		n++
	}

	fmt.Println(newArr)
	fmt.Println(n)

	return len(newArr)
}

// 解法一
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				fmt.Println(nums)
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}

	return last + 1
}

// 0, 0, 1, 1, 1, 2, 2, 3, 3, 4
func removeDuplicates3(nums []int) int {
	// 双指针 快/慢
	var fast, slow = 1, 1

	for fast < len(nums) {
		// 前一个数据和当前数据比较
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
		fmt.Println(nums)
	}
	return len(nums[:slow])
}
