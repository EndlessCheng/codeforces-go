package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func cf1736C2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, p, x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := n * (n + 1) / 2
	right := make([]int, n)
	right2 := make([]int, n)
	for i := range right {
		right[i] = n
		right2[i] = n
	}
	var q, q2 []int // 双队列 todo 简化成变量
	for i, v := range a {
		v -= i + 1
		for len(q2) > 0 && -q2[0] > v {
			right2[q2[0]] = i // q2 队首的下下个更小元素是 a[i]
			q2 = q2[1:]
		}
		for len(q) > 0 && -q[0] > v {
			ans -= n - i
			right[q[0]] = i // q 队首的下一个更小元素是 a[i]
			q2 = append(q2, q[0])
			q = q[1:]
		}
		q = append(q, i)
	}

	sumR := make([]int, n+1)
	sumInc := make([]int, n+1)
	for i, v := range right {
		sumR[i+1] = sumR[i] + v
		sumInc[i+1] = sumInc[i] + right2[i] - v
	}

	Fscan(in, &m)
	for range m {
		Fscan(in, &p, &x)
		p--
		// 变为 tar
		tar := x - p - 1
		l := sort.SearchInts(right, p+1) // 第一个包含 p 的区间
		if x == a[p] {
			Fprintln(out, ans)
		} else if x < a[p] { // 变小
			if tar >= -l {
				Fprintln(out, ans) // 无影响
			} else {
				r := -tar
				// 左端点为 l,l+1,...,r-1 的区间的右开端点都从 right[i] 缩小为 p
				Fprintln(out, ans-(sumR[r]-sumR[l]-(r-l)*p))
			}
		} else {
			// 变大
			if p == 0 || right[p-1] < p-1 {
				Fprintln(out, ans) // 无影响
			} else {
				ll := max(sort.SearchInts(right, p), -tar)
				// ll,ll+1,...,l-1 都是受到 a[p] 影响的区间
				Fprintln(out, ans+sumInc[l]-sumInc[ll])
			}
		}
	}
}

func main() { cf1736C2(bufio.NewReader(os.Stdin), os.Stdout) }
