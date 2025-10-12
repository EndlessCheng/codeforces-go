首先遍历 $s$，计算总得分 $\textit{total}$。

然后再遍历 $s$，一边遍历一边计算前缀得分 $\textit{left}$，那么另一半的得分就是 $\textit{total}-\textit{left}$。

题目要求

$$
\textit{left} = \textit{total}-\textit{left}
$$

即

$$
\textit{left}\cdot 2 = \textit{total}
$$

如果遍历过程中符合上式，返回 $\texttt{true}$。否则返回 $\texttt{false}$。

根据 ASCII 表，计算字母 $c$ 在字母表中的位置，可以用 `c - 'a' + 1`，也可以用更简洁的 `c & 31`。

[本题视频讲解](https://www.bilibili.com/video/BV16E4uzLEdK/?t=20m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def scoreBalance(self, s: str) -> bool:
        total = sum(ord(b) & 31 for b in s)

        left = 0
        for b in s:  # 字母位置是正数，可以遍历到 s 末尾（末尾一定不满足要求）
            left += ord(b) & 31
            if left * 2 == total:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean scoreBalance(String S) {
        char[] s = S.toCharArray();
        int total = 0;
        for (char b : s) {
            total += b & 31;
        }

        int left = 0;
        for (char b : s) { // 字母位置是正数，可以遍历到 s 末尾（末尾一定不满足要求）
            left += b & 31;
            if (left * 2 == total) {
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
    bool scoreBalance(string s) {
        int total = 0;
        for (char b : s) {
            total += b & 31;
        }

        int left = 0;
        for (char b : s) { // 字母位置是正数，可以遍历到 s 末尾（末尾一定不满足要求）
            left += b & 31;
            if (left * 2 == total) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func scoreBalance(s string) bool {
	total := 0
	for _, b := range s {
		total += int(b & 31)
	}

	left := 0
	for _, b := range s { // 字母位置是正数，可以遍历到 s 末尾（末尾一定不满足要求）
		left += int(b & 31)
		if left*2 == total {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

下面动态规划题单的「**专题：前后缀分解**」。

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
