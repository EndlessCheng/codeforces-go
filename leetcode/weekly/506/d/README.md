由于 $n\le 1500$，我们可以枚举子数组的左右端点。

外层循环枚举左端点，内层循环枚举右端点，不断向右扩大子数组长度。

贪心地，**把子数组内较小的数与子数组外较大的数交换**。所以要用两个数据结构，分别维护：

- 子数组内的前 $k$ **小**元素和。
- 子数组外的前 $k$ **大**元素和。

这可以用**对顶堆**或者**值域树状数组**维护，题解用的后者。我们可以在值域树状数组上二分，从而求出第 $k$ 小以及前 $k$ 小元素和，原理见 [OI-wiki](https://oi-wiki.org/ds/fenwick/#%E5%8D%95%E7%82%B9%E4%BF%AE%E6%94%B9%E6%9F%A5%E8%AF%A2%E5%85%A8%E5%B1%80%E7%AC%AC-k-%E5%B0%8F)。

但是，如果子数组内的第 $k$ 小比子数组外的第 $k$ 大还要大，那么不能交换。这意味着，实际交换次数可能小于 $k$。

难道要二分交换次数吗？这样总体时间复杂度是 $\mathcal{O}(n^2\log ^2 n)$，太慢了。

注意到，当子数组长度增加一时，只有一个元素从外面进入子数组，交换次数至多增加 $1$ 或者减少 $1$。

设元素进入子数组前，需要交换 $\textit{needSwap}$ 次。

- 如果子数组内的第 $\textit{needSwap}+1$ 个数 $<$ 子数组外的第 $\textit{needSwap}+1$ 个数，那么需要再交换一次，把 $\textit{needSwap}$ 加一。
- 如果子数组内的第 $\textit{needSwap}$ 个数 $\ge$ 子数组外的第 $\textit{needSwap}$ 个数，那么这一对元素不应交换，把 $\textit{needSwap}$ 减一。

[本题视频讲解](https://www.bilibili.com/video/BV1ptJw6hENZ/?t=20m58s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
# ⚠ 超时了！请看优化后的代码
class FenwickTree:
    def __init__(self, sorted_nums: List[int]):
        self.n = n = len(sorted_nums)
        self.high_bit = 1 << (n.bit_length() - 1)
        self.sorted_nums = sorted_nums
        self.cnt = [0] * (n + 1)
        self.sum = [0] * (n + 1)

    # 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
    # 如果 num < 0，表示减少 -num 个 val
    def update(self, i: int, num: int, val: int) -> None:
        while i < self.n:
            self.cnt[i] += num
            self.sum[i] += val
            i += i & -i

    # 返回第 k 小的数（k 从 1 开始）
    def kth(self, k: int) -> int:
        i = 0
        b = self.high_bit
        while b > 0:
            nxt = i | b
            if nxt < self.n and self.cnt[nxt] < k:
                k -= self.cnt[nxt]
                i = nxt
            b >>= 1
        return self.sorted_nums[i]

    # 返回前 k 小的数之和（k 从 1 开始）
    def pre_sum(self, k: int) -> int:
        s = i = 0
        b = self.high_bit
        while b > 0:
            nxt = i | b
            if nxt < self.n and self.cnt[nxt] < k:
                k -= self.cnt[nxt]
                s += self.sum[nxt]
                i = nxt
            b >>= 1
        # 加上等于第 k 小的数
        return s + self.sorted_nums[i] * k

    def copy(self) -> FenwickTree:
        t = FenwickTree(self.sorted_nums)
        t.cnt[:] = self.cnt
        t.sum[:] = self.sum
        return t


class Solution:
    def maxSum(self, nums: list[int], k: int) -> int:
        # 离散化
        n = len(nums)
        sorted_nums = sorted(set(nums))
        rank = [0] * n  # rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        all_tree = FenwickTree(sorted_nums)  # 包含所有元素的树状数组
        total = 0
        for i, x in enumerate(nums):
            rank[i] = bisect_left(sorted_nums, x) + 1
            all_tree.update(rank[i], 1, x)
            total += x

        ans = -inf

        # 枚举子数组左端点
        for left in range(n):
            in_tree = FenwickTree(sorted_nums)
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

# ⚠ 超时了！请看优化后的代码
```

```java [sol-Java]
class FenwickTree {
    private final int highBit;
    private final int[] sorted;
    private final int[] cnt;
    private final long[] sum;

    public FenwickTree(int[] sorted) {
        int n = sorted.length;
        highBit = 1 << (31 - Integer.numberOfLeadingZeros(n));
        this.sorted = sorted;
        cnt = new int[n + 1];
        sum = new long[n + 1];
    }

    // 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
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
        for (int b = highBit; b > 0; b >>= 1) {
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
        long s = 0;
        int i = 0;
        for (int b = highBit; b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.length && cnt[nxt] < k) {
                k -= cnt[nxt];
                s += sum[nxt];
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return s + (long) sorted[i] * k;
    }

    public FenwickTree copy() {
        FenwickTree f = new FenwickTree(sorted);
        System.arraycopy(cnt, 0, f.cnt, 0, cnt.length);
        System.arraycopy(sum, 0, f.sum, 0, sum.length);
        return f;
    }
}

class Solution {
    public long maxSum(int[] nums, int k) {
        // 离散化
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        int[] rank = new int[n]; // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        FenwickTree allTree = new FenwickTree(sorted); // 包含所有元素的树状数组
        long total = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            rank[i] = Arrays.binarySearch(sorted, x) + 1;
            allTree.update(rank[i], 1, x);
            total += x;
        }

        long ans = Long.MIN_VALUE;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            FenwickTree inTree = new FenwickTree(sorted);
            FenwickTree outTree = allTree.copy();
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
class FenwickTree {
    const int high_bit;
    const vector<int>& sorted;
    vector<int> cnt;
    vector<long long> sum;

public:
    FenwickTree(const vector<int>& sorted) : 
        cnt(sorted.size() + 1), 
        sum(sorted.size() + 1), 
        sorted(sorted), 
        high_bit(1 << (bit_width(sorted.size()) - 1)) {}

    // 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
    // 如果 num < 0，表示减少 -num 个 val
    void update(int i, int num, int val) {
        for (; i < cnt.size(); i += i & -i) {
            cnt[i] += num;
            sum[i] += val;
        }
    }

    // 返回第 k 小的数（k 从 1 开始）
    int kth(int k) const {
        int i = 0;
        for (int b = high_bit; b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.size() && cnt[nxt] < k) {
                k -= cnt[nxt];
                i = nxt;
            }
        }
        return sorted[i];
    }

    // 返回前 k 小的数之和（k 从 1 开始）
    long long pre_sum(int k) const {
        long long s = 0;
        int i = 0;
        for (int b = high_bit; b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.size() && cnt[nxt] < k) {
                k -= cnt[nxt];
                s += sum[nxt];
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return s + 1LL * sorted[i] * k;;
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
        vector<int> rank(n); // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        FenwickTree all_tree(sorted); // 包含所有元素的树状数组
        long long total = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            rank[i] = ranges::lower_bound(sorted, x) - sorted.begin() + 1;
            all_tree.update(rank[i], 1, x);
            total += x;
        }

        long long ans = LLONG_MIN;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            FenwickTree in_tree(sorted);
            FenwickTree out_tree = all_tree;
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
type pair struct{ cnt, sum int }
type fenwick struct {
	t       []pair
	sorted  []int
	highBit int
}

func newFenwickTree(sorted []int) fenwick {
	n := len(sorted)
	return fenwick{
		t:       make([]pair, n+1),
		sorted:  sorted,
		highBit: 1 << (bits.Len(uint(n)) - 1),
	}
}

// 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
// 如果 num < 0，表示减少 -num 个 val
func (f *fenwick) update(i, num, val int) {
	for ; i < len(f.t); i += i & -i {
		f.t[i].cnt += num
		f.t[i].sum += val
	}
}

// 返回第 k 小的数（k 从 1 开始）
func (f *fenwick) kth(k int) int {
	i := 0
	for b := f.highBit; b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
			k -= f.t[nxt].cnt
			i = nxt
		}
	}
	return f.sorted[i]
}

// 返回前 k 小的数之和（k 从 1 开始）
func (f *fenwick) preSum(k int) (s int) {
	i := 0
	for b := f.highBit; b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
			k -= f.t[nxt].cnt
			s += f.t[nxt].sum
			i = nxt
		}
	}
	// 加上等于第 k 小的数
	s += f.sorted[i] * k
	return
}

func maxSum(nums []int, k int) int64 {
	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allTree := newFenwickTree(sorted) // 包含所有元素的树状数组
	total := 0
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		allTree.update(rank[i], 1, x)
		total += x
	}

	inTree := newFenwickTree(sorted)
	outTree := newFenwickTree(sorted)
	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		clear(inTree.t)
		copy(outTree.t, allTree.t)
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
				if inTree.kth(needSwap+1) < outTree.kth(n-sz-needSwap) {
					inc = true
					needSwap++
				}
			}

			if !inc && needSwap > 0 {
				// 是否要减少交换次数
				if inTree.kth(needSwap) >= outTree.kth(n-sz-needSwap+1) {
					needSwap--
				}
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			if needSwap > 0 {
				inSum := inTree.preSum(needSwap)
				outSum := total - subSum - outTree.preSum(n-sz-needSwap)
				delta = outSum - inSum
			}

			ans = max(ans, subSum+delta)
		}
	}

	return int64(ans)
}
```

## 优化

### 1) 特判

设 $\textit{nums}$ 中的正数个数为 $p$。跑一个 [定长滑窗](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)，如果存在一个长为 $p$ 的窗口，其中 $正数个数+k\ge p$，则说明可以在 $k$ 次操作内把窗口外的正数全部换入窗口内，从而把所有正数都聚在一起，此时直接返回所有正数之和（这是答案的上界）。

特别地，如果 $p=0$，没有正数，则返回 $\max(\textit{nums})$（题目不允许子数组为空）。

在下面的讨论中，$\textit{nums}$ 包含正数。

### 2) 只需考虑正数与负数的交换

首先，对于和最大的子数组 $[\ell,r]$，有如下简单的性质：

1. $\textit{nums}[\ell-1]$ 和 $\textit{nums}[r+1]$（如果有）一定都 $\le 0$。**反证法**：如果其中有正数，则子数组扩大更好，矛盾。
2. $\textit{nums}[\ell]$ 和 $\textit{nums}[r]$ 一定都 $\ge 0$。**反证法**：如果其中有负数，则子数组缩小更好，矛盾。

这些性质有什么用？

1. 与其交换子数组内的正数（以及 $0$）与子数组外的正数，不如把子数组外的正数与 $\textit{nums}[\ell-1]$ 或者 $\textit{nums}[r+1]$ 交换，然后扩大子数组，包含这个正数，可以得到更大的子数组和。
2. 与其交换子数组内的负数与子数组外的负数，不如把子数组内的负数与 $\textit{nums}[\ell]$ 或者 $\textit{nums}[r]$ 交换，然后缩小子数组，移除这个负数，可以得到更大的子数组和。

所以我们只需考虑子数组内的负数与子数组外的正数的交换。

这样计算交换次数 $\textit{needSwap}$ 就很简单了。设子数组内的负数个数为 $\textit{negCnt}$，子数组外的正数个数为 $\textit{posCnt}$，那么

$$
\textit{needSwap} = \min(\textit{negCnt}, \textit{posCnt}, k)
$$

```py [sol-Python3]
class FenwickTree:
    def __init__(self, sorted_nums: List[int]):
        self.n = n = len(sorted_nums)
        self.high_bit = 1 << (n.bit_length() - 1)
        self.sorted_nums = sorted_nums
        self.cnt = [0] * (n + 1)
        self.sum = [0] * (n + 1)

    # 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
    # 如果 num < 0，表示减少 -num 个 val
    def update(self, i: int, num: int, val: int) -> None:
        while i < self.n:
            self.cnt[i] += num
            self.sum[i] += val
            i += i & -i

    # 返回前 k 小的数之和（k 从 1 开始）
    def pre_sum(self, k: int) -> int:
        s = i = 0
        b = self.high_bit
        while b > 0:
            nxt = i | b
            if nxt < self.n and self.cnt[nxt] < k:
                k -= self.cnt[nxt]
                s += self.sum[nxt]
                i = nxt
            b >>= 1
        # 加上等于第 k 小的数
        return s + self.sorted_nums[i] * k

    def copy(self) -> FenwickTree:
        t = FenwickTree(self.sorted_nums)
        t.cnt[:] = self.cnt
        t.sum[:] = self.sum
        return t


