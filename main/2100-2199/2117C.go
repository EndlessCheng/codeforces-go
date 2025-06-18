package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2117C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		ans := 0
		pre := map[int]struct{}{}
		cur := map[int]struct{}{}
		Fscan(in, &n)
		for range n {
			Fscan(in, &v)
			cur[v] = struct{}{}
			delete(pre, v)
			if len(pre) == 0 {
				ans++
				pre = cur
				cur = map[int]struct{}{}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2117C(bufio.NewReader(os.Stdin), os.Stdout) }
