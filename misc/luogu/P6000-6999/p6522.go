package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p6522(in io.Reader, out io.Writer) {
	const mod = 1_000_000_009
	var n, d int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	ans := 1
	j := 0
	for i, v := range a {
		for a[j] < v-d {
			j++
		}
		ans = ans * (i - j + 1) % mod
	}
	Fprint(out, ans)
}

//func main() { p6522(bufio.NewReader(os.Stdin), os.Stdout) }
