package problem8

import "testing"

func BenchmarkProblem8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem8()
	}
}

func TestProblem8(t *testing.T) {
	want := 23514624000
	got := problem8()

	if want != got {
		t.Errorf("problem 8: want=%v got=%v", want, got)
	}
}
