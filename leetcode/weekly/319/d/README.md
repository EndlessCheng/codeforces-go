计算每个子串是否回文，可以用**中心扩展法**，思路参考 [647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/) 的 [官方题解](https://leetcode.cn/problems/palindromic-substrings/solutions/379987/hui-wen-zi-chuan-by-leetcode-solution/)，下面只讲解 DP 部分。

定义 $f[i+1]$ 表示 $s[0]$ 到 $s[i]$ 中的不重叠回文子字符串的最大数目。这里 $+1$ 是因为我们用 $f[0]$ 表示空串。

**转移方程一**：如果 $s[i]$ 不在回文子串中，那么 $f[i+1] = f[i]$。

**转移方程二**：采用中心扩展法，如果 $s[l]$ 到 $s[r]$ 是回文子串，且 $r-l+1\ge k$，那么更新

$$
f[r+1] = \max(f[r+1], f[l]+1)
$$

**初始值**：$f[0] = 0$，表示空串。

**答案**：$f[n]$。

代码实现时，由于长度为 $m$ 的回文子串一定包含长为 $m-2$ 的回文子串，所以回文子串的长度 $\ge k$ 就可以退出循环了。

[本题视频讲解](https://www.bilibili.com/video/BV13841187gz/)（第四题）

```py [sol-Python3]
class Solution:
    def maxPalindromes(self, s: str, k: int) -> int:
        n = len(s)
        f = [0] * (n + 1)
        for i in range(2 * n - 1):
            l, r = i // 2, (i + 1) // 2
            f[l + 1] = max(f[l + 1], f[l])
            while l >= 0 and r < n and s[l] == s[r]:
                if r - l + 1 >= k:
                    f[r + 1] = max(f[r + 1], f[l] + 1)
                    break
                l -= 1
                r += 1
        return f[n]
```

```java [sol-Java]
class Solution {
    public int maxPalindromes(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int[] f = new int[n + 1];
        for (int i = 0; i < 2 * n - 1; i++) {
            int l = i / 2;
            int r = (i + 1) / 2;
            f[l + 1] = Math.max(f[l + 1], f[l]);
            for (; l >= 0 && r < n && s[l] == s[r]; l--, r++) {
                if (r - l + 1 >= k) {
                    f[r + 1] = Math.max(f[r + 1], f[l] + 1);
                    break;
                }
            }
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPalindromes(string s, int k) {
        int n = s.size();
        vector<int> f(n + 1);
        for (int i = 0; i < 2 * n - 1; i++) {
            int l = i / 2, r = (i + 1) / 2;
            f[l + 1] = max(f[l + 1], f[l]);
            for (; l >= 0 && r < n && s[l] == s[r]; l--, r++) {
                if (r - l + 1 >= k) {
                    f[r + 1] = max(f[r + 1], f[l] + 1);
                    break;
                }
            }
        }
        return f[n];
    }
};
```

```go [sol-Go]
func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		f[l+1] = max(f[l+1], f[l])
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 >= k {
				f[r+1] = max(f[r+1], f[l]+1)
				break
			}
			l--
			r++
		}
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面动态规划题单中的「**§5.2 最优划分**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
