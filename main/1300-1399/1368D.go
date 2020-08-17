package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1368D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	cnt := [20]int{}
	var n, v int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j := range cnt {
			cnt[j] += v >> j & 1
		}
	}
	s := int64(0)
	for ; n > 0; n-- {
		v = 0
		for i, c := range cnt {
			if c > 0 {
				cnt[i]--
				v |= 1 << i
			}
		}
		s += int64(v) * int64(v)
	}
	Fprint(out, s)
}

//func main() { CF1368D(os.Stdin, os.Stdout) }
