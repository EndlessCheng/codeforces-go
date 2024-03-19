[视频讲解](https://www.bilibili.com/video/BV1cV4y157BY) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

## 方法一：二分答案

「最小化最大值」就是二分答案的代名词。

二分答案 $\textit{limit}$，那么我们可以从后往前模拟：如果 $\textit{nums}[i]>\textit{limit}$，那么应当去掉多余的 $\textit{extra}=\textit{nums}[i]-\textit{limit}$ 加到 $\textit{nums}[i-1]$ 上，最后如果 $\textit{nums}[0]\le\textit{limit}$，则二分判定成功。

代码实现时可以不用修改 $\textit{nums}$，而是维护 $\textit{extra}$ 变量。

```py [sol-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        def check(limit: int) -> bool:
            extra = 0
            for i in range(len(nums) - 1, 0, -1):
                extra = max(nums[i] + extra - limit, 0)
            return nums[0] + extra <= limit
        return bisect_left(range(max(nums)), True, key=check)
```

```java [sol-Java]
class Solution {
    public int minimizeArrayValue(int[] nums) {
        int left = -1;
        int right = 0;
        for (int x : nums) {
            right = Math.max(right, x);
        }
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, mid)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
    private boolean check(int[] nums, int limit) {
        long extra = 0;
        for (int i = nums.length - 1; i > 0; i--) {
            extra = Math.max(nums[i] + extra - limit, 0);
        }
        return nums[0] + extra <= limit;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeArrayValue(vector<int> &nums) {
        auto check = [&](int limit) -> bool {
            long long extra = 0;
            for (int i = nums.size() - 1; i > 0; i--) {
                extra = max(nums[i] + extra - limit, 0LL);
            }
            return nums[0] + extra <= limit;
        };
        // 开区间二分，原理见 https://www.bilibili.com/video/BV1AP41137w7/
        int left = -1, right = ranges::max(nums);
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func minimizeArrayValue(nums []int) int {
	return sort.Search(slices.Max(nums), func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干变量。

## 方法二：分类讨论

从 $\textit{nums}[0]$ 开始讨论：

- 如果数组中只有 $\textit{nums}[0]$，那么最大值为 $\textit{nums}[0]$。
- 再考虑 $\textit{nums}[1]$，如果 $\textit{nums}[1]\le\textit{nums}[0]$，最大值还是 $\textit{nums}[0]$；否则可以平均这两个数，平均后的最大值为平均值的上取整，即 $\left\lceil\dfrac{\textit{nums}[0]+\textit{nums}[1]}{2}\right\rceil$。
- 再考虑 $\textit{nums}[2]$，如果 $\textit{nums}[2]\le$ 前面算出的最大值，或者这三个数的平均值不超过前面算出的最大值，那么最大值不变；否则可以平均这三个数，做法同上。
- 以此类推直到最后一个数。
- 过程中的最大值为答案。

```py [sol-Python3]
class Solution:
    def minimizeArrayValue(self, nums: List[int]) -> int:
        return max((s + i) // (i + 1) for i, s in enumerate(accumulate(nums)))
```

```java [sol-Java]
class Solution {
    public int minimizeArrayValue(int[] nums) {
        long ans = 0;
        long s = 0;
        for (int i = 0; i < nums.length; i++) {
            s += nums[i];
            ans = Math.max(ans, (s + i) / (i + 1));
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeArrayValue(vector<int> &nums) {
        long long ans = 0, s = 0;
        for (int i = 0; i < nums.size(); i++) {
            s += nums[i];
            ans = max(ans, (s + i) / (i + 1));
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimizeArrayValue(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

### 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。
