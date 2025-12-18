## 题意

把 $\textit{customers}$ 分成两段，第一段计算 $\texttt{N}$ 的个数 $\textit{preN}$，第二段计算 $\texttt{Y}$ 的个数 $\textit{sufY}$。

计算 $\textit{preN}+\textit{sufY}$ 的最小值对应的最小分割位置 $i$，其中 $[0,i-1]$ 是第一段，$[i,n-1]$ 是第二段。

注：也可以不分割，相当于其中一段是空的。

## 思路

先假设所有字母都在第二段，统计 $\textit{customers}$ 中 $\texttt{Y}$ 的个数 $\textit{sufY}$。

想象一根分割线在从左到右移动，$\textit{customers}[i]$ 原来在第二段，现在在第一段。如果 $\textit{customers}[i] = \texttt{N}$，那么把 $\textit{preN}$ 加一（$\textit{preN}$ 初始值为 $0$），否则把 $\textit{sufY}$ 减一。

答案为 $\textit{preN}+\textit{sufY}$ 最小时对应的分割位置。

代码实现时，$\textit{preN}+\textit{sufY}$ 可以合并为一个变量 $\textit{penalty}$。

## 优化前

```py [sol-Python3]
class Solution:
    def bestClosingTime(self, customers: str) -> int:
        min_penalty = penalty = customers.count('Y')
        ans = 0  # [0,n-1] 是第二段
        for i, c in enumerate(customers):
            penalty += 1 if c == 'N' else -1
            if penalty < min_penalty:
                min_penalty = penalty
                ans = i + 1  # [0,i] 是第一段，[i+1,n-1] 是第二段
        return ans
```

```java [sol-Java]
class Solution {
    public int bestClosingTime(String customers) {
        char[] s = customers.toCharArray();
        int penalty = 0;
        for (char c : s) {
            if (c == 'Y') {
                penalty++;
            }
        }

        int minPenalty = penalty;
        int ans = 0; // [0,n-1] 是第二段
        for (int i = 0; i < s.length; i++) {
            penalty += s[i] == 'N' ? 1 : -1;
            if (penalty < minPenalty) {
                minPenalty = penalty;
                ans = i + 1; // [0,i] 是第一段，[i+1,n-1] 是第二段
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int bestClosingTime(string customers) {
        int penalty = ranges::count(customers, 'Y');
        int min_penalty = penalty;
        int ans = 0; // [0,n-1] 是第二段
        for (int i = 0; i < customers.size(); i++) {
            penalty += customers[i] == 'N' ? 1 : -1;
            if (penalty < min_penalty) {
                min_penalty = penalty;
                ans = i + 1; // [0,i] 是第一段，[i+1,n-1] 是第二段
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func bestClosingTime(customers string) int {
	penalty := strings.Count(customers, "Y")
	minPenalty := penalty
	ans := 0 // [0,n-1] 是第二段
	for i, c := range customers {
		if c == 'N' {
			penalty++
		} else {
			penalty--
		}
		if penalty < minPenalty {
			minPenalty = penalty
			ans = i + 1 // [0,i] 是第一段，[i+1,n-1] 是第二段
		}
	}
	return ans
}
```

## 优化：一次遍历

第一次遍历（统计 $\texttt{Y}$ 的个数）是多余的。

打个比方，绝对温度下冰点为 $273\,\mathrm{K}$，摄氏温度下冰点为 $0~^\circ\mathrm{C}$。如果只关心哪一天最冷，用哪个单位计算都可以。

或者说，把折线图往下平移（第一个点移到坐标零点），最低点的横坐标仍然不变。

```py [sol-Python3]
class Solution:
    def bestClosingTime(self, customers: str) -> int:
        min_penalty = penalty = ans = 0
        for i, c in enumerate(customers):
            penalty += 1 if c == 'N' else -1
            if penalty < min_penalty:
                min_penalty = penalty
                ans = i + 1
        return ans
```

```java [sol-Java]
class Solution {
    public int bestClosingTime(String customers) {
        int penalty = 0;
        int minPenalty = 0;
        int ans = 0;
        for (int i = 0; i < customers.length(); i++) {
            penalty += customers.charAt(i) == 'N' ? 1 : -1;
            if (penalty < minPenalty) {
                minPenalty = penalty;
                ans = i + 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int bestClosingTime(string customers) {
        int penalty = 0, min_penalty = 0, ans = 0;
        for (int i = 0; i < customers.size(); i++) {
            penalty += customers[i] == 'N' ? 1 : -1;
            if (penalty < min_penalty) {
                min_penalty = penalty;
                ans = i + 1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func bestClosingTime(customers string) (ans int) {
	penalty, minPenalty := 0, 0
	for i, c := range customers {
		if c == 'N' {
			penalty++
		} else {
			penalty--
		}
		if penalty < minPenalty {
			minPenalty = penalty
			ans = i + 1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{customers}$ 的长度。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
