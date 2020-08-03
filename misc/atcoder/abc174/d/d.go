package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	var s []byte
	Fscan(in, &n, &s)
	i, j := 0, n-1
	for i < j {
		for ; i < j && s[i] == 'R'; i++ {
		}
		for ; i < j && s[j] == 'W'; j-- {
		}
		if i >= j {
			break
		}
		ans++
		i++
		j--
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
