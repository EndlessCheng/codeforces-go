[本题视频讲解](https://www.bilibili.com/video/BV1Rw411P72r/)

## 错误思路：贪心

错误思路：从左到右合并数字，只要不比前面小就停止合并。如果最后一个数小于倒数第二个数，就和倒数第二个数合并。

反例：$[1,2,1,3,3]$。

错误的合并方式：$[1,2,7]$。

正确的合并方式：$[1,3,3,3]$。

## 正确思路：数据结构优化 DP

定义 $f[i]$ 表示操作 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 所得到的最长非递减数组的长度。

同时记录 $\textit{last}[i]$ 表示在 $f[i]$ 尽量大的前提下，$\textit{nums}[0]$ 到 $\textit{nums}[i]$ 操作后的最后一个数的最小值。为什么要记录最小值？因为这个数越小，越有利于转移。

如果 $\textit{nums}[j+1] + \cdots + \textit{nums}[i] \ge \textit{last}[j]$，那么就可以从 $f[j]$ 转移过来，也就是

$$
f[i] = \max_{j} f[j] + 1
$$

答案为 $f[n-1]$。

> 注意：代码中的下标是从 $1$ 开始的。

以 $[6,5,1,9]$ 为例（合并后为 $[6,6,9]$），列出 $-1$ 是为了方便看出转移。

|  $i$ | $\textit{nums}[i]$ | $f[i]$  | $\textit{last}[i]$  |
|---|---|---|---|
| $-1$ |  -   | $0$  |  $0$ |
|  $0$ |  $6$ | $1$  |  $6$ |
|  $1$ |  $5$ | $1$  | $11$ |
|  $2$ |  $1$ | $2$  |  $6$ |
|  $3$ |  $9$ | $3$  |  $9$ |

由于总是可以把 $\textit{nums}[i]$ 合并到前一个数中，所以 $f[i]$ 不会小于 $f[i-1]$，所以 $f$ 数组是非递减的。

对于 $\textit{last}$ 数组来说，就不一定是非递减的了，这也是为什么某些其它 DP 写法是错误的（例如滑窗优化 DP）。

一个「粗暴」的思路是用前缀和+权值线段树（或者树状数组）来优化转移方程。定义

$$
s[i] = \textit{nums}[0] + \cdots + \textit{nums}[i]
$$

那么 

$$
\textit{nums}[j+1] + \cdots + \textit{nums}[i]  = s[i] - s[j] \ge \textit{last}[j]
$$

变形得

$$
s[i]\ge s[j] + \textit{last}[j]
$$

所以用权值线段树维护 $s[j] + \textit{last}[j]$ 对应的 $f[j]$ 的最大值，就可以用「查询 $\le s[i]$ 中的最大值」来快速计算转移了。

在满足 $s[i]\ge s[j] + \textit{last}[j]$ 的前提下，$j$ 越大，$\textit{last}[i]$ 就越小，所以还需要额外获取到最大值对应的最大下标 $j$。

这样写有点麻烦，有没有更简单的做法呢？

考虑两个转移来源 $j$ 和 $k$，如果 $j<k$ 且 $s[j] + \textit{last}[j] \ge s[k] + \textit{last}[k]$，这意味着：

- 如果能从 $f[j]$ 转移到 $f[i]$，那么也一定能从 $f[k]$ 转移到 $f[i]$。

又由于 $f[j]\le f[k]$，所以**永远不需要**从 $f[j]$ 转移到 $f[i]$。

所以可以用单调队列来维护 $j$，满足从队首到队尾的 $j$ 和 $s[j] + \textit{last}[j]$ 都是严格递增的。

原理请看 [单调队列【基础算法精讲 27】](https://www.bilibili.com/video/BV1bM411X72E/)

单调队列需要思考清楚三步：

1. **转移之前，去掉队首无用数据**：由于 $i$ 越大 $s[i]$ 越大，满足 $s[j] + \textit{last}[j]\le s[i]$ 的 $j$ 也越大，转移来源 $f[j]$ 也越大，所以在转移来源 $j$ 左边的下标，可以从队首出队。
2. **计算转移**：从单调队列中找到最大的 $j$，满足 $s[j] + \textit{last}[j]\le s[i]$，并据此算出 $f[i]=f[j]+1$ 和 $\textit{last}[i] = s[i] - s[j]$。去掉队首无用数据之后，直接取队首即可。
3. **去掉队尾无用数据**：把 $i$ 加入队尾，在此之前，弹出 $s[j] + \textit{last}[j] \ge s[i] + \textit{last}[i]$ 的 $j$。

```py [sol-Python3]
class Solution:
    def findMaximumLength(self, nums: List[int]) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))
        f = [0] * (n + 1)
        last = [0] * (n + 1)
        q = deque([0])
        for i in range(1, n + 1):
            # 1. 去掉队首无用数据（计算转移时，直接取队首）
            while len(q) > 1 and s[q[1]] + last[q[1]] <= s[i]:
                q.popleft()
            
            # 2. 计算转移
            f[i] = f[q[0]] + 1
            last[i] = s[i] - s[q[0]]
            
            # 3. 去掉队尾无用数据
            while q and s[q[-1]] + last[q[-1]] >= s[i] + last[i]:
                q.pop()
            q.append(i)
        return f[n]
```

```java [sol-Java]
class Solution {
    public int findMaximumLength(int[] nums) {
        int n = nums.length;
        long[] s = new long[n + 1];
        int[] f = new int[n + 1];
        long[] last = new long[n + 1];
        int[] q = new int[n + 1]; // 数组模拟队列
        int front = 0, rear = 0;
        for (int i = 1; i <= n; i++) {
            s[i] = s[i - 1] + nums[i - 1];
            
            // 1. 去掉队首无用数据（计算转移时，直接取队首）
            while (front < rear && s[q[front + 1]] + last[q[front + 1]] <= s[i]) {
                front++;
            }
            
            // 2. 计算转移
            f[i] = f[q[front]] + 1; 
            last[i] = s[i] - s[q[front]];
            
            // 3. 去掉队尾无用数据
            while (rear >= front && s[q[rear]] + last[q[rear]] >= s[i] + last[i]) {
                rear--;
            }
            q[++rear] = i;
        }
        return f[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findMaximumLength(vector<int> &nums) {
        int n = nums.size();
        vector<long long> s(n + 1), last(n + 1);
        vector<int> f(n + 1), q(n + 1); // 数组模拟队列
        int front = 0, rear = 0;
        for (int i = 1; i <= n; i++) {
            s[i] = s[i - 1] + nums[i - 1];

            // 1. 去掉队首无用数据（计算转移时，直接取队首）
            while (front < rear && s[q[front + 1]] + last[q[front + 1]] <= s[i]) {
                front++;
            }

            // 2. 计算转移
            f[i] = f[q[front]] + 1;
            last[i] = s[i] - s[q[front]];

            // 3. 去掉队尾无用数据
            while (rear >= front && s[q[rear]] + last[q[rear]] >= s[i] + last[i]) {
                rear--;
            }
            q[++rear] = i;
        }
        return f[n];
    }
};
```

```go [sol-Go]
func findMaximumLength(nums []int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	f := make([]int, n+1)
	last := make([]int, n+1)
	q := []int{0}
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]

		// 1. 去掉队首无用数据（计算转移时，直接取队首）
		for len(q) > 1 && s[q[1]]+last[q[1]] <= s[i] {
			q = q[1:]
		}

		// 2. 计算转移
		f[i] = f[q[0]] + 1
		last[i] = s[i] - s[q[0]]

		// 3. 去掉队尾无用数据
		for len(q) > 0 && s[q[len(q)-1]]+last[q[len(q)-1]] >= s[i]+last[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。每个下标至多入队出队各一次。
- 空间复杂度：$\mathcal{O}(n)$。
