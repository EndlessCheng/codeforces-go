package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1158A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m uint
	var v, ans, fi, se uint64
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &v)
		ans += v
		if v > fi {
			fi, se = v, fi
		} else if v > se {
			se = v
		}
	}
	ans = (ans - fi) * uint64(m)
	min := uint64(1e9)
	for ; m > 0; m-- {
		Fscan(in, &v)
		ans += v
		if v < min {
			min = v
		}
	}
	if fi > min {
		Fprint(out, -1)
	} else {
		if fi < min {
			ans += fi - se
		}
		Fprint(out, ans)
	}
}

//func main() { CF1158A(os.Stdin, os.Stdout) }
