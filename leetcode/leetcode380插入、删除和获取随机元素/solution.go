/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-13 22:18:36
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-13 22:56:08
 */
package main

import "math/rand"

type RandomizedSet struct {
	nums      []int
	randomMap map[int]int
}

func Constructor() RandomizedSet {

	return RandomizedSet{
		nums:      make([]int, 0),
		randomMap: make(map[int]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.randomMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.randomMap[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.randomMap[val]; ok {
		lenNums := len(this.nums)
		indexVal := this.randomMap[val]
		this.nums[indexVal] = this.nums[lenNums-1]
		this.randomMap[this.nums[indexVal]] = indexVal
		this.nums = this.nums[:lenNums-1]
		delete(this.randomMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(this.nums))
	return this.nums[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
