package main

type Foo struct {
	a int64
	b string
	c bool
}

type NodeWithSlices struct {
	id int
	s1 []Foo
	s2 []Foo
}

func (n NodeWithSlices) DeepCopy() NodeWithSlices {
	sl1 := len(n.s1)
	scp1 := make([]Foo, sl1, sl1)
	copy(scp1, n.s1)

	sl2 := len(n.s2)
	scp2 := make([]Foo, sl2, sl2)
	copy(scp2, n.s2)

	return NodeWithSlices{
		id: n.id,
		s1: scp1,
		s2: scp2,
	}
}

type NodeWithMaps struct {
	id int
	m1 map[string]Foo
	m2 map[string]Foo
}

func (n NodeWithMaps) DeepCopy() NodeWithMaps {
	m1 := make(map[string]Foo)
	for k, v := range n.m1 {
		m1[k] = v
	}
	m2 := make(map[string]Foo)
	for k, v := range n.m2 {
		m2[k] = v
	}

	return NodeWithMaps{
		id: n.id,
		m1: m1,
		m2: m2,
	}
}

func LoopDeepCopySlices(s []NodeWithSlices) []NodeWithSlices {
	cp := make([]NodeWithSlices, len(s))
	for i := range s {
		cp[i] = s[i].DeepCopy()
	}
	return cp
}

func LoopDeepCopyMap(s []NodeWithMaps) []NodeWithMaps {
	cp := make([]NodeWithMaps, len(s))
	for i := range s {
		cp[i] = s[i].DeepCopy()
	}
	return cp
}
