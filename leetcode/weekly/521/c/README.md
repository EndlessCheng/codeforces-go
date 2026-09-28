子数组越长，元素越多，越容易满足 $\textit{nums}[i] + \textit{nums}[j] = \textit{nums}[k]$。

反之，子数组越短，元素越少，越难满足上式。

有这样的性质，可以用**滑动窗口**解决，原理请看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

元素 $x$ 滑入窗口前，先判断：

- $x$ 能否作为 $\textit{nums}[k]$？如果窗口中存在**两数之和**等于 $x$，那么必须增大左端点，缩小窗口。
- $x$ 能否作为 $\textit{nums}[i]$ 或者 $\textit{nums}[j]$？如果窗口中存在**两数之差**等于 $x$，那么必须增大左端点，缩小窗口。

为此，定义：

- $\textit{cntS}[v]$ 表示窗口中的两数之和为 $v$ 的元素对的个数。
- $\textit{cntD}[v]$ 表示窗口中的两数之差为 $v$ 的元素对的个数。由于本题没有负数，这里算差只需用大的减去小的，即**绝对差**。

在元素进入和离开窗口时，维护 $\textit{cntS}$ 和 $\textit{cntD}$。

[本题视频讲解](https://www.bilibili.com/video/BV12gah6UE9b/?t=11m35s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def maxSubarray(self, nums: list[int]) -> int:
        mx = max(nums)
        cnt_s = [0] * (mx * 2 + 1)
        cnt_d = [0] * (mx + 1)
        ans = left = 0

        # 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for i, x in enumerate(nums):
            # x 进入窗口前，先判断：
            # 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while cnt_s[x] > 0 or cnt_d[x] > 0:
                y = nums[left]
                left += 1
                for z in nums[left: i]:
                    cnt_s[y + z] -= 1
                    cnt_d[abs(y - z)] -= 1

            # x 进入窗口
            for y in nums[left: i]:
                cnt_s[x + y] += 1
                cnt_d[abs(x - y)] += 1

            # 用子数组 [left, i] 的长度更新答案的最大值
            ans = max(ans, i - left + 1)

        return ans
```

```java [sol-Java]
class Solution {
    public int maxSubarray(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cntS = new int[mx * 2 + 1];
        int[] cntD = new int[mx + 1];
        int left = 0;
        int ans = 0;

        // 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // x 进入窗口前，先判断：
            // 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while (cntS[x] > 0 || cntD[x] > 0) {
                int y = nums[left];
                left++;
                for (int j = left; j < i; j++) {
                    int z = nums[j];
                    cntS[y + z]--;
                    cntD[Math.abs(y - z)]--;
                }
            }

            // x 进入窗口
            for (int j = left; j < i; j++) {
                int y = nums[j];
                cntS[x + y]++;
                cntD[Math.abs(x - y)]++;
            }

            // 用子数组 [left, i] 的长度更新答案的最大值
            ans = Math.max(ans, i - left + 1);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubarray(vector<int>& nums) {
        int mx = ranges::max(nums);
        vector<int> cnt_s(mx * 2 + 1);
        vector<int> cnt_d(mx + 1);
        int left = 0;
        int ans = 0;

        // 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];

            // x 进入窗口前，先判断：
            // 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while (cnt_s[x] > 0 || cnt_d[x] > 0) {
                int y = nums[left];
                left++;
                for (int j = left; j < i; j++) {
                    int z = nums[j];
                    cnt_s[y + z]--;
                    cnt_d[abs(y - z)]--;
                }
            }

            // x 进入窗口
            for (int j = left; j < i; j++) {
                int y = nums[j];
                cnt_s[x + y]++;
                cnt_d[abs(x - y)]++;
            }

            // 用子数组 [left, i] 的长度更新答案的最大值
            ans = max(ans, i - left + 1);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxSubarray(nums []int) (ans int) {
	mx := slices.Max(nums)
	cntS := make([]int, mx*2+1)
	cntD := make([]int, mx+1)
	left := 0

	// 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
	for i, x := range nums {
		// x 进入窗口前，先判断：
		// 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
		for cntS[x] > 0 || cntD[x] > 0 {
			y := nums[left]
			left++
			for _, z := range nums[left:i] {
				cntS[y+z]--
				cntD[abs(y-z)]--
			}
		}

		// 元素 x 进入窗口
		for _, y := range nums[left:i] {
			cntS[x+y]++
			cntD[abs(x-y)]++
		}

		// 用子数组 [left, i] 的长度更新答案的最大值
		ans = max(ans, i-left+1)
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 写法二

定义 $\textit{cnt}[x] = \textit{cntS}[x] + \textit{cntD}[x]$，把两个数组合并。

$\textit{cntS}[x]$ 和 $\textit{cntD}[x]$ 至少一个大于 $0$，等价于 $\textit{cnt}[x] > 0$。

```py [sol-Python3]
class Solution:
    def maxSubarray(self, nums: list[int]) -> int:
        mx = max(nums)
        cnt = [0] * (mx * 2 + 1)
        ans = left = 0

        # 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for i, x in enumerate(nums):
            # x 进入窗口前，先判断：
            # 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while cnt[x] > 0:
                y = nums[left]
                left += 1
                for z in nums[left: i]:
                    cnt[y + z] -= 1
                    cnt[abs(y - z)] -= 1

            # x 进入窗口
            for y in nums[left: i]:
                cnt[x + y] += 1
                cnt[abs(x - y)] += 1

            # 用子数组 [left, i] 的长度更新答案的最大值
            ans = max(ans, i - left + 1)

        return ans
```

```java [sol-Java]
class Solution {
    public int maxSubarray(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int[] cnt = new int[mx * 2 + 1];
        int left = 0;
        int ans = 0;

        // 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // x 进入窗口前，先判断：
            // 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while (cnt[x] > 0) {
                int y = nums[left];
                left++;
                for (int j = left; j < i; j++) {
                    int z = nums[j];
                    cnt[y + z]--;
                    cnt[Math.abs(y - z)]--;
                }
            }

            // x 进入窗口
            for (int j = left; j < i; j++) {
                int y = nums[j];
                cnt[x + y]++;
                cnt[Math.abs(x - y)]++;
            }

            // 用子数组 [left, i] 的长度更新答案的最大值
            ans = Math.max(ans, i - left + 1);
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSubarray(vector<int>& nums) {
        int mx = ranges::max(nums);
        vector<int> cnt(mx * 2 + 1);
        int left = 0;
        int ans = 0;

        // 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];

            // x 进入窗口前，先判断：
            // 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
            while (cnt[x] > 0) {
                int y = nums[left];
                left++;
                for (int j = left; j < i; j++) {
                    int z = nums[j];
                    cnt[y + z]--;
                    cnt[abs(y - z)]--;
                }
            }

            // x 进入窗口
            for (int j = left; j < i; j++) {
                int y = nums[j];
                cnt[x + y]++;
                cnt[abs(x - y)]++;
            }

            // 用子数组 [left, i] 的长度更新答案的最大值
            ans = max(ans, i - left + 1);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxSubarray(nums []int) (ans int) {
	mx := slices.Max(nums)
	cnt := make([]int, mx*2+1)
	left := 0

	// 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
	for i, x := range nums {
		// x 进入窗口前，先判断：
		// 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
		for cnt[x] > 0 {
			y := nums[left]
			left++
			for _, z := range nums[left:i] {
				cnt[y+z]--
				cnt[abs(y-z)]--
			}
		}

		// 元素 x 进入窗口
		for _, y := range nums[left:i] {
			cnt[x+y]++
			cnt[abs(x-y)]++
		}

		// 用子数组 [left, i] 的长度更新答案的最大值
		ans = max(ans, i-left+1)
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 + U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。对于外面的二重循环，由于 `left++` 至多执行 $n$ 次，所以外面的二重循环的总循环次数是 $\mathcal{O}(n)$ 的。此外，创建大小为 $U$ 的数组需要 $\mathcal{O}(U)$ 的时间。
- 空间复杂度：$\mathcal{O}(U)$。

**注**：本题是一个 3Sum 问题（$x+y-z=0$），目前不存在 $\mathcal{O}(n^{2-\epsilon})$ 的算法。

## 专题训练

见下面滑动窗口题单的「**§2.1 越短越合法/求最长/最大**」。

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
