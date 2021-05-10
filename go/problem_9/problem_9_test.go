package problem9

import "testing"

func BenchmarkProblem9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem9()
	}
}

func TestProblem9(t *testing.T) {
	want := 31875000
	got := problem9()

	if want != got {
		t.Errorf("problem 9: want=%v got=%v", want, got)
	}
}
