package sort

import (
	"fmt"
	"log"
	"math/rand"
)

/*
QuickSort



*/

type Quick struct {
}

func LaunchQuick() {
	c := []Comparable{}
	c = append(c, CString("f"))
	c = append(c, CString("d"))
	c = append(c, CString("s"))
	c = append(c, CString("w"))
	c = append(c, CString("v"))
	c = append(c, CString("d"))
	c = append(c, CString("b"))
	fmt.Println("Quick sort start")
	fmt.Println("get the array")

	e := Quick{}
	c = shuffle(c)
	e.show(c)
	fmt.Println("start to sort")

	e.Sort(c)
	if !e.IsSorted(c) {
		fmt.Printf("array %v not sorted", c)
		return
	}

	e.show(c)
}

func (e *Quick) Sort(c []Comparable) {
	e.sort(c, 0, len(c)-1)
}

func (e *Quick) sort(c []Comparable, lo int, hi int) {
	if lo >= hi {
		return
	}

	point := e.partition(c, lo, hi)
	e.sort(c, lo, point-1)
	e.sort(c, point+1, hi)
}

func (e *Quick) partition(c []Comparable, lo int, hi int) (point int) {
	i := lo + 1
	j := hi
	v := c[lo]
	for {
		// 左边越界
		for e.less(c[i], v) {
			if i >= hi {
				break
			}
			log.Printf("left index: %d", i)
			i++
		}

		// 右边越界
		for e.less(v, c[j]) {
			if j <= lo {
				break
			}
			log.Printf("right index: %d", i)
			j--
		}

		// 左右指针相遇
		if i >= j {
			log.Printf("jump index: %d", i)
			break
		}

		// point 左边值大于右边值
		e.exch(c, i, j)
	}
	e.exch(c, j, lo)
	return j
}

func (e *Quick) IsSorted(c []Comparable) bool {
	for i := 1; i < len(c); i++ {
		if e.less(c[i], c[i-1]) {
			return false
		}
	}
	return true
}

func (e *Quick) show(c []Comparable) {
	for i := 0; i < len(c); i++ {
		fmt.Printf("%s, ", c[i])
	}
	fmt.Println()
}

func (e *Quick) less(v Comparable, w Comparable) bool {
	return v.CompareTo(w) < 0
}

func (e *Quick) exch(c []Comparable, i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

func shuffle(src []Comparable) (dest []Comparable) {
	dest = make([]Comparable, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return
}
