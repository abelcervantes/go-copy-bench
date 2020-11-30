package main

type Node struct {
	id    string
	value int
}

func (n Node) copy() Node {
	return n
}

func LoopCopy(s []Node) []Node {
	cp := make([]Node, 0)
	for _, n := range s {
		cp = append(cp, n)
	}
	return cp
}

func LoopCopyPre(s []Node) []Node {
	cp := make([]Node, len(s))
	for i, n := range s {
		cp[i] = n
	}
	return cp
}

func LoopCopyPreCap(s []Node) []Node {
	cp := make([]Node, 0, len(s))
	for i, n := range s {
		cp[i] = n
	}
	return cp
}

func Copy(s []Node) []Node {
	cp := make([]Node, len(s))
	copy(cp, s)
	return cp
}

func Append(s []Node) []Node {
	return append(make([]Node, 0, len(s)), s...)
}

func main() {
}

func LoopCopyPointer(s []*Node) []*Node {
	cp := make([]*Node, 0)
	for _, n := range s {
		cp = append(cp, n)
	}
	return cp
}

func LoopCopyPrePointer(s []*Node) []*Node {
	cp := make([]*Node, len(s))
	for i, n := range s {
		cp[i] = n
	}
	return cp
}