### 提示 1

在 $[\textit{left}, \textit{right}]$ 之间的字符，删除是不影响得分的，且删除后更有机会让剩余部分是 $s$ 的子序列。

因此只需考虑删除的是 $t$ 的**子串**，而不是子序列。

### 提示 2

删除子串后，剩余部分是 $t$ 的一个前缀和一个后缀。

假设前缀匹配的是 $s$ 的一个前缀 $s[:i]$，后缀匹配的是 $s$ 的一个后缀 $s[i:]$。这里匹配指子序列匹配。

那么枚举 $i$，分别计算能够与 $s[:i]$ 和 $s[i:]$ 匹配的 $t$ 的最长前缀和最长后缀，就知道要删除的子串的最小值了。这个技巧叫做「前后缀分解」。

### 提示 3

具体来说：

定义 $\textit{pre}[i]$ 为 $s[:i]$ 对应的 $t$ 的最长前缀的结束下标。

定义 $\textit{suf}[i]$ 为 $s[i:]$ 对应的 $t$ 的最长后缀的开始下标。

计算方式见 [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/)。

删除的子串是从 $\textit{pre}[i]+1$ 到 $\textit{suf}[i]-1$ 这段，答案就是 $\textit{suf}[i]-\textit{pre}[i]-1$ 的最小值。

代码实现时，可以先计算 $\textit{suf}$，然后一边计算 $\textit{pre}$，一边更新最小值，所以 $\textit{pre}$ 可以省略。

[视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)

```py [sol-Python3]
class Solution:
    def minimumScore(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        suf = [m] * (n + 1)
        j = m - 1
        for i in range(n - 1, -1, -1):
            if s[i] == t[j]:
                j -= 1
            if j < 0:  # t 是 s 的子序列
                return 0
            suf[i] = j + 1

        ans = suf[0]  # 删除 t[:suf[0]]
        j = 0
        for i, c in enumerate(s):
            if c == t[j]:  # 注意上面判断了 t 是 s 子序列的情况，这里 j 不会越界
                j += 1
                ans = min(ans, suf[i + 1] - j)  # 删除 t[j:suf[i+1]]
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumScore(String S, String T) {
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;
        int m = t.length;

        int[] suf = new int[n + 1];
        suf[n] = m;
        int j = m - 1;
        for (int i = n - 1; i >= 0; i--) {
            if (s[i] == t[j]) {
                j--;
            }
            if (j < 0) { // t 是 s 的子序列
                return 0;
            }
            suf[i] = j + 1;
        }

        int ans = suf[0]; // 删除 t[:suf[0]]
        j = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == t[j]) { // 注意上面判断了 t 是 s 子序列的情况，这里 j 不会越界
                j++;
                ans = Math.min(ans, suf[i + 1] - j); // 删除 t[j:suf[i+1]]
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumScore(string s, string t) {
        int n = s.length(), m = t.length();
        vector<int> suf(n + 1);
        suf[n] = m;
        for (int i = n - 1, j = m - 1; i >= 0; i--) {
            if (s[i] == t[j]) {
                j--;
            }
            if (j < 0) { // t 是 s 的子序列
                return 0;
            }
            suf[i] = j + 1;
        }
        
        int ans = suf[0]; // 删除 t[:suf[0]]
        for (int i = 0, j = 0; i < n; i++) {
            if (s[i] == t[j]) { // 注意上面判断了 t 是 s 子序列的情况，这里 j 不会越界
                j++;
                ans = min(ans, suf[i + 1] - j); // 删除 t[j:suf[i+1]]
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumScore(s, t string) int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if s[i] == t[j] {
			j--
		}
		if j < 0 { // t 是 s 的子序列
			return 0
		}
		suf[i] = j + 1
	}

	ans := suf[0] // 删除 t[:suf[0]]
	j := 0
	for i := range s {
		if s[i] == t[j] { // 注意上面判断了 t 是 s 子序列的情况，这里 j 不会越界
			j++
			ans = min(ans, suf[i+1]-j) // 删除 t[j:suf[i+1]]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。注意时间复杂度与 $t$ 的长度无关。
- 空间复杂度：$\mathcal{O}(n)$。

## 变形题

删除 $s$ 中的子串，使得 $t$ 仍然是 $s$ 的子序列。可以删除的最长子串长度是多少？

见 [CF1203D2](https://codeforces.com/problemset/problem/1203/D2)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
