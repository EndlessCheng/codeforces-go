package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/big"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, b0, b1 int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &b0, &b1, &s)
		Fprintln(out, b0, string(s))
		for i, b := range s {
			if b&64 > 0 {
				s[i] ^= 32
			}
		}
		v, _ := new(big.Int).SetString(string(s), b0)
		t := []byte(v.Text(b1))
		for i, b := range t {
			if b&64 > 0 {
				t[i] ^= 32
			}
		}
		Fprintf(out, "%d %s\n\n", b1, t)
	}
}

func main() { run(os.Stdin, os.Stdout) }
