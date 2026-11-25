## 方法一：二分答案 + 多项式哈希 

如果子数组 $A$ 是唯一的，那么更长的包含 $A$ 的子数组也是唯一的。

如果子数组 $A$ 不是唯一的，那么 $A$ 的子数组也不是唯一的。

据此，可以**二分猜答案**。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题转化成一个判定性问题：

- 给定子数组长度 $\textit{length}$，判断 $\textit{nums}$ 是否存在长为 $\textit{length}$ 的唯一子数组。

**核心思路**：计算每个长为 $\textit{length}$ 的子数组的哈希值，用一个哈希表统计出现次数，判断是否存在出现次数为 $1$ 的哈希值。

推荐把 [2156. 查找给定哈希值的子串](https://leetcode.cn/problems/find-substring-with-given-hash-value/) 和 [3756. 连接非零数字并乘以其数字和 II](https://leetcode.cn/problems/concatenate-non-zero-digits-and-multiply-by-sum-ii/) 做了，对理解多项式哈希的计算方法有帮助。

### 细节

下面代码采用开区间二分。使用闭区间或者半闭半开区间也是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$0$。题目要求子数组非空，所以子数组长度不可能是 $0$。
- 开区间右端点初始值：$n$。长为 $n$ 的子数组只有一个，即 $\textit{nums}$，一定满足题目要求。

> 注 1：对于开区间写法，简单来说 `check(mid) == true` 成立时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。
> 
> 注 2：本题子数组长度固定，也可以用定长滑动窗口计算，见 2156 题。

```py [sol-Python3]
class Solution:
    def smallestUniqueSubarray(self, nums: list[int]) -> int:
        n = len(nums)
        # 多项式哈希（方便计算子数组哈希值）
        # 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
        MOD = 1_070_777_777
        BASE = randint(8 * 10 ** 8, 9 * 10 ** 8)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(nums[:i])
        for i, x in enumerate(nums):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + x) % MOD  # 秦九韶算法计算多项式哈希

        # 计算子数组 nums[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
        # 计算方法类似前缀和
        def sub_hash(l: int, r: int) -> int:
            return (pre_hash[r] - pre_hash[l] * pow_base[r - l]) % MOD

        def check(length: int) -> bool:
            cnt = defaultdict(int)
            for i in range(length, n + 1):
                h = sub_hash(i - length, i)
                cnt[h] += 1
            return 1 in cnt.values()

        # 开区间二分
        left, right = 0, n
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                right = mid
            else:
                left = mid
        return right
```

```py [sol-Python3 库函数二分]
class Solution:
    def smallestUniqueSubarray(self, nums: list[int]) -> int:
        n = len(nums)
        # 多项式哈希（方便计算子数组哈希值）
        # 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
        MOD = 1_070_777_777
        BASE = randint(8 * 10 ** 8, 9 * 10 ** 8)  # 随机 BASE，防止 hack
        pow_base = [1] + [0] * n  # pow_base[i] = BASE^i
        pre_hash = [0] * (n + 1)  # 前缀哈希值 pre_hash[i] = hash(nums[:i])
        for i, x in enumerate(nums):
            pow_base[i + 1] = pow_base[i] * BASE % MOD
            pre_hash[i + 1] = (pre_hash[i] * BASE + x) % MOD  # 秦九韶算法计算多项式哈希

        # 计算子数组 nums[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
        # 计算方法类似前缀和
        def sub_hash(l: int, r: int) -> int:
            return (pre_hash[r] - pre_hash[l] * pow_base[r - l]) % MOD

        def check(length: int) -> bool:
            cnt = defaultdict(int)
            for i in range(length, n + 1):
                h = sub_hash(i - length, i)
                cnt[h] += 1
            return 1 in cnt.values()

        return bisect_left(range(n), True, lo=1, key=check)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_070_777_777;
    private static final int BASE = (int) 8e8 + new Random().nextInt((int) 1e8); // 随机 base，防止 hack

    private int[] powBase;
    private int[] preHash;

    // 计算子数组 nums[l,r-1] 的哈希值
    // 计算方法类似前缀和
    private int subHash(int l, int r) {
        return (int) ((preHash[r] - (long) preHash[l] * powBase[r - l] % MOD + MOD) % MOD);
    }

    public int smallestUniqueSubarray(int[] nums) {
        int n = nums.length;
        // 多项式哈希（方便计算子数组哈希值）
        // 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
        powBase = new int[n + 1]; // powBase[i] = base^i
        preHash = new int[n + 1]; // 前缀哈希值 preHash[i] = hash(nums[0,i-1])
        powBase[0] = 1;
        for (int i = 0; i < n; i++) {
            powBase[i + 1] = (int) ((long) powBase[i] * BASE % MOD);
            preHash[i + 1] = (int) (((long) preHash[i] * BASE + nums[i]) % MOD); // 秦九韶算法计算多项式哈希
        }

        // 开区间二分
        int left = 0;
        int right = n;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(mid, n)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int length, int n) {
        HashMap<Integer, Integer> cnt = new HashMap<>();
        for (int i = length; i <= n; i++) {
            int h = subHash(i - length, i);
            cnt.merge(h, 1, Integer::sum); // cnt[h]++
        }
        return cnt.containsValue(1);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestUniqueSubarray(vector<int>& nums) {
        int n = nums.size();
        // 多项式哈希（方便计算子数组哈希值）
        // 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
        const int MOD = 1'070'777'777;
        mt19937 rng(chrono::steady_clock::now().time_since_epoch().count());
        const int BASE = uniform_int_distribution<>(8e8, 9e8)(rng); // 随机 base，防止 hack
        vector<int> pow_base(n + 1); // pow_base[i] = base^i
        vector<int> pre_hash(n + 1); // 前缀哈希值 pre_hash[i] = hash(nums[0,i-1])
        pow_base[0] = 1;
        for (int i = 0; i < n; i++) {
            pow_base[i + 1] = 1LL * pow_base[i] * BASE % MOD;
            pre_hash[i + 1] = (1LL * pre_hash[i] * BASE + nums[i]) % MOD; // 秦九韶算法计算多项式哈希
        }

        // 计算子数组 nums[l,r-1] 的哈希值
        // 计算方法类似前缀和
        auto sub_hash = [&](int l, int r) {
            return ((pre_hash[r] - 1LL * pre_hash[l] * pow_base[r - l]) % MOD + MOD) % MOD;
        };

        auto check = [&](int length) -> bool {
            unordered_map<int, int> cnt;
            for (int i = length; i <= n; i++) {
                int h = sub_hash(i - length, i);
                cnt[h]++;
            }
            for (auto& [_, c] : cnt) {
                if (c == 1) {
                    return true;
                }
            }
            return false;
        };

        // 开区间二分
        int left = 0, right = n;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 多项式哈希（方便计算子数组哈希值）
	// 哈希函数 hash(s) = nums[0] * base^(n-1) + nums[1] * base^(n-2) + ... + nums[n-2] * base + nums[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8)
	powBase := make([]int, n+1) // powBase[i] = base^i
	preHash := make([]int, n+1) // 前缀哈希值 preHash[i] = hash(nums[:i])
	powBase[0] = 1
	for i, x := range nums {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + x) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子数组 nums[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	ans := sort.Search(n, func(length int) bool {
		cnt := map[int]int{}
		for i := length; i <= n; i++ {
			h := subHash(i-length, i)
			cnt[h]++
		}
		for _, c := range cnt {
			if c == 1 {
				return true
			}
		}
		return false
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：后缀数组

计算 $\textit{nums}$ 的**后缀数组** $\textit{sa}$ 和**高度数组** $\textit{height}$。

其中 $\textit{height}[i]$ 定义为后缀 $\textit{sa}[i]$ 和后缀 $\textit{sa}[i-1]$ 的 LCP。

对于后缀 $\textit{sa}[i]$，它与其他后缀的最长公共前缀，来自后缀 $\textit{sa}[i]$ 和 $\textit{sa}[i-1]$ 的 LCP，以及后缀 $\textit{sa}[i]$ 和 $\textit{sa}[i+1]$ 的 LCP，这二者的最大值，即

$$
\max(\textit{height}[i],\textit{height}[i+1])
$$

这个值再加一，即为后缀 $\textit{sa}[i]$ 中的最短**唯一**前缀（可以用反证法证明，如果不唯一，那么 $\textit{height}[i]$ 或 $\textit{height}[i+1]$ 还能更大，矛盾）。

注意 $\max(\textit{height}[i],\textit{height}[i+1])+1$ 不能超过后缀 $\textit{sa}[i]$ 的长度，即

$$
\max(\textit{height}[i],\textit{height}[i+1])+1 \le n - \textit{sa}[i]
$$

如果上式成立，用 $\max(\textit{height}[i],\textit{height}[i+1])+1$ 更新答案的最小值。

```go
func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
	// 把 1 个 int 拆成 3 个 byte（题目保证 nums[i] <= 1e5），从而可以调用库函数计算后缀数组
	tmp := make([]byte, 0, n*3)
	for _, x := range nums {
		tmp = append(tmp, byte(x>>16), byte(x>>8), byte(x))
	}

	type _tp struct {
		_  []byte
		sa []int32
	}
	_sa := (*_tp)(unsafe.Pointer(suffixarray.New(tmp))).sa

	sa := make([]int32, 0, n)
	for _, p := range _sa {
		if p%3 == 0 { // 是 3 的倍数的 _sa[i] 就对应着 nums 的 sa[i]
			sa = append(sa, p/3)
		}
	}

	// 计算后缀名次数组
	// 后缀 nums[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 nums 在后缀字典序中的排名，rank[n-1] 即 nums[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, n)
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(nums[sa[i]:], nums[sa[i-1]:])  (i > 0)
	// 获取 nums[i] 所在位置的高度：height[rank[i]]
	height := make([]int, n)
	h := 0
	// 计算 nums 与 nums[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 nums[1:] 与 nums[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 nums[2:] 与 nums[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 nums[n-1:] 与 nums[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 nums[0] 和 nums[sa[rank[0]-1]] 这两个数
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && nums[i+h] == nums[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	ans := n
	for i, h := range height {
		// 对于后缀 nums[sa[i]:]，其长为 uniqueLength 的前缀是唯一的
		uniqueLength := h + 1
		if i < n-1 {
			uniqueLength = max(h, height[i+1]) + 1
		}
		// 注意 uniqueLength 不能超过后缀 nums[sa[i]:] 的长度
		if uniqueLength <= n-int(sa[i]) {
			ans = min(ans, uniqueLength)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。注：库函数用的是 SA-IS 算法。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面字符串题单的「**四、字符串哈希**」和「**八、后缀数组/后缀自动机**」。

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
