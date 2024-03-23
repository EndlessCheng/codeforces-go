package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1703F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		id := []int{}
		ans := 0
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if v < i {
				ans += sort.SearchInts(id, v)
				id = append(id, i)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1703F(os.Stdin, os.Stdout) }