class Solution:
    def maxSum(self, nums: list[int], k: int) -> int:
        # O(n) 特判：能否把正数都聚在一起
        all_pos_sum = all_pos_cnt = 0
        for x in nums:
            if x > 0:
                all_pos_sum += x
                all_pos_cnt += 1
        if all_pos_cnt == 0:  # 没有正数
            return max(nums)
        # 定长滑动窗口模板，窗口长度为 all_pos_cnt
        cnt = 0
        for i, x in enumerate(nums):
            if x > 0:
                cnt += 1
            left = i - all_pos_cnt + 1
            if left < 0:
                continue
            if cnt + k >= all_pos_cnt:  # 可以把正数都聚在一起
                return all_pos_sum
            if nums[left] > 0:
                cnt -= 1

        # 离散化
        n = len(nums)
        sorted_nums = sorted(set(nums))
        rank = [0] * n  # rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        all_pos_tree = FenwickTree(sorted_nums)  # 包含所有正数的树状数组
        for i, x in enumerate(nums):
            rank[i] = bisect_left(sorted_nums, x) + 1
            if x > 0:
                all_pos_tree.update(rank[i], 1, x)

        ans = -inf

        # 枚举子数组左端点
        for left in range(n):
            neg_tree = FenwickTree(sorted_nums)
            pos_tree = all_pos_tree.copy()
            pos_sum = all_pos_sum
            pos_cnt = all_pos_cnt
            neg_cnt = 0
            sub_sum = 0

            # 枚举子数组右端点
            for right in range(left, n):
                # x 从子数组外移到子数组内
                x = nums[right]
                rk = rank[right]
                sub_sum += x
                if x > 0:
                    pos_tree.update(rk, -1, -x)
                    pos_sum -= x
                    pos_cnt -= 1
                elif x < 0:
                    neg_tree.update(rk, 1, x)
                    neg_cnt += 1

                # 计算通过交换导致的元素和的增量
                delta = 0
                need_swap = min(neg_cnt, pos_cnt, k)
                if need_swap > 0:
                    in_sum = neg_tree.pre_sum(need_swap)
                    out_sum = pos_sum - pos_tree.pre_sum(pos_cnt - need_swap)
                    delta = out_sum - in_sum

                ans = max(ans, sub_sum + delta)

        return ans
