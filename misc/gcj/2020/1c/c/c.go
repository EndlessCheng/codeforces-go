package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

type pair struct{ x, y int }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func frac(a, b int) pair {
	g := gcd(a, b)
	return pair{a / g, b / g}
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) (ans int) {
		var n, needs int
		Fscan(in, &n, &needs)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)

		// 把每块薄片等分地切 [0,needs-1] 刀，对应地得到 [1,needs] 个薄片，最终的目标尺寸一定是其中一个
		mp := map[pair][]int{}
		for i, angle := range a {
			for cut := 1; cut <= needs; cut++ {
				sz := frac(angle, cut)
				mp[sz] = append(mp[sz], i)
			}
		}

		ans = 1e9
		used := make([]bool, n)
		// 对每个目标尺寸，我们检查所有 ≥ 目标尺寸的薄片：
		// 当薄片的尺寸是目标尺寸的倍数时，等分地切 k 刀会得到 k+1 个有效薄片，否则切 k 刀只能得到 k 个有效薄片
		// 在没超出 needs 个薄片前，这样的薄片相当于节省了一刀
		for sz, ids := range mp {
			x, y, made, cuts := sz.x, sz.y, 0, needs
			for _, i := range ids {
				md := a[i] * y / x
				if made+md <= needs {
					cuts--
				}
				made += md
				used[i] = true
			}
			// 检查剩余薄片能否凑足 needs 个薄片
			// 每个薄片能切分出 a[i]/(x/y) = a[i]*y/x 个
			for i := n - 1; i >= 0 && made < needs && a[i]*y >= x; i-- {
				if !used[i] {
					made += a[i] * y / x
				}
			}
			if made >= needs && cuts < ans {
				ans = cuts
			}
			for _, i := range ids {
				used[i] = false
			}
		}
		return
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: %d\n", _case, solve(_case))
	}
}

func main() { run(os.Stdin, os.Stdout) }
