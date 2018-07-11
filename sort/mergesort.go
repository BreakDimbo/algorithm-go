package sort

import (
	"fmt"
)

/*
MergeSort



*/

type Merge struct {
	aux []Comparable
}

func LaunchMerge() {
	c := []Comparable{}
	c = append(c, CString("f"))
	c = append(c, CString("d"))
	c = append(c, CString("s"))
	c = append(c, CString("w"))
	c = append(c, CString("v"))
	c = append(c, CString("d"))
	c = append(c, CString("b"))
	fmt.Println("merge sort start")
	fmt.Println("get the array")

	e := Merge{}
	e.show(c)
	fmt.Println("start to sort")

	e.Sort(c)
	if !e.IsSorted(c) {
		fmt.Printf("array %v not sorted", c)
		return
	}

	e.show(c)
}

func (e *Merge) Sort(c []Comparable) {
	e.aux = make([]Comparable, len(c))
	e.sort(c, 0, len(c)-1)
}

func (e *Merge) BUSort(c []Comparable) {
	length := len(c)
	e.aux = make([]Comparable, length)

	for sz := 1; sz < length; sz = sz + sz {
		for lo := 0; lo < length-sz; lo = lo + sz + sz {
			e.merge(c, lo, lo+sz-1, min(lo+sz+sz-1, length-1))
		}
	}
}

func (e *Merge) sort(c []Comparable, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	e.sort(c, lo, mid)
	e.sort(c, mid+1, hi)
	e.merge(c, lo, mid, hi)
}

func (e *Merge) merge(c []Comparable, lo int, mid int, hi int) {
	i := lo
	j := mid + 1

	for k, v := range c {
		e.aux[k] = v
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			// 左半边用尽，从右边取
			c[k] = e.aux[j]
			j++
		} else if j > hi {
			// 右半边用尽，从左边取
			c[k] = e.aux[i]
			i++
		} else if e.less(e.aux[i], e.aux[j]) {
			// 左边元素小于右边元素，从左边取
			c[k] = e.aux[i]
			i++
		} else {
			// 右边元素小于等于左边元素，从右边取
			c[k] = e.aux[j]
			j++
		}
	}
}

func (e *Merge) IsSorted(c []Comparable) bool {
	for i := 1; i < len(c); i++ {
		if e.less(c[i], c[i-1]) {
			return false
		}
	}
	return true
}

func (e *Merge) show(c []Comparable) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s, ", c[i])
	}
	fmt.Println()
}

func (e *Merge) less(v Comparable, w Comparable) bool {
	return v.CompareTo(w) < 0
}

func (e *Merge) exch(c []Comparable, i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
