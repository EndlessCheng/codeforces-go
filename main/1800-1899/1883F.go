package main

import (
	. "fmt"
	"io"
)

func cf1883F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		r := map[int]int{}
		for i := range a {
			Fscan(in, &a[i])
			r[a[i]] = i
		}

		ans := 0
		vis := map[int]bool{}
		rs := len(r)
		for i, v := range a {
			if !vis[v] {
				vis[v] = true
				ans += rs
			}
			if r[v] == i {
				rs--
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1883F(bufio.NewReader(os.Stdin), os.Stdout) }
