package heap

import (
	"fmt"
	"math"
)

type Node interface {
}

type Heap[T any] struct {
	nodes   []T
	compare func(target T, with T) bool
}

func New[T Node](compare func(target, with T) bool) Heap[T] {
	return Heap[T]{
		[]T{},
		compare,
	}
}

func (h *Heap[Node]) Push(node Node) {
	h.nodes = append(h.nodes, node)

	i := len(h.nodes) - 1

	for i > 0 && h.compare(node, h.parent(i)) {
		h.nodes[i] = h.parent(i)
		i = (i - 1) / 2
	}

	h.nodes[i] = node
}

func (h *Heap[Node]) Pop() (*Node, error) {
	if len(h.nodes) == 0 {
		return nil, fmt.Errorf("heap is empty")
	}

	node := h.nodes[0]
	last := h.nodes[len(h.nodes)-1]

	h.nodes = h.nodes[:len(h.nodes)-1] // 마지막 노드를 삭제

	if len(h.nodes) == 0 {
		return &node, nil
	}

	i := 0

	for h.leftIdx(i) < len(h.nodes) { // 왼쪽 자식이 있는 동안
		if h.rightIdx(i) < len(h.nodes) { // 오른쪽 자식이 있을 때
			if h.compare(h.rightChild(i), h.leftChild(i)) { // 왼쪽보다 오른쪽 자식이 더 크면
				if h.compare(h.rightChild(i), last) {
					h.swap(h.rightIdx(i), i)
					i = h.rightIdx(i)
				} else {
					break
				}
			} else { // 왼쪽 자식이 더 크면
				if h.compare(h.rightChild(i), last) {
					h.swap(h.leftIdx(i), i)
					i = h.leftIdx(i)
				} else {
					break
				}
			}
		} else { // 오른쪽 자식이 없을 때 왼쪽만 검사
			if h.compare(h.leftChild(i), last) {
				h.swap(h.leftIdx(i), i)
				i = h.leftIdx(i)
			} else {
				break
			}
		}
	}

	h.nodes[i] = last

	return &node, nil
}

func (h *Heap[Node]) PrintTree() {
	for i := 0; i < h.Height()+1; i++ {
		square := int(math.Pow(2, float64(i)))
		for j := square - 1; j < square*2-1; j++ {
			if j < len(h.nodes) {
				fmt.Printf("%v ", h.nodes[j])
			} else {
				fmt.Printf("%s ", "{ }")
			}
		}
		fmt.Println()
	}
}

func (h *Heap[T]) Height() int {
	return int(math.Log2(float64(len(h.nodes))))
}

func (h *Heap[Node]) parent(idx int) Node {
	return h.nodes[(idx-1)/2]
}

func (h *Heap[Node]) leftChild(idx int) Node {
	return h.nodes[idx*2+1]
}

func (h *Heap[Node]) rightChild(idx int) Node {
	return h.nodes[idx*2+2]
}

func (h *Heap[Node]) leftIdx(idx int) int {
	return idx*2 + 1
}

func (h *Heap[Node]) rightIdx(idx int) int {
	return idx*2 + 2
}

func (h *Heap[Node]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}
