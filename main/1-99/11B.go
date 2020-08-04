package __99

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF11B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var x, s, i int
	Fscan(in, &x)
	if x < 0 {
		x = -x
	}
	for {
		s += i
		if s >= x && (s-x)&1 == 0 {
			Fprint(out, i)
			return
		}
		i++
	}
}

//func main() { CF11B(os.Stdin, os.Stdout) }
