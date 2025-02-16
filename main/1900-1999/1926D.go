package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1926D(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 0
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			w := 1<<31 - 1 ^ v
			if cnt[w] > 0 {
				cnt[w]--
			} else {
				ans++
				cnt[v]++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1926D(bufio.NewReader(os.Stdin), os.Stdout) }
