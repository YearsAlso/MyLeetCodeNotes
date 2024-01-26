package LCR030

import "math/rand"

type RandomizedSet struct {
	mapList []int
	// 顺序表
	numberMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		mapList:   make([]int, 0),
		numberMap: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numberMap[val]; ok {
		return false
	}

	this.mapList = append(this.mapList, val)
	this.numberMap[val] = len(this.mapList) - 1

	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numberMap[val]; !ok {
		return false
	}

	index := this.numberMap[val]

	if index == len(this.mapList)-1 {
		this.mapList = this.mapList[:len(this.mapList)-1]
		delete(this.numberMap, val)
		return true
	}

	// 如果不是最后一个元素，那么就将最后一个元素放到当前位置
	this.mapList[index] = this.mapList[len(this.mapList)-1]
	this.mapList = this.mapList[:len(this.mapList)-1]
	this.numberMap[this.mapList[index]] = index
	delete(this.numberMap, val)

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	// 获得随机数
	randNumber := rand.Intn(len(this.mapList))

	return this.mapList[randNumber]
}
