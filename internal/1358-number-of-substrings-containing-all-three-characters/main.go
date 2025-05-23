package _358_number_of_substrings_containing_all_three_characters

// https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/description/
func numberOfSubstrings(s string) int {
	l, ans, exit := 0, 0, [3]int{}
	for _, v := range s {
		exit[v-'a']++
		for exit[0] > 0 && exit[1] > 0 && exit[2] > 0 {
			exit[s[l]-'a']--
			l++
		}
		ans += l
	}
	return ans

}

//func numberOfSubstrings(s string) int {
//	ans, exit := 0, [3]int{}
//	for l := 0; l <= len(s)-3; l++ {
//		for r := l; r < len(s); r++ {
//			exit[s[r]-'a'] = 1
//			if exit[0] > 0 && exit[1] > 0 && exit[2] > 0 {
//				ans += len(s) - r
//				exit = [3]int{}
//				break
//			}
//		}
//	}
//	return ans
//}
