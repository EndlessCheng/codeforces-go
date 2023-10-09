package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1690E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			ans += a[i] / k
			a[i] %= k
		}
		sort.Ints(a)
		i, j := 0, n-1
		for i < j {
			if a[i]+a[j] >= k {
				ans++
				j--
			}
			i++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1690E(os.Stdin, os.Stdout) }
