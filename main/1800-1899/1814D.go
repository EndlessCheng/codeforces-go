package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1814D(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		f := make([]int, n)
		for i := range f {
			Fscan(in, &f[i])
		}
		d := make([]int, n)
		pw := make([]int, n)
		for i := range d {
			Fscan(in, &d[i])
			pw[i] = f[i] * d[i]
		}
		slices.Sort(pw)

		ans := n
		for i, v := range f {
			p := v * d[i]
			tar := p
			for tar > 0 && p <= tar+k {
				t := tar
				for _, fj := range f {
					if (tar+fj-1)/fj*fj > tar+k {
						tar = (tar / fj) * fj
					}
				}
				if tar == t {
					break
				}
			}
			if tar <= 0 || p > tar+k {
				continue
			}
			l := sort.SearchInts(pw, tar)
			r := sort.SearchInts(pw, tar+k+1)
			ans = min(ans, n-r+l)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1814D(bufio.NewReader(os.Stdin), os.Stdout) }
