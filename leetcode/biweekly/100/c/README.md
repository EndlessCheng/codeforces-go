### 本题视频讲解

见[【双周赛 100】](https://www.bilibili.com/video/BV1WM411H7UE/)。

# 方法一：带着下标去排序

把 $\textit{nums}[i]$ 及其下标绑定后，按照元素值从小到大排序，元素值相同的按照下标排序。

然后按照题目模拟，用一个 $\textit{vis}$ 数组来实现标记。

也可以生成一个下标数组，对下标排序。具体见 Java 和 C++ 的实现。

```py [sol1-Python3]
class Solution:
    def findScore(self, nums: List[int]) -> int:
        ans = 0
        vis = [False] * (len(nums) + 2)  # 保证下标不越界
        for i, x in sorted(enumerate(nums, 1), key=lambda p: p[1]):
            if not vis[i]:
                vis[i - 1] = vis[i + 1] = True  # 标记相邻的两个元素
                ans += x
        return ans
```

```java [sol1-Java]
class Solution {
    public long findScore(int[] nums) {
        int n = nums.length;
        var ids = new Integer[n];
        for (int i = 0; i < n; ++i) ids[i] = i;
        Arrays.sort(ids, (i, j) -> nums[i] - nums[j]);

        long ans = 0;
        var vis = new boolean[n + 2]; // 保证下标不越界
        for (int i : ids)
            if (!vis[i + 1]) { // 避免 -1，偏移一位
                vis[i] = vis[i + 2] = true;
                ans += nums[i];
            }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long findScore(vector<int> &nums) {
        int n = nums.size(), ids[n];
        iota(ids, ids + n, 0);
        stable_sort(ids, ids + n, [&](int i, int j) {
            return nums[i] < nums[j];
        });

        long long ans = 0;
        bool vis[n + 2]; // 保证下标不越界
        memset(vis, 0, sizeof(vis));
        for (int i : ids)
            if (!vis[i + 1]) { // 避免 -1，偏移一位
                vis[i] = vis[i + 2] = true;
                ans += nums[i];
            }
        return ans;
    }
};
```

```go [sol1-Go]
func findScore(nums []int) (ans int64) {
	type pair struct{ v, i int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, i + 1} // +1 保证下面 for 循环下标不越界
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.v < b.v || a.v == b.v && a.i < b.i
	})
	vis := make([]bool, len(nums)+2) // 保证下标不越界
	for _, p := range a {
		if !vis[p.i] {
			vis[p.i-1] = true
			vis[p.i+1] = true // 标记相邻的两个元素
			ans += int64(p.v)
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

# 方法二：转换 + 分组循环

把 $\textit{nums}$ 视作由若干严格递减子段组成的数组。

例如示例 1 可以看成 $[2,1]+[3]+[4]+[5,2]$，示例 2 可以看成 $[2]+[3]+[5,1]+[3,2]$。

从左到右遍历 $\textit{nums}$，严格递减子段的最小值 $\textit{nums}[i]$ 一定可以选，因为它比 $\textit{nums}[i-1]$ 小，且不会超过 $\textit{nums}[i+1]$。如果等于 $\textit{nums}[i+1]$，由于我们是从左到右遍历的，下标也是最小的。

$\textit{nums}[i]$ 选了，那么这一段左侧的 $\textit{nums}[i-2],\textit{nums}[i-4],\cdots$ 就可以一起选了（因为不能选相邻的），且 $\textit{nums}[i+1]$ 不能选。

于是遍历 $\textit{nums}$ 就可以算出答案了。

```py [sol2-Python3]
class Solution:
    def findScore(self, nums: List[int]) -> int:
        ans = 0
        i, n = 0, len(nums)
        while i < n:
            i0 = i
            while i + 1 < n and nums[i] > nums[i + 1]:  # 找到下坡的坡底
                i += 1
            for j in range(i, i0 - 1, -2):  # 从坡底 i 到坡顶 i0，每隔一个累加
                ans += nums[j]
            i += 2  # i 选了 i+1 不能选
        return ans
```

```java [sol2-Java]
class Solution {
    public long findScore(int[] nums) {
        long ans = 0;
        for (int i = 0, n = nums.length; i < n; i += 2) { // i 选了 i+1 不能选
            int i0 = i;
            while (i + 1 < n && nums[i] > nums[i + 1]) // 找到下坡的坡底
                ++i;
            for (int j = i; j >= i0; j -= 2) // 从坡底 i 到坡顶 i0，每隔一个累加
                ans += nums[j];
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long findScore(vector<int> &nums) {
        long long ans = 0;
        for (int i = 0, n = nums.size(); i < n; i += 2) { // i 选了 i+1 不能选
            int i0 = i;
            while (i + 1 < n && nums[i] > nums[i + 1]) // 找到下坡的坡底
                ++i;
            for (int j = i; j >= i0; j -= 2) // 从坡底 i 到坡顶 i0，每隔一个累加
                ans += nums[j];
        }
        return ans;
    }
};
```

```go [sol2-Go]
func findScore(nums []int) (ans int64) {
	for i, n := 0, len(nums); i < n; i += 2 { // i 选了 i+1 不能选
		i0 := i
		for i+1 < n && nums[i] > nums[i+1] { // 找到下坡的坡底
			i++
		}
		for j := i; j >= i0; j -= 2 { // 从坡底 i 到坡顶 i0，每隔一个累加
			ans += int64(nums[j])
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。注意代码中的 $i$ 只增不减，所以整个二重循环是 $O(n)$ 的。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
