## 提示 1：二分答案

$\textit{nums}$ 的唯一性数组有多少个？也就是 $\textit{nums}$ 的非空连续子数组的个数。

长为 $n$ 的子数组有 $1$ 个，长为 $n-1$ 的子数组有 $2$ 个，……，长为 $1$ 的子数组有 $n$ 个。
 
所有一共有 $m = 1+2+\cdots +n = \dfrac{n(n+1)}{2}$ 个非空连续子数组。

这 $m$ 个子数组，对应着 $m$ 个 $\text{distinct}$ 值。

中位数是这 $m$ 个数中的第 $k = \left\lceil\dfrac{m}{2}\right\rceil$ 小元素。例如 $m=4$ 时，中位数是其中第 $2$ 小元素。

考虑这 $m$ 个数中，小于等于某个定值 $\textit{upper}$ 的数有多少个。

由于 $\textit{upper}$ 越大，小于等于 $\textit{upper}$ 的数越多，有**单调性**，故可以**二分**中位数为 $\textit{upper}$，问题变成：

- $\text{distinct}$ 值 $\le \textit{upper}$ 的子数组有多少个？

设子数组的个数为 $\textit{cnt}$，如果 $\textit{cnt} < k$ 说明二分的 $\textit{upper}$ 小了，更新二分左边界 $\textit{left}$，否则更新二分右边界 $\textit{right}$。

