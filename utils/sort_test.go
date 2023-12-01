package sort

import (
	"reflect"
	"testing"
)

func Test_Max(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty list", args{nil}, 0},
		{"non-empty list", args{[]int{1, 2, 3}}, 3},
		{"negative values in list", args{[]int{1, 2, 3, -4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.list); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Sum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty list", args{nil}, 0},
		{"non-empty list", args{[]int{1, 2, 3}}, 6},
		{"negative values in list", args{[]int{1, 2, 3, -4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.array); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"empty", args{nil}, nil},
		{"length 1", args{[]int{1}}, []int{1}},
		{"non-empty", args{[]int{5, 1, 3}}, []int{1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		la []int
		lb []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"both lists are nil", args{nil, nil}, nil},
		{"left list is nil", args{nil, []int{1, 3, 5}}, []int{1, 3, 5}},
		{"right list is nil", args{[]int{2, 4, 6}, nil}, []int{2, 4, 6}},
		{"both non-empty lists", args{[]int{1, 3}, []int{2, 4}}, []int{1, 2, 3, 4}},
		{"smaller-sized right list", args{[]int{1, 3, 5, 7}, []int{2, 4}}, []int{1, 2, 3, 4, 5, 7}},
		{"smaller-sized left list", args{[]int{1, 3}, []int{2, 4, 6, 8}}, []int{1, 2, 3, 4, 6, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.la, tt.args.lb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
