package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1197C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v int
	Fscan(in, &n, &k, &v)
	ans, pre := -v, v
	b := make([]int, n-1)
	for i := range b {
		Fscan(in, &v)
		b[i] = pre - v
		pre = v
	}
	sort.Ints(b)
	for _, v := range b[:k-1] {
		ans += v
	}
	Fprint(out, ans+v)
}

//func main() { CF1197C(os.Stdin, os.Stdout) }
