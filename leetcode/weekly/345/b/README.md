## 前置知识：异或运算的性质

$a\oplus a = 0$。

$a\oplus b = c$ 两边同时异或 $a$，可以得到 $b = c\oplus a$。

## 思路

根据前置知识，把题目中的公式变形一下，得到

$$
\textit{original}[i+1] = \textit{original}[i]\oplus \textit{derived}[i]
$$

那么有

$$
\begin{aligned}
    & \textit{original}[n-1]      \\
={} & \textit{original}[n-2]\oplus \textit{derived}[n-2]        \\
={} & (\textit{original}[n-3]\oplus \textit{derived}[n-3] )\oplus \textit{derived}[n-2]        \\
\vdots \\
={} & \textit{original}[0] \oplus \textit{derived}[0] \oplus \textit{derived}[1]\oplus \cdots \oplus \textit{derived}[n-2]
\end{aligned}
$$

由于 

$$
\textit{original}[0]\oplus \textit{original}[n-1] =\textit{derived}[n-1]
$$

联立得

$$
\textit{derived}[0] \oplus \textit{derived}[1] \oplus\cdots \oplus \textit{derived}[n-1] = 0
$$

所以如果上式成立，$\textit{original}$ 必然存在。

[视频讲解](https://www.bilibili.com/video/BV1ka4y137ua/)

```py [sol-Python3]
class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        return reduce(xor, derived) == 0
```

```java [sol-Java]
class Solution {
    public boolean doesValidArrayExist(int[] derived) {
        int xor = 0;
        for (int x : derived) {
            xor ^= x;
        }
        return xor == 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool doesValidArrayExist(vector<int>& derived) {
        int xor_ = 0;
        for (int x : derived) {
            xor_ ^= x;
        }
        return xor_ == 0;
    }
};
```

```go [sol-Go]
func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{derived}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

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
