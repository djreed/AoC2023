package main

import (
	"testing"
)

func Test_FilterDigits(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"bookend digits", args{"1abc2"}, []int{1, 2}},
		{"two digits", args{"pqr3stu8vwx"}, []int{3, 8}},
		{"multiple digits", args{"a1b2c3d4e5f"}, []int{1, 2, 3, 4, 5}},
		{"single digit", args{"treb7uchet"}, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorStr := "FilterDigits() = %v, want %v"
			got := FilterDigits(tt.args.str)
			if len(got) != len(tt.want) {
				t.Errorf(errorStr, got, tt.want)
			}

			for i, v := range got {
				if tt.want[i] != v {
					t.Errorf("FilterDigits() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func TestFirstLastDigits(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"one digit", args{[]int{7}}, 7, 7},
		{"two digits", args{[]int{3, 8}}, 3, 8},
		{"multiple digits", args{[]int{1, 2, 3, 4, 5}}, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FirstLastDigits(tt.args.list)
			if got != tt.want {
				t.Errorf("FirstLastDigits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FirstLastDigits() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCombineDigits(t *testing.T) {
	type args struct {
		d1, d2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"two digits", args{1, 5}, 15},
		{"prefix zero", args{0, 1}, 1},
		{"suffix zero", args{9, 0}, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombineDigits(tt.args.d1, tt.args.d2)
			if got != tt.want {
				t.Errorf("FirstLastDigits() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertWordsToDigits(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"simple 1", args{"two1nine"}, []int{2, 1, 9}},
		{"simple 2", args{"abcone2threexyz"}, []int{1, 2, 3}},
		{"repeated 1", args{"oneoneone"}, []int{1, 1, 1}},
		{"overlapped words 1", args{"eightwothree"}, []int{8, 2, 3}},
		{"overlapped words 2", args{"xtwone3four"}, []int{2, 1, 3, 4}},
		{"overlapped words 3", args{"zoneight234"}, []int{1, 8, 2, 3, 4}},
		{"complex 1", args{"8sixcbqlfmcq14vnlmsixlhzrq"}, []int{8, 6, 1, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorStr := "Filtered ConvertWordsToDigits() = %v, want %v"
			got := ConvertWordsToDigits(tt.args.str)
			filtered := FilterDigits(got)
			if len(filtered) != len(tt.want) {
				t.Errorf(errorStr, filtered, tt.want)
			}

			for i, v := range filtered {
				if tt.want[i] != v {
					t.Errorf(errorStr, got, tt.want)
				}
			}
		})
	}
}
