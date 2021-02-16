package bc

import (
	"testing"
)

func TestCountString(t *testing.T) {
	b := []byte("{test},{test},{test},{test}")
	c := CountString(b, "},{")

	if c != 3 {
		t.Errorf("want: 3, got %v", c)
	}
}

func BenchmarkCountString(b *testing.B) {
	s := []byte("{test},{test},{test},{test}")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountString(s, "},{")
	}
}