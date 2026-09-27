枚举取反的元素是 $\textit{nums}[i]$。

随后做法类似 [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)，[我的题解](https://leetcode.cn/problems/subarray-sums-divisible-by-k/solutions/3815616/qian-zhui-he-yu-ha-xi-biao-shi-zi-bian-x-qxc5/)。改成维护前缀和模 $k$ 的首次出现的下标，从而计算最大子数组长度。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    # 做法类似 974. 和可被 K 整除的子数组
    def longestSubarrayDivByK(self, nums: list[int], k: int) -> int:
        first_pos = {0: -1}  # 前缀和 % k 首次出现的下标
        s = 0  # 前缀和
        res = 0
        for r, x in enumerate(nums):
            s = (s + x) % k
            if s in first_pos:
                res = max(res, r - first_pos[s])
            else:
                first_pos[s] = r
        return res

    def longestSubarray(self, nums: list[int], k: int) -> int:
        # 不取反
        ans = self.longestSubarrayDivByK(nums, k)

        # 枚举取反元素
        for i in range(len(nums)):
            nums[i] *= -1  # 取反
            ans = max(ans, self.longestSubarrayDivByK(nums, k))
            nums[i] *= -1  # 复原

        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarrayDivByK(int[] nums, int k) {
        Map<Integer, Integer> firstPos = new HashMap<>();
        firstPos.put(0, -1); // 前缀和 % k 首次出现的下标
        int sum = 0; // 前缀和
        int res = 0;
        for (int r = 0; r < nums.length; r++) {
            sum = (sum + nums[r] % k + k) % k; // 保证 sum 非负
            Integer l = firstPos.get(sum);
            if (l != null) {
                res = Math.max(res, r - l);
            } else {
                firstPos.put(sum, r);
            }
        }
        return res;
    }

    public int longestSubarray(int[] nums, int k) {
        // 不取反
        int ans = longestSubarrayDivByK(nums, k);

        // 枚举取反元素
        for (int i = 0; i < nums.length; i++) {
            nums[i] *= -1; // 取反
            ans = Math.max(ans, longestSubarrayDivByK(nums, k));
            nums[i] *= -1; // 复原
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> first_pos; // 哈希表超时了，改用 vector

    // 类似 974. 和可被 K 整除的子数组
    int longestSubarrayDivByK(vector<int>& nums, int k) {
        ranges::fill(first_pos, -2);
        first_pos[0] = -1;
        int s = 0; // 前缀和
        int res = 0;

        for (int r = 0; r < nums.size(); r++) {
            s = (s + nums[r] % k + k) % k; // 保证 s 非负
            int l = first_pos[s];
            if (l != -2) {
                res = max(res, r - l);
            } else {
                first_pos[s] = r;
            }
        }

        return res;
    }

public:
    int longestSubarray(vector<int>& nums, int k) {
        // 不取反
        first_pos.resize(k);
        int ans = longestSubarrayDivByK(nums, k);

        // 枚举取反元素
        for (int i = 0; i < nums.size(); i++) {
            nums[i] *= -1; // 取反
            ans = max(ans, longestSubarrayDivByK(nums, k));
            nums[i] *= -1; // 复原
        }

        return ans;
    }
};
```

```go [sol-Go]
// 做法类似 974. 和可被 K 整除的子数组
func longestSubarrayDivByK(nums []int, k int) (res int) {
	firstPos := map[int]int{0: -1} // 前缀和 % k 首次出现的下标
	sum := 0 // 前缀和
	for r, x := range nums {
		sum = (sum + x%k + k) % k // 保证 sum 非负
		l, ok := firstPos[sum]
		if ok {
			res = max(res, r-l)
		} else {
			firstPos[sum] = r
		}
	}
	return
}

func longestSubarray(nums []int, k int) int {
	// 不取反
	ans := longestSubarrayDivByK(nums, k)

	// 枚举取反元素
	for i := range nums {
		nums[i] *= -1 // 取反
		ans = max(ans, longestSubarrayDivByK(nums, k))
		nums[i] *= -1 // 复原
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 附

还可以枚举取反的元素，做到 $\mathcal{O}(n + \min(n,k)^2)$ 的时间复杂度，可以通过本题和下一题。

思路及多语言代码稍后补充。

```go [sol-Go]
func longestSubarray(nums []int, k int) (ans int) {
	// 记录前缀和 % k 最后一次出现的下标
	lastPos := make([]int, k)
	for i := range lastPos {
		lastPos[i] = -1
	}

	lastPos[0] = 0
	sum := 0
	for i, x := range nums {
		x = x%k + k
		nums[i] = x // 保证 nums[i] 非负
		sum = (sum + x) % k
		lastPos[sum] = i + 1
	}

	// 记录前缀和 % k 首次出现的下标
	firstPos := make([]int, k)
	for i := range firstPos {
		firstPos[i] = -1
	}
	type pair struct{ sum, l int }
	first := []pair{{}}

	visTime := make([]int, k)
	t := 0
	sum = 0
	for i, x := range nums {
		// 发现新的前缀和 % k
		if firstPos[sum] < 0 {
			firstPos[sum] = i
			first = append(first, pair{sum, i})
			t++
		}

		sum = (sum + x) % k
		if l := firstPos[sum]; l >= 0 {
			ans = max(ans, i+1-l) // 不取反的情况
		}

		y := x * 2 % k
		if visTime[y] == t {
			// 没有发现新的前缀和 % k，不考虑重复的 2x % k 
			continue
		}
		visTime[y] = t

		for _, p := range first {
			r := lastPos[(p.sum+y)%k]
			if r > i {
				ans = max(ans, r-p.l)
			}
		}
	}

	return
}
```

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
