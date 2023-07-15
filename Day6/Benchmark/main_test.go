// in main_test.go
package main

import (
	"fmt"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {  
		fmt.Sprintf("hello, world")   
		}
	}
