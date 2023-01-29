### 提示 1

枚举 $j$ 和 $k$ 这两个**中间**的，会更容易计算。

这个技巧在去年的周赛题 [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/) 出现过。

需要计算哪些信息？

### 提示 2

需要计算：

- 在 $k$ 右侧的比 $\textit{nums}[j]$ 大的元素个数，记作 $\textit{great}[k][\textit{nums}[j]]$；
- 在 $j$ 左侧的比 $\textit{nums}[k]$ 小的元素个数，记作 $\textit{less}[j][\textit{nums}[k]]$。

对于固定的 $j$ 和 $k$，根据乘法原理，对答案的贡献为

$$
\textit{less}[j][\textit{nums}[k]]\cdot \textit{great}[k][\textit{nums}[j]]
$$

如何维护个数？

### 提示 3

维护方法如下：

- 倒序遍历 $\textit{nums}$，设 $x < \textit{nums}[k+1]$，对于 $x$，大于它的数的个数加一，即 $\textit{great}[k][x]$ 加一；
- 正序遍历 $\textit{nums}$，设 $x > \textit{nums}[j-1]$，对于 $x$，小于它的数的个数加一，即 $\textit{less}[j][x]$ 加一。

代码实现时，可以在枚举 $j$ 的同时更新 $\textit{less}$，并且只需要一个数组。

附：[视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)

```py [sol1-Python3]
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        n = len(nums)
        great = [0] * n
        great[-1] = [0] * (n + 1)
        for k in range(n - 2, 1, -1):
            great[k] = great[k + 1][:]
            for x in range(1, nums[k + 1]):
                great[k][x] += 1  # x < nums[k+1]，对于 x，大于它的数的个数 +1

        ans = 0
        less = [0] * (n + 1)
        for j in range(1, n - 1):
            for x in range(nums[j - 1] + 1, n + 1):
                less[x] += 1  # x > nums[j-1]，对于 x，小于它的数的个数 +1
            for k in range(j + 1, n - 1):
                if nums[j] > nums[k]:
                    ans += less[nums[k]] * great[k][nums[j]]
        return ans
```

```java [sol1-Java]
class Solution {
    public long countQuadruplets(int[] nums) {
        int n = nums.length;
        int[][] great = new int[n][n + 1];
        for (int k = n - 2; k >= 2; k--) {
            great[k] = great[k + 1].clone();
            for (int x = nums[k + 1] - 1; x > 0; x--)
                great[k][x]++; // x < nums[k+1]，对于 x，大于它的数的个数 +1
        }

        long ans = 0;
        int[] less = new int[n + 1];
        for (int j = 1; j < n - 2; j++) {
            for (int x = nums[j - 1] + 1; x <= n; x++)
                less[x]++; // x > nums[j-1]，对于 x，小于它的数的个数 +1
            for (int k = j + 1; k < n - 1; k++)
                if (nums[j] > nums[k])
                    ans += less[nums[k]] * great[k][nums[j]];
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
    int great[4000][4001];
public:
    long long countQuadruplets(vector<int> &nums) {
        int n = nums.size(), less[n + 1];
        for (int k = n - 2; k >= 2; k--) {
            memcpy(great[k], great[k + 1], sizeof(great[k + 1]));
            for (int x = nums[k + 1] - 1; x > 0; x--)
                great[k][x]++; // x < nums[k+1]，对于 x，大于它的数的个数 +1
        }

        long ans = 0;
        memset(less, 0, sizeof(less));
        for (int j = 1; j < n - 2; j++) {
            for (int x = nums[j - 1] + 1; x <= n; x++)
                less[x]++; // x > nums[j-1]，对于 x，小于它的数的个数 +1
            for (int k = j + 1; k < n - 1; k++)
                if (nums[j] > nums[k])
                    ans += less[nums[k]] * great[k][nums[j]];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k >= 2; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
			great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		}
	}

	less := make([]int, n+1)
	for j := 1; j < n-2; j++ {
		for x := nums[j-1] + 1; x <= n; x++ {
			less[x]++ // x > nums[j-1]，对于 x，小于它的数的个数 +1
		}
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += int64(less[nums[k]] * great[k][nums[j]])
			}
		}
	}
	return
}
```

其实 $\textit{less}$ 数组是多余的。

设 $x=\textit{nums}[k]$。在 $j$ 右边有 $n-1-j$ 个数，其中 $\textit{great}[j][x]$ 个数比 $x$ 大，由于 $\textit{nums}$ 是一个 $[1,n]$ 的排列，因此 $j$ 右边有

$$
n-1-j-\textit{great}[j][x]
$$

个不超过 $x$ 的数。

同时，由于总共有 $x$ 个不超过 $x$ 的数，所以在 $j$ 左边有

$$
x - (n-1-j-\textit{great}[j][x])
$$

个不超过 $x$ 的数。又因为 $x$ 在 $j$ 右边，所以上式亦为 $j$ 左边的小于 $x$ 的数的个数。

这样就把 $\textit{less}$ 数组优化掉了。

```py [sol2-Python3]
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        n = len(nums)
        great = [0] * n
        great[-1] = [0] * (n + 1)
        for k in range(n - 2, 0, -1):
            great[k] = great[k + 1][:]
            for x in range(1, nums[k + 1]):
                great[k][x] += 1

        ans = 0
        for j in range(1, n - 1):
            for k in range(j + 1, n - 1):
                x = nums[k]
                if nums[j] > x:
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]]
        return ans
```

```java [sol2-Java]
class Solution {
    public long countQuadruplets(int[] nums) {
        int n = nums.length;
        int[][] great = new int[n][n + 1];
        for (int k = n - 2; k > 0; k--) {
            great[k] = great[k + 1].clone();
            for (int x = nums[k + 1] - 1; x > 0; x--)
                great[k][x]++;
        }

        long ans = 0;
        for (int j = 1; j < n - 2; j++)
            for (int k = j + 1; k < n - 1; k++) {
                int x = nums[k];
                if (nums[j] > x)
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]];
            }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
    int great[4000][4001];
public:
    long long countQuadruplets(vector<int> &nums) {
        int n = nums.size();
        for (int k = n - 2; k; k--) {
            memcpy(great[k], great[k + 1], sizeof(great[k + 1]));
            for (int x = nums[k + 1] - 1; x > 0; x--)
                great[k][x]++;
        }

        long ans = 0;
        for (int j = 1; j < n - 2; j++)
            for (int k = j + 1; k < n - 1; k++)
                if (int x = nums[k]; nums[j] > x)
                    ans += (x - n + 1 + j + great[j][x]) * great[k][nums[j]];
        return ans;
    }
};
```

```go [sol2-Go]
func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	great := make([][]int, n)
	great[n-1] = make([]int, n+1)
	for k := n - 2; k > 0; k-- {
		great[k] = append([]int(nil), great[k+1]...)
		for x := nums[k+1] - 1; x > 0; x-- {
			great[k][x]++
		}
	}

	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			x := nums[k]
			if nums[j] > x {
				ans += int64((x - n + 1 + j + great[j][x]) * great[k][nums[j]])
			}
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n^2)$。
