package problem6

import "testing"

func BenchmarkProblem6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem6()
	}
}

func TestProblem6(t *testing.T) {
	want := 25164150
	got := problem6()

	if want != got {
		t.Errorf("problem 6: want=%v got=%v", want, got)
	}
}
