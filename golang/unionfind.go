package main

type UnionFind struct {
	root []int
	size []int
}

func New(capacity int) *UnionFind {
	size := make([]int, capacity)
	root := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		root[i] = i
		size[i] = 1
	}
	uf := UnionFind{
		root: root,
		size: size,
	}
	return &uf
}

func (this *UnionFind) Root(p int) int {
	for p != this.root[p] {
		this.root[p] = this.root[this.root[p]]
		p = this.root[p]
	}
	return p
}

func (this *UnionFind) Union(a, b int) {
	ra := this.Root(a)
	rb := this.Root(b)
	if ra == rb {
		return
	}
	if this.size[ra] < this.size[rb] {
		this.root[ra] = this.root[rb]
		this.size[rb] += this.size[ra]
	} else {
		this.root[rb] = this.root[ra]
		this.size[ra] += this.size[rb]
	}
}
