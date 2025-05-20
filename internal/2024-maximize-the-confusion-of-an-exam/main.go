package _024_maximize_the_confusion_of_an_exam

// https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/
func maxConsecutiveAnswers(answerKey string, k int) int {
	l, ans, cnt := 0, 0, make(map[uint8]int)
	for r, v := range answerKey {
		cnt[uint8(v)]++
		for cnt['T'] > k && cnt['F'] > k {
			cnt[answerKey[l]]--
			l++
		}

		ans = max(ans, r-l+1)
	}
	return ans
}
