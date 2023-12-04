package main

import (
	"reflect"
	"testing"
)

func TestParseCard(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want ScratchCard
	}{
		{"Full Example", args{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, ScratchCard{Id: "Card 1", Winning: []int{41, 48, 83, 86, 17}, Numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}},
		{"No Winners", args{"Card 2: | 1 2 3 4"}, ScratchCard{Id: "Card 2", Winning: nil, Numbers: []int{1, 2, 3, 4}}},
		{"All Winners", args{"Card 3: 1 2 3 4 |"}, ScratchCard{Id: "Card 3", Winning: []int{1, 2, 3, 4}, Numbers: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCard(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchingNumbers(t *testing.T) {
	type args struct {
		card ScratchCard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Full Example", args{ScratchCard{Id: "Card 1", Winning: []int{41, 48, 83, 86, 17}, Numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}}, 4},
		{"No Winners", args{ScratchCard{Id: "Card 2", Winning: nil, Numbers: []int{1, 2, 3, 4}}}, 0},
		{"No Numbers", args{ScratchCard{Id: "Card 3", Winning: []int{1, 2, 3, 4}, Numbers: nil}}, 0},
		{"No Matches", args{ScratchCard{Id: "Card 0", Winning: []int{1, 2, 3}, Numbers: []int{4, 5, 6}}}, 0},
		{"1 Match", args{ScratchCard{Id: "Card 1", Winning: []int{1}, Numbers: []int{1}}}, 1},
		{"2 Matches", args{ScratchCard{Id: "Card 2", Winning: []int{1, 2}, Numbers: []int{1, 2}}}, 2},
		{"3 Matches", args{ScratchCard{Id: "Card 3", Winning: []int{1, 2, 3}, Numbers: []int{1, 2, 3}}}, 3},
		{"4 Matches", args{ScratchCard{Id: "Card 4", Winning: []int{1, 2, 3, 4}, Numbers: []int{1, 2, 3, 4}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchingNumbers(tt.args.card); got != tt.want {
				t.Errorf("MatchingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardValue(t *testing.T) {
	type args struct {
		card ScratchCard
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"No Winners", args{ScratchCard{Id: "Card 2", Winning: nil, Numbers: []int{1, 2, 3, 4}}}, 0},
		{"No Numbers", args{ScratchCard{Id: "Card 3", Winning: []int{1, 2, 3, 4}, Numbers: nil}}, 0},
		{"No Matches", args{ScratchCard{Id: "Card 0", Winning: []int{1, 2, 3}, Numbers: []int{4, 5, 6}}}, 0},
		{"1 Match", args{ScratchCard{Id: "Card 1", Winning: []int{1}, Numbers: []int{1}}}, 1},
		{"2 Matches", args{ScratchCard{Id: "Card 2", Winning: []int{1, 2}, Numbers: []int{1, 2}}}, 2},
		{"3 Matches", args{ScratchCard{Id: "Card 3", Winning: []int{1, 2, 3}, Numbers: []int{1, 2, 3}}}, 4},
		{"4 Matches", args{ScratchCard{Id: "Card 4", Winning: []int{1, 2, 3, 4}, Numbers: []int{1, 2, 3, 4}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CardValue(tt.args.card); got != tt.want {
				t.Errorf("CardValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
