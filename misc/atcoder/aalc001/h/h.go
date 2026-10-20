package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) update(i, val int) {
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
	return t.pre(r) - t.pre(l-1)
}

func (t fenwick) kth(k int) int {
	i := 0
	for bit := 1 << 17; bit > 0; bit >>= 1 {
		nxt := i | bit
		if nxt < len(t) && t[nxt] < k {
			k -= t[nxt]
			i = nxt
		}
	}
	return i + 1
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, x, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	t := make(fenwick, n+1)
	zeroT := make(fenwick, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		t[i] += a[i]
		if a[i] == 0 {
			zeroT[i]++
		}
		j := i + i&-i
		if j <= n {
			t[j] += t[i]
			zeroT[j] += zeroT[i]
		}
	}

	const mod = 998244353
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			if a[l] == 0 {
				zeroT.update(l, -1)
			}
			if r == 0 {
				zeroT.update(l, 1)
			}
			t.update(l, r-a[l])
			a[l] = r
		} else {
			Fscan(in, &x)
			lz, rz := zeroT.pre(l-1), zeroT.pre(r)
			if lz == rz {
				Fprintln(out, x+t.query(l, r))
			} else {
				i := zeroT.kth(rz)
				Fprintln(out, (x+t.query(l, i))%mod+t.query(i, r))
			}
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
