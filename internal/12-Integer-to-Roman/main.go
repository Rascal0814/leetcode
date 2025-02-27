package main

import "fmt"

func main() {
	fmt.Println(intToRoman(899))
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	i, res := 0, ""
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}

func intToRoman2(num int) string {
	a, b, c, d := num/1000%100%10, num/100%10, num/10%10, num/1%10
	ge := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	shi := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	bai := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	qian := []string{"", "M", "MM", "MMM"}
	str1 := ""
	str1 += qian[a] + bai[b] + shi[c] + ge[d]
	return str1
}
func intToRoman3(num int) string {
	res, i := "", 0
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for num != 0 {
		if values[i] > num {
			i++
			continue
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}
