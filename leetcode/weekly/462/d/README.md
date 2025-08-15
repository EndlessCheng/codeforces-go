由于 $0$ 出现 $0$ 次，不算，所以回文数由 $1$ 到 $9$ 中的数字组成。

如果回文数包含奇数，那么奇数出现奇数次，必须在回文中心放一个奇数。如果有多个奇数，无法对称放置，所以**至多有一个奇数**。

所以回文数由 $\{2,4,6,8\}$ 加上至多一个奇数组成，有 $2^4\cdot 6 - 1 = 95$ 种非空组合。其中 $6$ 表示从 $5$ 种奇数中选一个，或者不选奇数。这就很少了，考虑**暴力枚举**。

枚举集合 $\{1,2,3,\dots,9\}$ 的至多包含一个奇数的非空子集。

对于每个子集，**枚举回文数左半边的所有排列**。比如子集 $\{2,3,4\}$ 生成的最小回文数是 $234434432$，保留左半边（奇数长度不含回文中心）的序列是 $[2,3,4,4]$。枚举这个序列的 [47. 全排列 II](https://leetcode.cn/problems/permutations-ii/)，就枚举了子集 $\{2,3,4\}$ 生成的所有回文数。

如何枚举一个集合的子集？请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

**优化**：对于任意 $n$，设其十进制长度为 $m$，如果没有符合要求的长为 $m$ 的回文串，那么长为 $m+1$ 的回文串中一定有符合要求的。这是因为 $2,4,6,8$ 可以组成 $16$ 以内的任意偶数，加上 $1$ 可以组成 $16$ 以内的任意奇数。其中 $16$ 是因为本题 $n\le 10^{15}$。

具体请看 [本题视频讲解](https://www.bilibili.com/video/BV1QNbNzxEtZ/?t=27m24s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
ODD_MASK = 0x155
D = 9

# 预处理 size 数组，详细解释见视频讲解
size = [0] * (1 << D)
for mask in range(1, 1 << D):
    t = mask & ODD_MASK
    if t & (t - 1):  # 至少有两个奇数
        continue
    for i in range(D):
        if mask >> i & 1:
            size[mask] += i + 1

class Solution:
    def specialPalindrome(self, num: int) -> int:
        target_size = len(str(num))
        ans = inf
        for mask, sz in enumerate(size):
            if sz != target_size and sz != target_size + 1:
                continue

            # 构造排列 perm
            perm = []
            odd = 0
            for x in range(1, D + 1):
                if mask >> (x - 1) & 1:
                    perm.extend([x] * (x // 2))
                    if x % 2:
                        odd = x

            # 枚举 perm 的所有排列 p，生成对应的回文数
            for p in permutations(perm):
                pal = 0
                for v in p:
                    pal = pal * 10 + v
                v = pal
                if odd:
                    pal = pal * 10 + odd
                # 反转 pal 的左半，拼在 pal 后面
                while v:
                    v, d = divmod(v, 10)
                    pal = pal * 10 + d
                if pal >= ans:  # 最优性剪枝：答案不可能变小
                    break
                if pal > num:  # 满足要求
                    ans = pal
                    break
        return ans
```

```java [sol-Java]
class Solution {
    private static final int ODD_MASK = 0x155;
    private static final int D = 9;

    private static final int[] size = new int[1 << D];

    static {
        // 预处理 size 数组，详细解释见视频讲解
        for (int mask = 1; mask < (1 << D); mask++) {
            int t = mask & ODD_MASK;
            if ((t & (t - 1)) > 0) { // 至少有两个奇数
                continue;
            }
            for (int i = 0; i < D; i++) {
                if ((mask >> i & 1) != 0) {
                    size[mask] += i + 1;
                }
            }
        }
    }

    public long specialPalindrome(long num) {
        int targetSize = String.valueOf(num).length();
        for (int mask = 1; mask < (1 << D); mask++) {
            int sz = size[mask];
            if (sz != targetSize && sz != targetSize + 1) {
                continue;
            }

            // 构造排列 perm
            int[] perm = new int[sz / 2];
            int idx = 0;
            int odd = 0;
            for (int x = 1; x <= D; x++) {
                if ((mask >> (x - 1) & 1) > 0) {
                    for (int k = 0; k < x / 2; k++) {
                        perm[idx++] = x;
                    }
                    if (x % 2 != 0) {
                        odd = x;
                    }
                }
            }

            boolean[] onPath = new boolean[perm.length];
            // 枚举 perm 的所有排列，生成对应的回文数
            dfs(0, 0, onPath, perm, odd, num);
        }
        return ans;
    }

    private long ans = Long.MAX_VALUE;

    // i 表示当前要填 perm 的第几个数，res 表示回文数的左半边
    private boolean dfs(int i, long res, boolean[] onPath, int[] perm, int odd, long num) {
        if (i == perm.length) {
            long v = res;
            if (odd > 0) {
                res = res * 10 + odd;
            }
            // 反转 x 的左半，拼在 x 后面
            while (v > 0) {
                res = res * 10 + v % 10;
                v /= 10;
            }
            if (res >= ans) { // 最优性剪枝：答案不可能变小
                return true;
            }
            if (res > num) { // 满足要求
                ans = res;
                return true;
            }
            return false;
        }

        // 见 47. 全排列 II
        for (int j = 0; j < perm.length; j++) {
            if (onPath[j] || (j > 0 && perm[j] == perm[j - 1] && !onPath[j - 1])) {
                continue;
            }
            onPath[j] = true;
            if (dfs(i + 1, res * 10 + perm[j], onPath, perm, odd, num)) {
                return true;
            }
            onPath[j] = false;
        }
        return false;
    }
}
```

```cpp [sol-C++]
const int ODD_MASK = 0x155;
const int D = 9;

int size_[1 << D];

// 预处理 size_ 数组，详细解释见视频讲解
int init = []() {
    for (int mask = 1; mask < (1 << D); mask++) {
        int t = mask & ODD_MASK;
        if (t & (t - 1)) { // 至少有两个奇数
            continue;
        }
        for (int i = 0; i < D; i++) {
            if ((mask >> i) & 1) {
                size_[mask] += i + 1;
            }
        }
    }
    return 0;
}();

class Solution {
public:
    long long specialPalindrome(long long num) {
        int target_size = to_string(num).size();
        long long ans = LLONG_MAX;
        for (int mask = 1; mask < (1 << D); mask++) {
            int sz = size_[mask];
            if (sz != target_size && sz != target_size + 1) {
                continue;
            }

            // 构造排列 perm
            vector<int> perm;
            int odd = 0;
            for (int x = 1; x <= D; x++) {
                if ((mask >> (x - 1)) & 1) {
                    perm.insert(perm.end(), x / 2, x);
                    if (x % 2) {
                        odd = x;
                    }
                }
            }

            // 枚举 perm 的所有排列，生成对应的回文数
            do {
                long long pal = 0;
                for (int v : perm) {
                    pal = pal * 10 + v;
                }
                int v = pal;
                if (odd) {
                    pal = pal * 10 + odd;
                }
                // 反转 pal 的左半，拼在 pal 后面
                while (v) {
                    pal = pal * 10 + v % 10;
                    v /= 10;
                }
                if (pal >= ans) { // 最优性剪枝：答案不可能变小
                    break;
                }
                if (pal > num) { // 满足要求
                    ans = pal;
                    break;
                }
            } while (ranges::next_permutation(perm).found);
        }
        return ans;
    }
};
```

```go [sol-Go]
var size [512]int

// 预处理 size 数组，详细解释见视频讲解
func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}
		for s := uint(mask); s > 0; s &= s - 1 {
			size[mask] += bits.TrailingZeros(s) + 1
		}
	}
}

func specialPalindrome(Num int64) int64 {
	num := int(Num)
	targetSize := len(strconv.Itoa(num))
	ans := math.MaxInt
	for mask := 1; mask < 512; mask++ {
		sz := size[mask]
		if sz != targetSize && sz != targetSize+1 {
			continue
		}

		// 构造排列 perm
		perm := make([]int, 0, sz/2)
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}

		// 枚举 perm 的所有排列，生成对应的回文数
		permutations(len(perm), len(perm), func(idx []int) (Break bool) {
			pal := 0
			for _, i := range idx {
				pal = pal*10 + perm[i]
			}
			v := pal
			if odd > 0 {
				pal = pal*10 + odd
			}
			// 反转 pal 的左半，拼在 pal 后面
			for ; v > 0; v /= 10 {
				pal = pal*10 + v%10
			}
			if pal >= ans { // 最优性剪枝：答案不可能变小
				return true
			}
			if pal > num { // 满足要求
				ans = pal
				return true
			}
			return false
		})
	}
	return int64(ans)
}

