package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type intHeap struct {
	sort.IntSlice
}

func (h *intHeap) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *intHeap) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

type DinnerPlates struct {
	cap      int
	stack    []int
	cachePos *intHeap
	stackLen [200005]int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		cap:      capacity,
		cachePos: &intHeap{},
	}
}

func (d *DinnerPlates) Push(val int) {
	if d.cachePos.Len() > 0 {
		pos := heap.Pop(d.cachePos).(int)
		d.stack[pos] = val
		d.stackLen[pos/d.cap]++
		return
	}
	d.stackLen[len(d.stack)/d.cap]++
	d.stack = append(d.stack, val)
}

func (d *DinnerPlates) Pop() (val int) {
	if len(d.stack) == d.cachePos.Len() {
		d.cachePos = &intHeap{}
		d.stack = []int{}
		return -1
	}
	i := len(d.stack) - 1
	for ; i >= 0; i-- {
		if d.stack[i] != 0 {
			break
		}
		heap.Pop(d.cachePos)
		d.stack = d.stack[:len(d.stack)-1]
	}
	d.stack, val = d.stack[:len(d.stack)-1], d.stack[len(d.stack)-1]
	d.stackLen[len(d.stack)/d.cap]--
	return
}

func (d *DinnerPlates) PopAtStack(index int) (val int) {
	if d.stackLen[index] == 0 {
		return -1
	}
	pos := d.cap*index + d.stackLen[index] - 1
	heap.Push(d.cachePos, pos)
	val = d.stack[pos]
	d.stack[pos] = 0
	d.stackLen[index]--
	return
}

func main() {
	obj := Constructor(2)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	//obj.Push(4)
	// fmt.Println(obj.Pop())
	fmt.Println(obj.PopAtStack(0))
	//obj.Push(2)
	fmt.Println(obj.PopAtStack(0))
	obj.Push(1)
	fmt.Println(obj.PopAtStack(0))

	fmt.Println(obj.stackLen[:5])
}
