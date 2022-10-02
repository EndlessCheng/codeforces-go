下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

枚举因子，挨个判断能否整除 $a$ 和 $b$。

改进方案是枚举 $a$ 和 $b$ 的最大公因数的因子。

```py [sol1-Python3]
class Solution:
    def commonFactors(self, a: int, b: int) -> int:
        g = gcd(a, b)
        ans, i = 0, 1
        while i * i <= g:
            if g % i == 0:
                ans += 1
                if i * i < g:
                    ans += 1
            i += 1
        return ans
```

```go [sol1-Go]
func commonFactors(a, b int) (ans int) {
	g := gcd(a, b)
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(\sqrt{\min(a,b)})$。
- 空间复杂度：$O(1)$，仅用到若干变量。
