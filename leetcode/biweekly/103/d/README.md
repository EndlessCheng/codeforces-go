### 本题视频讲解

见[【力扣双周赛 103】](https://www.bilibili.com/video/BV1Ez4y1Y7Az/)第四题。

## 方法一：树状数组

直接模拟是 $\mathcal{O}(n^2)$ 的（最坏情况发生在降序数组），如何优化呢？

想象成用一个下标**反复遍历**数组，考虑下标的总共移动次数。

以 $\textit{nums}=[2,4,1,3]$ 为例：

- 初始时，下标指向第一个数字 $2$。
- 移动 $2$ 次到 $1$，删除 $1$，此时下标指向 $3$。
- 移动 $1$ 次到 $2$，删除 $2$，此时下标指向 $4$。
- 移动 $1$ 次到 $3$（注意 $1$ 已经被删除了），删除 $3$，此时下标指向 $4$（注意 $2$ 已经被删除了）。
- 无需移动，删除 $4$。
- 总共移动 $2+1+1=4$ 次，加上删除操作 $4$ 次，故答案为 $8$。

上述过程有两个要点：

1. 按照元素值从小到大的顺序删除。
2. 需要跳过已经删除的元素。换句话说，需要知道移动到下一个位置的途中，有多少元素被删除了。

由于要按照元素值从小到大的顺序删除，需要对数组排序。但是数组元素的顺序很重要，不能直接排序。可以创建一个下标数组，对下标数组按照 $\textit{nums}[i]$ 的大小，从小到大排序。

由于要全部删除，先把数组长度 $n$ 计入答案，这样后面只需要统计移动次数。

为方便计算，后面讨论的下标从 $1$ 开始。

设上一个被删除的数的位置为 $\textit{pre}$（初始为 $1$），当前需要删除的位置为 $i$。定义 $\text{query}(i,j)$ 表示下标在 $[i,j]$ 中的被删除的元素个数。考虑从 $\textit{pre}$ 到 $i$ 的移动次数：

- 如果 $\textit{pre}\le i$，那么移动次数等于 $i-\textit{pre}-\text{query}(\textit{pre},i)$。
- 如果 $\textit{pre}> i$，那么移动次数等于 $(n-\textit{pre}-\text{query}(\textit{pre},n))+(i-\text{query}(1, i))$。假设删除了 $k$ 个数，那么 $\text{query}(\textit{pre},n))+\text{query}(1,i) = k - \text{query}(i+1,\textit{pre}-1)=k - \text{query}(i,\textit{pre}-1)$，因为 $i$ 处的元素还没被删除。所以上式最终化简为 $i-\textit{pre}+n-k+\text{query}(i,\textit{pre}-1)$。

代码实现时，$\text{query}(i,j)$ 可以用树状数组实现。（也可以用线段树、名次树等数据结构。）

```py [sol1-Python3]
class Solution:
    def countOperationsToEmptyArray(self, nums: List[int]) -> int:
        ans = n = len(nums)  # 先把 n 计入答案
        t = BIT(n + 1)  # 下标从 1 开始
        pre = 1  # 上一个最小值的位置，初始为 1
        for k, i in enumerate(sorted(range(n), key=lambda i: nums[i])):
            i += 1  # 下标从 1 开始
            if i >= pre:  # 从 pre 移动到 i，跳过已经删除的数
                ans += i - pre - t.query(pre, i)
            else:  # 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
                ans += i - pre + n - k + t.query(i, pre - 1)
            t.inc(i)  # 删除 i
            pre = i
        return ans

# 树状数组模板
class BIT:
    def __init__(self, n):
        self.tree = [0] * n

    # 将下标 i 上的数加一
    def inc(self, i: int) -> None:
        while i < len(self.tree):
            self.tree[i] += 1
            i += i & -i

    # 返回闭区间 [1, i] 的元素和
    def sum(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

    # 返回闭区间 [left, right] 的元素和
    def query(self, left: int, right: int) -> int:
        return self.sum(right) - self.sum(left - 1)
```

```java [sol1-Java]
class Solution {
    public long countOperationsToEmptyArray(int[] nums) {
        int n = nums.length;
        var id = new Integer[n];
        for (int i = 0; i < n; ++i)
            id[i] = i;
        Arrays.sort(id, (i, j) -> nums[i] - nums[j]);

        long ans = n; // 先把 n 计入答案
        var t = new BIT(n + 1); // 下标从 1 开始
        int pre = 1; // 上一个最小值的位置，初始为 1
        for (int k = 0; k < n; ++k) {
            int i = id[k] + 1; // 下标从 1 开始
            if (i >= pre) // 从 pre 移动到 i，跳过已经删除的数
                ans += i - pre - t.query(pre, i);
            else // 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
                ans += n - pre + i - k + t.query(i, pre - 1);
            t.inc(i); // 删除 i
            pre = i;
        }
        return ans;
    }
}

// 树状数组模板
class BIT {
    private final int[] tree;

    public BIT(int n) {
        tree = new int[n];
    }

    // 将下标 i 上的数加一
    public void inc(int i) {
        while (i < tree.length) {
            ++tree[i];
            i += i & -i;
        }
    }

    // 返回闭区间 [1, i] 的元素和
    public int sum(int x) {
        int res = 0;
        while (x > 0) {
            res += tree[x];
            x &= x - 1;
        }
        return res;
    }

    // 返回闭区间 [left, right] 的元素和
    public int query(int left, int right) {
        return sum(right) - sum(left - 1);
    }
}
```

