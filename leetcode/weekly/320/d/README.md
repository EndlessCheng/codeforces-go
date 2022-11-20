定义 $f[i][j]$ 表示把 $s$ 的前 $j$ 个字符分割成 $i$ 段的方案数（每段需要满足题目的后两个要求）。

如果 $s[j]$ 不是质数，那么可以考虑枚举第 $i$ 段的起始位置 $j'$，如果 $j-j'+1\ge \textit{minLength}$ 且 $s[j']$ 是质数，那么就可以分割。

累加所有可以分割的 $f[i-1][j']$，记作 $\textit{sum}$，那么 $f[i][j]=\textit{sum}$。

我们可以一边计算 $\textit{sum}$，一边更新 $f[i][j]$，具体见代码。

初始值 $f[0][0] = 1$，表示空串的 $0$ 个分割算作一种方案。

答案为 $f[k][n]$，这里 $n$ 为 $s$ 的长度。

还有一些剪枝和循环次数优化的小技巧，具体见代码。

```py [sol1-Python3]
class Solution:
    def beautifulPartitions(self, s: str, k: int, l: int) -> int:
        MOD = 10 ** 9 + 7
        def is_prime(c: str) -> bool:
            return c in "2357"
        # 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
        def can_partition(j: int) -> bool:
            return j == 0 or j == n or not is_prime(s[j - 1]) and is_prime(s[j])

        n = len(s)
        if k * l > n or not is_prime(s[0]) or is_prime(s[-1]):  # 剪枝
            return 0
        f = [[0] * (n + 1) for _ in range(k + 1)]
        f[0][0] = 1
        for i in range(1, k + 1):
            sum = 0
            # 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
            for j in range(i * l, n - (k - i) * l + 1):
                if can_partition(j - l): sum = (sum + f[i - 1][j - l]) % MOD  # 长度至少是 l
                if can_partition(j): f[i][j] = sum
        return f[k][n]
```

```java [sol1-Java]
class Solution {
    private static final int MOD = (int) 1e9 + 7;

    public int beautifulPartitions(String S, int k, int l) {
        var s = S.toCharArray();
        var n = s.length;
        if (k * l > n || !isPrime(s[0]) || isPrime(s[n - 1])) // 剪枝
            return 0;
        var f = new int[k + 1][n + 1];
        f[0][0] = 1;
        for (var i = 1; i <= k; ++i) {
            var sum = 0;
            // 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
            for (var j = i * l; j + (k - i) * l <= n; j++) {
                if (canPartition(s, j - l)) sum = (sum + f[i - 1][j - l]) % MOD; // 长度至少是 l
                if (canPartition(s, j)) f[i][j] = sum;
            }
        }
        return f[k][n];
    }

    private boolean isPrime(char c) {
        return c == '2' || c == '3' || c == '5' || c == '7';
    }

    // 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
    private boolean canPartition(char[] s, int j) {
        return j == 0 || j == s.length || !isPrime(s[j - 1]) && isPrime(s[j]);
    }
}
```

```cpp [sol1-C++]
class Solution {
    const int MOD = 1e9 + 7;

    bool is_prime(char c) {
        return c == '2' || c == '3' || c == '5' || c == '7';
    }

    // 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
    bool can_partition(string &s, int j) {
        return j == 0 || j == s.length() || !is_prime(s[j - 1]) && is_prime(s[j]);
    }

public:
    int beautifulPartitions(string &s, int k, int l) {
        int n = s.length();
        if (k * l > n || !is_prime(s[0]) || is_prime(s[n - 1])) // 剪枝
            return 0;
        int f[k + 1][n + 1]; memset(f, 0, sizeof(f));
        f[0][0] = 1;
        for (int i = 1; i <= k; ++i) {
            int sum = 0;
            // 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
            for (int j = i * l; j + (k - i) * l <= n; j++) {
                if (can_partition(s, j - l)) sum = (sum + f[i - 1][j - l]) % MOD; // 长度至少是 l
                if (can_partition(s, j)) f[i][j] = sum;
            }
        }
        return f[k][n];
    }
};
```

```go [sol1-Go]
func beautifulPartitions(s string, k, l int) (ans int) {
	const mod int = 1e9 + 7
	isPrime := func(c byte) bool { return strings.IndexByte("2357", c) >= 0 }
	n := len(s)
	if k*l > n || !isPrime(s[0]) || isPrime(s[n-1]) { // 剪枝
		return
	}
	// 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
	canPartition := func(j int) bool { return j == 0 || j == n || !isPrime(s[j-1]) && isPrime(s[j]) }
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i := 1; i <= k; i++ {
		sum := 0
		// 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
		for j := i * l; j+(k-i)*l <= n; j++ {
			if canPartition(j - l) { // 长度至少是 l
				sum = (sum + f[i-1][j-l]) % mod
			}
			if canPartition(j) {
				f[i][j] = sum
			}
		}
	}
	return f[k][n]
}
```

#### 复杂度分析

- 时间复杂度：$O(kn)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(kn)$。注：也可以用滚动数组优化至 $O(n)$。
