把 $1$ 当作**车**，想象有一条长为 $n$ 的道路上有一些车。

我们要把所有的车都排到最右边，例如 $011010$ 最终要变成 $000111$。

下文用 $s\rightarrow t$ 表示车的位置变化。

如果我们优先操作右边的车，那么每辆车都只需操作一次：

$$
\begin{aligned}
    & 011010      \\
\rightarrow{} & 011001        \\
\rightarrow{} & 010011        \\
\rightarrow{} & 000111        \\
\end{aligned}
$$

一共需要操作 $3$ 次。

> 注意一次操作可以让一辆车移动多次。

如果我们优先操作左边的（能移动的）车，这会制造大量的「**堵车**」，那么每辆车的操作次数就会更多。

$$
\begin{aligned}
& 011010      \\
\rightarrow{} & 010110        \\
\rightarrow{} & 001110        \\
\rightarrow{} & 001101        \\
\rightarrow{} & 001011        \\
\rightarrow{} & 000111        \\
\end{aligned}
$$

一共需要操作 $5$ 次。

## 算法

1. 从左到右遍历 $s$，同时用一个变量 $\textit{cnt}_1$ 维护遍历到的 $1$ 的个数。
2. 如果 $s[i]$ 是 $1$，把 $\textit{cnt}_1$ 增加 $1$。
3. 如果 $s[i]$ 是 $0$ 且 $s[i-1]$ 是 $1$，意味着我们找到了一段道路，可以让 $i$ **左边的每辆车都操作一次**，把答案增加 $\textit{cnt}_1$。
4. 遍历结束，返回答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16Z421N7P2/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maxOperations(self, s: str) -> int:
        ans = cnt1 = 0
        for i, c in enumerate(s):
            if c == '1':
                cnt1 += 1
            elif i and s[i - 1] == '1':
                ans += cnt1
        return ans
```

```java [sol-Java]
class Solution {
    public int maxOperations(String S) {
        char[] s = S.toCharArray();
        int ans = 0;
        int cnt1 = 0;
        for (int i = 0; i < s.length; i++) {
            if (s[i] == '1') {
                cnt1++;
            } else if (i > 0 && s[i - 1] == '1') {
                ans += cnt1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxOperations(string s) {
        int ans = 0, cnt1 = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] == '1') {
                cnt1++;
            } else if (i && s[i - 1] == '1') {
                ans += cnt1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxOperations(s string) (ans int) {
	cnt1 := 0
	for i, c := range s {
		if c == '1' {
			cnt1++
		} else if i > 0 && s[i-1] == '1' {
			ans += cnt1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

你能构造一个 $s$，让返回值尽量大吗？

如果 $n=10^5$，答案最大能是多少？会不会超过 int 范围？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
