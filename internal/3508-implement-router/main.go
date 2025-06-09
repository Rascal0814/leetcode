package _508_implement_router

type Router struct {
	arr         [][3]int
	memoryLimit int
	dest        map[int][]int
	packetMap   map[[3]int]struct{}
}

func Constructor(memoryLimit int) Router {
	return Router{arr: make([][3]int, 0), memoryLimit: memoryLimit, dest: make(map[int][]int), packetMap: make(map[[3]int]struct{})}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	tmp := [3]int{source, destination, timestamp}
	if _, ok := this.packetMap[tmp]; ok {
		return false
	}
	this.packetMap[tmp] = struct{}{}
	this.dest[destination] = append(this.dest[destination], timestamp)
	if len(this.arr) >= this.memoryLimit {
		this.ForwardPacket()
	}
	this.arr = append(this.arr, tmp)
	return true

}

func (this *Router) ForwardPacket() []int {
	if len(this.arr) == 0 {
		return []int{}
	}
	delete(this.packetMap, this.arr[0])
	ans := []int{this.arr[0][0], this.arr[0][1], this.arr[0][2]}
	this.arr = this.arr[1:]
	this.dest[ans[1]] = this.dest[ans[1]][1:]
	return ans
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	lowerBound := func(n []int, target int) int {
		l, r := 0, len(n)-1
		for l <= r {
			if mid := (l + r) >> 1; n[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}
	if len(this.dest[destination]) == 0 {
		return 0
	}

	l1 := lowerBound(this.dest[destination], startTime)
	l2 := lowerBound(this.dest[destination], endTime+1)

	return l2 - l1
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
