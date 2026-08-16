枚举 $s$ 左旋了 $\textit{rot} = 0,1,2,\ldots,n-1$ 次。

$s$ 左旋 $\textit{rot}$ 次后的字符串，是 $s+s$ 的一个子串，左端点下标为 $\textit{rot}$，右端点下标为 $\textit{rot}+n-1$。

我们需要判断该子串能否通过递增操作变成回文串，即 $s+s$ 下标 $\textit{rot}+i$ 的字母要等于下标 $\textit{rot}+n-1-i$ 的字母。

对于字母 $x$ 和 $y$（$x\le y$），我们可以：

- 把 $x$ 增大到 $y$，操作 $y-x$ 次。
- 把 $y$ 增大绕回到 $x$，操作 $26-y+x$ 次。

两种情况取最小值。

> **注**：把 $x$ 和 $y$ 都操作一定不优：把 $x$ 和 $y$ 都少操作一次，两个字母最终仍然相同。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s: str) -> int:
        n = len(s)
        ans = inf
        for rot in range(n):
            op = rot
            for i in range(n // 2):
                d = abs(ord(s[(rot + i) % n]) - ord(s[(rot - 1 - i) % n]))
                op += min(d, 26 - d)  # 注：这里可以加个剪枝，如果 op >= ans 则 break
            ans = min(ans, op)
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = Integer.MAX_VALUE;
        for (int rot = 0; rot < n; rot++) {
            int op = rot;
            for (int i = 0; i < n / 2; i++) {
                int d = Math.abs(s[(rot + i) % n] - s[(rot + n - 1 - i) % n]);
                op += Math.min(d, 26 - d); // 注：这里可以加个剪枝，如果 op >= ans 则 break
            }
            ans = Math.min(ans, op);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s) {
        int n = s.size();
        int ans = INT_MAX;
        for (int rot = 0; rot < n; rot++) {
            int op = rot;
            for (int i = 0; i < n / 2; i++) {
                int d = abs(s[(rot + i) % n] - s[(rot + n - 1 - i) % n]);
                op += min(d, 26 - d); // 注：这里可以加个剪枝，如果 op >= ans 则 break
            }
            ans = min(ans, op);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(s string) int {
	n := len(s)
	ans := math.MaxInt
	for rot := range n {
		op := rot
		for i := range n / 2 {
			d := abs(int(s[(rot+i)%n]) - int(s[(rot+n-1-i)%n]))
			op += min(d, 26-d) // 注：这里可以加个剪枝，如果 op >= ans 则 break
		}
		ans = min(ans, op)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
