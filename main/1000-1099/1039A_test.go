package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1039/A
// https://codeforces.com/problemset/status/1039/problem/A
func TestCF1039A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 10
4 6 8
2 2 3
outputCopy
Yes
16 17 21 
inputCopy
2 1
1 2
2 1
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 1, CF1039A)
}

func TestCompareCF1039A(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 3)
		rg.Int(0, 0)
		rg.NewLine()
		rg.IntSliceOrdered(n, 1, 6, true, true)
		rg.IntSlice(n, 1, n)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, t int
		Fscan(in, &n, &t)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] += t
		}
		x := make([]int, n)
		for i := range x {
			Fscan(in, &x[i])
			x[i]--
		}

		permutations := func(n, r int, do func(ids []int) (Break bool)) {
			ids := make([]int, n)
			for i := range ids {
				ids[i] = i
			}
			if do(ids[:r]) {
				return
			}
			cycles := make([]int, r)
			for i := range cycles {
				cycles[i] = n - i
			}
			for {
				i := r - 1
				for ; i >= 0; i-- {
					cycles[i]--
					if cycles[i] == 0 {
						tmp := ids[i]
						copy(ids[i:], ids[i+1:])
						ids[n-1] = tmp
						cycles[i] = n - i
					} else {
						j := cycles[i]
						ids[i], ids[n-j] = ids[n-j], ids[i]
						if do(ids[:r]) {
							return
						}
						break
					}
				}
				if i == -1 {
					return
				}
			}
		}
		max := func(a, b int) int {
			if b > a {
				return b
			}
			return a
		}

		b := make([]int, n)
		var f func(p int) bool
		f = func(p int) bool {
			if p == n {
				mx := make([]int, n)
				for i := range mx {
					mx[i] = -1
				}
				permutations(n, n, func(ids []int) (Break bool) {
					for i, id := range ids {
						if b[id] < a[i] {
							return
						}
					}
					for i, id := range ids {
						mx[i] = max(mx[i], id)
					}
					return
				})
				for i, v := range mx {
					if v != x[i] {
						return false
					}
				}
				return true
			}
			b[p] = 1
			if p > 0 {
				b[p] = b[p-1] + 1
			}
			for ; b[p] <= 11; b[p]++ {
				if f(p + 1) {
					return true
				}
			}
			return false
		}
		if !f(0) {
			Fprint(out, "No")
		} else {
			Fprintln(out, "Yes")
			for _, v := range b {
				Fprint(out, v, " ")
			}
		}
	}
	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1039A)
}

func TestCheckCF1039A(t *testing.T) {
	return
	assert := assert.New(t)
	_ = assert
	testutil.DebugTLE = 0

	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					if do(ids[:r]) {
						return
					}
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		n := rg.Int(4, 5)
		t := rg.Int(0, 0)
		rg.NewLine()
		a := rg.IntSliceOrdered(n, 1, 9, true, true)
		x := rg.IntSlice(n, 1, n)
		return rg.String(), func(myOutput string) (_b bool) {
			in := strings.NewReader(myOutput)

			var yn string
			Fscan(in, &yn)
			if yn == "No" {
				return true
			}

			b := make([]int, n)
			for i := range b {
				Fscan(in, &b[i])
				if i > 0 && b[i] <= b[i-1] {
					return
				}
			}
			mx := make([]int, n)
			for i := range mx {
				mx[i] = -1
			}
			permutations(n, n, func(ids []int) (Break bool) {
				for i, id := range ids {
					if b[id] < a[i]+t {
						return
					}
				}
				for i, id := range ids {
					mx[i] = max(mx[i], id+1)
				}
				return
			})
			if !assert.EqualValues(x, mx) {
				return
			}
			return true
		}
	}

	target := 0
	testutil.CheckRunResultsInfWithTarget(t, inputGenerator, target, CF1039A)
}
