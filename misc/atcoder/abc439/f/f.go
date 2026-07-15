package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) add(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	if l > r {
		return 0
	}
	return t.pre(r) - t.pre(l-1)
}

func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, v int
	Fscan(in, &n)
	single := make(fenwick, n+1)
	inc := make(fenwick, n+1)
	dec := make(fenwick, n+1)
	for range n {
		Fscan(in, &v)
		single.add(v, 1)
		inc.add(v, (inc.pre(v-1)+dec.pre(v-1)+1)%mod)
		dec.add(v, (dec.query(v+1, n)+inc.query(v+1, n)-single.query(v+1, n))%mod)
	}
	Fprint(out, (dec.pre(n)+mod)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
