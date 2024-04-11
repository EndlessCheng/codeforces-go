package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF988C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var k, n int
	Fscan(in, &k)
	type pair struct{ i, j int }
	mp := map[int]pair{}
	for i := 1; i <= k; i++ {
		Fscan(in, &n)
		a := make([]int, n)
		sum := 0
		for j := range a {
			Fscan(in, &a[j])
			sum += a[j]
		}
		for j, v := range a {
			p := mp[sum-v]
			if p.i > 0 && p.i < i {
				Fprintln(out, "YES", p.i, p.j+1, i, j+1)
				return
			}
			mp[sum-v] = pair{i, j}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF988C(os.Stdin, os.Stdout) }
