package main

import (
	"fmt"
	"testing"
)

func BenchmarkHelloWorldParallel(b *testing.B) {
	b.SetParallelism(4) // 设置并发度
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Sprintf("hello world")
		}
	})
}
