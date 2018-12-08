package search

import (
	"testing"
)

func Test_bsearch(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearch(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchR(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchR(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bsearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchFirstEqual(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
		{"find 14 of index 8", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 14, 14}, 14}, 8},
		{"find first 6 of index 4", args{[]int{1, 2, 3, 4, 6, 6, 8}, 6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchFirstEqual(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bSearchFirstEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchLastEqual(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
		{"find 14 of index 8", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 14, 14}, 14}, 10},
		{"find last 6 of index 5", args{[]int{1, 2, 3, 6, 6, 6, 8}, 6}, 5},
		{"find last 6 of index 5", args{[]int{1, 2, 3, 6, 6, 6, 6}, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchLastEqual(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bSearchLastEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchFirstGreaterOrEqual(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
		{"find first greater and equal than 6 of index 3", args{[]int{1, 2, 3, 6, 6, 6, 8}, 6}, 3},
		{"find first greater and equal than 6 of index 5", args{[]int{1, 2, 3, 4, 6, 6, 6}, 6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchFirstGreaterOrEqual(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bSearchFirstGreaterOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bSearchLastLessOrEqual(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"find 5 of index 4", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 5}, 4},
		{"find 15 of index 9", args{[]int{1, 2, 3, 4, 5, 6, 7, 11, 14, 15, 17}, 15}, 9},
		{"find last less and equal than 6 of index 5", args{[]int{1, 2, 3, 6, 6, 6, 8}, 6}, 5},
		{"find first less and equal than 6 of index 6", args{[]int{1, 2, 3, 4, 6, 6, 6}, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bSearchLastLessOrEqual(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("bSearchLastLessOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
