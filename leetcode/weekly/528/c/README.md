做法和 [1406. 石子游戏 III](https://leetcode.cn/problems/stone-game-iii/) 是一样的，[我的题解](https://leetcode.cn/problems/stone-game-iii/solutions/4001520/san-chong-fang-fa-zui-da-hua-de-fen-zhi-4gcvw/)。

本题和 1406 的区别是，当前玩家可以继续行动。这种情况只需把「减去 $\textit{dfs}$ 的返回值」改成「加上 $\textit{dfs}$ 的返回值」，相当于加上自己继续游戏的得分（得分差）。

[本题视频讲解](https://www.bilibili.com/video/BV1yrsn6wEHP/)，欢迎点赞关注~

## 答疑

**问**：题目说两位玩家要「最大化自己的得分」，为什么这等价于「最大化得分之差」？

**答**：自己的得分 + 对手的得分 = 数组总和。

得分之差 = 自己的得分 - 对手的得分 = 自己的得分 - (数组总和 - 自己的得分) = 自己的得分 × 2 - 数组总和。

由于数组总和是固定的，所以自己的得分越大，得分之差越大。

## 记忆化搜索

```py [sol-Python3]
class Solution:
    def maxDifference(self, nums: list[int]) -> int:
        n = len(nums)

        # 返回在剩余元素为 nums[i:] 的情况下，(当前玩家得分 - 对手得分) 的最大值
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int) -> int:
            if i == n:
                return 0

            x = nums[i]
            if x % 3 == 0:
                res = x - dfs(i + 1)  # 下一回合是对手
            else:
                res = x + dfs(i + 1)  # 下一回合还是我

            if i < n - 1:
                s = x + nums[i + 1]
                if s % 3 == 0:
                    s -= dfs(i + 2)  # 下一回合是对手
                else:
                    s += dfs(i + 2)  # 下一回合还是我
                res = max(res, s)

            return res

        return dfs(0)
```

```java [sol-Java]
class Solution {
    public long maxDifference(int[] nums) {
        long[] memo = new long[nums.length];
        Arrays.fill(memo, Long.MIN_VALUE);
        return dfs(0, nums, memo);
    }

    // 返回在剩余元素为 nums[i,n-1] 的情况下，(当前玩家得分 - 对手得分) 的最大值
    private long dfs(int i, int[] nums, long[] memo) {
        if (i == nums.length) {
            return 0;
        }

        if (memo[i] != Long.MIN_VALUE) {
            return memo[i];
        }

        int x = nums[i];
        long res;
        if (x % 3 == 0) {
            res = x - dfs(i + 1, nums, memo); // 下一回合是对手
        } else {
            res = x + dfs(i + 1, nums, memo); // 下一回合还是我
        }

        if (i < nums.length - 1) {
            long s = x + nums[i + 1];
            if (s % 3 == 0) {
                s -= dfs(i + 2, nums, memo); // 下一回合是对手
            } else {
                s += dfs(i + 2, nums, memo); // 下一回合还是我
            }
            res = Math.max(res, s);
        }

        return memo[i] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxDifference(vector<int>& nums) {
        int n = nums.size();
        vector<long long> memo(n, LLONG_MIN);

        // 返回在剩余元素为 nums[i,n-1] 的情况下，(当前玩家得分 - 对手得分) 的最大值
        auto dfs = [&](this auto&& dfs, int i) -> long long {
            if (i == n) {
                return 0;
            }

            long long& res = memo[i];
            if (res != LLONG_MIN) {
                return res;
            }

            int x = nums[i];
            if (x % 3 == 0) {
                res = x - dfs(i + 1); // 下一回合是对手
            } else {
                res = x + dfs(i + 1); // 下一回合还是我
            }

            if (i < n - 1) {
                long long s = x + nums[i + 1];
                if (s % 3 == 0) {
                    s -= dfs(i + 2); // 下一回合是对手
                } else {
                    s += dfs(i + 2); // 下一回合还是我
                }
                res = max(res, s);
            }

            return res;
        };

        return dfs(0);
    }
};
```

```go [sol-Go]
func maxDifference(nums []int) int64 {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = math.MinInt
	}

	// 返回在剩余元素为 nums[i:] 的情况下，(当前玩家得分 - 对手得分) 的最大值
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == n {
			return
		}

		p := &memo[i]
		if *p != math.MinInt {
			return *p
		}

		x := nums[i]
		if x%3 == 0 {
			res = x - dfs(i+1) // 下一回合是对手
		} else {
			res = x + dfs(i+1) // 下一回合还是我
		}

		if i < n-1 {
			s := x + nums[i+1]
			if s%3 == 0 {
				s -= dfs(i + 2) // 下一回合是对手
			} else {
				s += dfs(i + 2) // 下一回合还是我
			}
			res = max(res, s)
		}

		*p = res
		return
	}

	return int64(dfs(0))
}
```

## 1:1 翻译成递推

```py [sol-Python3]
class Solution:
    def maxDifference(self, nums: list[int]) -> int:
        n = len(nums)
        f = [0] * (n + 1)

        for i in range(n - 1, -1, -1):
            x = nums[i]
            f[i] = x + f[i + 1] if x % 3 else x - f[i + 1]
            if i < n - 1:
                s = x + nums[i + 1]
                s = s + f[i + 2] if s % 3 else s - f[i + 2]
                f[i] = max(f[i], s)

        return f[0]
```

```java [sol-Java]
class Solution {
    public long maxDifference(int[] nums) {
        int n = nums.length;
        long[] f = new long[n + 1];

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            f[i] = x % 3 == 0 ? x - f[i + 1] : x + f[i + 1];
            if (i < n - 1) {
                int s = x + nums[i + 1];
                long res = s % 3 == 0 ? s - f[i + 2] : s + f[i + 2];
                f[i] = Math.max(f[i], res);
            }
        }

        return f[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxDifference(vector<int>& nums) {
        int n = nums.size();
        vector<long long> f(n + 1);

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            f[i] = x % 3 ? x + f[i + 1] : x - f[i + 1];
            if (i < n - 1) {
                int s = x + nums[i + 1];
                long long res = s % 3 ? s + f[i + 2] : s - f[i + 2];
                f[i] = max(f[i], res);
            }
        }

        return f[0];
    }
};
```

```go [sol-Go]
func maxDifference(nums []int) int64 {
	n := len(nums)
	f := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		if x%3 == 0 {
			f[i] = x - f[i+1] // 下一回合是对手
		} else {
			f[i] = x + f[i+1] // 下一回合还是我
		}

		if i < n-1 {
			s := x + nums[i+1]
			if s%3 == 0 {
				s -= f[i+2] // 下一回合是对手
			} else {
				s += f[i+2] // 下一回合还是我
			}
			f[i] = max(f[i], s)
		}
	}

	return int64(f[0])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 空间优化

```py [sol-Python3]
class Solution:
    def maxDifference(self, nums: list[int]) -> int:
        f1, f2 = nums[-1], 0

        for i in range(len(nums) - 2, -1, -1):
            x = nums[i]
            f = x + f1 if x % 3 else x - f1

            s = x + nums[i + 1]
            s = s + f2 if s % 3 else s - f2
            f = max(f, s)

            f2 = f1
            f1 = f

        return f1
```

```java [sol-Java]
class Solution {
    public long maxDifference(int[] nums) {
        int n = nums.length;
        long f1 = nums[n - 1];
        long f2 = 0;

        for (int i = n - 2; i >= 0; i--) {
            int x = nums[i];
            long f = x % 3 == 0 ? x - f1 : x + f1;

            int s = x + nums[i + 1];
            long res = s % 3 == 0 ? s - f2 : s + f2;
            f = Math.max(f, res);

            f2 = f1;
            f1 = f;
        }

        return f1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxDifference(vector<int>& nums) {
        int n = nums.size();
        long long f1 = nums[n - 1], f2 = 0;

        for (int i = n - 2; i >= 0; i--) {
            int x = nums[i];
            long long f = x % 3 ? x + f1 : x - f1;

            int s = x + nums[i + 1];
            long long res = s % 3 ? s + f2 : s - f2;
            f = max(f, res);

            f2 = f1;
            f1 = f;
        }

        return f1;
    }
};
```

```go [sol-Go]
func maxDifference(nums []int) int64 {
	n := len(nums)
	f1, f2 := nums[n-1], 0

	for i := n - 2; i >= 0; i-- {
		x := nums[i]
		f := 0
		if x%3 == 0 {
			f = x - f1 // 下一回合是对手
		} else {
			f = x + f1 // 下一回合还是我
		}

		s := x + nums[i+1]
		if s%3 == 0 {
			s -= f2 // 下一回合是对手
		} else {
			s += f2 // 下一回合还是我
		}
		f = max(f, s)

		f2 = f1
		f1 = f
	}

	return int64(f1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**十四、博弈 DP**」。

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
