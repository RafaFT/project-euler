package problem1

import (
	"testing"
)

func BenchmarkProblem1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem()
	}
}

func TestProblem1(t *testing.T) {
	want := 233168
	got := problem()

	if got != want {
		t.Errorf("problem 1: want=%v got=%v", want, got)
	}
}
