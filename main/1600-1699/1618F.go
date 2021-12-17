package main

import (
	. "fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1618F(in io.Reader, out io.Writer) {
	var x, y uint64
	Fscan(in, &x, &y)
	if x == y {
		Fprint(out, "YES")
		return
	}

	t := strconv.FormatUint(y, 2)
	rev := []byte(t)
	for i, n := 0, len(rev); i < n/2; i++ {
		rev[i], rev[n-1-i] = rev[n-1-i], rev[i]
	}
	f := func(s string) bool {
		re := regexp.MustCompile("^1*" + s + "1*$")
		return re.MatchString(t) || re.Match(rev)
	}

	s := strconv.FormatUint(x, 2)
	if f(s+"1") || f(strings.TrimRight(s, "0")) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF1618F(os.Stdin, os.Stdout) }
