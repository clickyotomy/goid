package goid

import (
	"testing"
)

const MaxTestRuns = 1024

type compare struct {
	exp int64
	got int64
}

func TestGet(t *testing.T) {
	// Can't compare results.
	if !fastSupport() {
		return
	}

	ch := make(chan compare, MaxTestRuns)

	for i := 0; i < MaxTestRuns; i++ {
		go func(out chan compare) {
			out <- compare{
				exp: goidSlow(),
				got: goidFast(),
			}
		}(ch)
	}

	for i := 0; i < MaxTestRuns; i++ {
		check := <-ch

		if check.exp != check.got {
			t.Fatalf("error: expected: %d, but got: %d", check.exp, check.got)
		}
	}
}

func BenchmarkFast(b *testing.B) {
	if !fastSupport() {
		return
	}

	for i := 0; i < b.N; i++ {
		goidFast()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goidSlow()
	}
}
