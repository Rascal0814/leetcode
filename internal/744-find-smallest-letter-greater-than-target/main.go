package _44_find_smallest_letter_greater_than_target

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l <= r {
		if mid := (r + l) >> 1; letters[mid] < target+1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l >= len(letters) || letters[l] <= target {
		return letters[0]
	}
	return letters[l]
}
