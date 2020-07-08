package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestCF1027D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 2 10
1 3 4 3 3
outputCopy
3
inputCopy
4
1 10 2 10
2 4 2 2
outputCopy
10
inputCopy
7
1 1 1 1 1 1 1
2 2 2 3 6 7 6
outputCopy
2
inputCopy
4
1 1 1 1 
2 2 1 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1027D)
}

// 无尽对拍
func TestCF1027DInf(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		buf := &bytes.Buffer{}
		n := 4
		buf.WriteString(strconv.Itoa(n) + "\n")
		buf.WriteString(strings.Repeat("1 ", n) + "\n")
		for j := 0; j < n; j++ {
			buf.WriteString(strconv.Itoa(rand.Intn(n)+1) + " ")
		}
		return buf.String()
	}

	// AC 算法
	runAC := func(in io.Reader, out io.Writer) {
		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}

		var n, s int
		Fscan(in, &n)
		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
		}
		g := make([]int, n)
		deg := make([]int, n)
		for i := range g {
			Fscan(in, &g[i])
			g[i]--
			deg[g[i]]++
		}

		vis := make([]bool, n)
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			vis[v] = true
			q = q[1:]
			w := g[v]
			deg[w]--
			if deg[w] == 0 {
				q = append(q, w)
			}
		}

		var f func(int) int
		f = func(v int) int {
			vis[v] = true
			if vis[g[v]] {
				return c[v]
			}
			return min(f(g[v]), c[v])
		}
		for i, b := range vis {
			if !b {
				s += f(i)
			}
		}
		Fprint(out, s)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runAC, CF1027D)
}
