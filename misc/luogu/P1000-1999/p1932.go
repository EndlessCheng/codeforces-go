package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
)

// https://space.bilibili.com/206214
func p1932(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	v, w, res := new(big.Int), new(big.Int), new(big.Int)
	Fscan(in, v, w)
	Fprintln(out, res.Add(v, w))
	Fprintln(out, res.Sub(v, w))
	Fprintln(out, res.Mul(v, w))
	q, r := res.QuoRem(v, w, new(big.Int))
	Fprintln(out, q)
	Fprintln(out, r)
}

//func main() { p1932(os.Stdin, os.Stdout) }
