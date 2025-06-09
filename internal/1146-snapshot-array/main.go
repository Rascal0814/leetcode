package _146_snapshot_array

// https://leetcode.cn/problems/snapshot-array/
type SnapshotArray struct {
	arr   map[int][][2]int
	count int
}

func Constructor(int) SnapshotArray {
	sa := SnapshotArray{
		arr:   make(map[int][][2]int),
		count: 0,
	}
	return sa
}

func (this *SnapshotArray) Set(index int, val int) {
	this.arr[index] = append(this.arr[index], [2]int{val, this.count})
}

func (this *SnapshotArray) Snap() int {
	this.count++
	return this.count - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	lowerBound := func(n [][2]int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (r + l) >> 1; n[mid][1] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	i := lowerBound(this.arr[index], snap_id+1) - 1
	if i < 0 {
		return 0
	}

	return this.arr[index][i][0]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
