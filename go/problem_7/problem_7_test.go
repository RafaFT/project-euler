package problem7

import "testing"

func BenchmarkProblem7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem7()
	}
}

func TestProblem7(t *testing.T) {
	want := 104743
	got := problem7()

	if want != got {
		t.Errorf("problem 7: want=%v got=%v", want, got)
	}
}
