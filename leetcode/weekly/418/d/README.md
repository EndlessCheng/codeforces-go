## 每个 GCD 的出现次数

直接把每个 GCD 出现多少次计算出来，这样就方便回答询问了。

如何计算 GCD 等于 $2$ 的数对个数？

统计 $\textit{nums}$ 中的 $2$ 的倍数的个数，假设有 $5$ 个。那么我们可以从这 $5$ 个数中选 $2$ 个数，得到 $C_5^2=\dfrac{5(5-1)}{2}=10$ 个数对。

但这 $10$ 个数对，每个数对的 GCD 并不一定都是 $2$，比如 $8$ 和 $12$ 的 GCD 是 $4$。

只能说，这 $10$ 个数对的 GCD 都是 $2$ 的倍数。

我们可以从 $10$ 中减去 GCD 等于 $4,6,8,\cdots$ 的数对个数，就得到 GCD 恰好等于 $2$ 的数对个数了。

一般地，定义 $\textit{cntGcd}[i]$ 为 GCD 等于 $i$ 的数对个数。

枚举 $i$ 的倍数，统计 $\textit{nums}$ 中有多少个数等于 $i,2i,3i,\cdots$，记作 $c$。

这 $c$ 个数选 $2$ 个数，组成 $\dfrac{c(c-1)}{2}$ 个数对。

但是，这些数对的 GCD 只是 $i$ 的倍数，并不一定恰好等于 $i$。

减去其中 GCD 等于 $2i,3i,\cdots$ 的数对个数，得

$$
\textit{cntGcd}[i] = \dfrac{c(c-1)}{2} - \textit{cntGcd}[2i] - \textit{cntGcd}[3i] - \cdots
$$

为了完成这一计算，需要**倒序枚举** $i$。

## 回答询问

比如 $\textit{gcdPairs}=[1,1,2,2,3,3,3]$，对应的 $\textit{gcdCnt}=[0,2,2,3]$，计算其前缀和，得 $s=[0,2,4,7]$。

- $q=0,1$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $1$。
- $q=2,3$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $2$。
- $q=4,5,6$，答案都是 $s$ 中的大于 $q$ 的第一个数的下标，即 $3$。

所以在 $s$ 中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 大于 $q$ 的第一个数的下标即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15y1iYUE2h/) 第四题，欢迎点赞关注~

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
        partial_sum(cnt_gcd.begin(), cnt_gcd.end(), cnt_gcd.begin()); // 原地求前缀和

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

更多相似题目，见下面数学题单中 GCD 相关的小节。

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
