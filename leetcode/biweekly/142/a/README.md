根据题意：

- 如果所有相邻字母都互不相同，那么 Alice 不可能犯错，所以方案数只有 $1$ 种。
- 如果有 $1$ 对相邻字母相同，那么 Alice 可以在这里犯错一次，例如 $\texttt{abb}$，一开始想要输入的可能是 $\texttt{abb}$，也可能是 $\texttt{ab}$，其中 $\texttt{b}$ 多打了一次，所以方案数有 $2$ 种。
- 如果有 $2$ 对相邻字母相同，那么一开始想要输入的字符串会再多一种。例如 $\texttt{abbb}$，一开始想要输入的可能是 $\texttt{abbb}$，也可能是 $\texttt{abb}$，也可能是 $\texttt{ab}$，所以方案数有 $3$ 种。
- 依此类推，方案数等于相邻相同字母的个数加一。

具体请看 [视频讲解](https://www.bilibili.com/video/BV13J1MYwEGM/?t=16m7s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def possibleStringCount(self, word: str) -> int:
        ans = 1
        for x, y in pairwise(word):
            if x == y:
                ans += 1
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def possibleStringCount(self, word: str) -> int:
        return 1 + sum(x == y for x, y in pairwise(word))
```

```java [sol-Java]
class Solution {
    public int possibleStringCount(String word) {
        int ans = 1;
        for (int i = 1; i < word.length(); i++) {
            if (word.charAt(i - 1) == word.charAt(i)) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int possibleStringCount(string word) {
        int ans = 1;
        for (int i = 1; i < word.length(); i++) {
            if (word[i - 1] == word[i]) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
