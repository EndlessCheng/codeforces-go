package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol490C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s string
	var a, b int
	Fscan(in, &s, &a, &b)
	n := len(s)

	l := int(s[0]-'0') % a
	validPos := []int{}
	for i := 1; i < n; i++ {
		if s[i] != '0' && l == 0 {
			validPos = append(validPos, i)
		}
		l = (10*l + int(s[i]-'0')) % a
	}
	if len(validPos) == 0 {
		Fprint(out, "NO")
		return
	}

	j := n - 2
	pow10 := 1
	r := int(s[n-1]-'0') % b
	for i := len(validPos) - 1; i >= 0; i-- {
		pos := validPos[i]
		for ; j >= pos; j-- {
			pow10 = pow10 * 10 % b
			r = (pow10*int(s[j]-'0') + r) % b
		}
		if r == 0 {
			Fprintf(out, "YES\n%s\n%s", s[:pos], s[pos:])
			return
		}
	}
	Fprint(out, "NO")
}

func main() {
	Sol490C(os.Stdin, os.Stdout)
}
