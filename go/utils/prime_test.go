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
		{41, true, nil},
		{97, true, nil},
		{99, false, nil},
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
