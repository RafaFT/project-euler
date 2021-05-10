package problem10

import "testing"

func BenchmarkProblem10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem10()
	}
}

func TestProblem10(t *testing.T) {
	want := 142913828922
	got := problem10()

	if want != got {
		t.Errorf("problem 10: want=%v got=%v", want, got)
	}
}
