package problem12

import "testing"

func BenchmarkProblem12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem12()
	}
}

func BenchmarkNumDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDivisors(2147483646)
	}
}

func TestProblem12(t *testing.T) {
	want := 76576500
	got := problem12()

	if want != got {
		t.Errorf("problem 12: want=%v got=%v", want, got)
	}
}

func TestNumDivisors(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{10, 4},
		{12, 6},
		{25, 3},
		{28, 6},
	}

	for i, test := range tests {
		got := numDivisors(test.input)

		if test.want != got {
			t.Errorf("%d: input=%d want=%d got=%d", i, test.input, test.want, got)
		}
	}
}
