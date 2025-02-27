package main

func two(height []int) int {
	max, l, r := 0, 0, len(height)-1

	for l <= r {
		c := r - l
		h := min(height[l], height[r])
		if max < c*h {
			max = c * h
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max

}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
