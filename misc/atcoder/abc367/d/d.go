package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, ans, sl, sr int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	cnt := map[int]int{}
	for i := 0; i < n*2-1; i++ {
		sr += a[i%n]
		if i >= n-1 {
			ans += cnt[sr%m]
			sl += a[i-n+1]
			cnt[sl%m]--
		}
		cnt[sr%m]++
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
