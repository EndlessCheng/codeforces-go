## 前置题目

请先完成如下两题：

1. [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)，[我的题解](https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/)。
2. [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)，[我的题解](https://leetcode.cn/problems/subarray-sums-divisible-by-k/solutions/3815616/qian-zhui-he-yu-ha-xi-biao-shi-zi-bian-x-qxc5/)。

## 分析

如果本题不要求子数组互不相同，那么就是 974 题。

什么情况下，会有相同的子数组？

如果子数组包含多种不同元素，比如 $[2,2,2,3],[2,2,3,3],[2,3,3,3]$ 等。由于相同的子数组，长度一定相同，可以视作一个 [定长滑窗](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)。在子数组递增且包含多种不同元素的情况下，**进入窗口的元素一定比离开窗口的元素大**，所以这些子数组一定互不相同。

所以**只有当子数组只包含一种元素时，才会出现相同的子数组**。

## 方法一：减去重复的子数组个数

首先，用 974 题的方法，计算所有元素和能被 $k$ 整除的子数组个数。

这里面有重复的，减掉就是答案。

比如某一段 $a = [2,2,2,2,2]$，找能被 $k=4$ 整除的子数组。

子数组 $[2,2]$ 的和是 $k$ 的倍数，由于 $a$ 中所有元素都相同，此时我们要算的是：

- 计算 $a$ 中长为 $2$ 的子数组的个数。

子数组右端点可以从 $1$ 到 $4$，所以有 $4$ 个 $[2,2]$，其中有 $3$ 个是重复的，从答案中减掉。

子数组 $[2,2,2,2]$ 的和是 $k$ 的倍数，由于 $a$ 中所有元素都相同，此时我们要算的是：

- 计算 $a$ 中长为 $4$ 的子数组的个数。

子数组右端点可以从 $3$ 到 $4$，所以有 $2$ 个 $[2,2,2,2]$，其中有 $1$ 个是重复的，从答案中减掉。

```py [sol-Python3]
class Solution:
    # 974. 和可被 K 整除的子数组
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        cnt = defaultdict(int)
        ans = s = 0
        for x in nums:
            cnt[s] += 1
            s = (s + x) % k
            ans += cnt[s]
        return ans

    def numGoodSubarrays(self, nums: List[int], k: int) -> int:
        ans = self.subarraysDivByK(nums, k)
        start = 0
        for i, x in enumerate(nums):
            if i < len(nums) - 1 and x == nums[i + 1]:
                continue
            # 遍历到了连续相同元素段的末尾
            size = i - start + 1  # 这一段的长度
            for sz in range(1, size + 1):
                if x * sz % k == 0:
                    # 长为 sz 的子数组元素和能被 k 整除
                    # 一共有 size-sz+1 个长为 sz 的子数组，其中有 size-sz 个重复的
                    ans -= size - sz
            start = i + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long numGoodSubarrays(int[] nums, int k) {
        long ans = subarraysDivByK(nums, k);
        int n = nums.length;
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            // 遍历到了连续相同元素段的末尾
            int size = i - start + 1; // 这一段的长度
            for (int sz = 1; sz <= size; sz++) {
                if ((long) x * sz % k == 0) {
                    // 长为 sz 的子数组元素和能被 k 整除
                    // 一共有 size-sz+1 个长为 sz 的子数组，其中有 size-sz 个重复的
                    ans -= size - sz;
                }
            }
            start = i + 1;
        }
        return ans;
    }

    // 974. 和可被 K 整除的子数组
    // 由于本题 nums 没有负数，无需调整
    private long subarraysDivByK(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>(nums.length, 1); // 预分配空间
        int s = 0;
        long ans = 0;
        for (int x : nums) {
            cnt.merge(s, 1, Integer::sum); // cnt[s]++
            s = (s + x) % k;
            ans += cnt.getOrDefault(s, 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 974. 和可被 K 整除的子数组
    // 由于本题 nums 没有负数，无需调整
    long long subarraysDivByK(vector<int>& nums, int k) {
        unordered_map<int, int> cnt;
        int s = 0;
        long long ans = 0;
        for (int x : nums) {
            cnt[s]++;
            s = (s + x) % k;
            ans += cnt[s];
        }
        return ans;
    }
    
public:
    long long numGoodSubarrays(vector<int>& nums, int k) {
        long long ans = subarraysDivByK(nums, k);
        int n = nums.size();
        int start = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            // 遍历到了连续相同元素段的末尾
            int size = i - start + 1; // 这一段的长度
            for (int sz = 1; sz <= size; sz++) {
                if (1LL * x * sz % k == 0) {
                    // 长为 sz 的子数组元素和能被 k 整除
                    // 一共有 size-sz+1 个长为 sz 的子数组，其中有 size-sz 个重复的
                    ans -= size - sz;
                }
            }
            start = i + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
// 974. 和可被 K 整除的子数组
// 由于本题 nums 没有负数，无需调整
func subarraysDivByK(nums []int, k int) (ans int64) {
	cnt := make(map[int]int, len(nums)) // 预分配空间
	s := 0
	for _, x := range nums {
		cnt[s]++
		s = (s + x) % k
		ans += int64(cnt[s])
	}
	return
}

func numGoodSubarrays(nums []int, k int) int64 {
	ans := subarraysDivByK(nums, k)
	start := 0
	for i, x := range nums {
		if i < len(nums)-1 && x == nums[i+1] {
			continue
		}
		// 遍历到了连续相同元素段的末尾
		size := i - start + 1 // 这一段的长度
		for sz := 1; sz <= size; sz++ {
			if x*sz%k == 0 {
				// 长为 sz 的子数组元素和能被 k 整除
				// 一共有 size-sz+1 个长为 sz 的子数组，其中有 size-sz 个重复的
				ans -= int64(size - sz)
			}
		}
		start = i + 1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个二重循环，但内层循环的总循环次数不超过 $n$，所以总的循环次数不超过 $2n$。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：避免重复统计

直接在 974 题的做法上修改。

例如在 $\textit{nums}=[1,2,2,3,3,3]$ 中找 $[3,3]$，只要限定子数组的左端点在第一个 $3$ 或者更靠左的位置，就不会错误地统计两个 $[3,3]$ 子数组了。

换句话说，对于连续相同元素段，我们要保证哈希表暂时不包含这一段对应的前缀和，**等遍历完这一段，再把对应的前缀和加到哈希表中**。

例如 $\textit{nums}=[1,2,2,3,3,3]$，在遍历 $3$ 的过程中，哈希表只保存 $[1,2,2]$ 中的前缀和，此时在哈希表中查询，就相当于限定子数组的左端点在第一个 $3$ 或者更靠左的位置。这样就能保证统计的子数组无重复。

[本题视频讲解](https://www.bilibili.com/video/BV1eqxNzXE8v/?t=29m05s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def numGoodSubarrays(self, nums: List[int], k: int) -> int:
        cnt = defaultdict(int)
        cnt[0] = 1  # 见 560 题
        pre_sum = 0  # 前缀和
        last_start = 0  # 上一个连续相同段的起始下标
        ans = 0

        for i, x in enumerate(nums):
            if i and x != nums[i - 1]:
                # 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
                v = nums[i - 1]
                s = pre_sum
                for _ in range(i - last_start):
                    cnt[s % k] += 1
                    s -= v
                last_start = i

            pre_sum += x
            ans += cnt[pre_sum % k]

        return ans
```

```java [sol-Java]
class Solution {
    public long numGoodSubarrays(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>(nums.length, 1);
        cnt.put(0, 1); // 见 560 题
        long sum = 0; // 前缀和
        int lastStart = 0; // 上一个连续相同段的起始下标
        long ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (i > 0 && x != nums[i - 1]) {
                // 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
                long s = sum;
                for (int t = i - lastStart; t > 0; t--) {
                    cnt.merge((int) (s % k), 1, Integer::sum); // cnt[s % k]++
                    s -= nums[i - 1];
                }
                lastStart = i;
            }
            sum += x;
            ans += cnt.getOrDefault((int) (sum % k), 0);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numGoodSubarrays(vector<int>& nums, int k) {
        unordered_map<int, int> cnt = {{0, 1}}; // 为什么加个 0？见 560 题
        long long sum = 0; // 前缀和
        int last_start = 0; // 上一个连续相同段的起始下标
        long long ans = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (i > 0 && x != nums[i - 1]) {
                // 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
                long long s = sum;
                for (int t = i - last_start; t > 0; t--) {
                    cnt[s % k]++;
                    s -= nums[i - 1];
                }
                last_start = i;
            }
            sum += x;
            ans += cnt[sum % k];
        }
        return ans;
    }
};
```

```go [sol-Go]
func numGoodSubarrays(nums []int, k int) (ans int64) {
	cnt := map[int]int{0: 1} // 为什么加个 0？见 560 题
	sum := 0 // 前缀和
	lastStart := 0 // 上一个连续相同段的起始下标
	for i, x := range nums {
		if i > 0 && x != nums[i-1] {
			// 上一个连续相同段结束，可以把上一段对应的前缀和添加到 cnt
			s := sum
			for range i - lastStart {
				cnt[s%k]++
				s -= nums[i-1]
			}
			lastStart = i
		}
		sum += x
		ans += int64(cnt[sum%k])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个二重循环，但内层循环的总循环次数不超过 $n$，所以总的循环次数不超过 $2n$。
- 空间复杂度：$\mathcal{O}(n)$。

## 进阶问题

如果没有「非降序」要怎么做？

阅读：[本质不同子串](https://oi-wiki.org/string/sa/#%E4%B8%8D%E5%90%8C%E5%AD%90%E4%B8%B2%E7%9A%84%E6%95%B0%E7%9B%AE)。

## 专题训练

1. 下面数据结构题单的「**§1.2 前缀和与哈希表**」。
2. 下面双指针题单的「**六、分组循环**」。

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
