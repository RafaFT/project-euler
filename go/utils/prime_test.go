package utils

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input int
		want  bool
		err   error
	}{
		{-1, false, ErrNegativeInteger},
		{0, false, nil},
		{1, false, nil},
		{2, true, nil},
		{3, true, nil},
		{4, false, nil},
		{5, true, nil},
		{7, true, nil},
		{9, false, nil},
		{41, true, nil},
		{97, true, nil},
		{99, false, nil},
		{123456789987654321, false, nil},
	}

	for _, test := range tests {
		got, gotErr := IsPrime(test.input)

		if got != test.want || reflect.TypeOf(gotErr) != reflect.TypeOf(gotErr) {
			t.Errorf("want=%v got=%v err=%v gotErr=%v", test.want, got, test.err, gotErr)
		}
	}
}

func TestNextPrime(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{-1, 2},
		{0, 2},
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 5},
		{5, 7},
		{7, 11},
		{41, 43},
		{97, 101},
		{99, 101},
	}

	for _, test := range tests {
		got := NextPrime(test.input)

		if got != test.want {
			t.Errorf("want=%v got=%v", test.want, got)
		}
	}
}

func TestPlusOne(t *testing.T) {
	if got := findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}); !reflect.DeepEqual(got, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}) {
		t.Errorf("want=%v got=%v", []int{1, 2, 4, 7, 5, 3, 6, 8, 9}, got)
	}
}

func TestPlusOne1(t *testing.T) {
	if got := findDiagonalOrder([][]int{{1, 2}, {3, 4}}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("want=%v got=%v", []int{1, 2, 3, 4}, got)
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}

func TestCoordinatesClosure(t *testing.T) {
	f := coordinatesClosure(3)

	tests := []struct {
		row int
		col int
	}{
		{
			row: 0,
			col: 0,
		},
		{
			row: 0,
			col: 1,
		},
		{
			row: 1,
			col: 0,
		},
		{
			row: 2,
			col: 0,
		},
		{
			row: 1,
			col: 1,
		},
		{
			row: 0,
			col: 2,
		},
		{
			row: 1,
			col: 2,
		},
		{
			row: 2,
			col: 1,
		},
		{
			row: 2,
			col: 2,
		},
	}

	for _, test := range tests {
		row, col := f()

		if row != test.row || col != test.col {
			t.Errorf("(%d, %d) != (%d, %d)", row, col, test.row, test.col)
		}
	}
}
