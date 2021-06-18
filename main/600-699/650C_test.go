package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/650/C
// https://codeforces.com/problemset/status/650/problem/C
func TestCF650C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
1 2
3 4
outputCopy
1 2
2 3
inputCopy
4 3
20 10 30
50 40 30
50 60 70
90 80 70
outputCopy
2 1 3
5 4 3
5 6 7
9 8 7`
	testutil.AssertEqualCase(t, rawText, 0, CF650C)
}

type cell struct {
	r, c, a        int
	rl, rh, cl, ch *cell
	v              int
}

type byA []*cell

func (s byA) Len() int           { return len(s) }
func (s byA) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byA) Less(i, j int) bool { return s[i].a < s[j].a }

func fill(c *cell, v int) {
	if c.v != 0 {
		return
	}
	c.v = v
	for _, try := range []*cell{c.rl, c.rh, c.cl, c.ch} {
		if try != nil && try.a == c.a {
			fill(try, v)
		}
	}
}

func TestCompareCF650C(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		m := rg.Int(1, 9)
		rg.NewLine()
		rg.IntSlice(n*m, 1, 9)
		return rg.String()
	}

	runBF := func(in io.Reader, out io.Writer) {
		var n, m int
		fmt.Fscan(in, &n, &m)

		byRow := make([][]*cell, n)
		byCol := make([][]*cell, m)

		for r := range byRow {
			for c := range byCol {
				cell := &cell{r: r, c: c}
				fmt.Fscan(in, &cell.a)
				byRow[r] = append(byRow[r], cell)
				byCol[c] = append(byCol[c], cell)
			}
		}

		for _, s := range byRow {
			sort.Sort(byA(s))
			for i := range s {
				if i > 0 {
					s[i].rl = s[i-1]
				}
				if i+1 < m {
					s[i].rh = s[i+1]
				}
			}
		}
		for _, s := range byCol {
			sort.Sort(byA(s))
			for i := range s {
				if i > 0 {
					s[i].cl = s[i-1]
				}
				if i+1 < n {
					s[i].ch = s[i+1]
				}
			}
		}

		numVertices := 0
		for _, row := range byRow {
			for _, cell := range row {
				if cell.v == 0 {
					numVertices++
					fill(cell, numVertices)
				}
			}
		}

		edges := make([][]int, numVertices+1)
		for _, row := range byRow {
			for i := range row[1:] {
				if row[i].a < row[i+1].a {
					v := row[i+1].v
					edges[v] = append(edges[v], row[i].v)
				}
			}
		}
		for _, col := range byCol {
			for i := range col[1:] {
				if col[i].a < col[i+1].a {
					v := col[i+1].v
					edges[v] = append(edges[v], col[i].v)
				}
			}
		}

		vertices := make([]int, numVertices+1)
		var val func(int) int
		val = func(v int) int {
			if vertices[v] != 0 {
				return vertices[v]
			}
			ans := 1
			for _, u := range edges[v] {
				if ch := val(u); ans <= ch {
					ans = ch + 1
				}
			}
			vertices[v] = ans
			return ans
		}

		res := make([][]int, n)
		for i := range res {
			res[i] = make([]int, m)
		}
		for _, row := range byRow {
			for _, cell := range row {
				res[cell.r][cell.c] = val(cell.v)
			}
		}

		for _, row := range res {
			for i, a := range row {
				if i == 0 {
					fmt.Fprint(out, a)
				} else {
					fmt.Fprint(out, " ", a)
				}
			}
			fmt.Fprintln(out)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF650C)
}
