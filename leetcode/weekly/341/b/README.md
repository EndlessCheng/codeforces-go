## 优化前

遍历 $d=\textit{divisors}[i]$，暴力枚举 $\textit{nums}$ 中的数，统计能被 $d$ 整除的数的个数 $\textit{cnt}$。取 $\textit{cnt}$ 最大的 $d$ 作为答案。如果有多个最大的 $\textit{cnt}$，取其中最小的 $d$。

```py [sol-Python3]
class Solution:
    def maxDivScore(self, nums: List[int], divisors: List[int]) -> int:
        max_cnt, ans = -1, 0
        for d in divisors:
            cnt = sum(1 for x in nums if x % d == 0)
            if cnt > max_cnt or cnt == max_cnt and d < ans:
                max_cnt, ans = cnt, d
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDivScore(int[] nums, int[] divisors) {
        int ans = 0;
        int maxCnt = -1;
        for (int d : divisors) {
            int cnt = 0;
            for (int x : nums) {
                if (x % d == 0) {
                    cnt++;
                }
            }
            if (cnt > maxCnt || cnt == maxCnt && d < ans) {
                maxCnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDivScore(vector<int>& nums, vector<int>& divisors) {
        int max_cnt = -1, ans = 0;
        for (int d : divisors) {
            int cnt = ranges::count_if(nums, [&](int x) { return x % d == 0; });
            if (cnt > max_cnt || cnt == max_cnt && d < ans) {
                max_cnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDivScore(nums, divisors []int) (ans int) {
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}
```

