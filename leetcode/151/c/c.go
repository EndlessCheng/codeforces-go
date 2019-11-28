package main

import "fmt"

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

func main() {
	fmt.Println(removeZeroSolve([]int{0, 1, 2, -3, 3, 1}))
	fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, 4}))
	fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, -2}))
	fmt.Println(removeZeroSolve([]int{0, 0, 1}))
	fmt.Println(removeZeroSolve([]int{0, 1, 2, 3, -3, 3}))
}
