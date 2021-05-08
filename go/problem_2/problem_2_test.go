package problem2

import "testing"

func BenchmarkProblem2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem()
	}
}

func TestProblem2(t *testing.T) {
	want := 4613732
	got := problem()

	if want != got {
		t.Errorf("problem 2: want=%v got=%v", want, got)
	}
}
