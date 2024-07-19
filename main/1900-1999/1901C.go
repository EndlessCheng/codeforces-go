package main

import (
	. "fmt"
	"io"
	"math/bits"
	"strings"
)

func cf1901C(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		mn, mx := int(1e9), 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			mn = min(mn, v)
			mx = max(mx, v)
		}
		if mn == mx {
			Fprintln(out, 0)
			continue
		}
		u := 1<<bits.Len(uint(mx)) - 1
		d := u - mx
		ans := 1 + bits.Len(uint((mn+d)/2^u/2))
		Fprintln(out, ans)
		if ans <= n {
			Fprintln(out, d, strings.Repeat("0 ", ans-1))
		}
	}
}

//func main() { cf1901C(bufio.NewReader(os.Stdin), os.Stdout) }
