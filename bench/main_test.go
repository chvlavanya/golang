/*
go test -bench=.
goos: linux
goarch: amd64
pkg: example.com/alta3/benchmark/bench
cpu: Intel(R) Xeon(R) CPU E5-2680 v3 @ 2.50GHz
BenchmarkCalculate-2    1000000000               0.4068 ns/op
PASS
ok      example.com/alta3/benchmark/bench       0.459s

*/


package main

import (
    "testing"
)

func BenchmarkCalculate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Calculate(2)
    }
}

