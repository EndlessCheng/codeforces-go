package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s string
	var n, p, need [3]int64
	var m int64
	Fscan(in, &s, &n[0], &n[1], &n[2], &p[0], &p[1], &p[2], &m)
	for _, b := range s {
		for i, t := range "BSC" {
			if b == t {
				need[i]++
			}
		}
	}
	l, r := int64(1), m+101
	for l < r {
		v := (l + r) >> 1
		cm := int64(0)
		for i, c := range need {
			if c*v > n[i] {
				cm += (c*v - n[i]) * p[i]
			}
		}
		if cm > m {
			r = v
		} else {
			l = v + 1
		}
	}
	Fprint(out, l-1)
}

func main() { run(os.Stdin, os.Stdout) }
