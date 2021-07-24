package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ i, j int }

	var k, n int
	Fscan(in, &k)
	pos := map[int]pair{}
	for i := 1; i <= k; i++ {
		Fscan(in, &n)
		a := make([]int, n)
		sum := 0
		for j := range a {
			Fscan(in, &a[j])
			sum += a[j]
		}
		for j, v := range a {
			if p, has := pos[sum-v]; has && p.i < i {
				Fprintln(out, "YES")
				Fprintln(out, p.i, p.j+1)
				Fprintln(out, i, j+1)
				return
			}
			pos[sum-v] = pair{i, j} // 记录删掉该元素后的位置信息
		}
	}
	Fprint(out, "NO")
}

func main() { run(os.Stdin, os.Stdout) }
