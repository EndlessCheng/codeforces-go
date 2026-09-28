枚举取反的元素是 $\textit{nums}[i]$。

随后做法类似 [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)，[我的题解](https://leetcode.cn/problems/subarray-sums-divisible-by-k/solutions/3815616/qian-zhui-he-yu-ha-xi-biao-shi-zi-bian-x-qxc5/)。改成维护前缀和模 $k$ 的首次出现的下标，从而计算最大子数组长度。

[本题视频讲解](https://www.bilibili.com/video/BV1v9a86hEkp/?t=6m34s)，欢迎点赞关注~

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

    // 做法类似 974. 和可被 K 整除的子数组
    private int longestSubarrayDivByK(int[] nums, int k) {
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
}
```

```cpp [sol-C++]
class Solution {
    vector<int> first_pos; // 哈希表超时了，改用 vector

    // 做法类似 974. 和可被 K 整除的子数组
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
- 空间复杂度：$\mathcal{O}(\min(n,k))$。

## 更快的做法

原理见下一题 [我的题解](https://leetcode.cn/problems/longest-subarray-divisible-by-k-with-at-most-one-negation-ii/solution/mei-ju-qian-zhui-he-pythonjavacgo-by-end-nlum/) 中的「优化：无需二分」。虽然复杂度没有差别，但实际运行时间快很多。

这里只把代码搬过来。

```py [sol-Python3]
class Solution:
    def longestSubarray(self, nums: list[int], k: int) -> int:
        first_sum = [0]
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
                first_sum.append(s)
            else:
                # 不取反时的最大长度
                ans = max(ans, r - l)
            last_pos[s] = r

        last_x2 = [-1] * k
        sr = 0

        # 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
        for i, x in enumerate(nums):
            last_x2[x * 2 % k] = i
            sr = (sr + x) % k
            r = i + 1
            if last_pos[sr] != r:  # 只考虑 s[r]%k 最后一次出现的位置
                continue
            for sl in first_sum:
                l = first_pos[sl]
                if r - l <= ans:  # 最优性优化：ans 无法增大
                    break
                j = last_x2[(sr - sl) % k]
                if j >= l:
                    ans = r - l

        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums, int k) {
        List<Integer> firstSum = new ArrayList<>(); // 改成数组更快，见【Java 写法二】
        firstSum.add(0);
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
                firstSum.add(sum);
            } else {
                // 不取反时的最大长度
                ans = Math.max(ans, r - l);
            }
            lastPos[sum] = r;
        }

        int[] lastX2 = new int[k];
        Arrays.fill(lastX2, -1);
        int sr = 0;

        // 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            lastX2[x * 2 % k] = i;
            sr = (sr + x) % k;
            int r = i + 1;
            if (lastPos[sr] != r) { // 只考虑 s[r]%k 最后一次出现的位置
                continue;
            }
            for (int sl : firstSum) {
                int l = firstPos[sl];
                if (r - l <= ans) { // 最优性优化：ans 无法增大
                    break;
                }
                int j = lastX2[(sr - sl + k) % k]; // +k 保证结果非负
                if (j >= l) {
                    ans = r - l;
                }
            }
        }

        return ans;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public int longestSubarray(int[] nums, int k) {
        int[] firstSum = new int[k];
        int firstSumSize = 1;
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
                firstSum[firstSumSize++] = sum;
            } else {
                // 不取反时的最大长度
                ans = Math.max(ans, r - l);
            }
            lastPos[sum] = r;
        }

        int[] lastX2 = new int[k];
        Arrays.fill(lastX2, -1);
        int sr = 0;

        // 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            lastX2[x * 2 % k] = i;
            sr = (sr + x) % k;
            int r = i + 1;
            if (lastPos[sr] != r) { // 只考虑 s[r]%k 最后一次出现的位置
                continue;
            }
            for (int idx = 0; idx < firstSumSize; idx++) {
                int sl = firstSum[idx];
                int l = firstPos[sl];
                if (r - l <= ans) { // 最优性优化：ans 无法增大
                    break;
                }
                int j = lastX2[(sr - sl + k) % k]; // +k 保证结果非负
                if (j >= l) {
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
        vector<int> first_sum = {0};
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
                first_sum.push_back(sum);
            } else {
                // 不取反时的最大长度
                ans = max(ans, r - l);
            }
            last_pos[sum] = r;
        }

        vector<int> last_x2(k, -1);
        int sr = 0;

        // 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            last_x2[x * 2 % k] = i;
            sr = (sr + x) % k;
            int r = i + 1;
            if (last_pos[sr] != r) { // 只考虑 s[r]%k 最后一次出现的位置
                continue;
            }
            for (int sl : first_sum) {
                int l = first_pos[sl];
                if (r - l <= ans) { // 最优性优化：ans 无法增大
                    break;
                }
                int j = last_x2[(sr - sl + k) % k]; // +k 保证结果非负
                if (j >= l) {
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
	firstSum := []int{0}
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
			firstSum = append(firstSum, sum)
		} else {
			// 不取反时的最大长度
			ans = max(ans, r-l)
		}
		lastPos[sum] = r
	}

	lastX2 := make([]int, k)
	for i := range lastX2 {
		lastX2[i] = -1
	}
	sr := 0

	// 枚举 s[r]%k 和 s[l]%k，判断是否存在满足要求的 i
	for i, x := range nums {
		lastX2[x*2%k] = i
		sr = (sr + x) % k
		r := i + 1
		if lastPos[sr] != r { // 只考虑 s[r]%k 最后一次出现的位置
			continue
		}
		for _, sl := range firstSum {
			l := firstPos[sl]
			if r-l <= ans { // 最优性优化：ans 无法增大
				break
			}
			j := lastX2[(sr-sl+k)%k] // +k 保证结果非负
			if j >= l {
				ans = r - l
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
