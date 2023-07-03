## 方法一：枚举

由于 $n$ 至多为 $1000$，我们可以从左到右枚举子数组左端点 $i$，然后从 $i+1$ 开始向右枚举子数组右端点 $j$。一边枚举 $j$，一边维护不平衡度 $\textit{cnt}$：

- 如果 $x=\textit{nums}[j]$ 之前出现过，那么子数组排序后必然会和另一个 $x$ 相邻，$\textit{cnt}$ 不变；
- 如果 $x=\textit{nums}[j]$ 之前没出现过，那么看 $x-1$ 和 $x+1$ 是否出现过：
  - 都没有，$\textit{cnt}$ 加一；
  - 只有一个，$\textit{cnt}$ 不变；
  - 两个都有，$\textit{cnt}$ 减一。

遍历过程中，累加 $\textit{cnt}$，即为答案。

```py [sol-Python3]
class Solution:
    def sumImbalanceNumbers(self, nums: List[int]) -> int:
        ans, n = 0, len(nums)
        for i, x in enumerate(nums):
            vis = [False] * (n + 2)
            vis[x] = True
            cnt = 0
            for j in range(i + 1, n):
                x = nums[j]
                if not vis[x]:
                    cnt += 1 - vis[x - 1] - vis[x + 1]
                    vis[x] = True
                ans += cnt
        return ans
```

```java [sol-Java]
class Solution {
    public int sumImbalanceNumbers(int[] nums) {
        int ans = 0, n = nums.length;
        var vis = new boolean[n + 2];
        for (int i = 0; i < n; i++) {
            Arrays.fill(vis, false);
            vis[nums[i]] = true;
            int cnt = 0;
            for (int j = i + 1; j < n; j++) {
                int x = nums[j];
                if (!vis[x]) {
                    cnt++;
                    if (vis[x - 1]) cnt--;
                    if (vis[x + 1]) cnt--;
                    vis[x] = true;
                }
                ans += cnt;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumImbalanceNumbers(vector<int> &nums) {
        int ans = 0, n = nums.size();
        bool vis[n + 2];
        for (int i = 0; i < n; i++) {
            memset(vis, 0, sizeof(vis));
            vis[nums[i]] = true;
            int cnt = 0;
            for (int j = i + 1; j < n; j++) {
                int x = nums[j];
                if (!vis[x]) {
                    cnt += 1 - vis[x - 1] - vis[x + 1];
                    vis[x] = true;
                }
                ans += cnt;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumImbalanceNumbers(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums {
		vis := make([]int, n+2)
		vis[x] = 1
		cnt := 0
		for j := i + 1; j < n; j++ {
			if x := nums[j]; vis[x] == 0 {
				cnt += 1 - vis[x-1] - vis[x+1]
				vis[x] = 1
			}
			ans += cnt
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：贡献法

视频讲解见[【周赛 352】](https://www.bilibili.com/video/BV1ej411m7zV/)第四题。

考察 $x=\textit{nums}[i]$ 能为答案贡献多少不平衡数字。换句话说，有多少个包含 $x$ 的子数组，

为了避免重复统计，不妨规定只考察 $x$ 和比 $x$ 小的最大的数，它们之间是否会产生贡献。

> 如果还考察 $x$ 和比 $x$ 大的最小的数 $y$，那么在考察 $y$ 的时候，$y$ 和 $x$ 之间又统计了一遍。

有以下几种情况：

1. $x$ 是子数组的最小值，这种情况不能计入贡献。
2. 如果子数组中有 $x-1$，这种情况不能计入贡献。
3. 如果子数组中没有 $x-1$ 和 $x$，那么 $x$ 会产生一个贡献。
4. 如果子数组中没有 $x-1$，但是有 $x$ 呢？比如某个子数组为 $[1,3,3]$，第一个 $3$ 会和 $1$ 算一次贡献，第二个 $3$ 如果又和 $1$ 又算了一次贡献，那么就重复统计了。为了避免重复统计，下面规定每个 $3$ 向右计算子数组范围时，只


```py [sol-Python3]
class Solution:
    def sumImbalanceNumbers(self, nums: List[int]) -> int:
        n = len(nums)
        right = [0] * n  # nums[i] 右侧的 x 和 x-1 的最近下标（不存在时为 n）
        idx = [n] * (n + 1)
        for i in range(n - 1, -1, -1):
            x = nums[i]
            right[i] = min(idx[x], idx[x - 1])
            idx[x] = i

        ans = 0
        idx = [-1] * (n + 1)
        for i, (x, r) in enumerate(zip(nums, right)):
            # 统计 x 能产生多少贡献
            ans += (i - idx[x - 1]) * (r - i)  # 子数组左端点个数 * 子数组右端点个数
            idx[x] = i
        # 上面计算的时候，每个子数组的最小值必然可以作为贡献，而这是不合法的
        # 所以每个子数组都多算了 1 个不合法的贡献
        return ans - n * (n + 1) // 2
