package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1313C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type pair struct{ v, i int }
	posL := make([]int, n)
	stack := []pair{{0, -1}}
	for i, v := range a {
		for {
			if top := stack[len(stack)-1]; top.v <= v {
				posL[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}
	posR := make([]int, n)
	stack = []pair{{0, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		for {
			if top := stack[len(stack)-1]; top.v <= v {
				posR[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	dpl := make([]int64, n)
	dpl[0] = int64(a[0])
	for i := 1; i < n; i++ {
		l := posL[i]
		dpl[i] = int64(a[i]) * int64(i-l)
		if l >= 0 {
			dpl[i] += dpl[l]
		}
	}
	dpr := make([]int64, n)
	dpr[n-1] = int64(a[n-1])
	for i := n - 2; i >= 0; i-- {
		r := posR[i]
		dpr[i] = int64(a[i]) * int64(r-i)
		if r < n {
			dpr[i] += dpr[r]
		}
	}

	maxAns, maxPos := int64(0), 0
	for i, v := range a {
		if ans := dpl[i] + dpr[i] - int64(v); ans > maxAns {
			maxAns, maxPos = ans, i
		}
	}

	ans := make([]interface{}, n)
	curMax := a[maxPos]
	for j := maxPos; j >= 0; j-- {
		curMax = min(curMax, a[j])
		ans[j] = curMax
	}
	curMax = a[maxPos]
	for j := maxPos + 1; j < n; j++ {
		curMax = min(curMax, a[j])
		ans[j] = curMax
	}
	Fprint(out, ans...)
}

//func main() { CF1313C2(os.Stdin, os.Stdout) }
