package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type xorBasis []int

func (b xorBasis) insert(v int) {
	for i := len(b) - 1; i >= 0; i-- {
		if v>>i == 0 {
			continue
		}
		if b[i] == 0 {
			b[i] = v
			return
		}
		v ^= b[i]
	}
}

func (b xorBasis) maxXor() (res int) {
	for i := len(b) - 1; i >= 0; i-- {
		res = max(res, res^b[i])
	}
	return
}

func run(in io.Reader, out io.Writer) {
	var n, xor int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		xor ^= a[i]
	}
	b := make(xorBasis, 60)
	for _, v := range a {
		b.insert(v &^ xor)
	}
	Fprint(out, xor+b.maxXor()*2)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
