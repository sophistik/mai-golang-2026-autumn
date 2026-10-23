package benchdemo

import (
	"sync/atomic"
	"testing"
)

var counter int64

func BenchmarkSumParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&counter, int64(sumTo(10)))
		}
	})
}
