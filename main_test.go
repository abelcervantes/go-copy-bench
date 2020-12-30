package main

import (
	"reflect"
	"runtime"
	"testing"
)

func provideNodes() (ns []Node) {
	ns = make([]Node, 10)
	for i := 0; i < 10; i++ {
		children := make([]Node, 5)
		for c := range children {
			children[c] = Node{
				id:    i,
				slice: []int{1,2,3},
			}
		}

		ns[i] = Node{
			id:       i,
			slice:    []int{1, 2, 3},
			children: children,
		}
	}

	//for _, n := range ns { fmt.Println(n) }

	return
}

type valueCopy func([]Node) []Node

func TestValueCopy(t *testing.T) {
	fs := []valueCopy{
		LoopCopy,
		LoopCopyPreLen,
		LoopCopyPreCap,
		LoopCopyPreLenCap,
		Copy,
		CopyPreLenCap,
		Append,
		LoopDeepCopy,
		LoopDeepCopyPointerReceiver,
	}

	ns := provideNodes()
	for _, f := range fs {
		cp := f(ns)

		fName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		if &ns == &cp {
			t.Errorf("%s output -> same reference: %p %p", fName, &ns, &cp)
		}

		if len(ns) != len(cp) {
			t.Errorf("%s output -> dif size: %d %d", fName, len(ns), len(cp))
		}

		for i  := range cp {
			if cp[i].id != ns[i].id {
				t.Errorf("%s output -> dif element(id) %d %d", fName, ns[i].id, cp[i].id)
			}

			if &cp[i].slice == &ns[i].slice {
				t.Errorf("%s output -> same ref(slice) %p %p", fName, &ns[i].slice, &cp[i].slice)
			}

			if &cp[i].children == &ns[i].children {
				t.Errorf("%s output -> same ref(children) %p %p", fName, &ns[i].children, &cp[i].children)
			}
		}
	}
}

func BenchmarkLoop(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopy(ns)
	}
}

func BenchmarkLoopPreLen(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreLen(ns)
	}
}

func BenchmarkLoopPreCap(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreCap(ns)
	}
}

func BenchmarkLoopPreLenCap(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreLenCap(ns)
	}
}

func BenchmarkCopy(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(ns)
	}
}

func BenchmarkCopyPreLenCap(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyPreLenCap(ns)
	}
}

func BenchmarkAppend(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Append(ns)
	}
}

func BenchmarkLoopDeepCopy(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopDeepCopy(ns)
	}
}

func BenchmarkDeepCopyPointerReceiver(b *testing.B) {
	ns := provideNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopDeepCopyPointerReceiver(ns)
	}
}



func providePointerNodes() (ns []*Node) {
	ns = make([]*Node, 10)
	for i := 0; i < 10; i++ {
		children := make([]Node, 5)
		for c := range children {
			children[c] = Node{
				id:    i,
				slice: []int{1,2,3},
			}
		}

		ns[i] = &Node{
			id:       i,
			slice:    []int{1, 2, 3},
			children: children,
		}
	}

	//for _, n := range ns { fmt.Println(n) }

	return
}

type pointerCopy func([]*Node) []*Node

func TestPointerCopy(t *testing.T) {
	fs := []pointerCopy{
		LoopCopyPointer,
		LoopCopyPreLenPointer,
		LoopCopyPreCapPointer,
		LoopCopyPreLenCapPointer,
		LoopDeepCopyPointer,
	}

	ns := providePointerNodes()
	for _, f := range fs {
		cp := f(ns)

		fName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		if &ns == &cp {
			t.Errorf("%s output -> same reference: %p %p", fName, &ns, &cp)
		}

		if len(ns) != len(cp) {
			t.Errorf("%s output -> dif size: %d %d", fName, len(ns), len(cp))
		}

		for i  := range cp {
			if &cp[i] == &ns[i] {
				t.Errorf("%s output -> same ref(node) %p %p", fName, &ns[i], &cp[i])
			}

			if cp[i].id != ns[i].id {
				t.Errorf("%s output -> dif element(id) %d %d", fName, ns[i].id, cp[i].id)
			}

			if &cp[i].slice == &ns[i].slice {
				t.Errorf("%s output -> same ref(slice) %p %p", fName, &ns[i].slice, &cp[i].slice)
			}
		}
	}
}

func BenchmarkLoopPointer(b *testing.B) {
	ns := providePointerNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPointer(ns)
	}
}

func BenchmarkLoopPrePointer(b *testing.B) {
	ns := providePointerNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreLenPointer(ns)
	}
}

func BenchmarkLoopPreCapPointer(b *testing.B) {
	ns := providePointerNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreCapPointer(ns)
	}
}

func BenchmarkLoopPreLenCapPointer(b *testing.B) {
	ns := providePointerNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopCopyPreLenCapPointer(ns)
	}
}

func BenchmarkLoopDeepCopyPointer(b *testing.B) {
	ns := providePointerNodes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopDeepCopyPointer(ns)
	}
}
