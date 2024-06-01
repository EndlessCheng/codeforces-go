package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func cf985C(in io.Reader, out io.Writer) {
	var n, k, l, ans int
	Fscan(in, &n, &k, &l)
	a := make([]int, n*k)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	i := sort.SearchInts(a, a[0]+l+1)
	if i < n {
		Fprint(out, 0)
		return
	}

	x := 0
	if k > 1 {
		x = (i - n + k - 2) / (k - 1)
	}
	for j := 0; j <= (x-1)*k; j += k {
		ans += a[j]
	}
	for _, v := range a[i-n+x : i] {
		ans += v
	}
	Fprint(out, ans)
}

//func main() { cf985C(bufio.NewReader(os.Stdin), os.Stdout) }
