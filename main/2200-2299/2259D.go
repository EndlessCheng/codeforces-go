package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2259D(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := bytes.Repeat([]byte{'A'}, n)
		cnt0 := 0
		for i := range n {
			Fscan(in, &v)
			if v == 0 {
				if cnt0 == 0 {
					ans[i] = 'B'
				} else {
					ans[i] = 'C'
				}
				cnt0++
			}
		}

		if cnt0 == 1 {
			Fprintln(out, "NO")
		} else {
			Fprintf(out, "YES\n%s\n", ans)
		}
	}
}

//func main() { cf2259D(bufio.NewReader(os.Stdin), os.Stdout) }
