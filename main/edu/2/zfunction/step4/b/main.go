package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var T int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		Fprintln(out, strings.Index(s+s, t)) // or KMP or zSearch
	}
}
