package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf898D(in io.Reader, out io.Writer) {
	var n, m, k, cnt, l, ans int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	for i, x := range a {
		cnt++
		for x-a[l] >= m {
			if a[l] > 0 {
				cnt--
			}
			l++
		}
		if cnt == k {
			a[i] = 0
			cnt--
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf898D(bufio.NewReader(os.Stdin), os.Stdout) }
