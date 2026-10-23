package benchdemo

import "testing"

var sink int // защита от оптимизатора

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = sumTo(100)
	}
}
