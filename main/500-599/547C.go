package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF547C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 5e5
	ds := [mx + 1][]int{}
	mu := [mx + 1]int{1: 1}
	for i := 1; i <= mx; i++ {
		ds[i] = append(ds[i], i)
		for j := i * 2; j <= mx; j += i {
			ds[j] = append(ds[j], i)
			mu[j] -= mu[i]
		}
	}
	cnt := [mx + 1]int{}

	var n, q, id int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for ans := int64(0); q > 0; q-- {
		Fscan(in, &id)
		id--
		v, put := a[id], true
		if v < 0 {
			v, put = -v, false
		}
		a[id] = -a[id]
		for _, d := range ds[v] {
			if put {
				ans += int64(mu[d] * cnt[d])
				cnt[d]++
			} else {
				cnt[d]--
				ans -= int64(mu[d] * cnt[d])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF547C(os.Stdin, os.Stdout) }
