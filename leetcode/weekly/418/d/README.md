## 每个 GCD 的出现次数

直接把每个 GCD 分别出现多少次算出来，就方便回答询问了。

以 $2$ 为例，如何计算 GCD 恰好等于 $2$ 的数对个数？

统计 $\textit{nums}$ 中的 $2$ 的倍数的个数，假设有 $5$ 个。从这 $5$ 个数中选 $2$ 个数，可以得到 $\dbinom 5 2 =\dfrac{5(5-1)}{2}=10$ 个数对。

但这 $10$ 个数对的 GCD 并不都恰好等于 $2$，比如数对 $(8,12)$ 的 GCD 为 $4$。

只能说，这 $10$ 个数对的 GCD 都是 $2$ 的倍数。

我们可以从 $10$ 中减去 GCD 等于 $4,6,8,\ldots$ 的数对个数，就得到 GCD **恰好**等于 $2$ 的数对个数了。

一般地，定义 $\textit{cntGcd}[i]$ 为 GCD 等于 $i$ 的数对个数。

枚举 $i$ 的倍数，统计 $\textit{nums}$ 中有多少个数等于 $i,2i,3i,\ldots$ 把个数记作 $c$。

这 $c$ 个数选 $2$ 个数，组成 $\dfrac{c(c-1)}{2}$ 个数对。

但是，这些数对的 GCD 只是 $i$ 的倍数，并不都恰好等于 $i$。

减去其中 GCD 等于 $2i,3i,\ldots$ 的数对个数，得到如下**递推式**

$$
\textit{cntGcd}[i] = \dfrac{c(c-1)}{2} - \textit{cntGcd}[2i] - \textit{cntGcd}[3i] - \cdots
$$

> 注：由于计算 $\textit{cntGcd}[i]$ 需要知道 $\textit{cntGcd}[2i]$ 的值，所以代码实现时，要**倒序枚举** $i$。

## 回答询问

比如 $\textit{gcdPairs}=[1,1,2,2,3,3,3]$，对应的 $\textit{cntGcd}=[0,2,2,3]$。

计算 $\textit{cntGcd}$ 的前缀和数组 $s=[0,2,4,7]$。

- $q=0,1$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $1$。
- $q=2,3$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $2$。
- $q=4,5,6$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $3$。

一般地，$\textit{gcdPairs}$ 中等于 $g$ 的数，有 $\textit{cntGcd}[g]$ 个，这些数在 $\textit{gcdPairs}$ 中的下标为闭区间 $[s[g-1], s[g]-1]$。换句话说，对于 $\textit{gcdPairs}$ 的下标 $q$，我们有 $s[g-1]\le q < s[g]$。于是，在 $s$ 中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 严格大于 $q$ 的第一个数的下标，即为 $g$。

> 注：由于 $\textit{cntGcd}[i]\ge 0$，所以 $s$ 是有序数组，可以在 $s$ 上二分查找。

[本题视频讲解](https://www.bilibili.com/video/BV15y1iYUE2h/?t=17m58s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def gcdValues(self, nums: List[int], queries: List[int]) -> List[int]:
        mx = max(nums)
        cnt_x = [0] * (mx + 1)
        for x in nums:
            cnt_x[x] += 1

        cnt_gcd = [0] * (mx + 1)
        for i in range(mx, 0, -1):
            c = 0
            for j in range(i, mx + 1, i):
                c += cnt_x[j]
                cnt_gcd[i] -= cnt_gcd[j]  # gcd 是 2i,3i,4i,... 的数对不能统计进来
            cnt_gcd[i] += c * (c - 1) // 2  # c 个数选 2 个，组成 c*(c-1)/2 个数对

        s = list(accumulate(cnt_gcd))  # 前缀和
        return [bisect_right(s, q) for q in queries]
```

```java [sol-Java]
class Solution {
    public int[] gcdValues(int[] nums, long[] queries) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        int[] cntX = new int[mx + 1];
        for (int x : nums) {
            cntX[x]++;
        }

        long[] cntGcd = new long[mx + 1];
        for (int i = mx; i > 0; i--) {
            int c = 0;
            for (int j = i; j <= mx; j += i) {
                c += cntX[j];
                cntGcd[i] -= cntGcd[j]; // gcd 是 2i,3i,4i,... 的数对不能统计进来
            }
            cntGcd[i] += (long) c * (c - 1) / 2; // c 个数选 2 个，组成 c*(c-1)/2 个数对
        }

        for (int i = 2; i <= mx; i++) {
            cntGcd[i] += cntGcd[i - 1]; // 原地求前缀和
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = upperBound(cntGcd, queries[i]);
        }
        return ans;
    }

    private int upperBound(long[] nums, long target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = (left + right) >>> 1;
            if (nums[mid] > target) {
                right = mid; // 二分范围缩小到 (left, mid)
            } else {
                left = mid; // 二分范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> gcdValues(vector<int>& nums, vector<long long>& queries) {
        int mx = ranges::max(nums);
        vector<int> cnt_x(mx + 1);
        for (int x : nums) {
            cnt_x[x]++;
        }

        vector<long long> cnt_gcd(mx + 1);
        for (int i = mx; i > 0; i--) {
            int c = 0;
            for (int j = i; j <= mx; j += i) {
                c += cnt_x[j];
                cnt_gcd[i] -= cnt_gcd[j]; // gcd 是 2i,3i,4i,... 的数对不能统计进来
            }
            cnt_gcd[i] += (long long) c * (c - 1) / 2; // c 个数选 2 个，组成 c*(c-1)/2 个数对
        }

        // 原地求前缀和
        partial_sum(cnt_gcd.begin(), cnt_gcd.end(), cnt_gcd.begin());

        vector<int> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            ans[i] = ranges::upper_bound(cnt_gcd, queries[i]) - cnt_gcd.begin();
        }
        return ans;
    }
};
```

```go [sol-Go]
func gcdValues(nums []int, queries []int64) []int {
	mx := slices.Max(nums)
	cntX := make([]int, mx+1)
	for _, x := range nums {
		cntX[x]++
	}

	cntGcd := make([]int, mx+1)
	for i := mx; i > 0; i-- {
		c := 0
		for j := i; j <= mx; j += i {
			c += cntX[j]
			cntGcd[i] -= cntGcd[j] // gcd 是 2i,3i,4i,... 的数对不能统计进来
		}
		cntGcd[i] += c * (c - 1) / 2 // c 个数选 2 个，组成 c*(c-1)/2 个数对
	}

	for i := 2; i <= mx; i++ {
		cntGcd[i] += cntGcd[i-1] // 原地求前缀和
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sort.SearchInts(cntGcd, int(q)+1)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + (U+q)\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$q$ 是 $\textit{queries}$ 的长度。代码中的二重循环，根据调和级数可得，时间复杂度为 $\mathcal{O}(U\log U)$。
- 空间复杂度：$\mathcal{O}(U)$。返回值不计入。

## 变形题

1. 计算有多少个子序列的 GCD 恰好等于 $i$。见 [CF803F](https://codeforces.com/problemset/problem/803/F)。
2. 计算树上有多少条简单路径的点权 GCD 恰好等于 $i$。见 [CF990G](https://codeforces.com/problemset/problem/990/G)。

> 注：子数组 GCD 可以用 LogTrick 做，见位运算题单。

更多相似题目，见下面数学题单的「**§1.6 最大公约数（GCD）**」。

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
