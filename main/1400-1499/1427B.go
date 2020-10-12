package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1427B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		if !strings.Contains(s, "W") {
			if k == 0 {
				Fprintln(out, 0)
			} else {
				Fprintln(out, 2*(k-1)+1)
			}
			continue
		}

		lr := -2
		s = strings.TrimFunc(s, func(r rune) bool { lr++; return r == 'L' })
		ans := 0
		for _, s := range strings.Split(s, "L") {
			if s != "" {
				ans += 2*(len(s)-1) + 1
			}
		}
		cnt := []int{}
		for _, s := range strings.Split(s, "W") {
			if s != "" {
				cnt = append(cnt, len(s))
			}
		}
		sort.Ints(cnt)
		for _, c := range cnt {
			if k < c {
				ans += 2 * k
				k = 0
				break
			}
			ans += 2*c + 1
			k -= c
		}
		if k > 0 {
			if k > lr {
				k = lr
			}
			ans += 2 * k
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1427B(os.Stdin, os.Stdout) }
