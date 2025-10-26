## 前置题目

请先完成如下两题：

1. [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)，[我的题解](https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/)。
2. [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)，[我的题解](https://leetcode.cn/problems/subarray-sums-divisible-by-k/solutions/3815616/qian-zhui-he-yu-ha-xi-biao-shi-zi-bian-x-qxc5/)。

## 思路

如果本题不要求子数组互不相同，那么就是 974 题。

什么情况下，会有相同的子数组？

如果子数组包含多种不同元素，比如 $[2,2,2,3],[2,2,3,3],[2,3,3,3]$ 等。由于相同的子数组，长度一定相同，可以视作一个 [定长滑窗](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)。在子数组递增且包含多种不同元素的情况下，**进入窗口的元素一定比离开窗口的元素大**，所以这些子数组一定互不相同。

所以**只有当子数组只包含一种元素时，才会出现相同的子数组**。

我们可以在 974 题的做法上修改：对于连续相同元素段，要保证哈希表暂时不包含这一段对应的前缀和，**等我们遍历完这一段，再把对应的前缀和加到哈希表中**。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

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
        Map<Integer, Integer> cnt = new HashMap<>();
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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
