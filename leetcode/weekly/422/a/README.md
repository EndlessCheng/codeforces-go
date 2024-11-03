初始化 $s=0$。

遍历字符串，奇数下标数字加到 $s$ 中，偶数下标数字的相反数加到 $s$ 中。

如果最终 $s=0$，返回 $\texttt{true}$，否则返回 $\texttt{false}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hcS1YCETs/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def isBalanced(self, num: str) -> bool:
        s = 0
        for i, c in enumerate(map(int, num)):
            s += c if i % 2 else -c
        return s == 0
```

```py [sol-Python3 切片]
class Solution:
    def isBalanced(self, num: str) -> bool:
        a = list(map(int, num))
        return sum(a[::2]) == sum(a[1::2])
```

```java [sol-Java]
class Solution {
    boolean isBalanced(String num) {
        int s = 0;
        char[] digits = num.toCharArray();
        for (int i = 0; i < digits.length; i++) {
            int c = digits[i] - '0';
            s += i % 2 > 0 ? c : -c;
        }
        return s == 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isBalanced(string num) {
        int s = 0;
        for (int i = 0; i < num.length(); i++) {
            int c = num[i] - '0';
            s += i % 2 ? c : -c;
        }
        return s == 0;
    }
};
```

```go [sol-Go]
func isBalanced(num string) bool {
	s := 0
	for i, b := range num {
		s += (i%2*2 - 1) * int(b-'0')
	}
	return s == 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{num}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
