package gotest
import (
    "testing"
)

func Benchmark_Division(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
    b.StopTimer()
    // doing some job without any performance losses of your function
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}
