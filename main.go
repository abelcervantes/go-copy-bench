package main

import "fmt"

type Node struct {
	id       int
	slice    []int
	children []Node
}

func (n Node) DeepCopy() Node {
	sl := len(n.slice)
	scp := make([]int, sl, sl)
	copy(scp, n.slice)
	n.slice = scp

	nl := len(n.children)
	ccp := make([]Node, nl, nl)
	for i := range n.children {
		ccp[i] = n.children[i].DeepCopy()
	}
	n.children = ccp

	return n
}

func (n *Node) DeepCopyPointerReceiver() Node {
	sl := len(n.slice)
	scp := make([]int, sl, sl)
	copy(scp, n.slice)

	nl := len(n.children)
	ccp := make([]Node, nl, nl)
	for i := range n.children {
		ccp[i] = n.children[i].DeepCopy()
	}

	return Node{
		id:       n.id,
		slice:    scp,
		children: ccp,
	}
}

func (n *Node) DeepCopyPointer() *Node {
	sl := len(n.slice)
	scp := make([]int, sl, sl)
	copy(scp, n.slice)

	nl := len(n.children)
	ccp := make([]Node, nl, nl)
	for i := range n.children {
		ccp[i] = n.children[i].DeepCopy()
	}

	return &Node{
		id:       n.id,
		slice:    scp,
		children: ccp,
	}
}

// values
func LoopCopy(s []Node) []Node {
	cp := make([]Node, 0)
	for _, n := range s {
		cp = append(cp, n)
	}
	return cp
}

func LoopCopyPreLen(s []Node) []Node {
	cp := make([]Node, len(s)) // zeroed
	for i, n := range s {
		cp[i] = n
	}
	return cp
}

func LoopCopyPreCap(s []Node) []Node {
	cp := make([]Node, 0, len(s))
	for _, n := range s {
		cp = append(cp, n)
	}
	return cp
}

func LoopCopyPreLenCap(s []Node) []Node {
	l := len(s)
	cp := make([]Node, l, l)
	for i, n := range s {
		cp[i] = n
	}
	return cp
}

func Copy(s []Node) []Node {
	cp := make([]Node, len(s)) // zeroed
	copy(cp, s)
	return cp
}

func CopyPreLenCap(s []Node) []Node {
	l := len(s)
	cp := make([]Node, l, l)
	copy(cp, s)
	return cp
}

func Append(s []Node) []Node {
	return append(make([]Node, 0, len(s)), s...)
}

func LoopDeepCopy(s []Node) []Node {
	l := len(s)
	cp := make([]Node, l , l)
	for i := range s {
		cp[i] = s[i].DeepCopy()
	}
	return cp
}

func LoopDeepCopyPointerReceiver(s []Node) []Node {
	l := len(s)
	cp := make([]Node, l , l)
	for i := range s {
		cp[i] = s[i].DeepCopyPointerReceiver()
	}
	return cp
}


// pointers
func LoopCopyPointer(s []*Node) []*Node {
	cp := make([]*Node, 0)
	for i := range s {
		cp = append(cp, &Node{
			id:       s[i].id,
			slice:    s[i].slice,
			children: s[i].children,
		})
	}
	return cp
}

func LoopCopyPreLenPointer(s []*Node) []*Node {
	cp := make([]*Node, len(s))
	for i := range s {
		cp[i] = &Node{
			id:    s[i].id,
			slice: s[i].slice,
			children: s[i].children,
		}
	}
	return cp
}

func LoopCopyPreCapPointer(s []*Node) []*Node {
	cp := make([]*Node, 0, len(s))
	for i := range s {
		cp = append(cp, &Node{
			id:    s[i].id,
			slice: s[i].slice,
			children: s[i].children,
		})
	}
	return cp
}

func LoopCopyPreLenCapPointer(s []*Node) []*Node {
	l := len(s)
	cp := make([]*Node, l, l)
	for i := range s {
		cp[i] = &Node{
			id:    s[i].id,
			slice: s[i].slice,
			children: s[i].children,
		}
	}
	return cp
}

func LoopDeepCopyPointer(s []*Node) []*Node {
	l := len(s)
	cp := make([]*Node, l , l)
	for i := range s {
		cp[i] = s[i].DeepCopyPointer()
	}
	return cp
}


func main() {
	n := Node{
		id:     1,
		slice:  []int{1, 2, 3},
		children: []Node{
			{
				id:     11,
				slice:  []int{11,12,13},
			},
			{
				id:     22,
				slice:  []int{22,23,24},
			},
		},
	}

	//cp := n
	cp := n.DeepCopy()

	cp.id = 10
	cp.slice[0] = 11
	cp.children[0].id = 111
	cp.children[0].slice[0] = 99

	//cp := &Node{
	//	id:    n.id,
	//	value: n.value,
	//	slice: n.slice,
	//}
	//cp.slice = append(cp.slice, 4)

	fmt.Println(n)
	fmt.Println(cp)
}