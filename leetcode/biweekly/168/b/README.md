本文把 $\textit{num}$ 简记为 $n$。

首先判断无解的情况。如果所有位置都填 $9$，数位和 $9n$ 仍然小于 $\textit{sum}$，那么无解，返回空串。

否则有解。为了最大化答案的字典序，我们可以：

- 先填 $\left\lfloor\dfrac{\textit{sum}}{9}\right\rfloor$ 个 $9$。
- 如果 $\textit{sum}\bmod 9 > 0$，那么再填一个 $\textit{sum}\bmod 9$。
- 最后填入 $n-|\textit{ans}|$ 个 $0$，其中 $|\textit{ans}|$ 表示当前答案的长度。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def maxSumOfSquares(self, n: int, s: int) -> str:
        if n * 9 < s:
            return ""
        ans = '9' * (s // 9)
        if s % 9:
            ans += digits[s % 9]
        return ans + '0' * (n - len(ans))
```

```java [sol-Java]
class Solution {
    public String maxSumOfSquares(int n, int sum) {
        if (n * 9 < sum) {
            return "";
        }
        StringBuilder ans = new StringBuilder(n).repeat('9', sum / 9);
        if (sum % 9 > 0) {
            ans.append((char) ('0' + sum % 9));
        }
        ans.repeat('0', n - ans.length());
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string maxSumOfSquares(int n, int sum) {
        if (n * 9 < sum) {
            return "";
        }
        string ans(sum / 9, '9');
        if (sum % 9) {
            ans += '0' + sum % 9;
        }
        return ans + string(n - ans.size(), '0');
    }
};
```

```go [sol-Go]
func maxSumOfSquares(n, sum int) string {
	if n*9 < sum {
		return ""
	}
	ans := strings.Repeat("9", sum/9)
	if sum%9 > 0 {
		ans += string('0' + byte(sum%9))
	}
	return ans + strings.Repeat("0", n-len(ans))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，返回值不计入。

## 专题训练

见下面贪心题单的「**§3.1 字典序最小/最大**」。

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
