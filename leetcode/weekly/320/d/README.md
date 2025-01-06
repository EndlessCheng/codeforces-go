[视频讲解](https://www.bilibili.com/video/BV1A3411f7H3/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

定义 $f[i][j]$ 表示把 $s$ 的前 $j$ 个字符分割成 $i$ 段的方案数（每段需要满足题目的后两个要求）。

定义 $j$ 为分割点，等价于 $s[j]$ 不是质数且 $s[j+1]$ 是质数。

如果 $j$ 是分割点，那么可以考虑枚举第 $i-1$ 段与第 $i$ 段的分割点 $j'$，需满足 $j-j'\ge \textit{minLength}$。

累加所有 $f[i-1][j']$，记作 $\textit{sum}$，那么 $f[i][j]=\textit{sum}$。

每个 $f[i][j]$ 都要这样累加就太慢了，需要优化。

我们可以从小到大同时遍历 $j'$ 和 $j$，一边更新 $\textit{sum}$，一边计算 $f[i][j]$，具体见代码。

为方便计算，定义初始值 $f[0][0] = 1$，表示空串的 $0$ 个分割算作一种方案。**因为这个原因，要把所有下标 $j$ 向后移动一位。**

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
                if can_partition(j - l): sum = (sum + f[i - 1][j - l]) % MOD  # j'=j-l 双指针
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
                if (canPartition(s, j - l)) sum = (sum + f[i - 1][j - l]) % MOD; // j'=j-l 双指针
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
                if (can_partition(s, j - l)) sum = (sum + f[i - 1][j - l]) % MOD; // j'=j-l 双指针
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
			if canPartition(j - l) { // j'=j-l 双指针
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

- 时间复杂度：$O(k(n-kl))$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(kn)$。注：也可以用滚动数组优化至 $O(n)$。

#### 相似题目

- [1977. 划分数字的方案数](https://leetcode.cn/problems/number-of-ways-to-separate-numbers/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
