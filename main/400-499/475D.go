package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF475D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, x, q int
	Fscan(in, &n)
	cnt := map[int]int{}
	type pair struct{ v, cnt int }
	opRes := []pair{}
	for i := 0; i < n; i++ {
		Fscan(in, &x)
		for j, p := range opRes {
			opRes[j].v = gcd(p.v, x)
		}
		opRes = append(opRes, pair{x, 1})

		k := 1
		for j := 1; j < len(opRes); j++ {
			if opRes[j].v != opRes[j-1].v {
				opRes[k] = opRes[j]
				k++
			} else {
				opRes[k-1].cnt += opRes[j].cnt
			}
		}
		opRes = opRes[:k]

		for _, p := range opRes {
			cnt[p.v] += p.cnt
		}
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &x)
		Fprintln(out, cnt[x])
	}
}

//func main() { CF475D(os.Stdin, os.Stdout) }
