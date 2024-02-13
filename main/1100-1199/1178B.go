package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1178B(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	var f0, f1, f2 int
	for i := 1; i < len(s); i++ {
		if s[i] == 'o' {
			f1 += f0
		} else if s[i-1] == 'v' {
			f2 += f1
			f0++
		}
	}
	Fprint(out, f2)
}

func cf1178B2(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	var pre, suf, ans int
	n := len(s)
	for i := 1; i < n-1; i++ {
		if s[i] == 'v' && s[i+1] == 'v' {
			suf++
		}
	}
	for i := 1; i < n-2; i++ {
		if s[i] == 'o' {
			ans += pre * suf
		} else {
			if s[i-1] == 'v' {
				pre++
			}
			if s[i+1] == 'v' {
				suf--
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1178B(os.Stdin, os.Stdout) }
