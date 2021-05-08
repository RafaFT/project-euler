package problem3

import "testing"

func BenchmarkProblem3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem()
	}
}

func TestProblem3(t *testing.T) {
	want := 6857
	got := problem()

	if want != got {
		t.Errorf("problem 3: want=%v got=%v", want, got)
	}
}
