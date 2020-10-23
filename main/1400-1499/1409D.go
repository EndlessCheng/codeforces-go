package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1409D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	f := func(x int64) (s int64) {
		for ; x > 0; x /= 10 {
			s += x % 10
		}
		return
	}

	var T, up, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &up)
		if f(n) <= up {
			Fprintln(out, 0)
			continue
		}
		for p10 := int64(10); ; p10 *= 10 {
			if v := (n/p10 + 1) * p10; f(v) <= up {
				Fprintln(out, v-n)
				break
			}
		}
	}
}

//func main() { CF1409D(os.Stdin, os.Stdout) }
