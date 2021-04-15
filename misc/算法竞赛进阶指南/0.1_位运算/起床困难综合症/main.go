package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, z, atk, ans int
	var op string
	o := 1<<31 - 1
	for Fscan(in, &n, &m); n > 0; n-- {
		if Fscan(in, &op, &v); op[0] == 'O' {
			z |= v
			o |= v
		} else if op[0] == 'X' {
			z ^= v
			o ^= v
		} else {
			z &= v
			o &= v
		}
	}
	for i := 29; i >= 0; i-- {
		if atk|1<<i <= m && z>>i&1 == 0 && o>>i&1 > 0 {
			atk |= 1 << i
			ans |= 1 << i
		} else {
			ans |= z >> i & 1 << i
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
