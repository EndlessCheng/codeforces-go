设当前剩余资源为 $\textit{left}$，设当前遍历到的元素为 $x = \textit{nums}[i]$。

如果 $\textit{left} < x$，设需要操作 $\textit{op}$ 次，那么有

$$
\textit{left} + \textit{op}\cdot k \ge x
$$

解得最小的 $\textit{op}$ 为

$$
\left\lceil\dfrac{x-\textit{left}}{k}\right\rceil = \left\lfloor\dfrac{x-\textit{left}-1}{k}\right\rfloor + 1
$$

上式右边见 [上取整下取整转换公式的证明](https://zhuanlan.zhihu.com/p/1890356682149838951)。

累加所有操作次数，得到总操作次数 $\textit{sum}$。根据等差数列求和公式，最终答案为

$$
1+2+\cdots+\textit{sum} = \dfrac{\textit{sum}(\textit{sum}+1)}{2}
$$

代码实现时，注意取模。为什么可以在**中途取模**？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

> **注**：由于 $\textit{sum}(\textit{sum}+1)$ 一定是偶数，除以 $2$ 无需用逆元。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minimumCost(self, nums: list[int], k: int) -> int:
        MOD = 1_000_000_007
        s = 0  # 总操作次数
        left = k
        for x in nums:
            if left < x:
                op = (x - left - 1) // k + 1  # 把 left 增大到 >= x，至少操作 op 次
                s += op
                left += op * k
            left -= x

        # 1 + 2 + ... + s
        return s * (s + 1) // 2 % MOD
```

```java [sol-Java]
class Solution {
    public int minimumCost(int[] nums, int k) {
        final int MOD = 1_000_000_007;
        long sum = 0; // 总操作次数
        int left = k;
        for (int x : nums) {
            if (left < x) {
                int op = (x - left - 1) / k + 1; // 把 left 增大到 >= x，至少操作 op 次
                sum += op;
                left += op * k;
            }
            left -= x;
        }

        // 1 + 2 + ... + sum
        sum %= MOD;
        return (int) (sum * (sum + 1) / 2 % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumCost(vector<int>& nums, int k) {
        constexpr int MOD = 1'000'000'007;
        long long sum = 0; // 总操作次数
        int left = k;
        for (int x : nums) {
            if (left < x) {
                int op = (x - left - 1) / k + 1; // 把 left 增大到 >= x，至少操作 op 次
                sum += op;
                left += op * k;
            }
            left -= x;
        }

        // 1 + 2 + ... + sum
        sum %= MOD;
        return sum * (sum + 1) / 2 % MOD;
    }
};
```

```go [sol-Go]
func minimumCost(nums []int, k int) int {
	const mod = 1_000_000_007
	sum := 0 // 总操作次数
	left := k
	for _, x := range nums {
		if left < x {
			op := (x-left-1)/k + 1 // 把 left 增大到 >= x，至少操作 op 次
			sum += op
			left += op * k
		}
		left -= x
	}

	// 1 + 2 + ... + sum
	sum %= mod
	return sum * (sum + 1) / 2 % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
