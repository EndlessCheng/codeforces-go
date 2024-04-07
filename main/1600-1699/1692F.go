package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1692F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		a := []int{}
		cnt := [10]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			v %= 10
			if cnt[v] < 3 {
				cnt[v]++
				a = append(a, v)
			}
		}
		for i, x := range a {
			for j, y := range a[:i] {
				for _, z := range a[:j] {
					if (x+y+z)%10 == 3 {
						Fprintln(out, "YES")
						continue o
					}
				}
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1692F(os.Stdin, os.Stdout) }
