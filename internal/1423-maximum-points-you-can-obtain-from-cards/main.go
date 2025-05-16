package _423_maximum_points_you_can_obtain_from_cards

// https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/
func maxScore(cardPoints []int, k int) int {
	var total, sum, n = 0, 0, len(cardPoints)
	for i, v := range cardPoints {
		if i < n-k {
			sum += v
		}
		total += v
	}
	var res = total - sum
	for i := n - k; i < len(cardPoints); i++ {
		sum += cardPoints[i] - cardPoints[i-n+k]
		res = max(res, total-sum)
	}

	return res
}
