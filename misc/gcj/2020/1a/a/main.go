package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) {
		var n int
		var s string
		pres := []string{}
		sufs := []string{}
		mid := []byte{}
		Fscan(in, &n)
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			l, r := 0, len(s)
			if s[0] != '*' {
				l = strings.IndexByte(s, '*')
				pres = append(pres, s[:l])
			}
			if s[len(s)-1] != '*' {
				r = strings.LastIndexByte(s, '*')
				sufs = append(sufs, s[r+1:])
			}
			for _, b := range s[l:r] {
				if b != '*' {
					mid = append(mid, byte(b))
				}
			}
		}
		mxLen := 0
		pre, suf := "", ""
		for _, s := range pres {
			if len(s) > mxLen {
				mxLen = len(s)
				pre = s
			}
		}
		for _, s := range pres {
			if !strings.HasPrefix(pre, s) {
				Fprintln(out, "*")
				return
			}
		}
		mxLen = 0
		for _, s := range sufs {
			if len(s) > mxLen {
				mxLen = len(s)
				suf = s
			}
		}
		for _, s := range sufs {
			if !strings.HasSuffix(suf, s) {
				Fprintln(out, "*")
				return
			}
		}
		Fprintln(out, pre+string(mid)+suf)
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
