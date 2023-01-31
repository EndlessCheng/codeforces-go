正难则反，只有全部顺时针和全部逆时针才不会碰撞。

在不考虑碰撞时，由于每个猴子都可以往两边走，那么总共有 $2^n$ 种移动方法。

答案所有移动方法减去不会碰撞的移动方法，即 $2^n-2$。用快速幂计算。

> 不了解快速幂的同学可以看看 [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)。

注意为了避免负数，需要在 $-2$ 后再转换到非负数上。

附：[视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)

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
