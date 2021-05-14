package problem14

import "testing"

func BenchmarkProblem14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem14()
	}
}

func TestProblem14(t *testing.T) {
	want := 837799
	got := problem14()

	if want != got {
		t.Errorf("problem 14: want=%v got=%v", want, got)
	}
}
