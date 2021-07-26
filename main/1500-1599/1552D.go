package main

import (
	. "fmt"
	"io"
)

/*
由于 j 和 k 没有顺序关系
我们可以将 a 中的部分元素取反，记其为 a'
设 b[0]=0，构造 b 为 a'[:n-1] 的前缀和
若 a'[:n-1] 中存在一个子区间，其区间和等于 a'[n-1]
由于 b 就是 a'[:n-1] 的前缀和，所以 a'[n-1] 也能用 b 的两个元素相减表示
所以 a' 需要满足 a'[l]+...+a'[r] = a'[n-1]
移项得 a'[l]+...+a'[r] - a'[n-1] = 0
实际上，由于 a' 元素的顺序没有关系，这其实就是 a' 的一个子集，其元素和为 0
因此得出结论：需要在 a 中找到一个子集，该子集的部分元素取反后，能够使该子集元素和为 0，若能找到这样的子集则输出 YES，否则输出 NO

进一步地，我们只需要判断 a 中是否存在两个子集和相等就行了
设这两个子集为 A 和 B，则有 A[0]+A[1]+... = B[0]+B[1]+...
移项得 A[0]+A[1]+... -B[0]-B[1]-... = 0
注意，若两个子集有相同的元素，会在上式中消去
消去后的剩余部分就是 a 的一个子集，其中部分元素取反
*/

// github.com/EndlessCheng/codeforces-go
func CF1552D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		has := map[int]bool{}
		for i := 0; i < 1<<n; i++ {
			sum := 0
			for j, v := range a {
				sum += i >> j & 1 * v
			}
			has[sum] = true
		}
		if len(has) < 1<<n {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1552D(os.Stdin, os.Stdout) }
