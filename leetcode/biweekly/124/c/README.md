由于只能从两端删除，所以子问题是「从两侧向内缩小的」，使用区间 DP 解决。

这个套路我在 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)中讲过，欢迎收看。

如果确定了第一次删除的元素和，那么后续删除的元素和也就确定了，所以至多有三种不同的元素和，分别用三次区间 DP 解决。

## 方法一：记忆化搜索

定义 $\textit{dfs}(i,j)$ 表示当前剩余元素从 $\textit{nums}[i]$ 到 $\textit{nums}[j]$，此时最多可以进行的操作次数。

枚举三种操作方式，分别从 $\textit{dfs}(i+2,j)+1, \textit{dfs}(i,j-2)+1, \textit{dfs}(i+1,j-1)+1$ 转移过来，取最大值，即为 $\textit{dfs}(i,j)$。如果三种操作方式都不行，那么 $\textit{dfs}(i,j)=0$。

递归终点：如果 $i\ge j$，此时至多剩下一个数，无法操作，返回 $0$。

递归入口：根据三种初始操作，分别为 $\textit{dfs}(2,n-1), \textit{dfs}(0,n-3), \textit{dfs}(1,n-2)$。三者取最大值再加一（加上第一次操作），即为答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Sm411U7cR/) 第三题。

```py [sol-Python3]
class Solution:
    def maxOperations(self, nums: List[int]) -> int:
        @cache
        def dfs(i: int, j: int, target: int) -> int:
            if i >= j:
                return 0
            res = 0
            if nums[i] + nums[i + 1] == target:  # 最前面两个
                res = max(res, dfs(i + 2, j, target) + 1)
            if nums[j - 1] + nums[j] == target:  # 最后两个
                res = max(res, dfs(i, j - 2, target) + 1)
            if nums[i] + nums[j] == target:  # 第一个和最后一个
                res = max(res, dfs(i + 1, j - 1, target) + 1)
            return res

        n = len(nums)
        res1 = dfs(2, n - 1, nums[0] + nums[1])  # 最前面两个
        res2 = dfs(0, n - 3, nums[-2] + nums[-1])  # 最后两个
        res3 = dfs(1, n - 2, nums[0] + nums[-1])  # 第一个和最后一个
        return max(res1, res2, res3) + 1  # 加上第一次操作
```

```java [sol-Java]
class Solution {
    private int[] nums;
    private int[][] memo;

    public int maxOperations(int[] nums) {
        int n = nums.length;
        this.nums = nums;
        memo = new int[n][n];
        int res1 = helper(2, n - 1, nums[0] + nums[1]); // 最前面两个
        int res2 = helper(0, n - 3, nums[n - 2] + nums[n - 1]); // 最后两个
        int res3 = helper(1, n - 2, nums[0] + nums[n - 1]); // 第一个和最后一个
        return Math.max(Math.max(res1, res2), res3) + 1; // 加上第一次操作
    }

    private int helper(int i, int j, int target) {
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }
        return dfs(i, j, target);
    }

    private int dfs(int i, int j, int target) {
        if (i >= j) {
            return 0;
        }
        if (memo[i][j] != -1) {
            return memo[i][j];
        }
        int res = 0;
        if (nums[i] + nums[i + 1] == target) {
            res = Math.max(res, dfs(i + 2, j, target) + 1);
        }
        if (nums[j - 1] + nums[j] == target) {
            res = Math.max(res, dfs(i, j - 2, target) + 1);
        }
        if (nums[i] + nums[j] == target) {
            res = Math.max(res, dfs(i + 1, j - 1, target) + 1);
        }
        return memo[i][j] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxOperations(vector<int> &nums) {
        int n = nums.size();
        int memo[n][n];
        auto helper = [&](int i, int j, int target) -> int {
            memset(memo, -1, sizeof(memo));
            function<int(int, int)> dfs = [&](int i, int j) -> int {
                if (i >= j) return 0;
                int &res = memo[i][j]; // 注意这里是引用
                if (res != -1) return res;
                res = 0;
                if (nums[i] + nums[i + 1] == target) res = max(res, dfs(i + 2, j) + 1);
                if (nums[j - 1] + nums[j] == target) res = max(res, dfs(i, j - 2) + 1);
                if (nums[i] + nums[j] == target) res = max(res, dfs(i + 1, j - 1) + 1);
                return res;
            };
            return dfs(i, j);
        };
        int res1 = helper(2, n - 1, nums[0] + nums[1]); // 最前面两个
        int res2 = helper(0, n - 3, nums[n - 2] + nums[n - 1]); // 最后两个
        int res3 = helper(1, n - 2, nums[0] + nums[n - 1]); // 第一个和最后一个
        return max({res1, res2, res3}) + 1; // 加上第一次操作
    }
};
```

```go [sol-Go]
func maxOperations(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:], nums[0]+nums[1])       // 最前面两个
	res2 := helper(nums[:n-2], nums[n-2]+nums[n-1]) // 最后两个
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 第一个和最后一个
	return max(res1, res2, res3) + 1                // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if a[i]+a[i+1] == target { // 最前面两个
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 最后两个
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 第一个和最后一个
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res
		return
	}
	return dfs(0, n-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：1:1 翻译成递推

注：由于递推会计算大量无效状态，效率是不如记忆化搜索的。可能记忆化搜索中算到每个状态发现无法往下递归（三种操作方法都不行），但是递推会把所有的状态都算出来。

```py [sol-Python3]
class Solution:
    def maxOperations(self, nums: List[int]) -> int:
        n = len(nums)

        def helper(start: int, end: int, target: int) -> int:
            f = [[0] * (n + 1) for _ in range(n + 1)]
            for i in range(end, start - 1, -1):
                for j in range(i + 1, end + 1):
                    if nums[i] + nums[i + 1] == target:  # 最前面两个
                        f[i][j] = max(f[i][j], f[i + 2][j] + 1)
                    if nums[j - 1] + nums[j] == target:  # 最后两个
                        f[i][j] = max(f[i][j], f[i][j - 2] + 1)
                    if nums[i] + nums[j] == target:  # 第一个和最后一个
                        f[i][j] = max(f[i][j], f[i + 1][j - 1] + 1)
            return f[start][end]

        res1 = helper(2, n - 1, nums[0] + nums[1])  # 最前面两个
        res2 = helper(0, n - 3, nums[-2] + nums[-1])  # 最后两个
        res3 = helper(1, n - 2, nums[0] + nums[-1])  # 第一个和最后一个
        return max(res1, res2, res3) + 1  # 加上第一次操作
```

```java [sol-Java]

```

```cpp [sol-C++]

```

```go [sol-Go]

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
