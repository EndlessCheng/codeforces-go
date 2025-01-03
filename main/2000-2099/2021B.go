package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2021B(in io.Reader, out io.Writer) {
	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		cnt := map[int]int{}
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}
		for i := 0; ; i++ {
			if cnt[i] == 0 {
				Fprintln(out, i)
				break
			}
			cnt[i+x] += cnt[i] - 1
		}
	}
}

//func main() { cf2021B(bufio.NewReader(os.Stdin), os.Stdout) }
