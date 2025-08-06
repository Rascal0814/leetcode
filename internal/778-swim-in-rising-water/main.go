package _78_swim_in_rising_water

//https://leetcode.cn/problems/swim-in-rising-water/description/

func swimInWater(grid [][]int) int {
	l, r := -1, len(grid)*len(grid)
	for l+1 < r {
		if mid := (l + r) >> 1; bfs(grid, mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func bfs(grid [][]int, k int) bool {
	type pairs struct{ x, y int }
	dirs := []pairs{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	var n, m = len(grid), len(grid[0])
	var vis = make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	if grid[0][0] > k {
		return false
	}

	queue := []pairs{{}}

	for len(queue) != 0 {
		cnt := queue[0]
		if cnt.x == n-1 && cnt.y == m-1 {
			return true
		}
		for _, d := range dirs {
			c := pairs{d.x + cnt.x, d.y + cnt.y}
			if c.x >= 0 && c.y >= 0 && c.x < n && c.y < m && !vis[c.x][c.y] && grid[c.x][c.y] <= k {
				queue = append(queue, c)
				vis[c.x][c.y] = true
			}
		}
		queue = queue[1:]
	}

	return false
}
