package _80_insert_delete_getrandom_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	valMap   map[int]int
	exitList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valMap:   make(map[int]int),
		exitList: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valMap[val]; ok {
		return false
	}
	this.valMap[val] = len(this.exitList)
	this.exitList = append(this.exitList, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.valMap[val]; ok {
		lastId := len(this.exitList) - 1
		lastVal := this.exitList[lastId]

		this.valMap[lastVal] = index // 替换map中val index的值

		this.exitList[index] = lastVal         // 交换key
		this.exitList = this.exitList[:lastId] // 去除最末尾元素

		delete(this.valMap, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.exitList[rand.Intn(len(this.exitList))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
