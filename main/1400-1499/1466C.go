package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1466C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := 0
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] || i > 1 && s[i] == s[i-2] {
				s[i] = 0
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1466C(bufio.NewReader(os.Stdin), os.Stdout) }
