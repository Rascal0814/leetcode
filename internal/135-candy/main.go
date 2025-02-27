package _35_candy

const _ = `
135. 分发糖果

n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。`

func candy(ratings []int) int {
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var candyLeftList = make([]int, len(ratings))
	var candyRightList = make([]int, len(ratings))
	candyLeftList[0] = 1
	for i := 1; i < len(ratings); i++ {
		// 计算左侧需要的candy数量
		if ratings[i] > ratings[i-1] {
			candyLeftList[i] = candyLeftList[i-1] + 1
		} else {
			candyLeftList[i] = 1
		}
	}

	candyRightList[len(ratings)-1] = 1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyRightList[i] = candyRightList[i+1] + 1
		} else {
			candyRightList[i] = 1
		}
	}
	//先算左边需要的糖果数再算右边 在取两者的最大值
	total := 0
	for i := 0; i < len(ratings); i++ {
		total += maxInt(candyLeftList[i], candyRightList[i])
	}

	return total

}
