package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol632C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	s := make([]string, n)
	for i := range s {
		Fscan(in, &s[i])
	}
	sort.Slice(s, func(i, j int) (less bool) { return s[i]+s[j] < s[j]+s[i] })
	for _, val := range s {
		Fprint(out, val)
	}
}

//func main() { Sol632C(os.Stdin, os.Stdout) }
