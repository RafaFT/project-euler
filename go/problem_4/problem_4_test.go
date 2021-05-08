package problem4

import "testing"

func BenchmarkProblem4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem4()
	}
}

func TestProblem4(t *testing.T) {
	want := 906609
	got := problem4()

	if want != got {
		t.Errorf("problem 4: want=%v got=%v", want, got)
	}
}
