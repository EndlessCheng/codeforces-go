设 $m$ 为 $n$ 的二进制长度。

对于二进制长度小于 $m$ 的回文数，其左半边可以随便填：把左半边翻转，得到右半边，把左半右半拼起来，就是一个回文数了。

注意当回文数长为奇数时，回文中心也可以随便填，算到左半边中。

枚举二进制长度 $i=1,2,\dots,m-1$，左半的长度为 $\left\lceil\dfrac{i}{2}\right\rceil$。由于最高位一定填 $1$，所以可以随便填的位置有 $k = \left\lceil\dfrac{i}{2}\right\rceil - 1$ 个，每个位置填 $0$ 还是 $1$ 都可以，得到 $2^k$ 个回文数。

对于二进制长度等于 $m$ 的回文数：

- 左半比特位的编号，从高到低为 $m-1,m-2,\dots, \left\lfloor\dfrac{m}{2}\right\rfloor$。
- 最高位填 $1$。
- 对于左半的其余位，枚举 $i=m-2,\dots, \left\lfloor\dfrac{m}{2}\right\rfloor$，从高到低思考：
   - 如果 $n$ 的这个位是 $1$，那么回文数这个位填 $0$ 后，回文数左半边的剩余位置可以随便填。设 $k = i - \left\lfloor\dfrac{m}{2}\right\rfloor$，得到 $2^k$ 个回文数。
   - 然后，只需考虑回文数这个位填 $1$ 的情况，继续循环。

最后，判断 $n$ 的左半边翻转后得到的回文数是否 $\le n$，如果是，把答案加一。

注意 $0$ 也是回文数。

特判 $n=0$ 的情况，返回 $1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countBinaryPalindromes(self, n: int) -> int:
        if n == 0:
            return 1

        m = n.bit_length()  # n 的二进制长度

        # 二进制长度小于 m，随便填
        ans = 1  # 0 也是回文数
        # 枚举二进制长度，最高位填 1，回文数左半的其余位置随便填
        for i in range(1, m):
            ans += 1 << ((i - 1) // 2)

        # 最高位一定是 1，从次高位开始填
        for i in range(m - 2, m // 2 - 1, -1):
            if n >> i & 1:
                # 这一位可以填 0，那么回文数左半的剩余位置可以随便填
                ans += 1 << (i - m // 2)
            # 在后续循环中，这一位填 1

        pal = n >> (m // 2)
        # 左半反转到右半
        # 如果 m 是奇数，那么去掉回文中心再反转
        v = pal >> (m % 2)
        while v:
            pal = pal * 2 + v % 2
            v //= 2
        if pal <= n:
            ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    public int countBinaryPalindromes(long n) {
        if (n == 0) {
            return 1;
        }

        // n 的二进制长度
        int m = 64 - Long.numberOfLeadingZeros(n);

        // 二进制长度小于 m，随便填
        int ans = 1; // 0 也是回文数
        // 枚举二进制长度，最高位填 1，回文数左半的其余位置随便填
        for (int i = 1; i < m; i++) {
            ans += 1 << ((i - 1) / 2);
        }

        // 最高位一定是 1，从次高位开始填
        for (int i = m - 2; i >= m / 2; i--) {
            if ((n >> i & 1) > 0) {
                // 这一位可以填 0，那么回文数左半的剩余位置可以随便填
                ans += 1 << (i - m / 2);
            }
            // 在后续循环中，这一位填 1
        }

        long pal = n >> (m / 2);
        // 左半反转到右半
        // 如果 m 是奇数，那么去掉回文中心再反转
        for (long v = pal >> (m % 2); v > 0; v /= 2) {
            pal = pal * 2 + v % 2;
        }
        if (pal <= n) {
            ans++;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countBinaryPalindromes(long long n) {
        if (n == 0) {
            return 1;
        }

        int m = bit_width((uint64_t) n); // n 的二进制长度

        // 二进制长度小于 m，随便填
        int ans = 1; // 0 也是回文数
        // 枚举二进制长度，最高位填 1，回文数左半的其余位置随便填
        for (int i = 1; i < m; i++) {
            ans += 1 << ((i - 1) / 2);
        }

        // 最高位一定是 1，从次高位开始填
        for (int i = m - 2; i >= m / 2; i--) {
            if (n >> i & 1) {
                // 这一位可以填 0，那么回文数左半的剩余位置可以随便填
                ans += 1 << (i - m / 2);
            }
            // 在后续循环中，这一位填 1
        }

        long long pal = n >> (m / 2);
        // 左半反转到右半
        // 如果 m 是奇数，那么去掉回文中心再反转
        for (long long v = pal >> (m % 2); v > 0; v /= 2) {
            pal = pal * 2 + v % 2;
        }
        if (pal <= n) {
            ans++;
        }

        return ans;
    }
};
```

```go [sol-Go]
func countBinaryPalindromes(n int64) int {
	if n == 0 {
		return 1
	}

	m := bits.Len(uint(n)) // n 的二进制长度

	// 二进制长度小于 m，随便填
	ans := 1 // 0 也是回文数
	// 枚举二进制长度，最高位填 1，回文数左半的其余位置随便填
	for i := 1; i < m; i++ {
		ans += 1 << ((i - 1) / 2)
	}

	// 最高位一定是 1，从次高位开始填
	for i := m - 2; i >= m/2; i-- {
		if n>>i&1 > 0 {
			// 这一位可以填 0，那么回文数左半的剩余位置可以随便填
			ans += 1 << (i - m/2)
		}
		// 在后续循环中，这一位填 1
	}

	pal := n >> (m / 2)
	// 左半反转到右半
	// 如果 m 是奇数，那么去掉回文中心再反转
	for v := pal >> (m % 2); v > 0; v /= 2 {
		pal = pal*2 + v%2
	}
	if pal <= n {
		ans++
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 动态规划题单的「**十、数位 DP**」。
2. 数学题单的「**§7.1 回文数**」。

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
