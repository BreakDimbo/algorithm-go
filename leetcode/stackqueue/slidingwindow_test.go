package stackqueue

import (
	"fmt"
	"testing"
)

func Test_MaxSlidingWindow(t *testing.T) {
	fmt.Println(MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
