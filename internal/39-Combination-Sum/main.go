package main

import (
	"fmt"
	"sort"
)

const q = `
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 

对于给定的输入，保证和为 target 的不同组合数少于 150 个。


示例 1：
	
	输入：candidates = [2,3,6,7], target = 7
	输出：[[2,2,3],[7]]
	解释：
	2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
	7 也是一个候选， 7 = 7 。
	仅有这两种组合。
	示例 2：
	
	输入: candidates = [2,3,5], target = 8
	输出: [[2,2,2,2],[2,3,3],[3,5]]
	示例 3：
	
	输入: candidates = [2], target = 1
	输出: []
 

提示：
	1 <= candidates.length <= 30
	2 <= candidates[i] <= 40
	candidates 的所有元素 互不相同
	1 <= target <= 40
`

// 回溯算法
func combinationSum(candidates []int, target int) [][]int {
	// 先对 candidates 排序，便于后面处理和去重
	sort.Ints(candidates)

	var res = make([][]int, 0)
	var comb = make([]int, 0)

	var dfs func(target, index int)
	dfs = func(target, index int) {
		if target == 0 {
			// 找到一个符合条件的组合 因为comb是一个切片 引用类型 全局使用 如果直接插入的话后续的操作会改变数组中的值
			temp := make([]int, len(comb))
			copy(temp, comb)
			res = append(res, comb)
			return
		}

		for i := index; i < len(candidates); i++ {
			// 如果当前的数字大于剩余的 target，那么之后的数字肯定也大于 target，直接跳出循环
			if candidates[i] > target {
				break
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i)

			// 回溯，从组合中移除当前的数字
			comb = comb[:len(comb)-1]
		}
	}

	dfs(target, 0)
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(combinationSum(nums, 3))
}
