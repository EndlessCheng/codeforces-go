提供一个纯树状数组做法。

用两棵树状数组，一棵维护 $a$ 的前缀和，另一棵（记作 $\textit{zeroT}$）维护 $a$ 的前缀中的 $0$ 的个数。

分类讨论：

- 如果 $[\ell,r]$ 中没有 $0$，那就是普通树状数组求子数组和。
- 否则，设 $i$ 是 $[\ell,r]$ 中最后一个 $0$ 的下标，用树状数组算出 $[\ell,i]$ 的和，以及 $[i,r]$ 的和（注意 $a[i]=0$，不影响和）。

如何求出最后一个 $0$ 的下标？设 $[1,r]$ 中的 $0$ 的个数为 $k$，在 $\textit{zeroT}$ 中做树状数组二分（$\texttt{kth}$），即可求出第 $k$ 个 $0$ 的下标，原理见 [OI Wiki](https://oi-wiki.org/ds/fenwick/#%E5%8D%95%E7%82%B9%E4%BF%AE%E6%94%B9%E6%9F%A5%E8%AF%A2%E5%85%A8%E5%B1%80%E7%AC%AC-k-%E5%B0%8F)。

AC 代码（Golang）：

```go
package main
import("bufio";."fmt"; "os")

type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func (t fenwick) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func (t fenwick) kth(k int) int {
	i := 0
	for bit := 1 << 17; bit > 0; bit >>= 1 {
		nxt := i | bit
		if nxt < len(t) && t[nxt] < k {
			k -= t[nxt]
			i = nxt
		}
	}
	return i + 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q, x, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	t := make(fenwick, n+1)
	zeroT := make(fenwick, n+1)
	// O(n) 建树
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		t[i] += a[i]
		if a[i] == 0 {
			zeroT[i]++
		}
		j := i + i&-i
		if j <= n {
			t[j] += t[i]
			zeroT[j] += zeroT[i]
		}
	}

	const mod = 998244353
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			if a[l] == 0 {
				zeroT.update(l, -1)
			}
			if r == 0 {
				zeroT.update(l, 1)
			}
			t.update(l, r-a[l])
			a[l] = r
		} else {
			Fscan(in, &x)
			lz, rz := zeroT.pre(l-1), zeroT.pre(r)
			if lz == rz { // 区间没有 0
				Fprintln(out, x+t.query(l, r))
			} else {
				i := zeroT.kth(rz) // 区间最后一个 0 的下标
				Fprintln(out, (x+t.query(l, i))%mod+t.query(i, r))
			}
		}
	}
}
```

