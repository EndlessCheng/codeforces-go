package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]struct{ x, y, z int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}
	for i := -1; i < 2; i += 2 {
		for j := -1; j < 2; j += 2 {
			for k := -1; k < 2; k += 2 {
				sort.Slice(a, func(_i, _j int) bool {
					a, b := a[_i], a[_j]
					return a.x*i+a.y*j+a.z*k > b.x*i+b.y*j+b.z*k
				})
				var s1, s2, s3 int
				for _, p := range a[:m] {
					s1 += p.x
					s2 += p.y
					s3 += p.z
				}
				ans = max(ans, abs(s1)+abs(s2)+abs(s3))
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
