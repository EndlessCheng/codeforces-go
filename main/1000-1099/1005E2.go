package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1005E2(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(m int) (res int) {
		cnt := make([]int, n*2+1)
		s := n
		gr := 0
		for _, v := range a {
			cnt[s]++
			if v >= m {
				gr += cnt[s]
				s++
			} else {
				s--
				gr -= cnt[s]
			}
			res += gr
		}
		return
	}
	Fprint(out, f(m)-f(m+1))
}

//func main() { cf1005E2(bufio.NewReader(os.Stdin), os.Stdout) }
