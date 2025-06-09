package _81_time_based_key_value_store

type TimeValue struct {
	Val  string
	Time int
}

type TimeMap struct {
	arr map[string][]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{arr: make(map[string][]TimeValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.arr[key] = append(this.arr[key], TimeValue{Time: timestamp, Val: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	lowerBound := func(n []TimeValue, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (l + r) >> 1; n[mid].Time < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	i := lowerBound(this.arr[key], timestamp+1)
	if i == 0 {
		return ""
	}
	return this.arr[key][i-1].Val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
