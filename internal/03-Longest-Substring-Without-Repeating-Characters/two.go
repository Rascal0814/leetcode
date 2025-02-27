package main

func two(s string) int {
	l, r, max := 0, 0, 0
	// tempbucket 用户判断该byte值是否存在 存在就缩小左边界直到不存在
	var tempbucket = make(map[byte]bool)
	for l < len(s) && r < len(s) {
		if tempbucket[s[r]] {
			tempbucket[s[l]] = false
			l++
		} else {
			tempbucket[s[r]] = true
			r++
		}

		if max < r-l {
			max = r - l
		}
	}

	return max
}
