package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf152C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, m int
	var s string
	Fscan(in, &n, &m)
	mask := make([]uint, m)
	for range n {
		Fscan(in, &s)
		for j, b := range s {
			mask[j] |= 1 << (b - 'A')
		}
	}
	ans := 1
	for _, v := range mask {
		ans = ans * bits.OnesCount(v) % mod
	}
	Fprint(out, ans)
}

//func main() { cf152C(bufio.NewReader(os.Stdin), os.Stdout) }
