## 本题视频讲解

见[【周赛 340】](https://www.bilibili.com/video/BV1iN411w7my/)。

## 方法一：相同元素分组+前缀和+二分查找

按照相同元素分组后再计算。

这里的思路和 [2602. 使数组元素全部相等的最少操作次数（题解）](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)是一样的。由于目标位置就是数组中的下标，无需二分。

```py [sol1-Python3]
class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        groups = defaultdict(list)
        for i, x in enumerate(nums):
            groups[x].append(i)  # 相同元素分到同一组，记录下标
        ans = [0] * len(nums)
        for a in groups.values():
            n = len(a)
            s = list(accumulate(a, initial=0))  # 前缀和
            for j, target in enumerate(a):
                left = target * j - s[j]  # 蓝色面积
                right = s[n] - s[j] - target * (n - j)  # 绿色面积
                ans[target] = left + right
        return ans
```

```java [sol1-Java]
class Solution {
    public long[] distance(int[] nums) {
        int n = nums.length;
        var groups = new HashMap<Integer, List<Integer>>();
        for (int i = 0; i < n; ++i) // 相同元素分到同一组，记录下标
            groups.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);

        var ans = new long[n];
        var s = new long[n + 1];
        for (var a : groups.values()) {
            int m = a.size();
            for (int i = 0; i < m; ++i)
                s[i + 1] = s[i] + a.get(i); // 前缀和
            for (int i = 0; i < m; ++i) {
                int target = a.get(i);
                long left = (long) target * i - s[i]; // 蓝色面积
                long right = s[m] - s[i] - (long) target * (m - i); // 绿色面积
                ans[target] = left + right;
            }
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<long long> distance(vector<int> &nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> groups;
        for (int i = 0; i < n; ++i)
            groups[nums[i]].push_back(i); // 相同元素分到同一组，记录下标

        vector<long long> ans(n);
        long long s[n + 1];
        s[0] = 0;
        for (auto &[_, a]: groups) {
            int m = a.size();
            for (int i = 0; i < m; ++i)
                s[i + 1] = s[i] + a[i]; // 前缀和
            for (int i = 0; i < m; ++i) {
                long long target = a[i];
                long long left = target * i - s[i]; // 蓝色面积
                long long right = s[m] - s[i] - target * (m - i); // 绿色面积
                ans[target] = left + right;
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := make([]int, n+1)
		for i, x := range a {
			s[i+1] = s[i] + x // 前缀和
		}
		for i, target := range a {
			left := target*i - s[i] // 蓝色面积
			right := s[n] - s[i] - target*(n-i) // 绿色面积
			ans[target] = int64(left + right)
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

## 方法二：相同元素分组+考虑增量

分组后，对于其中一个组 $a$，我们先暴力计算出 $a[0]$ 到其它元素的距离之和，设为 $s$。

然后计算 $a[1]$，我们不再暴力计算，而是思考：$s$ 增加了多少？（增量可以是负数）

设 $n$ 为 $a$ 的长度，从 $a[0]$ 到 $a[1]$，有 $1$ 个数的距离变大了 $a[1]-a[0]$，有 $n-1$ 个数的距离变小了 $a[1]-a[0]$。

所以对于 $a[1]$，我们可以 $O(1)$ 地知道它到其它元素的距离之和为

$$
s + (2-n) \cdot (a[1]-a[0])
$$

一般地，设 $a[i-1]$ 到其它元素的距离之和为 $s$，那么 $a[i]$ 到其它元素的距离之和为

$$
s + (2i-n) \cdot (a[i]-a[i-1])
$$

```py [sol2-Python3]
class Solution:
    def distance(self, nums: List[int]) -> List[int]:
        groups = defaultdict(list)
        for i, x in enumerate(nums):
            groups[x].append(i)  # 相同元素分到同一组，记录下标
        ans = [0] * len(nums)
        for a in groups.values():
            n = len(a)
            s = sum(x - a[0] for x in a)  # a[0] 到其它下标的距离之和
            ans[a[0]] = s
            for i in range(1, n):
                # 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
                s += (i * 2 - n) * (a[i] - a[i - 1])
                ans[a[i]] = s
        return ans
```

```java [sol2-Java]
class Solution {
    public long[] distance(int[] nums) {
        int n = nums.length;
        var groups = new HashMap<Integer, List<Integer>>();
        for (int i = 0; i < n; ++i) // 相同元素分到同一组，记录下标
            groups.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);

        var ans = new long[n];
        for (var a : groups.values()) {
            int m = a.size();
            long s = 0;
            for (int x : a)
                s += x - a.get(0); // a[0] 到其它下标的距离之和
            ans[a.get(0)] = s;
            for (int i = 1; i < m; ++i)
                // 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
                ans[a.get(i)] = s += (long) (i * 2 - m) * (a.get(i) - a.get(i - 1));
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    vector<long long> distance(vector<int> &nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> groups;
        for (int i = 0; i < n; ++i)
            groups[nums[i]].push_back(i); // 相同元素分到同一组，记录下标

        vector<long long> ans(n);
        for (auto &[_, a]: groups) {
            int m = a.size();
            long long s = 0;
            for (int x: a)
                s += x - a[0]; // a[0] 到其它下标的距离之和
            ans[a[0]] = s;
            for (int i = 1; i < m; ++i)
                // 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
                ans[a[i]] = s += (long long) (i * 2 - m) * (a[i] - a[i - 1]);
        }
        return ans;
    }
};
```

```go [sol2-Go]
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := int64(0)
		for _, x := range a {
			s += int64(x - a[0]) // a[0] 到其它下标的距离之和
		}
		ans[a[0]] = s
		for i := 1; i < n; i++ {
			// 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
			s += int64(i*2-n) * int64(a[i]-a[i-1])
			ans[a[i]] = s
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
