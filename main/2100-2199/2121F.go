package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2121F(in io.Reader, out io.Writer) {
	var T, n, tar, x, v, s int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &tar, &x)
		ans := 0
		cnt := map[int]int{}
		todo := []int{}
		for range n {
			Fscan(in, &v)
			if v > x {
				cnt = map[int]int{}
				todo = todo[:0]
				continue
			}
			todo = append(todo, s)
			if v == x {
				for _, s := range todo {
					cnt[s]++
				}
				todo = todo[:0]
			}
			s += v
			ans += cnt[s-tar]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2121F(bufio.NewReader(os.Stdin), os.Stdout) }
