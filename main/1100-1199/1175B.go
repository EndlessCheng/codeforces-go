package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func CF1175B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const up int64 = math.MaxUint32
	type pair struct {
		add bool
		v   int64
	}

	var q, n int
	var op string
	s := []pair{}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op)
		if op[0] == 'f' {
			Fscan(in, &n)
			s = append(s, pair{false, int64(n)})
		} else if op[0] == 'a' {
			s = append(s, pair{true, 1})
		} else {
			v := int64(0)
			for s[len(s)-1].add {
				v += s[len(s)-1].v
				s = s[:len(s)-1]
			}
			v *= s[len(s)-1].v
			s = s[:len(s)-1]
			if v > up {
				Fprint(out, "OVERFLOW!!!")
				return
			}
			s = append(s, pair{true, v})
		}
	}
	v := int64(0)
	for _, p := range s {
		v += p.v
	}
	if v > up {
		Fprint(out, "OVERFLOW!!!")
		return
	}
	Fprint(out, v)
}

//func main() { CF1175B(os.Stdin, os.Stdout) }
