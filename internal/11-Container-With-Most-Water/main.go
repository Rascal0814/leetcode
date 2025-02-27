package main

import "fmt"

//Given n non-negative integers a1, a2, …, an , where each represents a point at coordinate (i, ai). n vertical
//lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
//with x-axis forms a container, such that the container contains the most water.

//Note: You may not slant the container and n is at least 2.

//The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water
//(blue section) the container can contain is 49.
//Example 1:
//	Input: [1,8,6,2,5,4,8,3,7]
//	Output: 49

func main() {
	fmt.Println(two([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

// 首尾指针
func maxArea(height []int) int {
	head, end, max := 0, len(height)-1, 0

	for !(end == head) {
		if max < (end-head)*minInt(height, end, head) {
			max = (end - head) * minInt(height, end, head)
		}
		if height[end] > height[head] {
			head++
		} else {
			end--
		}

	}
	return max
}

func minInt(height []int, end, head int) int {
	if height[end] > height[head] {
		return height[head]
	}
	return height[end]
}

func maxArea2(height []int) int {
	l, r, maxWater := 0, len(height)-1, 0
	for l < r {
		nowWater := (r - l) * (min(height[l], height[r]))
		if nowWater > maxWater {
			maxWater = nowWater
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxWater
}
