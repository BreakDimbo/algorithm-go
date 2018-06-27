package sort

import (
	"fmt"
)

type Insertion struct {
}

func LaunchInsertion() {
	c := []Comparable{}
	c = append(c, CString("f"))
	c = append(c, CString("d"))
	c = append(c, CString("s"))
	c = append(c, CString("w"))
	c = append(c, CString("v"))
	c = append(c, CString("d"))
	c = append(c, CString("b"))
	fmt.Println("get the array")

	e := Insertion{}
	e.show(c)
	fmt.Println("start to sort")

	e.Sort(c)
	if !e.IsSorted(c) {
		fmt.Errorf("array %v not sorted", c)
		return
	}
	e.show(c)
}

func (e *Insertion) Sort(c []Comparable) {
	for i := 1; i < len(c); i++ {
		for j := i; j > 0 && e.less(c[j], c[j-1]); j-- {
			e.exch(c, j, j-1)
		}
	}
}

func (e *Insertion) IsSorted(c []Comparable) bool {
	for i := 1; i < len(c); i++ {
		if e.less(c[i], c[i-1]) {
			return false
		}
	}
	return true
}

func (e *Insertion) show(c []Comparable) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s, ", c[i])
	}
	fmt.Println()
}

func (e *Insertion) less(v Comparable, w Comparable) bool {
	return v.CompareTo(w) < 0
}

func (e *Insertion) exch(c []Comparable, i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}
