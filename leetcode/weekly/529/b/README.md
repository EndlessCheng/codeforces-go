枚举最终都变成 $\textit{nums}[i]$。

设包含 $\textit{nums}[i]$ 的极大连续段为 $[L,R]$，那么：

1. 在 $L$ 处操作前缀，成本为 $L\cdot \textit{nums}[i]$。如果 $L=0$ 则不操作，相当于成本为 $0$。
2. 在 $R$ 处操作后缀，成本为 $(n-R-1)\cdot \textit{nums}[i]$。如果 $R=n-1$ 则不操作，相当于成本为 $0$。

总成本为

$$
(L + n-R-1)\cdot \textit{nums}[i]
$$

用上式更新答案的最小值。

注意 $[L,R]$ 不能只有 $\textit{nums}[0]$ 或者只有 $\textit{nums}[n-1]$，这种情况无法操作。

[本题视频讲解](https://www.bilibili.com/video/BV1LoyGzZE9E/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minCost(self, nums: list[int]) -> int:
        n = len(nums)
        ans = inf
        start = 0
        for i, x in enumerate(nums):
            if i == n - 1 or x != nums[i + 1]:  # i 是连续相同段的末尾
                if i > 0 and start < n - 1:  # 这一段不能只有 nums[0] 或者只有 nums[n-1]
                    ans = min(ans, (start + n - i - 1) * x)
                start = i + 1  # 下一段的开始位置
        return ans
```

```java [sol-Java]
class Solution {
    public long minCost(int[] nums) {
        int n = nums.length;
        long ans = Long.MAX_VALUE;
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i == n - 1 || x != nums[i + 1]) { // i 是连续相同段的末尾
                if (i > 0 && start < n - 1) { // 这一段不能只有 nums[0] 或者只有 nums[n-1]
                    ans = Math.min(ans, (long) (start + n - i - 1) * x);
                }
                start = i + 1; // 下一段的开始位置
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int>& nums) {
        int n = nums.size();
        long long ans = LLONG_MAX;
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i == n - 1 || x != nums[i + 1]) { // i 是连续相同段的末尾
                if (i > 0 && start < n - 1) { // 这一段不能只有 nums[0] 或者只有 nums[n-1]
                    ans = min(ans, 1LL * (start + n - i - 1) * x);
                }
                start = i + 1; // 下一段的开始位置
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCost(nums []int) int64 {
	n := len(nums)
	ans := math.MaxInt
	start := 0
	for i, x := range nums {
		if i == n-1 || x != nums[i+1] { // i 是连续相同段的末尾
			if i > 0 && start < n-1 { // 这一段不能只有 nums[0] 或者只有 nums[n-1]
				ans = min(ans, (start+n-i-1)*x)
			}
			start = i + 1 // 下一段的开始位置
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面双指针题单的「**六、分组循环**」。

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
