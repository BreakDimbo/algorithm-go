package sort

/*
func insertSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				// 数据移动
				a[j+1] = a[j]
			} else {
				break
			}
		}
		// 插入数据
		a[j+1] = value
	}
}
*/

/*
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}
}
*/