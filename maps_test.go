package main

import "testing"

func provideNodesMap() (ns []NodeWithMaps) {
	ns = make([]NodeWithMaps, 10)
	for i := 0; i < 10; i++ {
		children := make(map[string]NodeWithMaps, 5)
		for c := range children {
			children[c] = NodeWithMaps{
				id: i,
				m1: provideTestMap(),
				m2: provideTestMap(),
			}
		}

		ns[i] = NodeWithMaps{
			id:       i,
			m1: provideTestMap(),
			m2: provideTestMap(),
		}
	}

	return
}

func provideNodesSlice() (ns []NodeWithSlices) {
	ns = make([]NodeWithSlices, 10)
	for i := 0; i < 10; i++ {
		children := make(map[string]NodeWithSlices, 5)
		for c := range children {
			children[c] = NodeWithSlices{
				id: i,
				s1: provideTestSlice(),
				s2: provideTestSlice(),
			}
		}

		ns[i] = NodeWithSlices{
			id:       i,
			s1: provideTestSlice(),
			s2: provideTestSlice(),
		}
	}

	return
}

func provideTestMap() map[string]Foo {
	return map[string]Foo{
		"a": {
			a: 1,
			b: "a",
			c: true,
		},
		"b": {
			a: 2,
			b: "b",
			c: false,
		},
		"c": {
			a: 3,
			b: "c",
			c: true,
		},
	}
}

func provideTestSlice() []Foo {
	return []Foo{
		{
			a: 1,
			b: "a",
			c: true,
		},
		{
			a: 2,
			b: "b",
			c: false,
		},
		{
			a: 3,
			b: "c",
			c: true,
		},
	}
}

func BenchmarkLoopDeepCopySlice(b *testing.B) {
	ns := provideNodesSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopDeepCopySlices(ns)
	}
}

func BenchmarkLoopDeepCopyMap(b *testing.B) {
	ns := provideNodesMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopDeepCopyMap(ns)
	}
}