package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF305C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, max int
	Fscan(in, &n)
	bit := make(map[int]bool, n)
	for ; n > 0; n-- {
		Fscan(in, &v)
		for bit[v] {
			delete(bit, v)
			v++
		}
		bit[v] = true
		if v > max {
			max = v
		}
	}
	Fprint(out, max+1-len(bit))
}

//func main() { CF305C(os.Stdin, os.Stdout) }
