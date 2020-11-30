package main

import (
	"fmt"
	"strconv"
	"testing"
)

func provideNodes() []Node {
	ns := make([]Node, 10)
	for i := 0; i < 10; i++ {
		ns[i] = Node{
			id:    strconv.FormatInt(int64(i), 10),
			value: i * 2,
		}
	}
	fmt.Printf("ns: %+v", ns)
}

func BenchmarkLoop(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopy(ns)
	}
}

func BenchmarkLoopPre(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPre(ns)
	}
}

func BenchmarkLoopPreCap(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreCap(ns)
	}
}

func BenchmarkCopy(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(ns)
	}
}

func BenchmarkAppend(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Append(ns)
	}
}



func BenchmarkLoopPointer(b *testing.B) {
}

func BenchmarkLoopPrePointer(b *testing.B) {
}

func BenchmarkLoopPreCapPointer(b *testing.B) {	
}

func BenchmarkCopyPointer(b *testing.B) {	
}

func BenchmarkAppendPointer(b *testing.B) {
}