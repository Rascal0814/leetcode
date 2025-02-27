package main

import (
	"fmt"
	"sort"
)

const q = `
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。 

示例 1:

	输入: candidates = [10,1,2,7,6,1,5], target = 8,
	输出:
	[
		[1,1,6],
		[1,2,5],
		[1,7],
		[2,6]
	]
示例 2:

	输入: candidates = [2,5,2,1,2], target = 5,
	输出:
	[
		[1,2,2],
		[5]
	]
 

提示:

	1 <= candidates.length <= 100
	1 <= candidates[i] <= 50
	1 <= target <= 30
`

func combinationSum2(candidates []int, target int) [][]int {
	// 排序
	sort.Ints(candidates)

	var (
		comb []int   // 单个结果
		res  [][]int // 结果
	)

	var dfs func(target, index int)

	dfs = func(target, index int) {
		if target == 0 {
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, temp)
			return
		}

		for i := index; i < len(candidates); i++ {
			// 去重
			if target < candidates[i] {
				break
			}

			// ？ 如果这已经拿出去一组了就得判断前后两个值是否一样 ？
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)

			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)

	return res
}

func main() {
	//nums := []int{10, 1, 2, 7, 6, 1, 5}
	//fmt.Println(combinationSum2(nums, 8))

	nums2 := []int{2, 5, 2, 1, 2}
	fmt.Println(combinationSum2(nums2, 5))
}
