package sort

import "testing"

func Test_insertSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		// TODO: Add test cases.
		{"test1", []int{3, 2, 1, 5, 7, 4}, []int{1, 2, 3, 4, 5, 7}},
		{"test2", []int{}, []int{}},
		{"test2", []int{2, 2, 2}, []int{2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertSort(tt.input)
			if !Equal(tt.input, tt.want) {
				t.Errorf("sort is not right, want %v, get %v", tt.want, tt.input)
			}
		})
	}
}
