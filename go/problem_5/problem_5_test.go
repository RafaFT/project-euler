package problem5

import "testing"

func BenchmarkProblem5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem5()
	}
}

func TestProblem5(t *testing.T) {
	want := 232_792_560
	got := problem5()

	if want != got {
		t.Errorf("problem 5: want=%v got=%v", want, got)
	}
}
