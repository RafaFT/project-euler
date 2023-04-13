package problem18

import "testing"

func BenchmarkProblem18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem18()
	}
}

func TestProblem18(t *testing.T) {
	want := -1
	got := problem18()

	if want != got {
		t.Errorf("problem 18: want=%v got=%v", want, got)
	}
}
