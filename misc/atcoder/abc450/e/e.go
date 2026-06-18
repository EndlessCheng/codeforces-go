package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s, t, c string
	var q, l, r int
	Fscan(in, &s, &t)

	init := func(s string) [][26]int {
		sum := make([][26]int, len(s)+1)
		for i, b := range s {
			sum[i+1] = sum[i]
			sum[i+1][b-'a']++
		}
		return sum
	}
	sumS := init(s)
	sumT := init(t)

	sz := []int{len(s), len(t)}
	ss := [][26]int{sumS[len(s)], sumT[len(t)]}
	for i := 1; sz[i] <= 1e18; i++ {
		sz = append(sz, sz[i]+sz[i-1])
		sum := ss[i]
		for j, v := range ss[i-1] {
			sum[j] += v
		}
		ss = append(ss, sum)
	}

	f := func(k int) (res int) {
		b := c[0] - 'a'
		if k <= sz[1] {
			return sumT[k][b]
		}
		i := sort.SearchInts(sz[1:], k) + 1
		for i > 1 {
			if k > sz[i-1] {
				k -= sz[i-1]
				res += ss[i-1][b]
				i -= 2
			} else {
				i--
			}
		}
		if i == 0 {
			res += sumS[k][b]
		} else {
			res += sumT[k][b]
		}
		return
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r, &c)
		Fprintln(out, f(r)-f(l-1))
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
