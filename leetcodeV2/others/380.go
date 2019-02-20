package others

import "math/rand"

/*
	solution 1: hash + slice(for random get)
*/

type RandomizedSet struct {
	m map[int]int
	s []int
}

/** Initialize your data structure here. */
func ConstructorXX() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		s: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.s)
		this.s = append(this.s, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.m[val]; ok {
		n := len(this.s)
		this.s[index], this.s[n-1] = this.s[n-1], this.s[index]
		this.m[this.s[index]] = index
		delete(this.m, val)
		this.s = this.s[:len(this.s)-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.s[rand.Intn(len(this.s))]
}
