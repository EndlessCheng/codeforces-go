由于 $n\le 1500$，我们可以枚举子数组的左右端点。

外层循环枚举左端点，内层循环枚举右端点，不断向右扩大子数组长度。

贪心地，**把子数组内较小的数与子数组外较大的数交换**。所以要用两个数据结构，分别维护：

- 子数组内的前 $k$ **小**元素和。
- 子数组外的前 $k$ **大**元素和。

这可以用**对顶堆**或者**值域树状数组**维护。题解用的后者，原理见 [OI-wiki](https://oi-wiki.org/ds/fenwick/#%E5%8D%95%E7%82%B9%E4%BF%AE%E6%94%B9%E6%9F%A5%E8%AF%A2%E5%85%A8%E5%B1%80%E7%AC%AC-k-%E5%B0%8F)。

但是，如果子数组内的第 $k$ 小比子数组外的第 $k$ 大还要大，那么不能交换。这意味着，实际交换次数可能小于 $k$。

难道要二分交换次数吗？这样总体时间复杂度是 $\mathcal{O}(n^2\log ^2 n)$，太慢了。

注意到，当子数组长度增加一时，只有一个元素从外面进入子数组，交换次数至多增加 $1$ 或者减少 $1$。

设元素进入子数组前，需要交换 $\textit{needSwap}$ 次。

- 如果子数组内的第 $\textit{needSwap}+1$ 个数 $<$ 子数组外的第 $\textit{needSwap}+1$ 个数，那么需要再交换一次，把 $\textit{needSwap}$ 加一。
- 如果子数组内的第 $\textit{needSwap}$ 个数 $\ge$ 子数组外的第 $\textit{needSwap}$ 个数，那么这一对元素不应交换，把 $\textit{needSwap}$ 减一。

[本题视频讲解](https://www.bilibili.com/video/BV1ptJw6hENZ/?t=20m58s)，欢迎点赞关注~

```py [sol-Python3]
# 超时了！请看优化后的代码
class Fenwick:
    def __init__(self, n: int, sorted: list[int], width: int):
        self.n = n
        self.width = width
        self.sorted = sorted
        self.cnt = [0] * n
        self.sum = [0] * n

    def copy(self) -> Fenwick:
        t = Fenwick(self.n, self.sorted, self.width)
        t.cnt = self.cnt.copy()
        t.sum = self.sum.copy()
        return t

    # 添加 num 个 val，其中 val 离散化后的值为 i
    # 如果 num < 0，表示减少 -num 个 val
    def update(self, i: int, num: int, val: int) -> None:
        while i < self.n:
            self.cnt[i] += num
            self.sum[i] += val
            i += i & -i

    # 返回第 k 小的数（k 从 1 开始）
    def kth(self, k: int) -> int:
        i = 0
        b = 1 << (self.width - 1)
        while b > 0:
            nxt = i | b
            if nxt < self.n and self.cnt[nxt] < k:
                k -= self.cnt[nxt]
                i = nxt
            b >>= 1
        return self.sorted[i]

    # 返回前 k 小的数之和（k 从 1 开始）
    def pre_sum(self, k: int) -> int:
        res = 0
        i = 0
        b = 1 << (self.width - 1)
        while b > 0:
            nxt = i | b
            if nxt < self.n and self.cnt[nxt] < k:
                k -= self.cnt[nxt]
                res += self.sum[nxt]
                i = nxt
            b >>= 1
        # 加上等于第 k 小的数
        return res + self.sorted[i] * k


class Solution:
    def maxSum(self, nums: list[int], k: int) -> int:
        # 离散化
        n = len(nums)
        sorted_ = nums.copy()
        sorted_.sort()
        m = len(sorted_)
        width_m = m.bit_length()
        rank = [0] * n  # rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        all_tree = Fenwick(m + 1, sorted_, width_m)  # 包含所有元素的树状数组
        total = 0
        for i, x in enumerate(nums):
            rank[i] = bisect_left(sorted_, x) + 1
            all_tree.update(rank[i], 1, x)
            total += x

        ans = -inf

        # 枚举子数组左端点
        for left in range(n):
            in_tree = Fenwick(m + 1, sorted_, width_m)
            out_tree = all_tree.copy()
            need_swap = sub_sum = 0

            # 枚举子数组右端点
            for right in range(left, n):
                # x 从子数组外移到子数组内
                x = nums[right]
                rk = rank[right]
                sub_sum += x
                in_tree.update(rk, 1, x)
                out_tree.update(rk, -1, -x)

                inc = False
                sz = right - left + 1
                if need_swap < k and need_swap < sz and need_swap < n - sz:
                    # 能否多交换一次
                    if in_tree.kth(need_swap + 1) < out_tree.kth(n - sz - need_swap):
                        inc = True
                        need_swap += 1

                if not inc and need_swap > 0:
                    # 是否要减少交换次数
                    if in_tree.kth(need_swap) >= out_tree.kth(n - sz - need_swap + 1):
                        need_swap -= 1

                # 计算通过交换导致的元素和的增量
                delta = 0
                if need_swap > 0:
                    in_sum = in_tree.pre_sum(need_swap)
                    out_sum = total - sub_sum - out_tree.pre_sum(n - sz - need_swap)
                    delta = out_sum - in_sum

                ans = max(ans, sub_sum + delta)

        return ans
```

```java [sol-Java]
class Fenwick {
    private final int width;
    private final int[] sorted;
    private final int[] cnt;
    private final long[] sum;

    public Fenwick(int n, int[] sorted, int width) {
        this.width = width;
        this.sorted = sorted;
        cnt = new int[n];
        sum = new long[n];
    }

    // 添加 num 个 val，其中 val 离散化后的值为 i
    // 如果 num < 0，表示减少 -num 个 val
    public void update(int i, int num, int val) {
        for (; i < cnt.length; i += i & -i) {
            cnt[i] += num;
            sum[i] += val;
        }
    }

    // 返回第 k 小的数（k 从 1 开始）
    public int kth(int k) {
        int i = 0;
        for (int b = 1 << (width - 1); b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.length && cnt[nxt] < k) {
                k -= cnt[nxt];
                i = nxt;
            }
        }
        return sorted[i];
    }

    // 返回前 k 小的数之和（k 从 1 开始）
    public long preSum(int k) {
        long res = 0;
        int i = 0;
        for (int b = 1 << (width - 1); b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.length && cnt[nxt] < k) {
                k -= cnt[nxt];
                res += sum[nxt];
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return res + 1L * sorted[i] * k;
    }

    public Fenwick copy() {
        Fenwick f = new Fenwick(cnt.length, sorted, width);
        System.arraycopy(this.cnt, 0, f.cnt, 0, cnt.length);
        System.arraycopy(this.sum, 0, f.sum, 0, sum.length);
        return f;
    }
}

class Solution {
    public long maxSum(int[] nums, int k) {
        // 离散化
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        int widthN = 32 - Integer.numberOfLeadingZeros(n);
        int[] rank = new int[n]; // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        Fenwick allTree = new Fenwick(n + 1, sorted, widthN); // 包含所有元素的树状数组
        long total = 0;
        for (int i = 0; i < n; i++) {
            rank[i] = Arrays.binarySearch(sorted, nums[i]) + 1;
            allTree.update(rank[i], 1, nums[i]);
            total += nums[i];
        }

        long ans = Long.MIN_VALUE;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            Fenwick inTree = new Fenwick(n + 1, sorted, widthN);
            Fenwick outTree = allTree.copy();
            int needSwap = 0;
            long subSum = 0;

            // 枚举子数组右端点
            for (int right = left; right < n; right++) {
                // x 从子数组外移到子数组内
                int x = nums[right];
                int rk = rank[right];
                subSum += x;
                inTree.update(rk, 1, x);
                outTree.update(rk, -1, -x);

                boolean inc = false;
                int sz = right - left + 1;
                if (needSwap < k && needSwap < sz && needSwap < n - sz) {
                    // 能否多交换一次
                    if (inTree.kth(needSwap + 1) < outTree.kth(n - sz - needSwap)) {
                        inc = true;
                        needSwap++;
                    }
                }

                if (!inc && needSwap > 0) {
                    // 是否要减少交换次数
                    if (inTree.kth(needSwap) >= outTree.kth(n - sz - needSwap + 1)) {
                        needSwap--;
                    }
                }

                // 计算通过交换导致的元素和的增量
                long delta = 0;
                if (needSwap > 0) {
                    long inSum = inTree.preSum(needSwap);
                    long outSum = total - subSum - outTree.preSum(n - sz - needSwap);
                    delta = outSum - inSum;
                }

                ans = Math.max(ans, subSum + delta);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Fenwick {
    const int width;
    const vector<int>& sorted;
    vector<int> cnt;
    vector<long long> sum;

public:
    Fenwick(int n, const vector<int>& sorted, int width) : cnt(n), sum(n), sorted(sorted), width(width) {}

    // 添加 num 个 val，其中 val 离散化后的值为 i
    // 如果 num < 0，表示减少 -num 个 val
    void update(int i, int num, int val) {
        for (; i < cnt.size(); i += i & -i) {
            cnt[i] += num;
            sum[i] += val;
        }
    }

    // 返回第 k 小的数（k 从 1 开始）
    int kth(int k) {
        int i = 0;
        for (int b = 1 << (width - 1); b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.size() && cnt[nxt] < k) {
                k -= cnt[nxt];
                i = nxt;
            }
        }
        return sorted[i];
    }

    // 返回前 k 小的数之和（k 从 1 开始）
    long long pre_sum(int k) {
        long long res = 0;
        int i = 0;
        for (int b = 1 << (width - 1); b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.size() && cnt[nxt] < k) {
                k -= cnt[nxt];
                res += sum[nxt];
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return res + 1LL * sorted[i] * k;;
    }
};

class Solution {
public:
    long long maxSum(vector<int>& nums, int k) {
        // 离散化
        int n = nums.size();
        vector<int> sorted = nums;
        ranges::sort(sorted);
        sorted.erase(ranges::unique(sorted).begin(), sorted.end());
        int m = sorted.size();
        int width_m = bit_width(1u * m);
        vector<int> rank(n); // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        Fenwick all_tree(m + 1, sorted, width_m); // 包含所有元素的树状数组
        long long total = 0;
        for (int i = 0; i < n; i++) {
            rank[i] = ranges::lower_bound(sorted, nums[i]) - sorted.begin() + 1;
            all_tree.update(rank[i], 1, nums[i]);
            total += nums[i];
        }

        long long ans = LLONG_MIN;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            Fenwick in_tree(m + 1, sorted, width_m);
            Fenwick out_tree = all_tree;
            int need_swap = 0;
            long long sub_sum = 0;

            // 枚举子数组右端点
            for (int right = left; right < n; right++) {
                // x 从子数组外移到子数组内
                int x = nums[right];
                int rk = rank[right];
                sub_sum += x;
                in_tree.update(rk, 1, x);
                out_tree.update(rk, -1, -x);

                bool inc = false;
                int sz = right - left + 1;
                if (need_swap < k && need_swap < sz && need_swap < n - sz) {
                    // 能否多交换一次
                    if (in_tree.kth(need_swap + 1) < out_tree.kth(n - sz - need_swap)) {
                        inc = true;
                        need_swap++;
                    }
                }

                if (!inc && need_swap > 0) {
                    // 是否要减少交换次数
                    if (in_tree.kth(need_swap) >= out_tree.kth(n - sz - need_swap + 1)) {
                        need_swap--;
                    }
                }

                // 计算通过交换导致的元素和的增量
                long long delta = 0;
                if (need_swap > 0) {
                    long long in_sum = in_tree.pre_sum(need_swap);
                    long long out_sum = total - sub_sum - out_tree.pre_sum(n - sz - need_swap);
                    delta = out_sum - in_sum;
                }

                ans = max(ans, sub_sum + delta);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
var widthM int

type pair struct{ cnt, sum int }
type fenwick []pair

// 添加 num 个 val，其中 val 离散化后的值为 i
// 如果 num < 0，表示减少 -num 个 val
func (t fenwick) update(i, num, val int) {
	for ; i < len(t); i += i & -i {
		t[i].cnt += num
		t[i].sum += val
	}
}

// 返回第 k 小的数（k 从 1 开始）
func (t fenwick) kth(k int, sorted []int) int {
	i := 0
	for b := 1 << (widthM - 1); b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			i = nxt
		}
	}
	return sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (t fenwick) preSum(k int, sorted []int) (res int) {
	i := 0
	for b := 1 << (widthM - 1); b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(t) && t[nxt].cnt < k {
			k -= t[nxt].cnt
			res += t[nxt].sum
			i = nxt
		}
	}
	// 加上等于第 k 小的数
	res += sorted[i] * k
	return
}

func maxSum(nums []int, k int) int64 {
	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	widthM = bits.Len(uint(m))
	rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allTree := make(fenwick, m+1) // 包含所有元素的树状数组
	total := 0
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		allTree.update(rank[i], 1, x)
		total += x
	}

	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		inTree := make(fenwick, m+1)
		outTree := slices.Clone(allTree)
		needSwap := 0
		subSum := 0

		// 枚举子数组右端点
		for right := left; right < n; right++ {
			// x 从子数组外移到子数组内
			x := nums[right]
			rk := rank[right]
			subSum += x
			inTree.update(rk, 1, x)
			outTree.update(rk, -1, -x)

			inc := false
			sz := right - left + 1
			if needSwap < k && needSwap < sz && needSwap < n-sz {
				// 能否多交换一次
				if inTree.kth(needSwap+1, sorted) < outTree.kth(n-sz-needSwap, sorted) {
					inc = true
					needSwap++
				}
			}

			if !inc && needSwap > 0 {
				// 是否要减少交换次数
				if inTree.kth(needSwap, sorted) >= outTree.kth(n-sz-needSwap+1, sorted) {
					needSwap--
				}
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			if needSwap > 0 {
				inSum := inTree.preSum(needSwap, sorted)
				outSum := total - subSum - outTree.preSum(n-sz-needSwap, sorted)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§5.7 对顶堆**」和「**§8.1 树状数组**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
