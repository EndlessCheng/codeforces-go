package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var s []byte
	Fscan(bufio.NewReader(_r), &s)
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n+1)
	h := 0
	for i, ri := range rank {
		if h > 0 {
			h--
		}
		if ri > 0 {
			for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[ri] = h
	}

	type pair struct{ v, i int }
	posL := make([]int, n+1)
	st := []pair{{-1, 0}}
	for i, v := range height {
		for {
			if top := st[len(st)-1]; top.v < v {
				posL[i] = top.i
				break
			}
			st = st[:len(st)-1]
		}
		st = append(st, pair{v, i})
	}
	posR := make([]int, n+1)
	st = []pair{{-1, n}}
	for i := n - 1; i >= 0; i-- {
		v := height[i]
		for {
			if top := st[len(st)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			st = st[:len(st)-1]
		}
		st = append(st, pair{v, i})
	}

	ans := int64(n) * int64(n+1) / 2
	usedSt := []int{0} // 关键：用一个单调栈判断是否已经计算了！
	for i, v := range height {
		for v < usedSt[len(usedSt)-1] {
			usedSt = usedSt[:len(usedSt)-1]
		}
		if v == usedSt[len(usedSt)-1] {
			continue
		}
		usedSt = append(usedSt, v)
		l, r := posL[i], posR[i]
		h, w := int64(v-max(height[l], height[r])), int64(r-l)
		ans += h * w * (w - 1) / 2
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
