package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if strings.Contains(s, "11") {
			Fprintln(out, "No")
			continue
		}
		for i, b := range s {
			if b == '0' {
				if (i == 0 || s[i-1] == '0') && (i == n-1 || s[i+1] == '0') {
					Fprintln(out, "No")
					continue o
				}
			}
		}
		Fprintln(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }
