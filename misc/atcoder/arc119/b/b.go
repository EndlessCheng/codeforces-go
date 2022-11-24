package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, ans, j int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &s, &t)
	if strings.Count(s, "0") != strings.Count(t, "0") {
		Fprint(out, -1)
		return
	}
	for i, b := range s {
		if b == '0' {
			for t[j] != '0' {
				j++
			}
			if i != j {
				ans++
			}
			j++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
