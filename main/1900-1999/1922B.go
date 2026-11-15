package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1922B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		ans, sum := 0, 0
		for _, c := range cnt {
			ans += c*(c-1)*(c-2)/6 + c*(c-1)/2*sum
			sum += c
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1922B(bufio.NewReader(os.Stdin), os.Stdout) }
