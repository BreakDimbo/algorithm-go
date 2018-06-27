package sort

import (
	"fmt"
)

/*
MergeSort



*/

type Merge struct {
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
	fmt.Println("get the array")

	e := Merge{}
	e.show(c)
	fmt.Println("start to sort")

	e.Sort(c)
	if !e.IsSorted(c) {
		fmt.Errorf("array %v not sorted", c)
		return
	}

	e.show(c)
}

func (e *Merge) Sort(c []Comparable) {
	// find the value of h
	length := len(c)
	h := 0
	for h < length/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && e.less(c[j], c[j-h]); j -= h {
				e.exch(c, j, j-h)
			}
		}
		h = h / 3
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
