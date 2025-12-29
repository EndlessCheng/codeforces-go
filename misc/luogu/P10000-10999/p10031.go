package main

import (
	. "fmt"
	"io"
)

/*
## 方法一：GCD 的性质

由于 gcd(n,i) = gcd(n,n-i)（更相减损术），所以 gcd(n,i) XOR gcd(n,n-i) = 0，两两抵消。

当 n 是奇数时，最终剩下 gcd(n,n) = n。
当 n 是偶数时，最终剩下 gcd(n,n/2) 和 gcd(n,n)，异或结果为 n/2 XOR n。

## 方法二：欧拉函数的奇偶性

如果 gcd(n,i) = d，那么 n 和 i 都是 d 的倍数。
这可以写成 gcd(n,k*d) = d，也就是 gcd(n/d,k) = 1，即 k 与 n/d 互质。
1 到 n/d 中有 φ(n/d) 个与 n/d 互质的数。

答案可以表示为：计算 d 的异或和，其中 d 是 n 的因子，且 φ(n/d) 为奇数。（偶数个 d 异或为 0）
变量代换，令 m = n/d，答案表示为：计算 n/m 的异或和，其中 m 是 n 的因子，且 φ(m) 为奇数。

哪些 φ(m) 是奇数？
如果 m 含有奇质数 p，那么 φ(m) 式子包含 p-1，φ(m) 为偶数。
如果 m = 2^k 且 k >= 2，那么 φ(2^k) = 2^(k-1) 是偶数。
所以只有 φ(1) 和 φ(2) 是奇数，得到和方法一同样的结论。
*/

// https://space.bilibili.com/206214
func p10031(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 == 0 {
			n ^= n / 2
		}
		Fprintln(out, n)
	}
}

//func main() { p10031(bufio.NewReader(os.Stdin), os.Stdout) }