```

```java [sol-Java]
class FenwickTree {
    private final int highBit;
    private final int[] sorted;
    private final int[] cnt;
    private final long[] sum;

    public FenwickTree(int[] sorted) {
        int n = sorted.length;
        highBit = 1 << (31 - Integer.numberOfLeadingZeros(n));
        this.sorted = sorted;
        cnt = new int[n + 1];
        sum = new long[n + 1];
    }

    // 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
    // 如果 num < 0，表示减少 -num 个 val
    public void update(int i, int num, int val) {
        for (; i < cnt.length; i += i & -i) {
            cnt[i] += num;
            sum[i] += val;
        }
    }

    // 返回前 k 小的数之和（k 从 1 开始）
    public long preSum(int k) {
        long s = 0;
        int i = 0;
        for (int b = highBit; b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < cnt.length && cnt[nxt] < k) {
                k -= cnt[nxt];
                s += sum[nxt];
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return s + (long) sorted[i] * k;
    }

    public FenwickTree copy() {
        FenwickTree f = new FenwickTree(sorted);
        System.arraycopy(cnt, 0, f.cnt, 0, cnt.length);
        System.arraycopy(sum, 0, f.sum, 0, sum.length);
        return f;
    }
}

class Solution {
    public long maxSum(int[] nums, int k) {
        int n = nums.length;
        // O(n) 特判：能否把正数都聚在一起
        long allPosSum = 0;
        int allPosCnt = 0;
        int mx = Integer.MIN_VALUE;
        for (int x : nums) {
            if (x > 0) {
                allPosSum += x;
                allPosCnt++;
            } else {
                mx = Math.max(mx, x);
            }
        }
        if (allPosCnt == 0) { // 没有正数
            return mx;
        }
        // 定长滑动窗口模板，窗口长度为 allPosCnt
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (x > 0) {
                cnt++;
            }
            int left = i - allPosCnt + 1;
            if (left < 0) {
                continue;
            }
            if (cnt + k >= allPosCnt) { // 可以把正数都聚在一起
                return allPosSum;
            }
            if (nums[left] > 0) {
                cnt--;
            }
        }

