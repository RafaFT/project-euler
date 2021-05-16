package problem20

import "testing"

func BenchmarkProblem20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem20()
	}
}

func TestProblem20(t *testing.T) {
	want := 648
	got := problem20()

	if want != got {
		t.Errorf("problem 20: want=%v got=%v", want, got)
	}
}