```cpp [sol1-C++]
// 树状数组模板
class BIT {
    vector<int> tree;
public:
    BIT(int n) : tree(n) {}

    // 将下标 i 上的数加一
    void inc(int i) {
        while (i < tree.size()) {
            ++tree[i];
            i += i & -i;
        }
    }

    // 返回闭区间 [1, i] 的元素和
    int sum(int x) {
        int res = 0;
        while (x > 0) {
            res += tree[x];
            x &= x - 1;
        }
        return res;
    }

    // 返回闭区间 [left, right] 的元素和
    int query(int left, int right) {
        return sum(right) - sum(left - 1);
    }
};

class Solution {
public:
    long long countOperationsToEmptyArray(vector<int> &nums) {
        int n = nums.size(), id[n];
        iota(id, id + n, 0);
        sort(id, id + n, [&](int i, int j) {
            return nums[i] < nums[j];
        });

        long long ans = n; // 先把 n 计入答案
        BIT t(n + 1); // 下标从 1 开始
        int pre = 1; // 上一个最小值的位置，初始为 1
        for (int k = 0; k < n; ++k) {
            int i = id[k] + 1; // 下标从 1 开始
            if (i >= pre) // 从 pre 移动到 i，跳过已经删除的数
                ans += i - pre - t.query(pre, i);
            else // 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
                ans += n - pre + i - k + t.query(i, pre - 1);
            t.inc(i); // 删除 i
            pre = i;
        }
        return ans;
    }
};
```

```go [sol1-Go]
// 树状数组模板
type BIT []int

// 将下标 i 上的数加一
func (t BIT) inc(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

// 返回闭区间 [1, i] 的元素和
func (t BIT) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

// 返回闭区间 [left, right] 的元素和
func (t BIT) query(left, right int) int {
	return t.sum(right) - t.sum(left-1)
}

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })

	ans := n // 先把 n 计入答案
	t := make(BIT, n+1) // 下标从 1 开始
	pre := 1 // 上一个最小值的位置，初始为 1
	for k, i := range id {
		i++ // 下标从 1 开始
		if i >= pre { // 从 pre 移动到 i，跳过已经删除的数
			ans += i - pre - t.query(pre, i)
		} else { // 从 pre 移动到 n，再从 1 移动到 i，跳过已经删除的数
			ans += n - pre + i - k + t.query(i, pre-1)
		}
		t.inc(i) // 删除 i
		pre = i
	}
	return int64(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：进一步挖掘性质

仍然先把数组长度 $n$ 计入答案，后面只统计移动次数。在统计移动次数时，遇到要删除的元素，相当于可以免费向后移动一步（因为删除操作已经计入答案）。试想一下，如果数组是单调递增的，就没有任何额外的移动次数。

再次考察 $\textit{nums}=[2,4,1,3]$：

- 从元素 $1$ 移动到元素 $2$，由于 $2$ 在 $1$ 左侧，说明必须走一整圈才能到 $2$，减去删除 $1$ 产生的免费移动，移动次数为 $n-1=4-1=3$。
- 从元素 $2$ 移动到元素 $3$，这里合并到下面计算。
- 从元素 $3$ 移动到元素 $4$，由于 $4$ 在 $3$ 左侧，说明必须再走一整圈才能到 $4$，减去删除 $2,3$ 产生的免费移动，减去跳过的 $1$，移动次数为 $n-3=4-3=1$。
- 总共移动 $3+1=4$ 次，加上删除操作 $4$ 次，故答案为 $8$。

> 这里说的「走一整圈」指从数组左端走到右端，再回到数组左端。

从上面的例子中可以发现，如果第 $k$ 次要删除的元素在第 $k-1$ 次要删除的元素的左侧，那么必须多走一整圈，移动次数为 $n-k$。累加，即为总的移动次数。

> 最后如果剩下若干递增元素，直接一股脑删除，无需任何移动次数。

```py [sol2-Python3]
class Solution:
    def countOperationsToEmptyArray(self, nums: List[int]) -> int:
        ans = n = len(nums)
        id = sorted(range(n), key=lambda x: nums[x])
        for k, (pre, i) in enumerate(pairwise(id), 1):
            if i < pre:  # 必须多走一整圈
                ans += n - k  # 减去前面删除的元素个数
        return ans
```

```java [sol2-Java]
class Solution {
    public long countOperationsToEmptyArray(int[] nums) {
        int n = nums.length;
        var id = new Integer[n];
        for (int i = 0; i < n; ++i)
            id[i] = i;
        Arrays.sort(id, (i, j) -> nums[i] - nums[j]);

        long ans = n; // 先把 n 计入答案
        for (int k = 1; k < n; ++k)
            if (id[k] < id[k - 1]) // 必须多走一整圈
                ans += n - k; // 减去前面删除的元素个数
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long countOperationsToEmptyArray(vector<int> &nums) {
        int n = nums.size(), id[n];
        iota(id, id + n, 0);
        sort(id, id + n, [&](int i, int j) { return nums[i] < nums[j]; });
        long long ans = n;
        for (int k = 1; k < n; ++k)
            if (id[k] < id[k - 1]) // 必须多走一整圈
                ans += n - k; // 减去前面删除的元素个数
        return ans;
    }
};
```

```go [sol2-Go]
func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })

	ans := n
	for k := 1; k < n; k++ {
		if id[k] < id[k-1] { // 必须多走一整圈
			ans += n - k // 减去前面删除的元素个数
		}
	}
	return int64(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。
