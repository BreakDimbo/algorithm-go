package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHeap(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *Heap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeap(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {
	type fields struct {
		data   []int
		length int
		count  int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				data:   tt.fields.data,
				length: tt.fields.length,
				count:  tt.fields.count,
			}
			h.Insert(tt.args.value)
		})
	}
}

func TestHeap_Pop(t *testing.T) {
	type fields struct {
		data   []int
		length int
		count  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				data:   tt.fields.data,
				length: tt.fields.length,
				count:  tt.fields.count,
			}
			h.Pop()
		})
	}
}

func Test_heapify(t *testing.T) {
	type args struct {
		a      []int
		length int
		i      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapify(tt.args.a, tt.args.length, tt.args.i)
		})
	}
}

func TestHeapSort(t *testing.T) {
	a := []int{0, 2, 3, 1, 4, 5, 12, 21}
	HeapSort(a)
	fmt.Println(a)
}

func Test_buildHeap(t *testing.T) {
	type args struct {
		a []int
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildHeap(tt.args.a, tt.args.n)
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		data []int
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.data, tt.args.i, tt.args.j)
		})
	}
}
