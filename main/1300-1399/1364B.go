package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1364B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := []any{a[0]}
		for i := 1; i < n-1; i++ {
			if a[i-1] < a[i] == (a[i] > a[i+1]) {
				ans = append(ans, a[i])
			}
		}
		ans = append(ans, a[len(a)-1])
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1364B(bufio.NewReader(os.Stdin), os.Stdout) }
