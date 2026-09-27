## 分析

首先，对于不取反的情况，做法类似 [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)，[我的题解](https://leetcode.cn/problems/subarray-sums-divisible-by-k/solutions/3815616/qian-zhui-he-yu-ha-xi-biao-shi-zi-bian-x-qxc5/)。改成维护前缀和模 $k$ 的首次出现的下标，从而计算最大子数组长度。

下面只讨论恰好取反一个元素的情况。

设 $\textit{nums}$ 的前缀和数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

$\textit{nums}$ 的子数组 $[\ell, r)$ 的元素和为 $s[r] - s[\ell]$，长为 $r - \ell$。

取反一个在 $[\ell, r)$ 中的元素 $\textit{nums}[i]$，子数组的和变成

$$
s[r] - s[\ell] - 2\cdot \textit{nums}[i]
$$

问题相当于：

- 对于下标三元组 $(\ell,i,r)$，计算 $r - \ell$ 的最大值，满足 $0\le \ell\le i < r\le n$ 且 $s[r] - s[\ell] - 2\cdot \textit{nums}[i]$ 是 $k$ 的倍数。其中 $n$ 是 $\textit{nums}$ 的长度。

## 方法一：枚举前缀和

如果直接枚举所有 $(\ell, r)$，这有 $\mathcal{O}(n^2)$ 个，太慢了。

本题 $k\le 3000$，我们可以枚举前缀和的**值**（而不是下标）。

枚举 $s[\ell]\bmod k$ 和 $s[r]\bmod k$，那么贪心地：

- 对于多个相同的 $s[\ell]\bmod k$，选择 $\ell$ 最小的，这样不仅子数组更长，也更能包含满足要求的 $i$。
- 对于多个相同的 $s[r]\bmod k$，选择 $r$ 最大的，这样不仅子数组更长，也更能包含满足要求的 $i$。

现在问题变成：

- 判断在 $[\ell, r)$ 中是否存在 $i$，满足 $\ell\le i < r$ 且 $s[r] - s[\ell] - 2\cdot \textit{nums}[i]$ 是 $k$ 的倍数。

根据 [同余理论](https://leetcode.cn/circle/discuss/mDfnkW/)，$s[r] - s[\ell] - 2\cdot \textit{nums}[i]$ 是 $k$ 的倍数等价于

$$
s[r] - s[\ell] \equiv 2\cdot \textit{nums}[i] \pmod k
$$

我们可以把 $2\cdot \textit{nums}[i]\bmod k$ 的所有出现位置保存到数组中，然后在这个数组中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) 第一个 $\ge \ell$ 的下标 $i$，如果发现 $i < r$，那么存在满足要求的 $i$，用子数组长度 $r - \ell$ 更新答案的最大值。**相似题目**：[2080. 区间内查询数字的频率](https://leetcode.cn/problems/range-frequency-queries/)。

**最优性优化**：在二分之前，如果发现 $r-\ell \le \textit{ans}$，那么 $\textit{ans}$ 不会变大，无需二分。

[本题视频讲解](https://www.bilibili.com/video/BV1v9a86hEkp/?t=14m2s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestSubarray(self, nums: list[int], k: int) -> int:
        num_pos = [[] for _ in range(k)]  # 2*nums[i] % k 出现的所有位置
        first_pos = [-1] * k  # 前缀和 % k 首次出现的位置
        last_pos = [-1] * k  # 前缀和 % k 最后一次出现的位置
        first_pos[0] = last_pos[0] = 0
        s = 0  # 前缀和
        ans = 0

        for i, x in enumerate(nums):
            num_pos[x * 2 % k].append(i)
            s = (s + x) % k
            r = i + 1
            l = first_pos[s]
            if l < 0:
                first_pos[s] = r
            else:
                # 不取反时的最大长度
                ans = max(ans, r - l)
            last_pos[s] = r

        # 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
        for sl, l in enumerate(first_pos):
            if l < 0:
                continue
            for sr, r in enumerate(last_pos):
                if r - l <= ans:  # 最优性优化：ans 无法增大
                    continue
                # 2*nums[i]%k = (sr-sl)%k
                pos = num_pos[(sr - sl) % k]
                idx = bisect_left(pos, l)
                if idx < len(pos) and pos[idx] < r:
                    ans = r - l

        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums, int k) {
        List<Integer>[] numPos = new ArrayList[k]; // 2*nums[i] % k 出现的所有位置
        Arrays.setAll(numPos, _ -> new ArrayList<>());
        int[] firstPos = new int[k]; // 前缀和 % k 首次出现的位置
        int[] lastPos = new int[k]; // 前缀和 % k 最后一次出现的位置
        Arrays.fill(firstPos, -1);
        firstPos[0] = lastPos[0] = 0;
        int sum = 0; // 前缀和
        int ans = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i] % k + k; // 保证 x 非负
            numPos[x * 2 % k].add(i);

            sum = (sum + x) % k;
            int r = i + 1;
            int l = firstPos[sum];
            if (l < 0) {
                firstPos[sum] = r;
            } else {
                // 不取反时的最大长度
                ans = Math.max(ans, r - l);
            }
            lastPos[sum] = r;
        }

        // 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
        for (int sl = 0; sl < k; sl++) {
            int l = firstPos[sl];
            if (l < 0) {
                continue;
            }
            for (int sr = 0; sr < k; sr++) {
                int r = lastPos[sr];
                if (r - l <= ans) { // 最优性优化：ans 无法增大
                    continue;
                }
                // 2*nums[i]%k = (sr-sl)%k
                List<Integer> pos = numPos[(sr - sl + k) % k]; // +k 保证结果非负
                int idx = Collections.binarySearch(pos, l);
                if (idx < 0) idx = ~idx; // 见 Collections.binarySearch 源码
                if (idx < pos.size() && pos.get(idx) < r) {
                    ans = r - l;
                }
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubarray(vector<int>& nums, int k) {
        vector<vector<int>> num_pos(k); // 2*nums[i] % k 出现的所有位置
        vector<int> first_pos(k, -1); // 前缀和 % k 首次出现的位置
        vector<int> last_pos(k); // 前缀和 % k 最后一次出现的位置
        first_pos[0] = last_pos[0] = 0;
        int sum = 0; // 前缀和
        int ans = 0;

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i] % k + k; // 保证 x 非负
            num_pos[x * 2 % k].push_back(i);

            sum = (sum + x) % k;
            int r = i + 1;
            int l = first_pos[sum];
            if (l < 0) {
                first_pos[sum] = r;
            } else {
                // 不取反时的最大长度
                ans = max(ans, r - l);
            }
            last_pos[sum] = r;
        }

        // 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
        for (int sl = 0; sl < k; sl++) {
            int l = first_pos[sl];
            if (l < 0) {
                continue;
            }
            for (int sr = 0; sr < k; sr++) {
                int r = last_pos[sr];
                if (r - l <= ans) { // 最优性优化：ans 无法增大
                    continue;
                }
                // 2*nums[i]%k = (sr-sl)%k
                auto& pos = num_pos[(sr - sl + k) % k]; // +k 保证结果非负
                auto it = ranges::lower_bound(pos, l);
                if (it != pos.end() && *it < r) {
                    ans = r - l;
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func longestSubarray(nums []int, k int) (ans int) {
	numPos := make([][]int, k) // 2*nums[i] % k 出现的所有位置
	firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
	for i := range firstPos {
		firstPos[i] = -1
	}
	lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
	firstPos[0] = 0
	lastPos[0] = 0
	sum := 0 // 前缀和

	for i, x := range nums {
		x = x%k + k // 保证 x 非负
		x2 := x * 2 % k
		numPos[x2] = append(numPos[x2], i)

		sum = (sum + x) % k
		r := i + 1
		l := firstPos[sum]
		if l < 0 {
			firstPos[sum] = r
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	// 枚举 s[l]%k 和 s[r]%k，判断是否存在满足要求的 i
	for sl, l := range firstPos {
		if l < 0 {
			continue
		}
		for sr, r := range lastPos {
			if r-l <= ans { // 最优性优化：ans 无法增大
				continue
			}
			// 2*nums[i]%k = (sr-sl)%k
			pos := numPos[(sr-sl+k)%k] // +k 保证结果非负
			idx := sort.SearchInts(pos, l)
			if idx < len(pos) && pos[idx] < r {
				ans = r - l
			}
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + k^2\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + k)$。

## 方法二：枚举取反元素

方法一相当于枚举 $\ell$ 和 $r$，找 $i$。

我们还可以枚举 $i$ 和 $\ell$，找 $r$。

具体地，枚举 $i=0,1,2,\ldots,n-1$，枚举 $s[\ell]\bmod k$，其中 $\ell\le i$。对于多个相同的 $s[\ell]\bmod k$，选择 $\ell$ 最小的。如果存在 $r$ 满足 $r > i$ 且 $s[r] - s[\ell] - 2\cdot \textit{nums}[i]$ 是 $k$ 的倍数，那么在多个相同的 $s[r]\bmod k$ 中取最大的 $r$，然后用 $r-\ell$ 更新答案的最大值。

如果直接枚举 $i$ 和 $s[\ell]\bmod k$，时间复杂度是 $\mathcal{O}(nk)$，太慢了。

问题在于，对于相同的 $x = (2\cdot \textit{nums}[i])\bmod k$，我们重复枚举了相同的 $s[\ell]\bmod k$。实际上，只需枚举位于上一个 $x$ 和当前 $x$ 之间的 $\ell$。这可以用**双指针**实现，用 $\textit{ptr}[x]$ 记录每个 $x$ 遍历到了哪个 $\ell$。

```py [sol-Python3]
class Solution:
    def longestSubarray(self, nums: list[int], k: int) -> int:
        firsts = [(0, 0)]
        first_pos = [-1] * k  # 前缀和 % k 首次出现的位置
        last_pos = [-1] * k  # 前缀和 % k 最后一次出现的位置
        first_pos[0] = last_pos[0] = 0
        s = 0  # 前缀和
        ans = 0

        for i, x in enumerate(nums):
            s = (s + x) % k
            r = i + 1
            l = first_pos[s]
            if l < 0:
                first_pos[s] = r
                firsts.append((r, s))
            else:
                # 不取反时的最大长度
                ans = max(ans, r - l)
            last_pos[s] = r

        # ptr[2*nums[i]%k] 表示 2*nums[i]%k 目前处理到的 firsts 的下标
        ptr = [0] * k

        # 枚举 2*nums[i]%k 和 s[l]%k，判断是否存在满足要求的 r
        for i, x in enumerate(nums):
            x2 = x * 2 % k
            p = ptr[x2]
            while p < len(firsts) and firsts[p][0] <= i:
                r = last_pos[(firsts[p][1] + x2) % k]
                if i < r:
                    ans = max(ans, r - firsts[p][0])
                # 后面遇到相同的 x2，可以从 ptr[x2] 开始继续枚举，从而避免重复枚举
                p += 1
            ptr[x2] = p

        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums, int k) {
        List<int[]> firsts = new ArrayList<>();
        firsts.add(new int[]{0, 0});
        int[] firstPos = new int[k]; // 前缀和 % k 首次出现的位置
        int[] lastPos = new int[k]; // 前缀和 % k 最后一次出现的位置
        Arrays.fill(firstPos, -1);
        firstPos[0] = lastPos[0] = 0;
        int sum = 0; // 前缀和
        int ans = 0;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i] % k + k; // 保证 x 非负
            nums[i] = x;

            sum = (sum + x) % k;
            int r = i + 1;
            int l = firstPos[sum];
            if (l < 0) {
                firstPos[sum] = r;
                firsts.add(new int[]{r, sum});
            } else {
                // 不取反时的最大长度
                ans = Math.max(ans, r - l);
            }
            lastPos[sum] = r;
        }

        // ptr[2*nums[i]%k] 表示 2*nums[i]%k 目前处理到的 firsts 的下标
        int[] ptr = new int[k];

        // 枚举 2*nums[i]%k 和 s[l]%k，判断是否存在满足要求的 r
        for (int i = 0; i < nums.length; i++) {
            int x2 = nums[i] * 2 % k;
            int p = ptr[x2];
            while (p < firsts.size() && firsts.get(p)[0] <= i) {
                int[] pair = firsts.get(p);
                int r = lastPos[(pair[1] + x2) % k];
                if (i < r) {
                    ans = Math.max(ans, r - pair[0]);
                }
                // 后面遇到相同的 y，可以从 ptr[y] 开始继续枚举，从而避免重复枚举
                p++;
            }
            ptr[x2] = p;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubarray(vector<int>& nums, int k) {
        vector<pair<int, int>> firsts = {{0, 0}};
        vector<int> first_pos(k, -1); // 前缀和 % k 首次出现的位置
        vector<int> last_pos(k); // 前缀和 % k 最后一次出现的位置
        first_pos[0] = last_pos[0] = 0;
        int sum = 0; // 前缀和
        int ans = 0;

        for (int i = 0; i < nums.size(); i++) {
            int& x = nums[i];
            x = x % k + k; // 保证 x 非负

            sum = (sum + x) % k;
            int r = i + 1;
            int l = first_pos[sum];
            if (l < 0) {
                first_pos[sum] = r;
                firsts.emplace_back(r, sum);
            } else {
                // 不取反时的最大长度
                ans = max(ans, r - l);
            }
            last_pos[sum] = r;
        }

        // ptr[2*nums[i]%k] 表示 2*nums[i]%k 目前处理到的 firsts 的下标
        vector<int> ptr(k);

        // 枚举 2*nums[i]%k 和 s[l]%k，判断是否存在满足要求的 r
        for (int i = 0; i < nums.size(); i++) {
            int x2 = nums[i] * 2 % k;
            int& p = ptr[x2];
            while (p < firsts.size() && firsts[p].first <= i) {
                int r = last_pos[(firsts[p].second + x2) % k];
                if (i < r) {
                    ans = max(ans, r - firsts[p].first);
                }
                // 后面遇到相同的 y，可以从 ptr[y] 开始继续枚举，从而避免重复枚举
                p++;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func longestSubarray(nums []int, k int) (ans int) {
	type pair struct{ l, sum int }
	firsts := []pair{{}}
	firstPos := make([]int, k) // 前缀和 % k 首次出现的位置
	for i := range firstPos {
		firstPos[i] = -1
	}
	lastPos := make([]int, k) // 前缀和 % k 最后一次出现的位置
	firstPos[0] = 0
	lastPos[0] = 0
	sum := 0 // 前缀和

	for i, x := range nums {
		x = x%k + k // 保证 x 非负
		nums[i] = x

		sum = (sum + x) % k
		r := i + 1
		l := firstPos[sum]
		if l < 0 {
			firstPos[sum] = r
			firsts = append(firsts, pair{r, sum})
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	// ptr[2*nums[i]%k] 表示 2*nums[i]%k 目前处理到的 firsts 的下标
	ptr := make([]int, k)

	// 枚举 2*nums[i]%k 和 s[l]%k，判断是否存在满足要求的 r
	for i, x := range nums {
		x2 := x * 2 % k
		p := &ptr[x2]
		for *p < len(firsts) && firsts[*p].l <= i {
			r := lastPos[(firsts[*p].sum+x2)%k]
			if i < r {
				ans = max(ans, r-firsts[*p].l)
			}
			// 后面遇到相同的 x2，可以从 ptr[x2] 开始继续枚举，从而避免重复枚举
			*p++
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + k + \min(n,k)^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。有 $\mathcal{O}(\min(n,k))$ 个不同的 $(2\cdot \textit{nums}[i])\bmod k$，每个 $(2\cdot \textit{nums}[i])\bmod k$ 我们枚举了 $\mathcal{O}(\min(n,k))$ 个不同的 $s[\ell]\bmod k$，所以二重循环的循环次数为 $\mathcal{O}(n + \min(n,k)^2)$。此外，创建大小为 $k$ 的数组需要 $\mathcal{O}(k)$ 的时间。
- 空间复杂度：$\mathcal{O}(k)$。

方法二也能通过前一题 [4063. 至多一次取反能被 K 整除的最长子数组 I](https://leetcode.cn/problems/longest-subarray-divisible-by-k-with-at-most-one-negation-i/)，那题 $n$ 小 $k$ 大。

**注**：本题是一个带位置约束的 3Sum 问题，目前不存在 $\mathcal{O}(\min(n,k)^{2-\epsilon})$ 的算法。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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
