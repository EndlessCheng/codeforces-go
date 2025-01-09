package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var lowS, highS string
	Fscan(in, &lowS, &highS)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][10]int, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, hd int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if hd == 0 {
				return 0
			}
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][hd]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		if hd == 0 {
			for d := lo; d <= hi; d++ {
				res += f(i+1, d, limitLow && d == lo, limitHigh && d == hi)
			}
		} else {
			for d := lo; d <= min(hi, hd-1); d++ {
				res += f(i+1, hd, limitLow && d == lo, limitHigh && d == hi)
			}
		}
		return
	}
	Fprint(out, f(0, 0, true, true))
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
