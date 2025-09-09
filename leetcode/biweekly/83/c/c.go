package main

import (
	"container/heap"
	"sort"
)

type NumberContainers struct {
	indexToNumber   map[int]int
	numberToIndices map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*hp{}}
}

func (n NumberContainers) Change(index, number int) {
	// 添加新数据
	n.indexToNumber[index] = number
	if _, ok := n.numberToIndices[number]; !ok {
		n.numberToIndices[number] = &hp{}
	}
	heap.Push(n.numberToIndices[number], index)
}

func (n NumberContainers) Find(number int) int {
	indices, ok := n.numberToIndices[number]
	if !ok {
		return -1
	}
	for indices.Len() > 0 && n.indexToNumber[indices.IntSlice[0]] != number {
		heap.Pop(indices) // 堆顶货不对板，说明是旧数据，删除
	}
	if indices.Len() == 0 {
		return -1
	}
	return indices.IntSlice[0]
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
