由于每次都有两个数贡献到总代价中，所以总代价是 $\textit{nums}$ 中的偶数个数之和。

$\textit{nums}[0]$ 和 $\textit{nums}[n-1]$ 一定是子数组的首尾，一定要选。

其他元素怎么选？

先来算算答案的下界（理论最小值）。把 $\textit{nums}[1,n-2]$ 从小到大排序，两个两个一组，只要两数之和小于 $0$，就选，否则跳出循环。

这个下界是可以达到的。例如我们从左到右选了 $a,b,c,d,e,f$ 六个数（在原数组中的顺序），那么按照 $[c,d],[b,e],[a,f]$ 的顺序**从内到外删除子数组**，就可以保证我们选到这六个数。

所以本题只需排序+遍历。

> 也可以先统计负数的个数，然后跑快速选择算法，这样可以做到 $\mathcal{O}(n)$，见 C++ 第二份代码。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minCost(self, nums: list[int]) -> int:
        ans = nums[0] + nums[-1]
        a = sorted(nums[1:-1])
        for i in range(1, len(a), 2):
            if a[i - 1] + a[i] >= 0:
                break
            ans += a[i - 1] + a[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long minCost(int[] nums) {
        int n = nums.length;
        Arrays.sort(nums, 1, n - 1);

        long ans = nums[0] + nums[n - 1];
        for (int i = 2; i < n - 1 && nums[i - 1] + nums[i] < 0; i += 2) {
            ans += nums[i - 1] + nums[i];
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
        ranges::sort(nums.begin() + 1, nums.end() - 1);

        long long ans = nums[0] + nums[n - 1];
        for (int i = 2; i < n - 1 && nums[i - 1] + nums[i] < 0; i += 2) {
            ans += nums[i - 1] + nums[i];
        }
        return ans;
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long minCost(vector<int>& nums) {
        auto l = nums.begin() + 1, r = nums.end() - 1;
        int neg = count_if(l, r, [](int x) { return x < 0; });
        nth_element(l, l + neg, r);

        long long ans = nums[0] + nums.back();
        ans += reduce(l, l + neg, 0LL); // 先把负数都加进来
        if (neg % 2) { // 如果有奇数个负数，要么再加一个非负数，要么去掉一个负数
            int mx = *max_element(l, l + neg);
            int mn = l + neg != r ? *min_element(l + neg, r) : INT_MAX;
            // 如果最大的负数 + 最小的非负数 < 0，那么加上最小的非负数，否则减去最大的负数
            ans += mx + mn < 0 ? mn : -mx;
        }

        return ans;
    }
};
```

```go [sol-Go]
func minCost(nums []int) int64 {
	n := len(nums)
	a := nums[1 : n-1]
	slices.Sort(a)

	ans := int64(nums[0] + nums[n-1])
	for i := 1; i < n-2 && a[i-1]+a[i] < 0; i += 2 {
		ans += int64(a[i-1] + a[i])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。可以用快速选择做到 $\mathcal{O}(n)$，见 C++ 第二份代码。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面思维题单的「**§5.2 脑筋急转弯**」。

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
