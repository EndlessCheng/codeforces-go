package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

type (
	input1365 struct{ n int }
	guess1365 struct{ ans []int64 }
	req1365   struct{ q []int }
	resp1365  struct{ or int64 }
)

// github.com/EndlessCheng/codeforces-go
func run1365(in input1365, Q func(req1365) resp1365) (gs guess1365) {
	n := in.n
	mask6 := make([]uint, n)
	qs := [13][]int{}
	x := uint(0)
	for i := range mask6 {
		for x++; bits.OnesCount(x) != 6; x++ {
		}
		mask6[i] = x
		for j := range qs {
			if x>>j&1 == 0 {
				qs[j] = append(qs[j], i+1)
			}
		}
	}
	or := [13]int64{}
	for i, q := range qs {
		if len(q) > 0 {
			or[i] = Q(req1365{q}).or
		}
	}
	a := make([]int64, n)
	for i, v := range mask6 {
		for j, o := range or {
			if v>>j&1 > 0 {
				a[i] |= o
			}
		}
	}
	gs.ans = a
	return
}

func ioq1365() {
	in := bufio.NewReader(os.Stdin)
	Q := func(req req1365) (resp resp1365) {
		Print("? ", len(req.q))
		for _, v := range req.q {
			Print(" ", v)
		}
		Println()
		Fscan(in, &resp.or)
		return
	}
	d := input1365{}
	Fscan(in, &d.n)
	gs := run1365(d, Q)
	Print("!")
	for _, v := range gs.ans {
		Print(" ", v)
	}
	Println()
}

//func main() { ioq1365() }
