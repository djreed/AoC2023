package main

import (
	"reflect"
	"testing"
)

func MS(b, r, g int) MarbleSet {
	return MarbleSet{
		Blue:  b,
		Red:   r,
		Green: g,
	}
}

func TestFilterGameRounds(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []MarbleSet
	}{
		{"Sample 1", args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			1, []MarbleSet{MS(3, 4, 0), MS(6, 1, 2), MS(0, 0, 2)}},
		{"Sample 2", args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"},
			2, []MarbleSet{MS(1, 0, 2), MS(4, 1, 3), MS(1, 0, 1)}},
		{"Sample 3", args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"},
			3, []MarbleSet{MS(6, 20, 8), MS(5, 4, 13), MS(0, 1, 5)}},
		{"Sample 4", args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"},
			4, []MarbleSet{MS(6, 3, 1), MS(0, 6, 3), MS(15, 14, 3)}},
		{"Sample 5", args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"},
			5, []MarbleSet{MS(1, 6, 3), MS(2, 1, 2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterGameRounds(tt.args.str)
			if got != tt.want {
				t.Errorf("FilterGameRounds() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FilterGameRounds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestValidRound(t *testing.T) {
	type args struct {
		restriction MarbleSet
		round       MarbleSet
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid", args{MS(14, 12, 13), MS(5, 6, 7)}, true},
		{"Invalid", args{MS(14, 12, 13), MS(20, 10, 7)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidRound(tt.args.restriction, tt.args.round); got != tt.want {
				t.Errorf("ValidRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