```

```java [sol-Java]
class Solution {
    public int sumImbalanceNumbers(int[] nums) {
        int n = nums.length;
        var right = new int[n];
        var idx = new int[n + 1];
        Arrays.fill(idx, n);
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            // right[i] 表示 nums[i] 右侧的 x 和 x-1 的最近下标（不存在时为 n）
            right[i] = Math.min(idx[x], idx[x - 1]);
            idx[x] = i;
        }

        int ans = 0;
        Arrays.fill(idx, -1);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 统计 x 能产生多少贡献
            ans += (i - idx[x - 1]) * (right[i] - i); // 子数组左端点个数 * 子数组右端点个数
            idx[x] = i;
        }
        // 上面计算的时候，每个子数组的最小值必然可以作为贡献，而这是不合法的
        // 所以每个子数组都多算了 1 个不合法的贡献
        return ans - n * (n + 1) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumImbalanceNumbers(vector<int> &nums) {
        int n = nums.size(), right[n], idx[n + 1];
        fill(idx, idx + n + 1, n);
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            // right[i] 表示 nums[i] 右侧的 x 和 x-1 的最近下标（不存在时为 n）
            right[i] = min(idx[x], idx[x - 1]);
            idx[x] = i;
        }

        int ans = 0;
        memset(idx, -1, sizeof(idx));
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            // 统计 x 能产生多少贡献
            ans += (i - idx[x - 1]) * (right[i] - i); // 子数组左端点个数 * 子数组右端点个数
            idx[x] = i;
        }
        // 上面计算的时候，每个子数组的最小值必然可以作为贡献，而这是不合法的
        // 所以每个子数组都多算了 1 个不合法的贡献
        return ans - n * (n + 1) / 2;
    }
};
```

```go [sol-Go]
func sumImbalanceNumbers(nums []int) (ans int) {
	n := len(nums)
	right := make([]int, n)
	idx := make([]int, n+1)
	for i := range idx {
		idx[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		// right[i] 表示 nums[i] 右侧的 x 和 x-1 的最近下标（不存在时为 n）
		right[i] = min(idx[x], idx[x-1])
		idx[x] = i
	}

	for i := range idx {
		idx[i] = -1
	}
	for i, x := range nums {
		// 统计 x 能产生多少贡献
		ans += (i - idx[x-1]) * (right[i] - i) // 子数组左端点个数 * 子数组右端点个数
		idx[x] = i
	}
	// 上面计算的时候，每个子数组的最小值必然可以作为贡献，而这是不合法的
	// 所以每个子数组都多算了 1 个不合法的贡献
	return ans - n*(n+1)/2
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

把 `sarr[i+1] - sarr[i] > 1` 改成 `sarr[i+1] - sarr[i] > k` 要怎么做？欢迎在评论区发表你的思路。
