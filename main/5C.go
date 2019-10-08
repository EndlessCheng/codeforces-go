package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func Sol5C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	bs := []byte(s)

	size := 0
	for i, c := range s {
		if c == '(' {
			size++
		} else {
			size--
			if size < 0 {
				bs[i] = ' '
				size = 0
			}
		}
	}
	size = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			size++
		} else {
			size--
			if size < 0 {
				bs[i] = ' '
				size = 0
			}
		}
	}
	maxLen := map[int]int{0: 1}
	for _, ss := range strings.Split(string(bs), " ") {
		if ss != "" {
			maxLen[len(ss)]++
		}
	}
	ans := 0
	for l := range maxLen {
		if l > ans {
			ans = l
		}
	}
	Fprint(out, ans, maxLen[ans])
}

func main() {
	Sol5C(os.Stdin, os.Stdout)
}
