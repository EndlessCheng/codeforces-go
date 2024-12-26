设 $x=s[i-1]$，$y=s[i]$，那么 $x+y$ 是 $s$ 的一个长为 $2$ 的子串。

**核心思路**：$x+y$ 在反转串 $\text{reverse}(s)$ 中，等价于 $y+x$ 在 $s$ 中。

用一个 $26\times 26$ 的布尔数组（或哈希表）$\textit{vis}$，其中 $\textit{vis}[x][y]=\texttt{true}$ 表示遇到了子串 $x+y$。

如果 $\textit{vis}[y][x]$ 为 $\texttt{true}$，则说明 $x+y$ 在反转串 $\text{reverse}(s)$ 中。

[视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/)。

```py [sol-Python3]
class Solution:
    def isSubstringPresent(self, s: str) -> bool:
        vis = set()
        for x, y in pairwise(s):
            vis.add(x + y)
            if y + x in vis:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean isSubstringPresent(String S) {
        char[] s = S.toCharArray();
        boolean[][] vis = new boolean[26][26];
        for (int i = 1; i < s.length; i++) {
            int x = s[i - 1] - 'a';
            int y = s[i] - 'a';
            vis[x][y] = true;
            if (vis[y][x]) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isSubstringPresent(string s) {
        bool vis[26][26]{};
        for (int i = 1; i < s.length(); i++) {
            int x = s[i - 1] - 'a', y = s[i] - 'a';
            vis[x][y] = true;
            if (vis[y][x]) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func isSubstringPresent(s string) bool {
	vis := [26][26]bool{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x][y] = true
		if vis[y][x] {
			return true
		}
	}
	return false
}
```

也可以用位运算优化，把 $\textit{vis}$ 数组的第二维压缩成二进制数。原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def isSubstringPresent(self, s: str) -> bool:
        vis = [0] * 26
        for x, y in pairwise(map(ord, s)):
            x -= ord('a')
            y -= ord('a')
            vis[x] |= 1 << y
            if vis[y] >> x & 1:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean isSubstringPresent(String S) {
        char[] s = S.toCharArray();
        int[] vis = new int[26];
        for (int i = 1; i < s.length; i++) {
            int x = s[i - 1] - 'a';
            int y = s[i] - 'a';
            vis[x] |= 1 << y;
            if ((vis[y] >> x & 1) > 0) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isSubstringPresent(string s) {
        int vis[26]{};
        for (int i = 1; i < s.length(); i++) {
            int x = s[i - 1] - 'a', y = s[i] - 'a';
            vis[x] |= 1 << y;
            if (vis[y] >> x & 1) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func isSubstringPresent(s string) bool {
	vis := [26]int{}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'a', s[i]-'a'
		vis[x] |= 1 << y
		if vis[y]>>x&1 > 0 {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|)$，其中 $n$ 为 $\textit{nums}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。注意创建一个大小为 $\mathcal{O}(|\Sigma|)$ 的数组需要 $\mathcal{O}(|\Sigma|)$ 的时间。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
