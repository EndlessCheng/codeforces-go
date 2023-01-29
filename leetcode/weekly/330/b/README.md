下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

正难则反，只有全部顺时针和全部逆时针才不会碰撞。

因此答案为 $2^n-2$。用快速幂计算。

注意为了避免负数，需要在 $-2$ 后再转换到非负数上。

```py [sol1-Python3]
class Solution:
    def monkeyMove(self, n: int) -> int:
        MOD = 10 ** 9 + 7
        return (pow(2, n, MOD) - 2) % MOD
```

```go [sol1-Go]
const mod int = 1e9 + 7

func monkeyMove(n int) int {
	return (pow(2, n) - 2 + mod) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

### 复杂度分析

- 时间复杂度：$O(\log n)$。
- 空间复杂度：$O(1)$，仅用到若干变量。
