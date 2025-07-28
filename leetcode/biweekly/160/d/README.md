一个月前，出过类似的题目 [3574. 最大子数组 GCD 分数](https://leetcode.cn/problems/maximize-subarray-gcd-score/)。

## 转化

看到**最小化最大值**，想到二分答案。关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

原问题转化成一个判定性问题：

- 给定上界 $\textit{upper}$，能否通过至多 $\textit{maxC}$ 次修改，让 $\textit{nums}$ 的稳定性因子（最长稳定子数组的长度）不超过 $\textit{upper}$？

如果可以做到，说明答案 $\le \textit{upper}$（上界可能更小），否则答案 $> \textit{upper}$（上界必须增大）。这也是为什么我们可以二分答案。

## 思路

**前置知识**：[LogTrick 入门教程](https://zhuanlan.zhihu.com/p/1933215367158830792)。

遍历 $\textit{nums}$，遍历到 $\textit{nums}[i]$ 时，计算以 $i$ 为右端点的所有子数组的最大公因数（GCD）。对于 GCD $\ge 2$ 的以 $i$ 为右端点的最长子数组，如果其长度 $> \textit{upper}$，那么必须修改。

修改谁呢？

如果修改 $i$ 左边的元素，那么 $\textit{nums}[i]$ 可能与后续元素组成不符合要求的子数组。

如果修改 $\textit{nums}[i]=1$，那么包含 $\textit{nums}[i]$ 的任意子数组的 GCD 均为 $1$。这样后续就不用考虑 $\textit{nums}[i]$ 了，相比修改 $i$ 左边的元素更好。

如果遍历结束，修改次数 $\le \textit{maxC}$，那么满足判定性问题的要求，否则不满足。

## 细节

### 1)

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的，喜欢哪种写法就用哪种。

- 开区间左端点初始值：$-1$。子数组长度不能为负，无法满足要求。
- 开区间右端点初始值：$n$。无需操作，一定满足要求。
- 开区间右端点初始值（优化）：思考这样一个问题：有 $n$ 个白球排成一行，把其中的 $\textit{maxC}$ 个白球涂成黑色。涂色后，最长连续白球的长度，最小是多少？我们把剩余的 $n-\textit{maxC}$ 个白球均分成 $\textit{maxC}+1$ 段，均分后，最长连续白球的长度为 $\left\lceil\dfrac{n-\textit{maxC}}{\textit{maxC}+1}\right\rceil$。其作为上界时，一定满足要求。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

### 2)

关于上取整的计算，当 $a$ 为非负整数，$b$ 为正整数时，有恒等式

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a+b-1}{b}\right\rfloor
$$

见 [上取整下取整转换公式的证明](https://zhuanlan.zhihu.com/p/1890356682149838951)。

所以有

$$
\left\lceil\dfrac{n-\textit{maxC}}{\textit{maxC}+1}\right\rceil = \left\lfloor\dfrac{n}{\textit{maxC}+1}\right\rfloor
$$

## 优化前

```py [sol-Python3]
class Solution:
    def minStable(self, nums: List[int], maxC: int) -> int:
        def check(upper: int) -> bool:
            intervals = []  # (子数组 GCD，最小左端点)
            c = maxC
            for i, x in enumerate(nums):
                # 计算以 i 为右端点的子数组 GCD
                for p in intervals:
                    p[0] = gcd(p[0], x)
                # nums[i] 单独一个数作为子数组
                intervals.append([x, i])

                # 去重（合并 GCD 相同的区间）
                idx = 1
                for j in range(1, len(intervals)):
                    if intervals[j][0] != intervals[j - 1][0]:
                        intervals[idx] = intervals[j]
                        idx += 1
                del intervals[idx:]

                # intervals 的性质：越靠左，GCD 越小

                # 我们只关心 GCD >= 2 的子数组
                if intervals[0][0] == 1:
                    intervals.pop(0)

                # intervals[0] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
                if intervals and i - intervals[0][1] + 1 > upper:
                    if c == 0:
                        return False
                    c -= 1
                    intervals.clear()  # 修改后 GCD 均为 1，直接清空
            return True

        return bisect_left(range(len(nums) // (maxC + 1)), True, key=check)
```

```java [sol-Java]
class Solution {
    public int minStable(int[] nums, int maxC) {
        int left = -1;
        int right = nums.length / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(nums, maxC, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int maxC, int upper) {
        List<int[]> intervals = new ArrayList<>(); // 每个元素是 (子数组 GCD，最小左端点)
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // 计算以 i 为右端点的子数组 GCD
            for (int[] interval : intervals) {
                interval[0] = gcd(interval[0], x);
            }
            // nums[i] 单独一个数作为子数组
            intervals.add(new int[]{x, i});

            // 去重（合并 GCD 相同的区间）
            int idx = 1;
            for (int j = 1; j < intervals.size(); j++) {
                if (intervals.get(j)[0] != intervals.get(j - 1)[0]) {
                    intervals.set(idx, intervals.get(j));
                    idx++;
                }
            }
            intervals.subList(idx, intervals.size()).clear();

            // intervals 的性质：越靠左，GCD 越小

            // 我们只关心 GCD >= 2 的子数组
            if (intervals.get(0)[0] == 1) {
                intervals.remove(0);
            }

            // intervals[0] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
            if (!intervals.isEmpty() && i - intervals.get(0)[1] + 1 > upper) {
                if (maxC == 0) {
                    return false;
                }
                maxC--;
                intervals.clear(); // 修改后 GCD 均为 1，直接清空
            }
        }
        return true;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minStable(vector<int>& nums, int maxC) {
        int n = nums.size();

        auto check = [&](int upper) -> bool {
            vector<pair<int, int>> intervals; // pair{子数组 GCD，最小左端点}
            int c = maxC;
            for (int i = 0; i < n; i++) {
                int x = nums[i];
                // 计算以 i 为右端点的子数组 GCD
                for (auto& [g, _] : intervals) {
                    g = gcd(g, x);
                }
                // nums[i] 单独一个数作为子数组
                intervals.emplace_back(x, i);

                // 去重（合并 GCD 相同的区间）
                int idx = 1;
                for (int j = 1; j < intervals.size(); j++) {
                    if (intervals[j].first != intervals[j - 1].first) {
                        intervals[idx++] = intervals[j];
                    }
                }
                intervals.resize(idx);

                // intervals 的性质：越靠左，GCD 越小

                // 我们只关心 GCD >= 2 的子数组
                if (intervals[0].first == 1) {
                    intervals.erase(intervals.begin());
                }

                // intervals[0] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
                if (!intervals.empty() && i - intervals[0].second + 1 > upper) {
                    if (c == 0) {
                        return false;
                    }
                    c--;
                    intervals.clear(); // 修改后 GCD 均为 1，直接清空
                }
            }
            return true;
        };

        int left = -1, right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minStable(nums []int, maxC int) int {
	ans := sort.Search(len(nums)/(maxC+1), func(upper int) bool {
		type interval struct{ gcd, l int } // 子数组 GCD，最小左端点
		intervals := []interval{}
		c := maxC
		for i, x := range nums {
			// 计算以 i 为右端点的子数组 GCD
			for j, p := range intervals {
				intervals[j].gcd = gcd(p.gcd, x)
			}
			// nums[i] 单独一个数作为子数组
			intervals = append(intervals, interval{x, i})

			// 去重（合并 GCD 相同的区间）
			idx := 1
			for j := 1; j < len(intervals); j++ {
				if intervals[j].gcd != intervals[j-1].gcd {
					intervals[idx] = intervals[j]
					idx++
				}
			}
			intervals = intervals[:idx]

			// intervals 的性质：越靠左，GCD 越小

			// 我们只关心 GCD >= 2 的子数组
			if intervals[0].gcd == 1 {
				intervals = intervals[1:]
			}

			// intervals[0] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
			if len(intervals) > 0 && i-intervals[0].l+1 > upper { // 必须修改 nums[i]=1
				if c == 0 {
					return false
				}
				c--
				intervals = intervals[:0] // 修改后 GCD 均为 1，直接清空
			}
		}
		return true
	})
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U\log M)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$M=n/\textit{maxC}$。每次二分是 $\mathcal{O}(n\log U)$。外层循环每次会增加一个区间，这个区间在整个算法过程中，要么合并到其他区间中（消失啦），要么其 GCD 一共减少 $\mathcal{O}(\log U)$ 次，所以每个区间的 GCD 的计算过程会贡献 $\mathcal{O}(\log U)$ 个循环次数，所以每次二分的循环次数是 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(\log U)$。$\textit{intervals}$ 的长度为 $\mathcal{O}(\log U)$。

## 优化（写法一）

每次二分都要算 logTrick，有点慢。其实我们只需要计算，对于每个 $i$，以 $i$ 为右端点的子数组 GCD $\ge 2$ 时，子数组的左端点的最小值，记作 $\textit{leftMin}[i]$。如果没有 GCD $\ge 2$ 的子数组，则 $\textit{leftMin}[i]=n$（或者任意大于 $i$ 的数）。这可以在二分之前用 logTrick 预处理出来。

然后二分。从第一个可能会修改的右端点 $i= \textit{upper}$ 开始，如果 $i-\textit{leftMin}[i]+1 > \textit{upper}$，那么修改 $i$，跳到下一个可能会修改的右端点 $i+\textit{upper}+1$，否则不修改，把 $i$ 加一。

```py [sol-Python3]
class Solution:
    def minStable(self, nums: List[int], maxC: int) -> int:
        n = len(nums)
        left_min = [n] * n
        intervals = [[1, 0]]  # 哨兵
        for i, x in enumerate(nums):
            # 计算以 i 为右端点的子数组 GCD
            for p in intervals:
                p[0] = gcd(p[0], x)
            # nums[i] 单独一个数作为子数组
            intervals.append([x, i])

            # 去重（合并 GCD 相同的区间）
            idx = 1
            for j in range(1, len(intervals)):
                if intervals[j][0] != intervals[j - 1][0]:
                    intervals[idx] = intervals[j]
                    idx += 1
            del intervals[idx:]

            # 由于我们添加了哨兵，intervals[1] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
            if len(intervals) > 1:
                left_min[i] = intervals[1][1]

        def check(upper: int) -> bool:
            c = maxC
            i = upper
            while i < n:
                if i - left_min[i] + 1 > upper:
                    if c == 0:
                        return False
                    c -= 1
                    i += upper + 1
                else:
                    i += 1
            return True

        return bisect_left(range(len(nums) // (maxC + 1)), True, key=check)
```

```java [sol-Java]
// 更快的写法见【Java 数组】
class Solution {
    public int minStable(int[] nums, int maxC) {
        int n = nums.length;
        int[] leftMin = new int[n];
        List<int[]> intervals = new ArrayList<>(); // 每个元素是 (子数组 GCD，最小左端点)
        intervals.add(new int[]{1, 0}); // 哨兵

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // 计算以 i 为右端点的子数组 GCD
            for (int[] interval : intervals) {
                interval[0] = gcd(interval[0], x);
            }
            // nums[i] 单独一个数作为子数组
            intervals.add(new int[]{x, i});

            // 去重（合并 GCD 相同的区间）
            int idx = 1;
            for (int j = 1; j < intervals.size(); j++) {
                if (intervals.get(j)[0] != intervals.get(j - 1)[0]) {
                    intervals.set(idx, intervals.get(j));
                    idx++;
                }
            }
            intervals.subList(idx, intervals.size()).clear();

            // 由于我们添加了哨兵，intervals[1] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
            leftMin[i] = intervals.size() > 1 ? intervals.get(1)[1] : n;
        }

        int left = -1;
        int right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(leftMin, maxC, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] leftMin, int maxC, int upper) {
        int i = upper;
        while (i < leftMin.length) {
            if (i - leftMin[i] + 1 > upper) {
                if (maxC == 0) {
                    return false;
                }
                maxC--;
                i += upper + 1;
            } else {
                i++;
            }
        }
        return true;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minStable(int[] nums, int maxC) {
        int n = nums.length;
        int[] leftMin = new int[n];
        int[][] intervals = new int[31][2];
        intervals[0][0] = 1; // 哨兵
        int size = 1;

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];

            // 计算以 i 为右端点的子数组 GCD
            for (int j = 1; j < size; j++) {
                intervals[j][0] = gcd(intervals[j][0], x);
            }
            // nums[i] 单独一个数作为子数组
            intervals[size][0] = x;
            intervals[size][1] = i;
            size++;

            // 去重（合并 GCD 相同的区间）
            int idx = 1;
            for (int j = 1; j < size; j++) {
                if (intervals[j][0] != intervals[j - 1][0]) {
                    intervals[idx][0] = intervals[j][0];
                    intervals[idx][1] = intervals[j][1];
                    idx++;
                }
            }
            size = idx;

            leftMin[i] = size > 1 ? intervals[1][1] : n;
        }

        int left = -1;
        int right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(leftMin, maxC, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] leftMin, int maxC, int upper) {
        int i = upper;
        while (i < leftMin.length) {
            if (i - leftMin[i] + 1 > upper) {
                if (maxC == 0) {
                    return false;
                }
                maxC--;
                i += upper + 1;
            } else {
                i++;
            }
        }
        return true;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minStable(vector<int>& nums, int maxC) {
        int n = nums.size();
        vector<int> left_min(n);
        vector<pair<int, int>> intervals = {{1, 0}}; // 哨兵

        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 计算以 i 为右端点的子数组 GCD
            for (auto& [g, _] : intervals) {
                g = gcd(g, x);
            }
            // nums[i] 单独一个数作为子数组
            intervals.emplace_back(x, i);

            // 去重（合并 GCD 相同的区间）
            int idx = 1;
            for (int j = 1; j < intervals.size(); j++) {
                if (intervals[j].first != intervals[j - 1].first) {
                    intervals[idx++] = intervals[j];
                }
            }
            intervals.resize(idx);

            // 由于我们添加了哨兵，intervals[1] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
            left_min[i] = intervals.size() > 1 ? intervals[1].second : n;
        }

        auto check = [&](int upper) -> bool {
            int c = maxC;
            int i = upper;
            while (i < n) {
                if (i - left_min[i] + 1 > upper) {
                    if (c == 0) {
                        return false;
                    }
                    c--;
                    i += upper + 1;
                } else {
                    i++;
                }
            }
            return true;
        };

        int left = -1, right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minStable(nums []int, maxC int) int {
	n := len(nums)
	leftMin := make([]int, n)
	type interval struct{ gcd, l int } // 子数组 GCD，最小左端点
	intervals := []interval{{1, 0}} // 哨兵
	for i, x := range nums {
		// 计算以 i 为右端点的子数组 GCD
		for j, p := range intervals {
			intervals[j].gcd = gcd(p.gcd, x)
		}
		// nums[i] 单独一个数作为子数组
		intervals = append(intervals, interval{x, i})

		// 去重（合并 GCD 相同的区间）
		idx := 1
		for j := 1; j < len(intervals); j++ {
			if intervals[j].gcd != intervals[j-1].gcd {
				intervals[idx] = intervals[j]
				idx++
			}
		}
		intervals = intervals[:idx]

		// 由于我们添加了哨兵，intervals[1] 的 GCD >= 2 且最长，取其区间左端点作为子数组的最小左端点
		if len(intervals) > 1 {
			leftMin[i] = intervals[1].l
		} else {
			leftMin[i] = n
		}
	}

	ans := sort.Search(n/(maxC+1), func(upper int) bool {
		c := maxC
		i := upper
		for i < n {
			if i-leftMin[i]+1 > upper {
				if c == 0 {
					return false
				}
				c--
				i += upper + 1
			} else {
				i++
			}
		}
		return true
	})
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + n\log M)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$M=n/\textit{maxC}$。计算 logTrick 时，外层循环每次会增加一个区间，这个区间在整个算法过程中，要么合并到其他区间中（消失啦），要么其 GCD 一共减少 $\mathcal{O}(\log U)$ 次，所以每个区间的 GCD 的计算过程会贡献 $\mathcal{O}(\log U)$ 个循环次数，所以 logTrick 的循环次数是 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 优化（写法二）

也可以用 [3171. 找到按位或最接近 K 的子数组](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/solutions/2798206/li-yong-and-de-xing-zhi-pythonjavacgo-by-gg4d/) 介绍的栈 + 滑动窗口预处理 $\textit{leftMin}$。

```py [sol-Python3]
class Solution:
    def minStable(self, nums: List[int], maxC: int) -> int:
        n = len(nums)
        left_min = [0] * n
        left = bottom = right_gcd = 0
        for i, x in enumerate(nums):
            right_gcd = gcd(right_gcd, x)
            while left <= i and gcd(nums[left], right_gcd) == 1:
                if bottom <= left:
                    # 重新构建一个栈
                    # 由于 left 即将移出窗口，只需计算到 left+1
                    for j in range(i - 1, left, -1):
                        nums[j] = gcd(nums[j], nums[j + 1])
                    bottom = i
                    right_gcd = 0
                left += 1
            left_min[i] = left

        def check(upper: int) -> bool:
            c = maxC
            i = upper
            while i < n:
                if i - left_min[i] + 1 > upper:
                    if c == 0:
                        return False
                    c -= 1
                    i += upper + 1
                else:
                    i += 1
            return True

        return bisect_left(range(len(nums) // (maxC + 1)), True, key=check)
```

```java [sol-Java]
class Solution {
    public int minStable(int[] nums, int maxC) {
        int n = nums.length;
        int[] leftMin = new int[n];
        int left = 0, bottom = 0, rightGcd = 0;
        for (int i = 0; i < n; i++) {
            rightGcd = gcd(rightGcd, nums[i]);
            while (left <= i && gcd(nums[left], rightGcd) == 1) {
                if (bottom <= left) {
                    // 重新构建一个栈
                    // 由于 left 即将移出窗口，只需计算到 left+1
                    for (int j = i - 1; j > left; j--) {
                        nums[j] = gcd(nums[j], nums[j + 1]);
                    }
                    bottom = i;
                    rightGcd = 0;
                }
                left++;
            }
            leftMin[i] = left;
        }

        left = -1;
        int right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (check(leftMin, maxC, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] leftMin, int maxC, int upper) {
        int i = upper;
        while (i < leftMin.length) {
            if (i - leftMin[i] + 1 > upper) {
                if (maxC == 0) {
                    return false;
                }
                maxC--;
                i += upper + 1;
            } else {
                i++;
            }
        }
        return true;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minStable(vector<int>& nums, int maxC) {
        int n = nums.size();
        vector<int> left_min(n);
        int bottom = 0, right_gcd = 0;
        for (int i = 0, left = 0; i < n; i++) {
            right_gcd = gcd(right_gcd, nums[i]);
            while (left <= i && gcd(nums[left], right_gcd) == 1) {
                if (bottom <= left) {
                    // 重新构建一个栈
                    // 由于 left 即将移出窗口，只需计算到 left+1
                    for (int j = i - 1; j > left; j--) {
                        nums[j] = gcd(nums[j], nums[j + 1]);
                    }
                    bottom = i;
                    right_gcd = 0;
                }
                left++;
            }
            left_min[i] = left;
        }

        auto check = [&](int upper) -> bool {
            int c = maxC;
            int i = upper;
            while (i < n) {
                if (i - left_min[i] + 1 > upper) {
                    if (c == 0) {
                        return false;
                    }
                    c--;
                    i += upper + 1;
                } else {
                    i++;
                }
            }
            return true;
        };

        int left = -1, right = n / (maxC + 1);
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minStable(nums []int, maxC int) int {
	n := len(nums)
	leftMin := make([]int, n)
	var left, bottom, rightGcd int
	for i, x := range nums {
		rightGcd = gcd(rightGcd, x)
		for left <= i && gcd(nums[left], rightGcd) == 1 {
			if bottom <= left {
				// 重新构建一个栈
				// 由于 left 即将移出窗口，只需计算到 left+1
				for j := i - 1; j > left; j-- {
					nums[j] = gcd(nums[j], nums[j+1])
				}
				bottom = i
				rightGcd = 0
			}
			left++
		}
		leftMin[i] = left
	}

	ans := sort.Search(n/(maxC+1), func(upper int) bool {
		c := maxC
		i := upper
		for i < n {
			if i-leftMin[i]+1 > upper {
				if c == 0 {
					return false
				}
				c--
				i += upper + 1
			} else {
				i++
			}
		}
		return true
	})
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + n\log M)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$M=n/\textit{maxC}$。见 3171 题我题解中的分析。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 二分题单的「**§2.4 最小化最大值**」。
2. 位运算题单的「**LogTrick**」。

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
