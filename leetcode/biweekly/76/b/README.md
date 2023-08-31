枚举买了 $i$ 支钢笔。

至多能买 $\left\lfloor\dfrac{\textit{total}}{\textit{cost}_1}\right\rfloor$ 支钢笔。

剩余钱数为 $\textit{total} - i * \textit{cost}_1$，至少可以买 $0$ 支铅笔，至多可以买 $\left\lfloor\dfrac{\textit{total} - i * \textit{cost}_1}{\textit{cost}_2}\right\rfloor$ 支铅笔，这一共有 

$$
1 + \left\lfloor\dfrac{\textit{total} - i * \textit{cost}_1}{\textit{cost}_2}\right\rfloor
$$ 

种不同的购买方案。

所以答案为

$$
\sum_{i=0}^{\lfloor \textit{total}/\textit{cost}_1\rfloor} 1 + \left\lfloor\dfrac{\textit{total} - i * \textit{cost}_1}{\textit{cost}_2}\right\rfloor
$$

即

$$
1+\left\lfloor\dfrac{\textit{total}}{\textit{cost}_1}\right\rfloor + \sum_{i=0}^{\lfloor \textit{total}/\textit{cost}_1\rfloor} \left\lfloor\dfrac{\textit{total} - i * \textit{cost}_1}{\textit{cost}_2}\right\rfloor
$$

```py [sol-Python3]
class Solution:
    def waysToBuyPensPencils(self, total: int, cost1: int, cost2: int) -> int:
        n = 1 + total // cost1
        return n + sum((total - cost1 * i) // cost2 for i in range(n))
```

```java [sol-Java]
class Solution {
    public long waysToBuyPensPencils(int total, int cost1, int cost2) {
        long n = 1 + total / cost1, ans = n;
        for (long i = 0; i < n; i++)
            ans += (total - cost1 * i) / cost2;
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long waysToBuyPensPencils(int total, int cost1, int cost2) {
        long long n = 1 + total / cost1, ans = n;
        for (long long i = 0; i < n; i++)
            ans += (total - cost1 * i) / cost2;
        return ans;
    }
};
```

```go [sol-Go]
func waysToBuyPensPencils(total, cost1, cost2 int) int64 {
	n := 1 + total/cost1
	ans := int64(n)
	for i := 0; i < n; i++ {
		ans += int64((total - cost1*i) / cost2)
	}
	return ans
}
```

```js [sol-JavaScript]
var waysToBuyPensPencils = function(total, cost1, cost2) {
    const n = 1 + Math.floor(total / cost1);
    let ans = n;
    for (let i = 0; i < n; i++)
        ans += Math.floor((total - cost1 * i) / cost2);
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(\left\lfloor\dfrac{\textit{total}}{\textit{cost}_1}\right\rfloor\right)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 附：类欧几里得算法

上面的和式可以直接套用类欧几里德算法解决。

参考：[类欧几里德算法](https://oi-wiki.org/math/number-theory/euclidean/)。

```go
func waysToBuyPensPencils(total, cost1, cost2 int) int64 {
	n := total/cost1 + 1
	return int64(n + floorSum(n, cost2, -cost1, total))
}

// 返回 sum(floor((a*i+b)/m)), i 从 0 到 n-1
func floorSum(n, m, a, b int) (res int) {
	if a < 0 {
		a2 := a%m + m
		res -= n * (n - 1) / 2 * ((a2 - a) / m)
		a = a2
	}
	if b < 0 {
		b2 := b%m + m
		res -= n * ((b2 - b) / m)
		b = b2
	}
	for {
		if a >= m {
			res += n * (n - 1) / 2 * (a / m)
			a %= m
		}
		if b >= m {
			res += n * (b / m)
			b %= m
		}
		yMax := a*n + b
		if yMax < m {
			break
		}
		n = yMax / m
		b = yMax % m
		m, a = a, m
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(\log\left\lfloor\dfrac{\textit{total}}{\textit{cost}_1}\right\rfloor\right)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

[我的其它题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