```js [sol-JavaScript]
var maxDivScore = function(nums, divisors) {
    let ans = 0;
    let maxCnt = -1;
    for (const d of divisors) {
        let cnt = 0;
        for (const x of nums) {
            if (x % d === 0) {
                cnt++;
            }
        }
        if (cnt > maxCnt || cnt === maxCnt && d < ans) {
            maxCnt = cnt;
            ans = d;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_div_score(nums: Vec<i32>, divisors: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut max_cnt = -1;
        for d in divisors {
            let cnt = nums.iter().filter(|&&x| x % d == 0).count() as i32;
            if cnt > max_cnt || cnt == max_cnt && d < ans {
                max_cnt = cnt;
                ans = d;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{divisors}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 优化 1

注意到，小于 $d$ 的正整数无法被 $d$ 整除。

把 $\textit{nums}$ 排序，从大到小遍历 $\textit{nums}$。只需遍历 $\ge d$ 的 $\textit{nums}[i]$，当 $\textit{nums}[i] < d$ 时，退出内层循环。

```py [sol-Python3]
class Solution:
    def maxDivScore(self, nums: List[int], divisors: List[int]) -> int:
        nums.sort(reverse=True)
        max_cnt, ans = -1, 0
        for d in divisors:
            cnt = 0
            for x in nums:
                if x < d:
                    break
                if x % d == 0:
                    cnt += 1
            if cnt > max_cnt or cnt == max_cnt and d < ans:
                max_cnt, ans = cnt, d
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDivScore(int[] nums, int[] divisors) {
        Arrays.sort(nums);
        int ans = 0;
        int maxCnt = -1;
        for (int d : divisors) {
            int cnt = 0;
            for (int i = nums.length - 1; i >= 0; i--) {
                int x = nums[i];
                if (x < d) {
                    break;
                }
                if (x % d == 0) {
                    cnt++;
                }
            }
            if (cnt > maxCnt || cnt == maxCnt && d < ans) {
                maxCnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDivScore(vector<int>& nums, vector<int>& divisors) {
        ranges::sort(nums, greater()); // 从大到小排序
        int max_cnt = -1, ans = 0;
        for (int d : divisors) {
            int cnt = 0;
            for (int x : nums) {
                if (x < d) {
                    break;
                }
                if (x % d == 0) {
                    cnt++;
                }
            }
            if (cnt > max_cnt || cnt == max_cnt && d < ans) {
                max_cnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDivScore(nums, divisors []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x < d {
				break
			}
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}
```

```js [sol-JavaScript]
var maxDivScore = function(nums, divisors) {
    nums.sort((a, b) => b - a);
    let ans = 0;
    let maxCnt = -1;
    for (const d of divisors) {
        let cnt = 0;
        for (const x of nums) {
            if (x < d) {
                break;
            }
            if (x % d === 0) {
                cnt++;
            }
        }
        if (cnt > maxCnt || cnt === maxCnt && d < ans) {
            maxCnt = cnt;
            ans = d;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_div_score(mut nums: Vec<i32>, divisors: Vec<i32>) -> i32 {
        nums.sort_unstable_by(|a, b| b.cmp(a));
        let mut ans = 0;
        let mut max_cnt = -1;
        for d in divisors {
            let mut cnt = 0;
            for &x in &nums {
                if x < d {
                    break;
                }
                if x % d == 0 {
                    cnt += 1;
                }
            }
            if cnt > max_cnt || cnt == max_cnt && d < ans {
                max_cnt = cnt;
                ans = d;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + nm)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{divisors}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

## 优化 2

在优化 1 的基础上，把 $\textit{divisors}$ 从小到大排序，并统计 $\textit{nums}$ 中的重复元素个数 $\textit{dup}$，例如 $\textit{nums}=[3,3,3,2,1,1]$，其中有 $\textit{dup}=3$ 个数是重复的。

遍历 $d=\textit{divisors}[i]$，如果

$$
(\textit{maxCnt}-\textit{dup} + 1)\cdot d > \max(\textit{nums})
$$

说明 $d$ 的倍数 $d,2d,3d,\cdots,(\textit{maxCnt}-\textit{dup} + 1)\cdot d$ 中的最大值已经超出了 $\textit{nums}$ 的最大值，即使把 $\textit{nums}$ 中的重复元素也算上，我们也无法统计出比 $\textit{maxCnt}$ 还多的倍数。由于我们已经把 $\textit{divisors}$ 从小到大排序了，当前的 $d$ 满足上面的不等式，那后面的更大的 $d$ 也同样满足上面的不等式，所以后面不可能找到一个比 $\textit{maxCnt}$ 更大的数，直接退出外层循环。

对于 C++ 等语言，为避免乘法溢出，可以改为判断

$$
\textit{maxCnt}-\textit{dup} + 1 > \left\lfloor\dfrac{\max(\textit{nums})}{d}\right\rfloor
$$

即

$$
\textit{maxCnt}-\textit{dup} \ge \left\lfloor\dfrac{\max(\textit{nums})}{d}\right\rfloor
$$

```py [sol-Python3]
class Solution:
    def maxDivScore(self, nums: List[int], divisors: List[int]) -> int:
        nums.sort(reverse=True)
        dup = sum(1 for x, y in pairwise(nums) if x == y)
        divisors.sort()
        max_cnt, ans = -1, 0
        for d in divisors:
            if (max_cnt - dup + 1) * d > nums[0]:
                break
            cnt = 0
            for x in nums:
                if x < d:
                    break
                if x % d == 0:
                    cnt += 1
            if cnt > max_cnt:
                max_cnt, ans = cnt, d
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDivScore(int[] nums, int[] divisors) {
        Arrays.sort(nums);
        int n = nums.length;
        int dup = 0;
        for (int i = 1; i < n; i++) {
            if (nums[i] == nums[i - 1]) {
                dup++;
            }
        }
        Arrays.sort(divisors);

        int ans = 0;
        int maxCnt = -1;
        for (int d : divisors) {
            if (maxCnt - dup >= nums[n - 1] / d) {
                break;
            }
            int cnt = 0;
            for (int i = n - 1; i >= 0; i--) {
                int x = nums[i];
                if (x < d) {
                    break;
                }
                if (x % d == 0) {
                    cnt++;
                }
            }
            if (cnt > maxCnt) {
                maxCnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxDivScore(vector<int>& nums, vector<int>& divisors) {
        ranges::sort(nums, greater()); // 从大到小排序
        int dup = 0;
        for (int i = 1; i < nums.size(); i++) {
            dup += nums[i] == nums[i - 1]; // 重复元素个数
        }
        ranges::sort(divisors);
        int max_cnt = -1, ans = 0;
        for (int d : divisors) {
            if (max_cnt - dup >= nums[0] / d) {
                break;
            }
            int cnt = 0;
            for (int x : nums) {
                if (x < d) {
                    break;
                }
                if (x % d == 0) {
                    cnt++;
                }
            }
            if (cnt > max_cnt) {
                max_cnt = cnt;
                ans = d;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDivScore(nums []int, divisors []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	dup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dup++
		}
	}
	slices.Sort(divisors)
	maxCnt := -1
	for _, d := range divisors {
		if (maxCnt-dup+1)*d > nums[0] {
			break
		}
		cnt := 0
		for _, x := range nums {
			if x < d {
				break
			}
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt {
			maxCnt, ans = cnt, d
		}
	}
	return
}
```

```js [sol-JavaScript]
var maxDivScore = function(nums, divisors) {
    nums.sort((a, b) => b - a);
    let dup = 0;
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] === nums[i - 1]) {
            dup++;
        }
    }
    divisors.sort((a, b) => a - b);

    let ans = 0;
    let maxCnt = -1;
    for (const d of divisors) {
        if ((maxCnt - dup + 1) * d > nums[0]) {
            break;
        }
        let cnt = 0;
        for (const x of nums) {
            if (x < d) {
                break;
            }
            if (x % d === 0) {
                cnt++;
            }
        }
        if (cnt > maxCnt) {
            maxCnt = cnt;
            ans = d;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_div_score(mut nums: Vec<i32>, mut divisors: Vec<i32>) -> i32 {
        nums.sort_unstable_by(|a, b| b.cmp(a));
        let dup = nums.windows(2).filter(|w| w[0] == w[1]).count() as i32;
        divisors.sort_unstable();
        let mut ans = 0;
        let mut max_cnt = -1;
        for d in divisors {
            if max_cnt - dup >= nums[0] / d {
                break;
            }
            let mut cnt = 0;
            for &x in &nums {
                if x < d {
                    break;
                }
                if x % d == 0 {
                    cnt += 1;
                }
            }
            if cnt > max_cnt {
                max_cnt = cnt;
                ans = d;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m\log m + nm)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{divisors}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
