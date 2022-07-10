下午 2 点在 B 站直播讲周赛和双周赛的题目（包含本题并查集和单调栈的原理），感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

## 方法一：并查集

#### 提示 1

数组中的元素越大越好，不妨从大往小考虑 $\textit{nums}[i]$。

#### 提示 2

子数组的长度 $k$ 越大，$\dfrac{\textit{threshold}}{k}$ 就越小，越能满足要求。

#### 提示 3

把考虑过的元素都串起来，这条链的长度就是 $k$。

于是关键在于如何动态维护每条链的长度，高效地串联两条链。

#### 提示 4

用并查集，遍历到 $\textit{nums}[i]$ 时，合并 $i$ 和 $i+1$，那么 $i$ 所处的连通块的大小减一就是 $k$。

```py [sol1-Python3]
class Solution:
    def validSubarraySize(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        fa = list(range(n + 1))
        sz = [1] * (n + 1)
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]
        for num, i in sorted(zip(nums, range(n)), reverse=True):
            j = find(i + 1)
            fa[i] = j  # 合并 i 和 i+1
            sz[j] += sz[i]
            if num > threshold // (sz[j] - 1):
                return sz[j] - 1
        return -1
```

```java [sol1-Java]
class Solution {
    int[] fa, sz;

    public int validSubarraySize(int[] nums, int threshold) {
        var n = nums.length;
        fa = new int[n + 1];
        for (var i = 0; i <= n; i++) fa[i] = i;
        sz = new int[n + 1];
        Arrays.fill(sz, 1);

        var ids = IntStream.range(0, n).boxed().toArray(Integer[]::new);
        Arrays.sort(ids, (i, j) -> nums[j] - nums[i]);
        for (var i : ids) {
            var j = find(i + 1);
            fa[i] = j; // 合并 i 和 i+1
            sz[j] += sz[i];
            if (nums[i] > threshold / (sz[j] - 1)) return sz[j] - 1;
        }
        return -1;
    }

    // 非递归并查集
    int find(int x) {
        var f = x;
        while (fa[f] != f) f = find(fa[f]);
        var tmp = x;
        while (fa[tmp] != f) {
            var t = tmp;
            tmp = fa[tmp];
            fa[t] = f;
        }
        return fa[x];
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int validSubarraySize(vector<int> &nums, int threshold) {
        int n = nums.size();
        int fa[n + 1], sz[n + 1];
        iota(fa, fa + n + 1, 0);
        for (int i = 0; i <= n; ++i) sz[i] = 1;
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };

        int ids[n];
        iota(ids, ids + n, 0);
        sort(ids, ids + n, [&](int i, int j) { return nums[i] > nums[j]; });
        for (int i : ids) {
            int j = find(i + 1);
            fa[i] = j; // 合并 i 和 i+1
            sz[j] += sz[i];
            if (nums[i] > threshold / (sz[j] - 1)) return sz[j] - 1;
        }
        return -1;
    }
};
```

```go [sol1-Go]
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i, v := range nums {
		a[i] = pair{v, i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for _, p := range a {
		i := p.i
		j := find(i + 1)
		fa[i] = j // 合并 i 和 i+1
		sz[j] += sz[i]
		if p.v > threshold/(sz[j]-1) {
			return sz[j] - 1
		}
	}
	return -1
}
```
 
#### 方法二：单调栈

#### 提示 1

枚举每个元素，假设它是子数组中的最小值。

#### 提示 2

子数组的左右边界最远能到哪？

#### 提示 3

用**单调栈**来计算左右边界。

不了解单调栈的同学可以看一下 [496. 下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)。本题求的是更小元素。

知道了左右边界也就知道了子数组的长度 $k$。

```py [sol1-Python3]
class Solution:
    def validSubarraySize(self, nums: List[int], threshold: int) -> int:
        n = len(nums)
        left, st = [-1] * n, []  # left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        for i, v in enumerate(nums):
            while st and nums[st[-1]] >= v: st.pop()
            if st: left[i] = st[-1]
            st.append(i)

        right, st = [n] * n, []  # right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        for i in range(n - 1, -1, -1):
            while st and nums[st[-1]] >= nums[i]: st.pop()
            if st: right[i] = st[-1]
            st.append(i)

        for i, num in enumerate(nums):
            k = right[i] - left[i] - 1
            if num > threshold // k:
                return k
        return -1
```

```java [sol1-Java]
class Solution {
    public int validSubarraySize(int[] nums, int threshold) {
        var n = nums.length;
        var left = new int[n]; // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        var st = new ArrayDeque<Integer>();
        for (var i = 0; i < n; i++) {
            while (!st.isEmpty() && nums[st.peek()] >= nums[i]) st.pop();
            left[i] = st.isEmpty() ? -1 : st.peek();
            st.push(i);
        }

        var right = new int[n]; // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        st = new ArrayDeque<>();
        for (var i = n - 1; i >= 0; i--) {
            while (!st.isEmpty() && nums[st.peek()] >= nums[i]) st.pop();
            right[i] = st.isEmpty() ? n : st.peek();
            st.push(i);
        }

        for (var i = 0; i < n; ++i) {
            var k = right[i] - left[i] - 1;
            if (nums[i] > threshold / k) return k;
        }
        return -1;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int validSubarraySize(vector<int> &nums, int threshold) {
        int n = nums.size();
        int left[n]; // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
        stack<int> s;
        for (int i = 0; i < n; ++i) {
            while (!s.empty() && nums[s.top()] >= nums[i]) s.pop();
            left[i] = s.empty() ? -1 : s.top();
            s.push(i);
        }

        int right[n]; // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
        s = stack<int>();
        for (int i = n - 1; i >= 0; --i) {
            while (!s.empty() && nums[s.top()] >= nums[i]) s.pop();
            right[i] = s.empty() ? n : s.top();
            s.push(i);
        }

        for (int i = 0; i < n; ++i) {
            int k = right[i] - left[i] - 1;
            if (nums[i] > threshold / k) return k;
        }
        return -1;
    }
};
```

```go [sol1-Go]
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	left := make([]int, n) // left[i] 为左侧小于 nums[i] 的最近元素位置（不存在时为 -1）
	st := []int{}
	for i, v := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n) // right[i] 为右侧小于 nums[i] 的最近元素位置（不存在时为 n）
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] >= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i, num := range nums {
		k := right[i] - left[i] - 1
		if num > threshold/k {
			return k
		}
	}
	return -1
}
```

