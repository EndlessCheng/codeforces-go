package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type xorBasis62 struct {
	b   [60]int
	num int
}

func (b *xorBasis62) insert(v int) bool {
	for i := len(b.b) - 1; i >= 0; i-- {
		if v>>i&1 == 0 {
			continue
		}
		if b.b[i] == 0 {
			b.b[i] = v
			b.num++
			return true
		}
		v ^= b.b[i]
	}
	return false
}

func cf662A(in io.Reader, out io.Writer) {
	var n, v, w, s int
	Fscan(in, &n)
	b := &xorBasis62{}
	for range n {
		Fscan(in, &v, &w)
		s ^= v
		b.insert(v ^ w)
	}
	if b.insert(s) {
		Fprint(out, "1/1")
	} else {
		Fprint(out, 1<<b.num-1, "/", 1<<b.num)
	}
}

//func main() { cf662A(bufio.NewReader(os.Stdin), os.Stdout) }
