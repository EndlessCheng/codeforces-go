package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf817B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	a := make([]int, 0, len(cnt))
	for k := range cnt {
		a = append(a, k)
	}
	sort.Ints(a)

	c0 := cnt[a[0]]
	if c0 > 2 {
		ans = c0 * (c0 - 1) * (c0 - 2) / 6
	} else if c1 := cnt[a[1]]; c0 > 1 {
		ans = c1
	} else if c1 > 1 {
		ans = c1 * (c1 - 1) / 2
	} else {
		ans = cnt[a[2]]
	}
	Fprint(out, ans)
}

//func main() { cf817B(os.Stdin, os.Stdout) }
