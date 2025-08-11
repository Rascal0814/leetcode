package _812_find_the_safest_path_in_a_grid

type pairs struct{ x, y int }

var dirs = []pairs{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func maximumSafenessFactor(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var thieves []pairs             // 小偷所在位置
	var distance = make([][]int, n) // 安全距离
	for i := range distance {
		distance[i] = make([]int, m)
	}
	for x, row := range grid {
		for y, col := range row {
			if col > 0 {
				thieves = append(thieves, pairs{x, y})
			} else {
				distance[x][y] = -1
			}
		}
	}

	for len(thieves) != 0 {
		cnt := thieves[0]
		thieves = thieves[1:]
		for _, d := range dirs {
			t := pairs{cnt.x + d.x, cnt.y + d.y}
			if t.x >= 0 && t.y >= 0 && t.x < n && t.y < m && distance[t.x][t.y] == -1 {
				distance[t.x][t.y] = distance[cnt.x][cnt.y] + 1
				thieves = append(thieves, t)
			}
		}
	}

	l, r := 0, len(grid)+len(grid[0])
	for l+1 < r {
		if mid := (l + r) >> 1; check(distance, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func check(distance [][]int, k int) bool {
	n, m := len(distance), len(distance[0])
	if distance[0][0] < k {
		return false
	}
	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, m)
	}
	vis[0][0] = true
	queue := []pairs{{}}
	for len(queue) != 0 {
		cnt := queue[0]
		if cnt.x == n-1 && cnt.y == m-1 {
			return true
		}
		for _, d := range dirs {
			t := pairs{cnt.x + d.x, cnt.y + d.y}
			if t.x >= 0 && t.y >= 0 && t.x < n && t.y < m && !vis[t.x][t.y] && distance[t.x][t.y] >= k {
				vis[t.x][t.y] = true
				queue = append(queue, pairs{t.x, t.y})
			}
		}
		queue = queue[1:]
	}
	return false
}
