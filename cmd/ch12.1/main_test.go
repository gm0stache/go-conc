package main

import (
	"sync/atomic"
	"testing"
)

func BenchmarkNormalVarUpdate(b *testing.B) {
	num := int64(0)
	for i := 0; i < 1000; i++ {
		num += 1
	}
}

func BenchmarkAtomicVarUpdate(b *testing.B) {
	num := int64(0)
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&num, 1)
	}
}
