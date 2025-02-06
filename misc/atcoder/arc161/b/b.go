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
	var T, n uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n < 7 {
			Fprintln(out, -1)
			continue
		}
		for bits.OnesCount(n) < 3 {
			n--
		}
		for bits.OnesCount(n) > 3 {
			n &= n - 1
		}
		Fprintln(out, n)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
