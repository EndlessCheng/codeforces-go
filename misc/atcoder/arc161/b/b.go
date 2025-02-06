package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const k = 3
	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n < 1<<k-1 {
			Fprintln(out, -1)
			continue
		}
		for bits.OnesCount(n) < k {
			n = n&(n+1) - 1
		}
		for bits.OnesCount(n) > k {
			n &= n - 1
		}
		Fprintln(out, n)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
