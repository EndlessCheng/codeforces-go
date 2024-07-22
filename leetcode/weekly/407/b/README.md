分类讨论：

- 如果 $s$ 不包含任何元音，小红输。
- 如果 $s$ 包含奇数个元音，小红可以直接把整个 $s$ 移除，小红赢。
- 如果 $s$ 包含正偶数个元音，由于**偶数减奇数等于奇数**，小红移除任意包含奇数个元音的子串后，剩余元音个数仍然为奇数。由于**奇数减偶数还是奇数**，所以无论小明怎么操作，仍然会剩下奇数个元音，此时小红可以直接把整个 $s$ 移除，小红赢。

所以只要 $s$ 包含元音，就返回 $\texttt{true}$，否则返回 $\texttt{false}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16Z421N7P2/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def doesAliceWin(self, s: str) -> bool:
        return any(c in "aeiou" for c in s)
```

```java [sol-Java]
class Solution {
    public boolean doesAliceWin(String s) {
        for (char c : s.toCharArray()) {
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
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
    bool doesAliceWin(string s) {
        return ranges::any_of(s, [](char c) {
            return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
        });
    }
};
```

```go [sol-Go]
func doesAliceWin(s string) bool {
	return strings.ContainsAny(s, "aeiou")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果小明先手呢？

欢迎在评论区发表你的思路/代码。

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
