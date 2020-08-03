package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m, x int
	Fscan(in, &m)
	vis := make([]bool, m)
	for i := 1; ; i++ {
		if x = (x*10 + 7) % m; x == 0 {
			Fprint(out, i)
			return
		}
		if vis[x] {
			break
		}
		vis[x] = true
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
