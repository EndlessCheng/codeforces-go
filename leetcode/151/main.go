package main

import (
	"container/heap"
	. "fmt"
	"sort"
	"strconv"
	"strings"
)

var _ = Print

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func ifElseI(cond bool, a, b int) int {
	if cond {
		return a
	}
	return b
}
func ifElseS(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

const mod int = 1e9 + 7

func invalidTransactions(transactions []string) (ans []string) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	n := len(transactions)
	bad := make([]bool, n)
	for i, ti := range transactions {
		splitsI := strings.Split(ti, ",")
		amount, _ := strconv.Atoi(splitsI[2])
		if amount > 1000 {
			bad[i] = true
		}
		timeI, _ := strconv.Atoi(splitsI[1])
		for j, tj := range transactions {
			splitsJ := strings.Split(tj, ",")
			if splitsJ[0] == splitsI[0] && splitsJ[3] != splitsI[3] {
				timeJ, _ := strconv.Atoi(splitsJ[1])
				if abs(timeI-timeJ) <= 60 {
					bad[i] = true
					bad[j] = true
				}
			}
		}
	}
	for i, b := range bad {
		if b {
			ans = append(ans, transactions[i])
		}
	}
	return
}

func numSmallerByFrequency(queries []string, words []string) (ans []int) {
	minW := make([]int, len(words))
	for i, w := range words {
		minC := int('z')
		for _, s := range w {
			minC = min(minC, int(s))
		}
		cnt := 0
		for _, s := range w {
			if int(s) == minC {
				cnt++
			}
		}
		minW[i] = cnt
	}

	ans = make([]int, len(queries))
	for i, q := range queries {
		minC := int('z')
		for _, s := range q {
			minC = min(minC, int(s))
		}
		cnt := 0
		for _, s := range q {
			if int(s) == minC {
				cnt++
			}
		}
		for _, cntw := range minW {
			if cnt < cntw {
				ans[i]++
			}
		}
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSolve(a []int) []bool {
	n := len(a)
	sum := make([]int, n)
	remove := make([]bool, n)
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + a[i]
		if sum[i] == 0 {
			for k := 1; k <= i; k++ {
				remove[k] = true
				sum[k] = 0
			}
			continue
		}
		for j := 1; j < i; j++ {
			if sum[i] == sum[j] {
				for k := j + 1; k <= i; k++ {
					remove[k] = true
					sum[k] = sum[i]
				}
			}
		}
	}
	return remove
}

func removeZeroSumSublists(o *ListNode) *ListNode {
	a := []int{0}
	for ; o != nil; o = o.Next {
		a = append(a, o.Val)
	}
	n := len(a)
	remove := removeZeroSolve(a)
	o = &ListNode{}
	head := o
	isFirst := true
	for i := 1; i < n; i++ {
		if !remove[i] {
			if isFirst {
				o.Val = a[i]
				isFirst = false
			} else {
				newNode := &ListNode{Val: a[i]}
				o.Next = newNode
				o = o.Next
			}
		}
	}
	if head.Val == 0 {
		return nil
	}

	return head
}

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
	//fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,100,beijing"}))
	//fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","alice,50,1200,mtv"}))
	//fmt.Println(invalidTransactions([]string{"alice,20,800,mtv","bob,50,1200,mtv"}))

	//fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
	//fmt.Println(removeZeroSolve([]int{0, 1, 2, -3, 3, 1}))
	//fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, 4}))
	//fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, -2}))
	//fmt.Println(removeZeroSolve([]int{0, 0, 1}))
	//fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, 3}))
	obj := Constructor(2)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	//obj.Push(4)
	// fmt.Println(obj.Pop())
	Println(obj.PopAtStack(0))
	//obj.Push(2)
	Println(obj.PopAtStack(0))
	obj.Push(1)
	Println(obj.PopAtStack(0))

	Println(obj.stackLen[:5])
}
