## 方法一：前后缀分解

推荐先完成本题的简单版本：[3738. 替换至多一个元素后最长非递减子数组](https://leetcode.cn/problems/longest-non-decreasing-subarray-after-replacing-at-most-one-element/)。

为方便描述，下文把 $\textit{nums}$ 简称为 $a$。

假设修改的是 $a[i]$，我们需要知道什么？

- 以 $i-1$ 结尾的最长等差子数组的长度，记作 $\textit{pre}[i-1]$。
- 以 $i+1$ 开头的最长等差子数组的长度，记作 $\textit{suf}[i+1]$。

这可以递推算出。以 $\textit{pre}$ 为例：

- $\textit{pre}[0] = 1$。
- 如果 $i = 1$ 或者 $a[i] - a[i-1] \ne a[i-1] - a[i-2]$，那么 $\textit{pre}[i]$ 只能是 $2$，也就是把 $a[i-1]$ 和 $a[i]$ 组成长为 $2$ 等差子数组。
- 否则可以把 $a[i]$ 拼接在以 $i-1$ 结尾的最长等差子数组的后面，即 $\textit{pre}[i] = \textit{pre}[i-1] + 1$。

然后来计算答案。

枚举修改的元素是 $a[i]$。修改后，如果满足如下条件，那么可以把左右两侧的等差子数组拼在一起：

- 设 $d = \dfrac{a[i+1] - a[i-1]}{2}$，这必须是个整数（题目要求修改后的数是整数）。把 $a[i]$ 改成 $d$。
- $a[i-1] - a[i-2] = d = a[i+2] - a[i+1]$。

拼接后的等差子数组的长度为

$$
\textit{pre}[i-1] + 1 + \textit{suf}[i+1]
$$

用上式更新答案的最大值。

还有四种特殊情况：

- 只满足 $a[i-1] - a[i-2] = d$，只能在 $\textit{pre}[i-1]$ 的后面拼接 $a[i]$ 和 $a[i+1]$，拼接后的等差子数组的长度为 $\textit{pre}[i-1] + 2$。
- 只满足 $a[i+2] - a[i+1] = d$，只能在 $\textit{suf}[i+1]$ 的前面拼接 $a[i]$ 和 $a[i-1]$，拼接后的等差子数组的长度为 $\textit{suf}[i+1] + 2$。
- 修改 $a[i]$，拼在 $\textit{pre}[i-1]$ 的后面，拼接后的等差子数组的长度为 $\textit{pre}[i-1] + 1$。
- 修改 $a[i]$，拼在 $\textit{suf}[i+1]$ 的前面，拼接后的等差子数组的长度为 $\textit{suf}[i+1] + 1$。

[本题视频讲解](https://www.bilibili.com/video/BV1DvwTzbE1n/?t=11m23s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def calc(self, nums: List[int]) -> List[int]:
        n = len(nums)
        pre = [0] * n
        pre[0] = 1
        pre[1] = 2
        for i in range(2, n):
            if nums[i - 2] + nums[i] == nums[i - 1] * 2:  # 三个数等差
                pre[i] = pre[i - 1] + 1
            else:
                pre[i] = 2
        return pre

    def longestArithmetic(self, nums: List[int]) -> int:
        n = len(nums)
        pre = self.calc(nums)
        ans = max(pre) + 1
        if ans >= n:  # 整个数组是等差的，或者修改端点元素后是等差的
            return n

        suf = self.calc(nums[::-1])[::-1]
        # 注意 max(pre) == max(suf)，无需重复计算

        for i in range(1, n - 1):
            # 把 nums[i] 改成 d2 / 2
            d2 = nums[i + 1] - nums[i - 1]
            if d2 % 2:  # d2 / 2 必须是整数
                continue

            ok_left = i > 1 and nums[i - 1] - nums[i - 2] == d2 // 2
            ok_right = i + 2 < n and nums[i + 2] - nums[i + 1] == d2 // 2

            if ok_left and ok_right:
                ans = max(ans, pre[i - 1] + 1 + suf[i + 1])
            elif ok_left:
                ans = max(ans, pre[i - 1] + 2)
            elif ok_right:
                ans = max(ans, suf[i + 1] + 2)

        return ans
```

```java [sol-Java]
class Solution {
    public int longestArithmetic(int[] nums) {
        int[] pre = calc(nums);
        int ans = 0;
        for (int x : pre) {
            ans = Math.max(ans, x);
        }
        ans++;

        int n = nums.length;
        if (ans >= n) { // 整个数组是等差的，或者修改端点元素后是等差的
            return n;
        }

        reverse(nums);
        int[] suf = calc(nums);
        reverse(suf);
        reverse(nums);
        // 注意 max(pre) == max(suf)，无需重复计算

        for (int i = 1; i < n - 1; i++) {
            // 把 nums[i] 改成 d2 / 2
            int d2 = nums[i + 1] - nums[i - 1];
            if (d2 % 2 != 0) { // d2 / 2 必须是整数
                continue;
            }

            boolean okLeft = i > 1 && nums[i - 1] - nums[i - 2] == d2 / 2;
            boolean okRight = i + 2 < n && nums[i + 2] - nums[i + 1] == d2 / 2;

            if (okLeft && okRight) {
                ans = Math.max(ans, pre[i - 1] + 1 + suf[i + 1]);
            } else if (okLeft) {
                ans = Math.max(ans, pre[i - 1] + 2);
            } else if (okRight) {
                ans = Math.max(ans, suf[i + 1] + 2);
            }
        }

        return ans;
    }

    private int[] calc(int[] nums) {
        int n = nums.length;
        int[] pre = new int[n];
        pre[0] = 1;
        pre[1] = 2;
        for (int i = 2; i < n; i++) {
            if (nums[i - 2] + nums[i] == nums[i - 1] * 2) { // 三个数等差
                pre[i] = pre[i - 1] + 1;
            } else {
                pre[i] = 2;
            }
        }
        return pre;
    }

    private void reverse(int[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre(n);
        pre[0] = 1;
        pre[1] = 2;
        for (int i = 2; i < n; i++) {
            if (nums[i - 2] + nums[i] == nums[i - 1] * 2) { // 三个数等差
                pre[i] = pre[i - 1] + 1;
            } else {
                pre[i] = 2;
            }
        }
        return pre;
    }

public:
    int longestArithmetic(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre = calc(nums);
        int ans = ranges::max(pre) + 1;
        if (ans >= n) { // 整个数组是等差的，或者修改端点元素后是等差的
            return n;
        }

        ranges::reverse(nums);
        vector<int> suf = calc(nums);
        ranges::reverse(suf);
        ranges::reverse(nums);
        // 注意 max(pre) == max(suf)，无需重复计算

        for (int i = 1; i < n - 1; i++) {
            // 把 nums[i] 改成 d2 / 2
            int d2 = nums[i + 1] - nums[i - 1];
            if (d2 % 2) { // d2 / 2 必须是整数
                continue;
            }

            bool ok_left = i > 1 && nums[i - 1] - nums[i - 2] == d2 / 2;
            bool ok_right = i + 2 < n && nums[i + 2] - nums[i + 1] == d2 / 2;

            if (ok_left && ok_right) {
                ans = max(ans, pre[i - 1] + 1 + suf[i + 1]);
            } else if (ok_left) {
                ans = max(ans, pre[i - 1] + 2);
            } else if (ok_right) {
                ans = max(ans, suf[i + 1] + 2);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func calc(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	pre[1] = 2
	for i := 2; i < n; i++ {
		if nums[i-2]+nums[i] == nums[i-1]*2 { // 三个数等差
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 2
		}
	}
	return pre
}

func longestArithmetic(nums []int) (ans int) {
	n := len(nums)
	pre := calc(nums)
	ans = slices.Max(pre) + 1
	if ans >= n { // 整个数组是等差的，或者修改端点元素后是等差的
		return n
	}

	slices.Reverse(nums)
	suf := calc(nums)
	slices.Reverse(suf)
	slices.Reverse(nums)
	// 注意 max(pre) == max(suf)，无需重复计算

	for i := 1; i < n-1; i++ {
		// 把 nums[i] 改成 d2/2
		d2 := nums[i+1] - nums[i-1]
		if d2%2 != 0 { // d2/2 必须是整数
			continue
		}

		okLeft := i > 1 && nums[i-1]-nums[i-2] == d2/2
		okRight := i+2 < n && nums[i+2]-nums[i+1] == d2/2

		if okLeft && okRight {
			ans = max(ans, pre[i-1]+1+suf[i+1])
		} else if okLeft {
			ans = max(ans, pre[i-1]+2)
		} else if okRight {
			ans = max(ans, suf[i+1]+2)
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：分组循环

**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

[例题讲解](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)

对于本题，$\textit{nums}$ 包含若干段等差子数组。修改等差子数组的中间元素是无意义的（值没变），所以修改操作只会发生在等差子数组的端点旁边。用分组循环找到等差子数组的端点。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def longestArithmetic(self, nums: List[int]) -> int:
        ans = 0
        n = len(nums)
        i = 1
        while True:
            # 枚举 i-1 和 i 作为等差子数组的前两项，且我们不改 nums[i-1] 和 nums[i]
            start = i - 1
            d = nums[i] - nums[i - 1]

            # 往右移动，直到 nums[i] 不满足等差
            i += 1
            while i < n and nums[i] - nums[i - 1] == d:
                i += 1

            # 现在 [start, i-1] 是等差子数组
            # 要想让子数组更长，要么改 nums[start-1]，要么改 nums[i]

            # 改 nums[start-1]
            if start >= 2 and nums[start] - nums[start - 2] == d * 2:  # 可以和 nums[start-2] 连起来
                ans = max(ans, i - start + 2)  # 等差子数组 [start-2, i-1]
                # 继续往左延长的情况等同于上一段继续往右延长，无需重复计算
            else:  # 子数组左端点最远只能到 max(start-1,0)
                ans = max(ans, i - max(start - 1, 0))  # 等差子数组 [max(start-1,0), i-1]

            if i == n:
                return ans

            # 改 nums[i]
            if i < n - 1 and nums[i + 1] - nums[i - 1] == d * 2:  # 可以和 nums[i+1] 连起来
                # 继续往右延长
                j = i + 2
                while j < n and nums[j] - nums[j - 1] == d:
                    j += 1
                ans = max(ans, j - start)  # 等差子数组 [start, j-1]
            else:  # 子数组右端点最远只能到 i
                ans = max(ans, i - start + 1)  # 等差子数组 [start, i]
```

```java [sol-Java]
class Solution {
    public int longestArithmetic(int[] nums) {
        int ans = 0;
        int n = nums.length;
        int i = 1;
        while (true) {
            // 枚举 i-1 和 i 作为等差子数组的前两项，且我们不改 nums[i-1] 和 nums[i]
            int start = i - 1;
            int d = nums[i] - nums[i - 1];

            // 往右移动，直到 nums[i] 不满足等差
            i++;
            while (i < n && nums[i] - nums[i - 1] == d) {
                i++;
            }

            // 现在 [start, i-1] 是等差子数组
            // 要想让子数组更长，要么改 nums[start-1]，要么改 nums[i]

            // 改 nums[start-1]
            if (start >= 2 && nums[start] - nums[start - 2] == d * 2) { // 可以和 nums[start-2] 连起来
                ans = Math.max(ans, i - start + 2); // 等差子数组 [start-2, i-1]
                // 继续往左延长的情况等同于上一段继续往右延长，无需重复计算
            } else { // 子数组左端点最远只能到 max(start-1,0)
                ans = Math.max(ans, i - Math.max(start - 1, 0)); // 等差子数组 [max(start-1,0), i-1]
            }

            if (i == n) {
                return ans;
            }

            // 改 nums[i]
            if (i < n - 1 && nums[i + 1] - nums[i - 1] == d * 2) { // 可以和 nums[i+1] 连起来
                // 继续往右延长
                int j = i + 2;
                while (j < n && nums[j] - nums[j - 1] == d) {
                    j++;
                }
                ans = Math.max(ans, j - start); // 等差子数组 [start, j-1]
            } else { // 子数组右端点最远只能到 i
                ans = Math.max(ans, i - start + 1); // 等差子数组 [start, i]
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestArithmetic(vector<int>& nums) {
        int ans = 0;
        int i = 1, n = nums.size();
        while (true) {
            // 枚举 i-1 和 i 作为等差子数组的前两项，且我们不改 nums[i-1] 和 nums[i]
            int start = i - 1;
            int d = nums[i] - nums[i - 1];

            // 往右移动，直到 nums[i] 不满足等差
            i++;
            while (i < n && nums[i] - nums[i - 1] == d) {
                i++;
            }

            // 现在 [start, i-1] 是等差子数组
            // 要想让子数组更长，要么改 nums[start-1]，要么改 nums[i]

            // 改 nums[start-1]
            if (start >= 2 && nums[start] - nums[start - 2] == d * 2) { // 可以和 nums[start-2] 连起来
                ans = max(ans, i - start + 2); // 等差子数组 [start-2, i-1]
                // 继续往左延长的情况等同于上一段继续往右延长，无需重复计算
            } else { // 子数组左端点最远只能到 max(start-1,0)
                ans = max(ans, i - max(start - 1, 0)); // 等差子数组 [max(start-1,0), i-1]
            }

            if (i == n) {
                return ans;
            }

            // 改 nums[i]
            if (i < n - 1 && nums[i + 1] - nums[i - 1] == d * 2) { // 可以和 nums[i+1] 连起来
                // 继续往右延长
                int j = i + 2;
                while (j < n && nums[j] - nums[j - 1] == d) {
                    j++;
                }
                ans = max(ans, j - start); // 等差子数组 [start, j-1]
            } else { // 子数组右端点最远只能到 i
                ans = max(ans, i - start + 1); // 等差子数组 [start, i]
            }
        }
    }
};
```

```go [sol-Go]
func longestArithmetic(nums []int) (ans int) {
	n := len(nums)
	for i := 1; ; {
		// 枚举 i-1 和 i 作为等差子数组的前两项，且我们不改 nums[i-1] 和 nums[i]
		start := i - 1
		d := nums[i] - nums[i-1]

		// 往右移动，直到 nums[i] 不满足等差
		for i++; i < n && nums[i]-nums[i-1] == d; i++ {
		}

		// 现在 [start, i-1] 是等差子数组
		// 要想让子数组更长，要么改 nums[start-1]，要么改 nums[i]

		// 改 nums[start-1]
		if start >= 2 && nums[start]-nums[start-2] == d*2 { // 可以和 nums[start-2] 连起来
			ans = max(ans, i-start+2) // 等差子数组 [start-2, i-1]
			// 继续往左延长的情况等同于上一段继续往右延长，无需重复计算
		} else { // 子数组左端点最远只能到 max(start-1,0)
			ans = max(ans, i-max(start-1, 0)) // 等差子数组 [max(start-1,0), i-1]
		}

		if i == n {
			return
		}

		// 改 nums[i]
		if i < n-1 && nums[i+1]-nums[i-1] == d*2 { // 可以和 nums[i+1] 连起来
			// 继续往右延长
			j := i + 2
			for ; j < n && nums[j]-nums[j-1] == d; j++ {
			}
			ans = max(ans, j-start) // 等差子数组 [start, j-1]
		} else { // 子数组右端点最远只能到 i
			ans = max(ans, i-start+1) // 等差子数组 [start, i]
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。每个数至多被遍历两次。
- 空间复杂度：$\mathcal{O}(1)$。

## 附：状态机 DP 做法

```py
class Solution:
    def longestArithmetic(self, nums: List[int]) -> int:
        # 返回以 i 结尾的最长等差子数组的长度，其中 nums[i] 不改
        # left 表示剩余操作次数
        @cache
        def dfs(i: int, left: int) -> int:
            if i <= 1:
                return i + 1  # 等差子数组 [0, i]，长为 i+1

            # 不改
            res = 2  # 等差子数组 [i-1, i]，长为 2
            if nums[i - 2] + nums[i] == nums[i - 1] * 2:  # 三个数等差
                res = dfs(i - 1, left) + 1  # 以 i-1 结尾的最长等差子数组 + [i, i]

            if left:
                res = max(res, 3)  # 改 nums[i-2]，得到等差子数组 [i-2, i]，长为 3
                if i >= 3:
                    # 改 nums[i-1]
                    if (nums[i - 2] - nums[i - 3]) * 2 == nums[i] - nums[i - 2]:
                        res = max(res, dfs(i - 2, 0) + 2)  # 以 i-2 结尾的最长等差子数组 + [i-1, i]

                    # 改 nums[i-2]
                    d = nums[i] - nums[i - 1]
                    if nums[i - 1] - nums[i - 3] == d * 2:
                        res = max(res, 4)  # 等差子数组 [i-3, i]，长为 4
                        if i >= 4 and nums[i - 3] - nums[i - 4] == d:
                            res = max(res, dfs(i - 3, 0) + 3)  # 以 i-3 结尾的最长等差子数组 + [i-2, i]

            return res

        n = len(nums)
        ans1 = max(dfs(i, 1) for i in range(n))  # 不改 nums[i]
        ans2 = max(dfs(i - 1, 0) for i in range(1, n)) + 1  # 改 nums[i]
        return max(ans1, ans2)
```

1:1 翻译成递推：

```py
# 手写 max 更快
fmax = lambda a, b: b if b > a else a

class Solution:
    def longestArithmetic(self, nums: List[int]) -> int:
        n = len(nums)
        f = [[0, 0] for _ in range(n)]
        f[0][0] = f[0][1] = 1
        f[1][0] = f[1][1] = 2
        for i in range(2, n):
            for left in range(2):
                # 不改
                res = 2  # 等差子数组 [i-1, i]，长为 2
                if nums[i - 2] + nums[i] == nums[i - 1] * 2:  # 三个数等差
                    res = f[i - 1][left] + 1  # 以 i-1 结尾的最长等差子数组 + [i, i]

                if left:
                    res = fmax(res, 3)  # 改 nums[i-2]，得到等差子数组 [i-2, i]，长为 3
                    if i >= 3:
                        # 改 nums[i-1]
                        if (nums[i - 2] - nums[i - 3]) * 2 == nums[i] - nums[i - 2]:
                            res = fmax(res, f[i - 2][0] + 2)  # 以 i-2 结尾的最长等差子数组 + [i-1, i]

                        # 改 nums[i-2]
                        d = nums[i] - nums[i - 1]
                        if nums[i - 1] - nums[i - 3] == d * 2:
                            res = fmax(res, 4)  # 等差子数组 [i-3, i]，长为 4
                            if i >= 4 and nums[i - 3] - nums[i - 4] == d:
                                res = fmax(res, f[i - 3][0] + 3)  # 以 i-3 结尾的最长等差子数组 + [i-2, i]

                f[i][left] = res

        ans1 = max(f[i][1] for i in range(n))  # 不改 nums[i]
        ans2 = max(f[i][0] for i in range(n - 1)) + 1  # 改 nums[i+1]
        return max(ans1, ans2)
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[3830. 移除至多一个元素后的最长交替子数组](https://leetcode.cn/problems/longest-alternating-subarray-after-removing-at-most-one-element/)

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
