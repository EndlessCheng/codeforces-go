package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1495B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, mx, mxCnt, pre int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ok := false
	for i := 1; i < n; {
		st := i
		inc := a[i-1] < a[i]
		for ; i < n && a[i-1] < a[i] == inc; i++ {
		}
		l := i - st + 1
		if l > mx {
			mx, mxCnt = l, 1
		} else if l == mx {
			mxCnt++
			if pre == mx {
				ok = !inc
			}
		}
		pre = l
	}
	if mxCnt == 2 && ok {
		Fprint(out, mx&1)
	} else {
		Fprint(out, 0)
	}
}

//func main() { CF1495B(os.Stdin, os.Stdout) }
