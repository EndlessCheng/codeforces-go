package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF627A(in io.Reader, out io.Writer) {
	var s, x, ans int64
	Fscan(in, &s, &x)
	if d := s - x; d >= 0 && d&1 == 0 && d/2|x == d/2+x {
		ans = int64(1) << bits.OnesCount64(uint64(x))
		if d == 0 {
			ans -= 2
		}
	}
	Fprint(out, ans)
}

//func main() { CF627A(os.Stdin, os.Stdout) }
