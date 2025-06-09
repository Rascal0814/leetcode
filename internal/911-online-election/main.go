package _11_online_election

type TopVotedCandidate struct {
	timeMap []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	res := TopVotedCandidate{times: times, timeMap: make([]int, len(persons))}
	var personMap = make(map[int]int)
	var maxPersonNum = persons[0]
	for i, _ := range times {
		personMap[persons[i]]++
		if personMap[persons[i]] >= personMap[maxPersonNum] {
			maxPersonNum = persons[i]
		}
		res.timeMap[i] = maxPersonNum
	}
	return res
}

func (this *TopVotedCandidate) Q(t int) int {
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

	left := lowerBound(this.times, t)
	if left >= len(this.times) || this.times[left] > t {
		left--
	}
	return this.timeMap[left]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
