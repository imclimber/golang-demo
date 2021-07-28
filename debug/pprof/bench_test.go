package main

import "testing"

/*
go test -bench="Fib$" -cpuprofile=cpu.pprof .
go tool pprof -text cpu.pprof
*/

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
