package _011_capacity_to_ship_packages_within_d_days

func shipWithinDays(weights []int, days int) int {
	var l, r int
	for _, weight := range weights {
		if l < weight {
			l = weight
		}
		r += weight
	}
	a := func(x int) int {
		sum, day := 0, 1
		for _, weight := range weights {
			sum += weight
			if sum > x {
				day++
				sum = weight
			}
		}
		return day
	}

	for l <= r {
		mid := (l + r) >> 1
		if a(mid) <= days {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
