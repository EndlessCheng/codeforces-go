形象理解，有一些人站在一维数轴上，$\texttt{b}$ 站在最左边，$\texttt{z}$ 站在最右边。（$\texttt{a}$ 不变，所以不考虑 $\texttt{a}$）

目标是让所有人都移动到 $\texttt{z}+1$ 的位置。

每次操作可以选择一个点上的**所有人**，全体向右移动一步。

问：最少要操作多少次。

答案就是最小的非 $\texttt{a}$ 字母到 $\texttt{z}+1$ 的距离。

比如最左边的人是 $\texttt{y}$，那么要操作 $2$ 次。

特别地，如果 $s$ 全是 $\texttt{a}$，无需操作，返回 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1heYGzWEUa/?t=2m47s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str) -> int:
        # 'z' 的下一个字符是 '{'
        min_c = min((c for c in s if c != 'a'), default='{')
        return ord('{') - ord(min_c)
```

```java [sol-Java]
class Solution {
    public int minOperations(String s) {
        int minC = 'z' + 1;
        for (char c : s.toCharArray()) {
            if (c != 'a') {
                minC = Math.min(minC, c);
            }
        }
        return 'z' + 1 - minC;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s) {
        char min_c = 'z' + 1;
        for (char c : s) {
            if (c != 'a') {
                min_c = min(min_c, c);
            }
        }
        return 'z' + 1 - min_c;
    }
};
```

```go [sol-Go]
func minOperations(s string) int {
	minC := 'z' + 1
	for _, c := range s {
		if c != 'a' {
			minC = min(minC, c)
		}
	}
	return int('z' + 1 - minC)
}
```

## 优化

如果发现最小的字母是 $\texttt{b}$，可以提前退出循环。

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str) -> int:
        if 'b' in s:
            return 25
        # 'z' 的下一个字符是 '{'
        min_c = min((c for c in s if c != 'a'), default='{')
        return ord('{') - ord(min_c)
```

```java [sol-Java]
class Solution {
    public int minOperations(String s) {
        int minC = 'z' + 1;
        for (char c : s.toCharArray()) {
            if (c != 'a') {
                minC = Math.min(minC, c);
                if (minC == 'b') {
                    break;
                }
            }
        }
        return 'z' + 1 - minC;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s) {
        char min_c = 'z' + 1;
        for (char c : s) {
            if (c != 'a') {
                min_c = min(min_c, c);
                if (min_c == 'b') {
                    break;
                }
            }
        }
        return 'z' + 1 - min_c;
    }
};
```

```go [sol-Go]
func minOperations(s string) int {
	minC := 'z' + 1
	for _, c := range s {
		if c != 'a' {
			minC = min(minC, c)
			if minC == 'b' {
				break
			}
		}
	}
	return int('z' + 1 - minC)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心与思维题单的「**§5.2 脑筋急转弯**」。

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
