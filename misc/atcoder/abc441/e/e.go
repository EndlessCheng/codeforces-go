package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	var s string
	Fscan(in, &n, &s)

	cnt := make([]int, n*2+1)
	cnt[n] = 1
	sum, f := n, 0
	for _, b := range s {
		if b == 'A' {
			f += cnt[sum]
			sum++
		} else if b == 'B' {
			sum--
			f -= cnt[sum]
		}
		ans += f
		cnt[sum]++
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
