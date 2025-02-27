package main

import "fmt"

/**
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one’s added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
*/

// 思路： 加法的话正反加都无所谓，所以得考虑减法，反着加记录上一个值然后比较当前值与上个值的大小关系。若当前值小于上一值就为减法。
func main() {
	fmt.Println(romanToInt2("DCCCXCIX"))
}

var tempMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum, lastnum, num := 0, 0, 0
	if len(s) == 0 {
		return sum
	}
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		fmt.Println(char)
		num = tempMap[char]
		if num < lastnum {
			sum = sum - num
		} else {
			sum = sum + num
		}
		lastnum = num
	}

	return sum
}

func romanToInt2(s string) int {
	var last, total int
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i : i+1]
		num := tempMap[char]
		if num < last {
			total -= num
		} else {
			total += num
		}
		last = num
	}
	return total
}
