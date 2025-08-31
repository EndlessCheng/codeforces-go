### 分析

设 $s$ 中有 $z$ 个 $0$，设一共操作了 $m$ 次。那么总翻转次数为 $mk$。

这 $z$ 个 $0$ 必须翻转奇数次，其余 $n-z$ 个 $1$ 必须翻转偶数次。

总翻转次数减去 $z$，剩下每个位置都必须翻转偶数次，所以 

$$
mk-z\ 是偶数
$$

下面计算 $m$ 的下界。只要能证明 $m$ 可以等于下界，问题就解决了。

要想把 $z$ 个 $0$ 变成 $1$，总翻转次数至少要是 $z$，即

$$
mk\ge z
$$

即

$$
m\ge \left\lceil\dfrac{z}{k}\right\rceil
$$

除此以外，还需要满足什么要求？

### 情况一：m 是偶数

由于 $mk-z$ 是偶数，如果 $m$ 是偶数，那么 $z$ 也必须是偶数。

$s$ 中的每个位置至多翻转 $m$ 次。但是，对于 $s$ 中的 $0$，由于要翻转奇数次，所以至多翻转 $m-1$ 次。

所以 $s$ 中的所有位置的翻转次数的上界是 $z(m-1)+(n-z)m$，其可能等于 $mk$，也可能比 $mk$ 大（因为是上界），所以有

$$
z(m-1)+(n-z)m\ge mk
$$

解得

$$
m\ge \left\lceil\dfrac{z}{n-k}\right\rceil
$$

与

$$
m\ge \left\lceil\dfrac{z}{k}\right\rceil
$$

联立得

$$
m\ge \max\left(\left\lceil\dfrac{z}{k}\right\rceil,\left\lceil\dfrac{z}{n-k}\right\rceil\right)
$$

### 情况二：m 是奇数

由于 $mk-z$ 是偶数，如果 $m$ 是奇数，那么 $z$ 和 $k$ 必须同为奇数，或者同为偶数（奇偶性相同）。

$s$ 中的每个位置至多翻转 $m$ 次。但是，对于 $s$ 中的 $1$，由于要翻转偶数次，所以至多翻转 $m-1$ 次。

所以 $s$ 中的所有位置的翻转次数的上界是 $zm+(n-z)(m-1)$，其可能等于 $mk$，也可能比 $mk$ 大（因为是上界），所以有

$$
zm+(n-z)(m-1)\ge mk
$$

解得

$$
m\ge \left\lceil\dfrac{n-z}{n-k}\right\rceil
$$

与

$$
m\ge \left\lceil\dfrac{z}{k}\right\rceil
$$

联立得

$$
m\ge \max\left(\left\lceil\dfrac{z}{k}\right\rceil,\left\lceil\dfrac{n-z}{n-k}\right\rceil\right)
$$

情况一和情况二取最小值。

如果两个情况都不满足要求，返回 $-1$。

### 下界可以取到

这可以用 **Gale-Ryser 定理**证明。

具体来说，我们需要判断是否存在一个 $m$ 行 $n$ 列的 $0\text{-}1$ 矩阵 $M$，第 $i$ 行对应着第 $i$ 次操作，其中 $M_{i,j} = 0$ 表示没有翻转 $s_j$，$M_{i,j} = 1$ 表示翻转 $s_j$。每一行的元素和都是 $k$，第 $j$ 列的元素和是 $s_j$ 的翻转次数 $a_j$。由于 $a_j\le m$ 且 $\sum\limits_{j} a_j\le mk$，由 Gale-Ryser 定理可得，这样的矩阵是存在的。

### 特殊情况

如果 $z=0$，那么无需操作，答案是 $0$。

由于下界公式中的分母 $n-k$ 不能为 $0$，我们需要特判 $n=k$ 的情况，此时每次操作只能翻转整个 $s$。

- 如果 $z=n$，即 $s$ 全为 $0$，那么只需操作 $1$ 次。
- 否则无论怎么操作，$s$ 中始终有 $0$，返回 $-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1aCaGzWEm4/?t=14m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str, k: int) -> int:
        n = len(s)
        z = s.count('0')
        if z == 0:
            return 0
        if n == k:
            return 1 if z == n else -1

        ans = inf
        # 情况一：操作次数 m 是偶数
        if z % 2 == 0:  # z 必须是偶数
            m = max((z + k - 1) // k, (z + n - k - 1) // (n - k))  # 下界
            ans = m + m % 2  # 把 m 往上调整为偶数

        # 情况二：操作次数 m 是奇数
        if z % 2 == k % 2:  # z 和 k 的奇偶性必须相同
            m = max((z + k - 1) // k, (n - z + n - k - 1) // (n - k))  # 下界
            ans = min(ans, m | 1)  # 把 m 往上调整为奇数

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minOperations(String s, int k) {
        int n = s.length();
        int z = 0;
        for (int i = 0; i < n; i++) {
            if (s.charAt(i) == '0') {
                z++;
            }
        }

        if (z == 0) {
            return 0;
        }
        if (n == k) {
            return z == n ? 1 : -1;
        }

        int ans = Integer.MAX_VALUE;
        // 情况一：操作次数 m 是偶数
        if (z % 2 == 0) { // z 必须是偶数
            int m = Math.max((z + k - 1) / k, (z + n - k - 1) / (n - k)); // 下界
            ans = m + m % 2; // 把 m 往上调整为偶数
        }

        // 情况二：操作次数 m 是奇数
        if (z % 2 == k % 2) { // z 和 k 的奇偶性必须相同
            int m = Math.max((z + k - 1) / k, (n - z + n - k - 1) / (n - k)); // 下界
            ans = Math.min(ans, m | 1); // 把 m 往上调整为奇数
        }

        return ans < Integer.MAX_VALUE ? ans : -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s, int k) {
        int n = s.size();
        int z = ranges::count(s, '0');
        if (z == 0) {
            return 0;
        }
        if (n == k) {
            return z == n ? 1 : -1;
        }

        int ans = INT_MAX;
        // 情况一：操作次数 m 是偶数
        if (z % 2 == 0) { // z 必须是偶数
            int m = max((z + k - 1) / k, (z + n - k - 1) / (n - k)); // 下界
            ans = m + m % 2; // 把 m 往上调整为偶数
        }

        // 情况二：操作次数 m 是奇数
        if (z % 2 == k % 2) { // z 和 k 的奇偶性必须相同
            int m = max((z + k - 1) / k, (n - z + n - k - 1) / (n - k)); // 下界
            ans = min(ans, m | 1); // 把 m 往上调整为奇数
        }

        return ans < INT_MAX ? ans : -1;
    }
};
```

```go [sol-Go]
func minOperations(s string, k int) int {
	n := len(s)
	z := strings.Count(s, "0")
	if z == 0 {
		return 0
	}
	if n == k {
		if z == n {
			return 1
		}
		return -1
	}

	ans := math.MaxInt
	// 情况一：操作次数 m 是偶数
	if z%2 == 0 { // z 必须是偶数
		m := max((z+k-1)/k, (z+n-k-1)/(n-k)) // 下界
		ans = m + m%2 // 把 m 往上调整为偶数
	}

	// 情况二：操作次数 m 是奇数
	if z%2 == k%2 { // z 和 k 的奇偶性必须相同
		m := max((z+k-1)/k, (n-z+n-k-1)/(n-k)) // 下界
		ans = min(ans, m|1) // 把 m 往上调整为奇数
	}

	if ans < math.MaxInt {
		return ans
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。瓶颈在遍历 $s$ 上，如果已知 $0$ 的个数，则时间复杂度是 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
