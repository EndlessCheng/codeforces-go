package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1062B(in io.Reader, out io.Writer) {
	var n, maxE int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 1, 0)
		return
	}

	rad := 1
	has := map[int]bool{}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			e := 0
			for ; n%i == 0; n /= i {
				e++
			}
			rad *= i
			has[e] = true
			if e > maxE {
				maxE = e
			}
		}
	}
	if n > 1 {
		rad *= n
		has[1] = true
		if maxE == 0 {
			maxE = 1
		}
	}
	op := bits.Len(uint(maxE - 1))
	if len(has) > 1 || maxE&(maxE-1) > 0 {
		op++
	}
	Fprint(out, rad, op)
}

//func main() { CF1062B(os.Stdin, os.Stdout) }