如果你没有想到二分答案，可以做做 [二分题单](https://leetcode.cn/circle/discuss/SqopEo/) 中的「**第 K 小/大**」。

## 提示 2：滑动窗口

怎么计算 $\text{distinct}$ 值 $\le \textit{upper}$ 的子数组个数？

由于子数组越长，不同元素个数（$\text{distinct}$ 值）不会变小，有**单调性**，故可以用**滑动窗口**计算子数组个数。

用一个哈希表 $\textit{freq}$ 统计窗口（子数组）内的元素及其出现次数。

枚举窗口右端点 $r$，把 $\textit{nums}[r]$ 加入 $\textit{freq}$（出现次数加一）。如果发现 $\textit{freq}$ 的大小超过 $\textit{upper}$，说明窗口内的元素过多，那么不断移出窗口左端点元素 $\textit{nums}[l]$（出现次数减一，如果出现次数等于 $0$ 就从 $\textit{freq}$ 中移除），直到 $\textit{freq}$ 的大小 $\le \textit{upper}$ 为止。

此时右端点为 $r$，左端点为 $l,l+1,l+2,\cdots,r$ 的子数组都是满足要求的（$\text{distinct}$ 值 $\le \textit{upper}$），一共有 $r-l+1$ 个，加到子数组个数 $\textit{cnt}$ 中。

如果你没有想到滑动窗口，可以做做 [滑动窗口题单](https://leetcode.cn/circle/discuss/0viNMK/) 中的「**不定长滑动窗口（求子数组个数）**」。

## 其它细节

开区间二分左边界：$0$，一定不满足要求，因为没有子数组的 $\text{distinct}$ 值是 $0$。

开区间二分右边界：$n$，或者 $\textit{nums}$ 中的不同元素个数，一定满足要求，因为所有子数组的 $\text{distinct}$ 值都不超过 $n$。

用开区间计算二分仅仅是个人喜好，你也可以使用闭区间或半闭半开区间计算二分，并无本质区别。

## 答疑

**问**：为什么二分出来的答案，一定是某个子数组的 $\text{distinct}$ 值？有没有可能，二分出来的答案不是任何子数组的 $\text{distinct}$ 值？

**答**：反证法。如果答案 $d$ 不是任何子数组的 $\text{distinct}$ 值，那么 $\text{distinct}$ 值 $\le d$ 和 $\le d-1$ 算出来的子数组个数是一样的。也就是说 $d-1$ 同样满足要求，即 `check(d - 1) == true`，这与循环不变量相矛盾。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

关于滑动窗口的原理，请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

[本题视频讲解](https://www.bilibili.com/video/BV1Pw4m1C79N/)（第四题），欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def medianOfUniquenessArray(self, nums: List[int]) -> int:
        n = len(nums)
        k = (n * (n + 1) // 2 + 1) // 2

        def check(upper: int) -> bool:
            cnt = l = 0
            freq = defaultdict(int)
            for r, in_ in enumerate(nums):
                freq[in_] += 1  # 移入右端点
                while len(freq) > upper:  # 窗口内元素过多
                    out = nums[l]
                    freq[out] -= 1  # 移出左端点
                    if freq[out] == 0:
                        del freq[out]
                    l += 1
                cnt += r - l + 1  # 右端点固定为 r 时，有 r-l+1 个合法左端点
                if cnt >= k:
                    return True
            return False

        return bisect_left(range(len(set(nums))), True, 1, key=check)
```

```java [sol-Java]
class Solution {
    public int medianOfUniquenessArray(int[] nums) {
        int n = nums.length;
        long k = ((long) n * (n + 1) / 2 + 1) / 2;
        int left = 0;
        int right = n;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, mid, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int upper, long k) {
        long cnt = 0;
        int l = 0;
        HashMap<Integer, Integer> freq = new HashMap<>();
        for (int r = 0; r < nums.length; r++) {
            freq.merge(nums[r], 1, Integer::sum); // 移入右端点
            while (freq.size() > upper) { // 窗口内元素过多
                int out = nums[l++];
                if (freq.merge(out, -1, Integer::sum) == 0) { // 移出左端点
                    freq.remove(out);
                }
            }
            cnt += r - l + 1; // 右端点固定为 r 时，有 r-l+1 个合法左端点
            if (cnt >= k) {
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
    int medianOfUniquenessArray(vector<int>& nums) {
        int n = nums.size();
        long long k = ((long long) n * (n + 1) / 2 + 1) / 2;

        auto check = [&](int upper) {
            long long cnt = 0;
            int l = 0;
            unordered_map<int, int> freq;
            for (int r = 0; r < n; r++) {
                freq[nums[r]]++; // 移入右端点
                while (freq.size() > upper) { // 窗口内元素过多
                    int out = nums[l++];
                    if (--freq[out] == 0) { // 移出左端点
                        freq.erase(out);
                    }
                }
                cnt += r - l + 1; // 右端点固定为 r 时，有 r-l+1 个合法左端点
                if (cnt >= k) {
                    return true;
                }
            }
            return false;
        };

        int left = 0, right = n;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2
	ans := 1 + sort.Search(n-1, func(upper int) bool {
		upper++
		cnt := 0
		l := 0
		freq := map[int]int{}
		for r, in := range nums {
			freq[in]++ // 移入右端点
			for len(freq) > upper { // 窗口内元素过多
				out := nums[l]
				freq[out]-- // 移出左端点
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			cnt += r - l + 1 // 右端点固定为 r 时，有 r-l+1 个合法左端点
			if cnt >= k {
				return true
			}
		}
		return false
	})
	return ans
}
```

```js [sol-JavaScript]
var medianOfUniquenessArray = function(nums) {
    const n = nums.length;
    const k = Math.floor((n * (n + 1) / 2 + 1) / 2);

    function check(upper) {
        const freq = new Map();
        let cnt = 0, l = 0;
        for (let r = 0; r < n; r++) {
            freq.set(nums[r], (freq.get(nums[r]) ?? 0) + 1); // 移入右端点
            while (freq.size > upper) { // 窗口内元素过多
                const out = nums[l++];
                const f = freq.get(out) - 1;
                if (f === 0) {
                    freq.delete(out); // 移出左端点
                } else {
                    freq.set(out, f);
                }
            }
            cnt += r - l + 1; // 右端点固定为 r 时，有 r-l+1 个合法左端点
            if (cnt >= k) {
                return true;
            }
        }
        return false;
    }

    let left = 0, right = n;
    while (left + 1 < right) {
        const mid = Math.floor((left + right) / 2);
        if (check(mid)) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn median_of_uniqueness_array(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let k = ((n * (n + 1) / 2 + 1) / 2) as i64;

        let check = |upper: usize| -> bool {
            let mut cnt = 0i64;
            let mut l = 0;
            let mut freq = HashMap::new();
            for (r, &x) in nums.iter().enumerate() {
                *freq.entry(x).or_insert(0) += 1; // 移入右端点
                while freq.len() > upper { // 窗口内元素过多
                    let e = freq.entry(nums[l]).or_insert(0);
                    *e -= 1;
                    if *e == 0 {
                        freq.remove(&nums[l]); // 移出左端点
                    }
                    l += 1;
                }
                cnt += (r - l + 1) as i64; // 右端点固定为 r 时，有 r-l+1 个合法左端点
                if cnt >= k {
                    return true;
                }
            }
            false
        };

        let mut left = 0;
        let mut right = n;
        while left + 1 < right {
            let mid = (left + right) / 2;
            if check(mid) {
                right = mid;
            } else {
                left = mid;
            }
        }
        right as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。二分 $\mathcal{O}(\log n)$ 次，每次会跑一个 $\mathcal{O}(n)$ 的滑动窗口。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
