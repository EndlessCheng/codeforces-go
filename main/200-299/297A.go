package _00_299

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF297A(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	if (strings.Count(s, "1")+1)/2*2 >= strings.Count(t, "1") {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF297A(os.Stdin, os.Stdout) }
