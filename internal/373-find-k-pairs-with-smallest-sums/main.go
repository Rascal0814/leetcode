package _73_find_k_pairs_with_smallest_sums

//https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// todo
	var cnt [][]int

	len1, len2, index1, index2 := len(nums1), len(nums2), 0, 0

	for len(cnt) < k {
		cnt = append(cnt, []int{nums1[index1], nums2[index2]})
		pre1 := index1 + 1
		pre2 := index2 + 1
		if pre1 >= len1 {
			index2 = pre2
			continue
		}

		if pre2 >= len2 {
			index1 = pre2
			continue
		}

		if pre1 < len1 && pre2 < len2 && nums2[pre2]+nums1[index1] < nums2[index2]+nums1[pre1] {
			index2 = pre2
		} else {
			index1 = pre1
		}

	}
	return cnt
}
