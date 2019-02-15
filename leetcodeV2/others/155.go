package others

import "math"

type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	v := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if v == this.min {
		this.min = math.MaxInt32
		for _, ele := range this.data {
			if ele < this.min {
				this.min = ele
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.min
}
