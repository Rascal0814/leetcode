package _631_path_with_minimum_effort

//https://leetcode.cn/problems/path-with-minimum-effort/description/
func minimumEffortPath(heights [][]int) int {
	l, r := -1, 1_000_000
	for l+1 < r {
		if mid := (l + r) >> 1; bfs(heights, mid) {
			r = mid
		} else {
			l = mid
		}
	}

	return r
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 广度优先遍历
func bfs(list [][]int, k int) bool {
	type pairs struct{ x, y int }
	dirs := []pairs{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} // 相邻点
	n, m := len(list), len(list[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	vis[0][0] = true
	var queue []pairs
	queue = append(queue, pairs{0, 0}) // 插入队列
	for len(queue) != 0 {
		cnt := queue[0]
		if cnt.x == n-1 && cnt.y == m-1 {
			return true
		}
		for _, v := range dirs {
			t := pairs{cnt.x + v.x, cnt.y + v.y}
			if t.x >= 0 && t.y >= 0 && t.x < n && t.y < m && !vis[t.x][t.y] && abs(list[cnt.x][cnt.y]-list[t.x][t.y]) <= k {
				queue = append(queue, t)
				vis[t.x][t.y] = true
			}
		}
		queue = queue[1:]
	}

	return false

}