        // 离散化
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        int[] rank = new int[n]; // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        FenwickTree allPosTree = new FenwickTree(sorted); // 包含所有正数的树状数组
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            rank[i] = Arrays.binarySearch(sorted, x) + 1;
            if (x > 0) {
                allPosTree.update(rank[i], 1, x);
            }
        }

        long ans = Long.MIN_VALUE;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            FenwickTree negTree = new FenwickTree(sorted);
            FenwickTree posTree = allPosTree.copy();
            long posSum = allPosSum;
            int posCnt = allPosCnt;
            int negCnt = 0;
            long subSum = 0;

            // 枚举子数组右端点
            for (int right = left; right < n; right++) {
                // x 从子数组外移到子数组内
                int x = nums[right];
                int rk = rank[right];
                subSum += x;
                if (x > 0) {
                    posTree.update(rk, -1, -x);
                    posSum -= x;
                    posCnt--;
                } else if (x < 0) {
                    negTree.update(rk, 1, x);
                    negCnt++;
                }

                // 计算通过交换导致的元素和的增量
                long delta = 0;
                int needSwap = Math.min(Math.min(negCnt, posCnt), k);
                if (needSwap > 0) {
                    long inSum = negTree.preSum(needSwap);
                    long outSum = posSum - posTree.preSum(posCnt - needSwap);
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
class FenwickTree {
    const int high_bit;
    const vector<int>& sorted;
    vector<pair<int, long long>> t; // {cnt, sum}

public:
    FenwickTree(const vector<int>& sorted) : 
        t(sorted.size() + 1), 
        sorted(sorted), 
        high_bit(1 << (bit_width(sorted.size()) - 1)) {}

    // 添加 num 个 val，其中 val 离散化后的值为 i
    // 如果 num < 0，表示减少 -num 个 val
    void update(int i, int num, int val) {
        for (; i < t.size(); i += i & -i) {
            t[i].first += num;
            t[i].second += val;
        }
    }

    // 返回前 k 小的数之和（k 从 1 开始）
    long long pre_sum(int k) const {
        long long s = 0;
        int i = 0;
        for (int b = high_bit; b > 0; b >>= 1) {
            int nxt = i | b;
            if (nxt < t.size() && t[nxt].first < k) {
                k -= t[nxt].first;
                s += t[nxt].second;
                i = nxt;
            }
        }
        // 加上等于第 k 小的数
        return s + 1LL * sorted[i] * k;;
    }
};

class Solution {
public:
    long long maxSum(vector<int>& nums, int k) {
        int n = nums.size();
        // O(n) 特判：能否把正数都聚在一起
        int all_pos_sum = 0;
        int all_pos_cnt = 0;
        for (int x : nums) {
            if (x > 0) {
                all_pos_sum += x;
                all_pos_cnt++;
            }
        }
        if (all_pos_cnt == 0) { // 没有正数
            return ranges::max(nums);
        }
        // 定长滑动窗口模板，窗口长度为 all_pos_cnt
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            cnt += x > 0;
            int left = i - all_pos_cnt + 1;
            if (left < 0) {
                continue;
            }
            if (cnt + k >= all_pos_cnt) { // 可以把正数都聚在一起
                return all_pos_sum;
            }
            cnt -= nums[left] > 0;
        }

        // 离散化
        auto sorted = nums;
        ranges::sort(sorted);
        sorted.erase(ranges::unique(sorted).begin(), sorted.end());
        vector<int> rank(n); // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
        FenwickTree all_pos_tree(sorted); // 包含所有正数的树状数组
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            rank[i] = ranges::lower_bound(sorted, x) - sorted.begin() + 1;
            if (x > 0) {
                all_pos_tree.update(rank[i], 1, x);
            }
        }

        long long ans = LLONG_MIN;

        // 枚举子数组左端点
        for (int left = 0; left < n; left++) {
            FenwickTree neg_tree(sorted);
            FenwickTree pos_tree = all_pos_tree;
            long long pos_sum = all_pos_sum;
            int pos_cnt = all_pos_cnt;
            int neg_cnt = 0;
            long long sub_sum = 0;

            // 枚举子数组右端点
            for (int right = left; right < n; right++) {
                // x 从子数组外移到子数组内
                int x = nums[right];
                int rk = rank[right];
                sub_sum += x;
                if (x > 0) {
                    pos_tree.update(rk, -1, -x);
                    pos_sum -= x;
                    pos_cnt--;
                } else if (x < 0) {
                    neg_tree.update(rk, 1, x);
                    neg_cnt++;
                }

                long long delta = 0;
                int need_swap = min({neg_cnt, pos_cnt, k});
                if (need_swap > 0) {
                    long long in_sum = neg_tree.pre_sum(need_swap);
                    long long out_sum = pos_sum - pos_tree.pre_sum(pos_cnt - need_swap);
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
type pair struct{ cnt, sum int }
type fenwick struct {
	t       []pair
	sorted  []int
	highBit int
}

func newFenwickTree(sorted []int) fenwick {
	n := len(sorted)
	return fenwick{
		t:       make([]pair, n+1),
		sorted:  sorted,
		highBit: 1 << (bits.Len(uint(n)) - 1),
	}
}

// 添加 num 个 val，其中 val 离散化后的值为 i（i 从 1 开始）
// 如果 num < 0，表示减少 -num 个 val
func (f *fenwick) update(i, num, val int) {
	for ; i < len(f.t); i += i & -i {
		f.t[i].cnt += num
		f.t[i].sum += val
	}
}

// 返回前 k 小的数之和（k 从 1 开始）
func (f *fenwick) preSum(k int) (s int) {
	i := 0
	for b := f.highBit; b > 0; b >>= 1 {
		if nxt := i | b; nxt < len(f.t) && f.t[nxt].cnt < k {
			k -= f.t[nxt].cnt
			s += f.t[nxt].sum
			i = nxt
		}
	}
	// 加上等于第 k 小的数
	s += f.sorted[i] * k
	return
}

func maxSum(nums []int, k int) int64 {
	// O(n) 特判：能否把正数都聚在一起
	allPosSum := 0
	allPosCnt := 0
	for _, x := range nums {
		if x > 0 {
			allPosSum += x
			allPosCnt++
		}
	}
	if allPosCnt == 0 { // 没有正数
		return int64(slices.Max(nums))
	}
	// 定长滑动窗口模板，窗口长度为 allPosCnt
	cnt := 0
	for i, x := range nums {
		if x > 0 {
			cnt++
		}
		left := i - allPosCnt + 1
		if left < 0 {
			continue
		}
		if cnt+k >= allPosCnt { // 可以把正数都聚在一起
			return int64(allPosSum)
		}
		if nums[left] > 0 {
			cnt--
		}
	}

	// 离散化
	n := len(nums)
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	rank := make([]int, n) // rank[i] 是 nums[i] 离散化后的值（从 1 开始）
	allPosTree := newFenwickTree(sorted) // 包含所有正数的树状数组
	for i, x := range nums {
		rank[i] = sort.SearchInts(sorted, x) + 1
		if x > 0 {
			allPosTree.update(rank[i], 1, x)
		}
	}

	negTree := newFenwickTree(sorted)
	posTree := newFenwickTree(sorted)
	ans := math.MinInt

	// 枚举子数组左端点
	for left := range nums {
		clear(negTree.t)
		copy(posTree.t, allPosTree.t)
		posSum := allPosSum
		posCnt := allPosCnt
		negCnt := 0
		subSum := 0

		// 枚举子数组右端点
		for right := left; right < n; right++ {
			// x 从子数组外移到子数组内
			x := nums[right]
			rk := rank[right]
			subSum += x
			if x > 0 {
				posTree.update(rk, -1, -x)
				posSum -= x
				posCnt--
			} else if x < 0 {
				negTree.update(rk, 1, x)
				negCnt++
			}

			// 计算通过交换导致的元素和的增量
			delta := 0
			needSwap := min(negCnt, posCnt, k)
			if needSwap > 0 {
				inSum := negTree.preSum(needSwap)
				outSum := posSum - posTree.preSum(posCnt-needSwap)
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
