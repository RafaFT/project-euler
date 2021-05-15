package problem16

import "testing"

func BenchmarkProblem16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem16()
	}
}

func TestProblem16(t *testing.T) {
	want := 1366
	got := problem16()

	if want != got {
		t.Errorf("problem 16: want=%v got=%v", want, got)
	}
}
