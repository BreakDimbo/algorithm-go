package sort

// func bubbleSort(a []int, n int) {
// 	if n <= 1 {
// 		return
// 	}

// 	for i := 0; i < len(a); i++ {
// 		isSwap := false
// 		for j := 0; j < len(a)-i-1; j++ {
// 			if a[j+1] < a[j] {
// 				a[j], a[j+1] = a[j+1], a[j]
// 				isSwap = true
// 			}
// 		}
// 		if !isSwap {
// 			return
// 		}
// 	}
// }

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		swap := false
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
