下午 2 点在 B 站直播讲周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

## 方法一：刷表法

根据题意，有两类人：

- A 类：知道秘密，但还不能分享；
- B 类：知道秘密，且可以分享。

定义 $f[i]$ 表示第 $i$ 天的 B 类人数，则有初始值 $f[1]=1$。

转移时，我们可以从 $f[i]$ 出发，把天数在区间 $[i+\textit{delay}, i+\textit{forget})$ 的 B 类人数都加上 $f[i]$。此外，如果 $i+\textit{delay}\ge n$，则可以把第 $n$ 天的 A 类人数（记作 $\textit{cntA}$）也加上 $f[i]$。

答案为 $\textit{cntA}+f[n]$。

代码实现时，下标可以从 $0$ 开始，从而优化掉一些 `+1` 的逻辑。

#### 复杂度分析

- 时间复杂度：$O(n^2)$。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def peopleAwareOfSecret(self, n: int, delay: int, forget: int) -> int:
        MOD = 10 ** 9 + 7
        cnt_a = 0
        f = [0] * n
        f[0] = 1
        for i in range(n):
            if i + delay >= n:
                cnt_a = (cnt_a + f[i]) % MOD
            for j in range(i + delay, min(i + forget, n)):
                f[j] = (f[j] + f[i]) % MOD
        return (cnt_a + f[-1]) % MOD
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7;

    public int peopleAwareOfSecret(int n, int delay, int forget) {
        var cntA = 0;
        var f = new int[n];
        f[0] = 1;
        for (var i = 0; i < n; ++i) {
            if (i + delay >= n) cntA = (cntA + f[i]) % MOD;
            for (int j = i + delay; j < Math.min(i + forget, n); ++j)
                f[j] = (f[j] + f[i]) % MOD;
        }
        return (cntA + f[n - 1]) % MOD;
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
public:
    int peopleAwareOfSecret(int n, int delay, int forget) {
        int cnt_a = 0;
        int f[n]; memset(f, 0, sizeof(f));
        f[0] = 1;
        for (int i = 0; i < n; ++i) {
            if (i + delay >= n) cnt_a = (cnt_a + f[i]) % MOD;
            for (int j = i + delay; j < min(i + forget, n); ++j)
                f[j] = (f[j] + f[i]) % MOD;
        }
        return (cnt_a + f[n - 1]) % MOD;
    }
};
```

```go [sol1-Go]
func peopleAwareOfSecret(n, delay, forget int) int {
	const mod int = 1e9 + 7
	cntA := 0
	f := make([]int, n)
	f[0] = 1
	for i, v := range f {
		if i+delay >= n {
			cntA = (cntA + v) % mod
		}
		for j := i + delay; j < i+forget && j < n; j++ {
			f[j] = (f[j] + v) % mod
		}
	}
	return (cntA + f[n-1]) % mod
}
```

## 方法二：填表法

另一种方法是把 $f[i]$ 定义成第 $i$ 天新增的知道秘密的人数，则有初始值 $f[1]=1$。

根据题意，我们可以从天数在 $[i-\textit{forget}+1, i-\textit{delay}]$ 内的 $f[j]$ 转移过来，即

$$
f[i] = \sum_{j=i-\textit{forget}+1}^{i-\textit{delay}} f[j]
$$

答案为天数在 $[n-\textit{forget}+1, n]$ 内的新增的知道秘密的人数之和，即

$$
\sum_{i=n-\textit{forget}+1}^{n} f[i]
$$

代码实现时，和式可以用前缀和优化。有了前缀和，$f$ 数组也可以省略。

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(n)$。

```py [sol1-Python3]
class Solution:
    def peopleAwareOfSecret(self, n: int, delay: int, forget: int) -> int:
        MOD = 10 ** 9 + 7
        sum = [0] * (n + 1)
        sum[1] = 1
        for i in range(2, n + 1):
            f = sum[max(i - delay, 0)] - sum[max(i - forget, 0)]
            sum[i] = (sum[i - 1] + f) % MOD
        return (sum[n] - sum[max(0, n - forget)]) % MOD
```

```java [sol1-Java]
class Solution {
    static final int MOD = (int) 1e9 + 7;

    public int peopleAwareOfSecret(int n, int delay, int forget) {
        var sum = new int[n + 1];
        sum[1] = 1;
        for (var i = 2; i <= n; i++) {
            var f = (sum[Math.max(i - delay, 0)] - sum[Math.max(i - forget, 0)]) % MOD;
            sum[i] = (sum[i - 1] + f) % MOD;
        }
        return ((sum[n] - sum[Math.max(0, n - forget)]) % MOD + MOD) % MOD; // 防止结果为负数
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;
public:
    int peopleAwareOfSecret(int n, int delay, int forget) {
        int sum[n + 1];
        sum[0] = 0, sum[1] = 1;
        for (int i = 2; i <= n; ++i) {
            int f = (sum[max(i - delay, 0)] - sum[max(i - forget, 0)]) % MOD;
            sum[i] = (sum[i - 1] + f) % MOD;
        }
        return ((sum[n] - sum[max(0, n - forget)]) % MOD + MOD) % MOD; // 防止结果为负数
    }
};
```

```go [sol1-Go]
func peopleAwareOfSecret(n, delay, forget int) int {
	const mod int = 1e9 + 7
	sum := make([]int, n+1)
	sum[1] = 1
	for i := 2; i <= n; i++ {
		f := (sum[max(i-delay, 0)] - sum[max(i-forget, 0)]) % mod
		sum[i] = (sum[i-1] + f) % mod
	}
	return ((sum[n]-sum[max(0, n-forget)])%mod + mod) % mod // 防止结果为负数
}

func max(a, b int) int { if b > a { return b }; return a }
```


