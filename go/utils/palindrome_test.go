package utils

import "testing"

func TestIsIntPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{0, true},
		{9, true},
		{10, false},
		{11, true},
		{12, false},
		{98766789, true},
		{987656789, true},
		{-987656789, true},
		{98765789, false},
		{-98765789, false},
	}

	for i, test := range tests {
		got := IsIntPalindrome(test.input)

		if test.want != got {
			t.Errorf("%d: input=%v want=%v got=%v", i, test.input, test.want, got)
		}
	}
}

func BenchmarkIsIntPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsIntPalindrome(123456789987654321)
	}
}
