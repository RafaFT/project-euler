package problem15

import "testing"

func BenchmarkProblem15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem15()
	}
}

func TestProblem15(t *testing.T) {
	want := 137846528820
	got := problem15()

	if want != got {
		t.Errorf("problem 15: want=%v got=%v", want, got)
	}
}
