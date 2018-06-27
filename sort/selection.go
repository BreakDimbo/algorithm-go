package sort

import (
	"fmt"
)

type Selection struct {
}

func LaunchSelect() {
	c := []Comparable{}
	c = append(c, CString("f"))
	c = append(c, CString("d"))
	c = append(c, CString("s"))
	c = append(c, CString("w"))
	c = append(c, CString("v"))
	c = append(c, CString("d"))
	c = append(c, CString("b"))
	fmt.Println("get the array")

	e := Selection{}
	e.show(c)
	fmt.Println("start to sort")

	e.Sort(c)
	if !e.IsSorted(c) {
		fmt.Errorf("array %v not sorted", c)
		return
	}
	e.show(c)
}

func (e *Selection) Sort(c []Comparable) {
	for i := 0; i < len(c); i++ {
		minIndex := i
		for j := i + 1; j < len(c); j++ {
			if e.less(c[j], c[minIndex]) {
				minIndex = j
			}
		}
		e.exch(c, i, minIndex)
	}
}

func (e *Selection) IsSorted(c []Comparable) bool {
	for i := 1; i < len(c); i++ {
		if e.less(c[i], c[i-1]) {
			return false
		}
	}
	return true
}

func (e *Selection) show(c []Comparable) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s, ", c[i])
	}
	fmt.Println()
}

func (e *Selection) less(v Comparable, w Comparable) bool {
	return v.CompareTo(w) < 0
}

func (e *Selection) exch(c []Comparable, i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}
