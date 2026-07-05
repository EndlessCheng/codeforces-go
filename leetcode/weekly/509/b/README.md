**前置题目**：[392. 判断子序列](https://leetcode.cn/problems/is-subsequence/)，[我的题解](https://leetcode.cn/problems/is-subsequence/solution/jian-ji-xie-fa-pythonjavaccgojsrust-by-e-mz22/)。

在 392 题中，我们定义 $j_0$ 表示在不修改的情况下，$s$ 的前缀 $[0, j_0-1]$ 是 $t$ 的（当前正在遍历的）前缀的子序列，且 $j_0$ 尽量大。

对于本题，我们定义 $j_1$ 表示在改过一次的情况下，$s$ 的前缀 $[0, j_1-1]$ 是 $t$ 的（当前正在遍历的）前缀的子序列，且 $j_1$ 尽量大。

对于 $j_1$，有两种情况：

- 普通匹配：如果 $s[j_1] = t[i]$，那么 $j_1$ 增加一。
- 修改：把 $s[j_0]$ 改成 $t[i]$，那么 $j_1$ 在 $j_0$ 的基础上加一。

两种情况取最大值。

对于 $j_0$，只能普通匹配：如果 $s[j_0] = t[i]$，那么 $j_0$ 增加一。

任意时刻，只要 $j_0=|s|$ 或者 $j_1=|s|$，则说明 $s$ 是 $t$ 的子序列。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def canMakeSubsequence(self, s: str, t: str) -> bool:
        n = len(s)
        j0 = 0  # 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        j1 = 0  # 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列

        for ch in t:
            # j1 普通匹配
            if s[j1] == ch:
                j1 += 1

            # 也可以修改 s[j0] 为 ch，强行匹配
            j1 = max(j1, j0 + 1)

            # j0 普通匹配
            if s[j0] == ch:
                j0 += 1

            if j0 == n or j1 == n:
                # s 是 t 的子序列
                return True

        return False
```

```java [sol-Java]
class Solution {
    public boolean canMakeSubsequence(String S, String t) {
        char[] s = S.toCharArray();
        int n = s.length;
        int j0 = 0; // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        int j1 = 0; // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
        for (char ch : t.toCharArray()) {
            // j1 普通匹配
            if (s[j1] == ch) {
                j1++;
            }

            // 也可以修改 s[j0] 为 ch，强行匹配
            j1 = Math.max(j1, j0 + 1);

            // j0 普通匹配
            if (s[j0] == ch) {
                j0++;
            }

            if (j0 == n || j1 == n) {
                // s 是 t 的子序列
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
    bool canMakeSubsequence(string s, string t) {
        int n = s.size();
        int j0 = 0; // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
        int j1 = 0; // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
        for (char ch : t) {
            // j1 普通匹配
            if (s[j1] == ch) {
                j1++;
            }

            // 也可以修改 s[j0] 为 ch，强行匹配
            j1 = max(j1, j0 + 1);

            // j0 普通匹配
            if (s[j0] == ch) {
                j0++;
            }

            if (j0 == n || j1 == n) {
                // s 是 t 的子序列
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func canMakeSubsequence(s, t string) bool {
	n := len(s)
	j0 := 0 // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
	j1 := 0 // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
	for _, ch := range t {
		// j1 普通匹配
		if s[j1] == byte(ch) {
			j1++
		}

		// 也可以修改 s[j0] 为 ch，强行匹配
		j1 = max(j1, j0+1)

		// j0 普通匹配
		if s[j0] == byte(ch) {
			j0++
		}

		if j0 == n || j1 == n {
			// s 是 t 的子序列
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 双指针题单的「**§4.2 判断子序列**」。
2. 动态规划题单的「**六、状态机 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
