package problem11

import "testing"

func BenchmarkProblem11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem11()
	}
}

func TestProblem11(t *testing.T) {
	want := 1
	got := problem11()

	if want != got {
		t.Errorf("problem 11: want=%v got=%v", want, got)
	}
}
