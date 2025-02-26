// LUOGU_RID: 204790319
package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, t, x, s1, s2, k int
	Fscan(in, &n, &m)
	a := [3][]int{}
	for ; n > 0; n-- {
		Fscan(in, &t, &x)
		a[t] = append(a[t], x)
	}
	cmp := func(a, b int) int { return b - a }
	slices.SortFunc(a[0], cmp)
	slices.SortFunc(a[1], cmp)
	slices.SortFunc(a[2], cmp)

	s0 := make([]int, len(a[0])+1)
	for i, v := range a[0] {
		s0[i+1] = s0[i] + v
	}

	ans := s0[min(m, len(a[0]))]
	for j, v := range a[1] {
		s1 += v
		m--
		for s2 <= j && k < len(a[2]) {
			s2 += a[2][k]
			k++
			m--
		}
		if s2 <= j || m < 0 {
			break
		}
		ans = max(ans, s1+s0[min(m, len(a[0]))])
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
