package _080_range_frequency_queries

type RangeFreqQuery struct {
	position map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{position: make(map[int][]int)}
	for i, v := range arr {
		r.position[v] = append(r.position[v], i)
	}
	return r
}

func (rr *RangeFreqQuery) Query(left int, right int, value int) int {
	lowerBound := func(n []int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	l1 := lowerBound(rr.position[value], left)
	l2 := lowerBound(rr.position[value], right+1)
	return l2 - l1
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