func permutations(n, r int, do func(ids []int) (Break bool)) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(K\cdot (m/2)!\cdot m)$，其中 $K=10$，$m=\log n$。至多有 $K$ 个子集，能生成长度恰好为 $m$ 的回文串。每个子集需要 $\mathcal{O}((m/2)!)$ 的时间枚举排列。对于长度等于 $m+1$ 的子集，算出最小的回文数后，就立刻结束了。最后，对于每个回文数，需要 $\mathcal{O}(m)$ 的时间计算（反转左半）。注：如果在回溯过程中同时计算反转后的结果，可以把 $\mathcal{O}(m)$ 优化为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(m)$。

## 写法二：预处理

预处理所有长度 $\le 16$ 的特殊数列表，这只有 $2296$ 个。

把列表排序后，在列表中**二分查找**大于 $n$ 的最小特殊数，原理见 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
ODD_MASK = 0x155
D = 9

special_numbers = []
for mask in range(1, 1 << D):
    t = mask & ODD_MASK
    if t & (t - 1):  # 至少有两个奇数
        continue

    # 构造排列 perm
    perm = []
    size = odd = 0
    for x in range(1, D + 1):
        if mask >> (x - 1) & 1:
            size += x
            perm.extend([x] * (x // 2))
            if x % 2:
                odd = x
    if size > 16:  # 回文串太长了
        continue

    # 枚举 perm 的所有排列 p，生成对应的回文数
    for p in permutations(perm):
        pal = 0
        for v in p:
            pal = pal * 10 + v
        v = pal
        if odd:
            pal = pal * 10 + odd
        # 反转 pal 的左半，拼在 pal 后面
        while v:
            v, d = divmod(v, 10)
            pal = pal * 10 + d
        special_numbers.append(pal)
special_numbers = sorted(set(special_numbers))

class Solution:
    def specialPalindrome(self, n: int) -> int:
        i = bisect_right(special_numbers, n)
        return special_numbers[i]
```

```java [sol-Java]
class Solution {
    private static final int ODD_MASK = 0x155;
    private static final int D = 9;

    private static final List<Long> specialNumbers = new ArrayList<>();

    static {
        for (int mask = 1; mask < (1 << D); mask++) {
            int t = mask & ODD_MASK;
            if ((t & (t - 1)) > 0) { // 至少有两个奇数
                continue;
            }

            int size = 0;
            for (int i = 0; i < D; i++) {
                if ((mask >> i & 1) != 0) {
                    size += i + 1;
                }
            }
            if (size > 16) { // 回文串太长了
                continue;
            }

            // 构造排列 perm
            int[] perm = new int[size / 2];
            int idx = 0;
            int odd = 0;
            for (int x = 1; x <= D; x++) {
                if ((mask >> (x - 1) & 1) > 0) {
                    for (int k = 0; k < x / 2; k++) {
                        perm[idx++] = x;
                    }
                    if (x % 2 != 0) {
                        odd = x;
                    }
                }
            }

            boolean[] onPath = new boolean[perm.length];
            // 枚举 perm 的所有排列，生成对应的回文数
            dfs(0, 0, onPath, perm, odd);
        }

        Collections.sort(specialNumbers);
    }

    // i 表示当前要填 perm 的第几个数，res 表示回文数的左半边
    private static void dfs(int i, long res, boolean[] onPath, int[] perm, int odd) {
        if (i == perm.length) {
            long v = res;
            if (odd > 0) {
                res = res * 10 + odd;
            }
            // 反转 x 的左半，拼在 x 后面
            while (v > 0) {
                res = res * 10 + v % 10;
                v /= 10;
            }
            specialNumbers.add(res);
            return;
        }

        // 见 47. 全排列 II
        for (int j = 0; j < perm.length; j++) {
            if (onPath[j] || (j > 0 && perm[j] == perm[j - 1] && !onPath[j - 1])) {
                continue;
            }
            onPath[j] = true;
            dfs(i + 1, res * 10 + perm[j], onPath, perm, odd);
            onPath[j] = false;
        }
    }

    public long specialPalindrome(long n) {
        int i = upperBound(specialNumbers, n);
        return specialNumbers.get(i);
    }

    // https://www.bilibili.com/video/BV1AP41137w7/
    // 返回 nums 中第一个大于 target 的数的下标（注意是大于，不是大于等于）
    // 如果这样的数不存在，则返回 nums.length
    // 时间复杂度 O(log nums.size())
    // 采用开区间写法实现
    private int upperBound(List<Long> nums, long target) {
        int left = -1, right = nums.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = left + (right - left) / 2;
            if (nums.get(mid) > target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
const int ODD_MASK = 0x155;
const int D = 9;

vector<long long> special_numbers;

int init = []() {
    for (int mask = 1; mask < (1 << D); mask++) {
        int t = mask & ODD_MASK;
        if (t & (t - 1)) { // 至少有两个奇数
            continue;
        }

        // 构造排列 perm
        vector<int> perm;
        int size = 0, odd = 0;
        for (int x = 1; x <= D; x++) {
            if ((mask >> (x - 1)) & 1) {
                size += x;
                perm.insert(perm.end(), x / 2, x);
                if (x % 2) {
                    odd = x;
                }
            }
        }
        if (size > 16) { // 回文串太长了
            continue;
        }

        // 枚举 perm 的所有排列，生成对应的回文数
        do {
            long long pal = 0;
            for (int v : perm) {
                pal = pal * 10 + v;
            }
            int v = pal;
            if (odd) {
                pal = pal * 10 + odd;
            }
            // 反转 pal 的左半，拼在 pal 后面
            while (v) {
                pal = pal * 10 + v % 10;
                v /= 10;
            }
            special_numbers.push_back(pal);
        } while (ranges::next_permutation(perm).found);
    }
    ranges::sort(special_numbers);
    return 0;
}();

class Solution {
public:
    long long specialPalindrome(long long n) {
        return *ranges::upper_bound(special_numbers, n);
    }
};
```

```go [sol-Go]
var specialNumbers []int

func init() {
	const oddMask = 0x155
	for mask := 1; mask < 512; mask++ {
		t := mask & oddMask
		if t&(t-1) > 0 { // 至少有两个奇数
			continue
		}

		// 构造排列 perm
		perm := []int{}
		size := 0
		odd := 0
		for s := uint(mask); s > 0; s &= s - 1 {
			x := bits.TrailingZeros(s) + 1
			size += x
			for range x / 2 {
				perm = append(perm, x)
			}
			if x%2 > 0 {
				odd = x
			}
		}
		if size > 16 { // 回文串太长了
			continue
		}

		permutations(len(perm), len(perm), func(idx []int) bool {
			pal := 0
			for _, i := range idx {
				pal = pal*10 + perm[i]
			}
			v := pal
			if odd > 0 {
				pal = pal*10 + odd
			}
			// 反转 pal 的左半，拼在 pal 后面
			for ; v > 0; v /= 10 {
				pal = pal*10 + v%10
			}
			specialNumbers = append(specialNumbers, pal)
			return false
		})
	}
	slices.Sort(specialNumbers)
	specialNumbers = slices.Compact(specialNumbers)
}

func specialPalindrome(n int64) int64 {
	i := sort.SearchInts(specialNumbers, int(n+1))
	return int64(specialNumbers[i])
}

func permutations(n, r int, do func(ids []int) (Break bool)) {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	if do(ids[:r]) {
		return
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}
	for {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				tmp := ids[i]
				copy(ids[i:], ids[i+1:])
				ids[n-1] = tmp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				ids[i], ids[n-j] = ids[n-j], ids[i]
				if do(ids[:r]) {
					return
				}
				break
			}
		}
		if i == -1 {
			return
		}
	}
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(\log N)$，其中 $N=2296$。
- 空间复杂度：$\mathcal{O}(1)$。

## 如果数据范围更大呢？

有类似数位 DP 的搜索做法，见题解下面我的评论。

## 相似题目

[564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)

## 专题训练

见下面回溯题单的「**§4.6 有重复元素的回溯**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
