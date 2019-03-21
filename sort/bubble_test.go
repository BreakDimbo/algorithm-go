package sort

import "testing"

func Test_bubbleSort(t *testing.T) {
	a := []int{4, 2, 1, 6, 5, 7, 9, 2}
	expect := []int{1, 2, 2, 4, 5, 6, 7, 9}
	bubbleSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != expect[i] {
			t.Errorf("sort a not right: %v", a)
		}
	}
}
