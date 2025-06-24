## 方法一：离线算法 + 双指针

> **离线算法**：把 $\textit{queries}$ 排序，通过改变回答询问的顺序，使问题更容易处理。
>
> **在线算法**：按照 $\textit{queries}$ 的顺序一个一个地回答询问。

暴力的做法是，对于每个询问 $q=\textit{queries}[i]$，遍历 $\textit{items}$，计算其中 $\textit{price}\le q$ 的最大 $\textit{beauty}$。

如何优化？

假如 $\textit{queries}$ 已经按照从小到大的顺序排好了，例如示例 1 $\textit{queries}=[1,2,3,4,5,6]$。

- 首先找所有 $\textit{price}\le \textit{queries}[0]=1$ 的物品，求得其中最大 $\textit{beauty}$ 为 $2$。
- 然后找所有 $\textit{price}\le \textit{queries}[1]=2$ 的物品，由于我们已经知道 $\textit{price}\le 1$ 的物品的最大 $\textit{beauty}$ 为 $2$。所以只需要求出 $\textit{price}$ 大于 $1$ 且小于等于 $2$ 的物品中的最大 $\textit{beauty}$，即 $4$，然后计算 $\max(4,2)=4$，即为所有 $\textit{price}\le 2$ 的物品中的最大 $\textit{beauty}$。
- 继续，找所有 $\textit{price}\le \textit{queries}[2]=3$ 的物品，由于我们已经知道 $\textit{price}\le 2$ 的物品的最大 $\textit{beauty}$ 为 $4$。所以只需要求出 $\textit{price}$ 大于 $2$ 且小于等于 $3$ 的物品中的最大 $\textit{beauty}$，即 $5$，然后计算 $\max(5,4)=5$，即为所有 $\textit{price}\le 3$ 的物品中的最大 $\textit{beauty}$。
- 依此类推，我们只需要「增量」地计算所有满足 $\textit{queries}[i-1] < \textit{price}\le \textit{queries}[i]$ 的物品中的最大 $\textit{beauty}$，然后和上一次计算出的最大 $\textit{beauty}$ 取最大值，即为所有 $\textit{price}\le \textit{queries}[i]$ 的物品中的最大 $\textit{beauty}$。

为此，需要做两件事情：

1. 把询问从小到大排序。但由于 $\textit{answer}$ 需要按照输入的顺序回答，可以额外创建一个下标数组，对下标数组排序。
2. 把物品按价格从小到大排序，这样就可以用**双指针**「增量」地遍历满足 $\textit{queries}[i-1] < \textit{price}\le \textit{queries}[i]$ 的物品。

```py [sol-Python3]
class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort(key=lambda item: item[0])
        idx = sorted(range(len(queries)), key=lambda i: queries[i])

        ans = [0] * len(queries)
        max_beauty = j = 0
        for i in idx:
            q = queries[i]
            # 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
            while j < len(items) and items[j][0] <= q:
                max_beauty = max(max_beauty, items[j][1])
                j += 1
            ans[i] = max_beauty
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, (a, b) -> a[0] - b[0]);
        Integer[] idx = new Integer[queries.length];
        Arrays.setAll(idx, i -> i);
        Arrays.sort(idx, (i, j) -> queries[i] - queries[j]);

        int[] ans = new int[queries.length];
        int maxBeauty = 0;
        int j = 0;
        for (int i : idx) {
            int q = queries[i];
            // 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
            while (j < items.length && items[j][0] <= q) {
                maxBeauty = Math.max(maxBeauty, items[j][1]);
                j++;
            }
            ans[i] = maxBeauty;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumBeauty(vector<vector<int>>& items, vector<int>& queries) {
        ranges::sort(items, {}, [](auto& item) { return item[0]; });
        vector<int> idx(queries.size());
        iota(idx.begin(), idx.end(), 0);
        ranges::sort(idx, {}, [&](int i) { return queries[i]; });

        vector<int> ans(queries.size());
        int max_beauty = 0, j = 0;
        for (int i : idx) {
            int q = queries[i];
            // 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
            while (j < items.size() && items[j][0] <= q) {
                max_beauty = max(max_beauty, items[j][1]);
                j++;
            }
            ans[i] = max_beauty;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumBeauty(items [][]int, queries []int) []int {
    slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
    idx := make([]int, len(queries))
    for i := range queries {
        idx[i] = i
    }
    slices.SortFunc(idx, func(i, j int) int { return queries[i] - queries[j] })

    ans := make([]int, len(queries))
    maxBeauty, j := 0, 0
    for _, i := range idx {
        q := queries[i]
        // 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
        for j < len(items) && items[j][0] <= q {
            maxBeauty = max(maxBeauty, items[j][1])
            j++
        }
        ans[i] = maxBeauty
    }
    return ans
}
```

```js [sol-JavaScript]
var maximumBeauty = function(items, queries) {
    items.sort((a, b) => a[0] - b[0]);
    const nq = queries.length;
    const idx = _.range(nq).sort((i, j) => queries[i] - queries[j]);

    const ans = Array(nq);
    let maxBeauty = 0, j = 0;
    for (const i of idx) {
        const q = queries[i];
        // 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
        while (j < items.length && items[j][0] <= q) {
            maxBeauty = Math.max(maxBeauty, items[j][1]);
            j++;
        }
        ans[i] = maxBeauty;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_beauty(mut items: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        items.sort_unstable_by_key(|item| item[0]);
        let mut idx = (0..queries.len()).collect::<Vec<_>>();
        idx.sort_unstable_by_key(|&i| queries[i]);

        let mut ans = vec![0; queries.len()];
        let mut max_beauty = 0;
        let mut j = 0;
        for i in idx {
            let q = queries[i];
            // 增量地遍历满足 queries[i-1] < price <= queries[i] 的物品
            while j < items.len() && items[j][0] <= q {
                max_beauty = max_beauty.max(items[j][1]);
                j += 1;
            }
            ans[i] = max_beauty;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n +m\log m)$，其中 $n$ 是 $\textit{items}$ 的长度，$m$ 是 $\textit{queries}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(m)$。

## 方法二：在线算法 + 二分查找

### 写法一：前缀最大值

示例 1 的 $\textit{items} = [[1,2],[3,2],[2,4],[5,6],[3,5]]$，将其按照 $\textit{price}$ 从小到大排序，得

$$
[[1,2],[2,4],[3,2],[3,5],[5,6]]
$$

然后原地计算其 $\textit{beauty}$ 的前缀最大值，得

$$
[[1,2],[2,4],[3,\underline{4}],[3,5],[5,6]]
$$

注意其中 $[3,2]$ 变成了 $[3,4]$，这里的 $4$ 就是前三个物品的最大 $\textit{beauty}$，即 $\max(2,4,2)=4$。

算好前缀最大值后，所有 $\textit{price}\le q$ 的物品的最大 $\textit{beauty}$，就保存在满足 $\textit{price}\le q$ 的**最右边**的那个物品中！

根据 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)中的技巧，我们可以二分 $\textit{price}>q$ 的第一个物品，它左边相邻物品就是 $\textit{price}\le q$ 的最后一个物品。（如果左边没有物品那么答案为 $0$）

```py [sol-Python3]
class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort(key=lambda item: item[0])
        for i in range(1, len(items)):
            # 原地计算 beauty 的前缀最大值
            items[i][1] = max(items[i][1], items[i - 1][1])

        for i, q in enumerate(queries):
            j = bisect_right(items, q, key=lambda item: item[0])
            queries[i] = items[j - 1][1] if j else 0
        return queries
```

```java [sol-Java]
class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, (a, b) -> a[0] - b[0]);
        for (int i = 1; i < items.length; i++) {
            // 原地计算 beauty 的前缀最大值
            items[i][1] = Math.max(items[i][1], items[i - 1][1]);
        }

        for (int i = 0; i < queries.length; i++) {
            int j = upperBound(items, queries[i]);
            queries[i] = j > 0 ? items[j - 1][1] : 0;
        }
        return queries;
    }

    // https://www.bilibili.com/video/BV1AP41137w7/
    // 返回 items 中第一个 price 大于 target 的物品的下标（注意是大于，不是大于等于）
    // 如果这样的数不存在，则返回 items.length
    // 时间复杂度 O(log items.length)
    // 采用开区间写法实现
    private int upperBound(int[][] items, int target) {
        int left = -1;
        int right = items.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // items[left][0] <= target
            // items[right][0] > target
            int mid = (left + right) >>> 1;
            if (items[mid][0] > target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumBeauty(vector<vector<int>>& items, vector<int>& queries) {
        ranges::sort(items, {}, [](auto& item) { return item[0]; });
        for (int i = 1; i < items.size(); i++) {
            // 原地计算 beauty 的前缀最大值
            items[i][1] = max(items[i][1], items[i - 1][1]);
        }

        for (int& q : queries) {
            int j = ranges::upper_bound(items, q, {}, [](auto& item) { return item[0]; }) - items.begin();
            q = j ? items[j - 1][1] : 0;
        }
        return queries;
    }
};
```

```go [sol-Go]
func maximumBeauty(items [][]int, queries []int) []int {
    slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
    for i := 1; i < len(items); i++ {
        // 原地计算 beauty 的前缀最大值
        items[i][1] = max(items[i][1], items[i-1][1])
    }

    for i, q := range queries {
        j := sort.Search(len(items), func(i int) bool { return items[i][0] > q })
        if j > 0 {
            queries[i] = items[j-1][1]
        } else {
            queries[i] = 0
        }
    }
    return queries
}
```

```js [sol-JavaScript]
var maximumBeauty = function(items, queries) {
    items.sort((a, b) => a[0] - b[0]);
    for (let i = 1; i < items.length; i++) {
        // 原地计算 beauty 的前缀最大值
        items[i][1] = Math.max(items[i][1], items[i - 1][1]);
    }

    for (let i = 0; i < queries.length; i++) {
        const j = upperBound(items, queries[i]);
        queries[i] = j > 0 ? items[j - 1][1] : 0;
    }
    return queries;
};

function upperBound(items, target) {
    let left = -1, right = items.length;
    while (left + 1 < right) {
        const mid = Math.floor((left + right) / 2);
        if (items[mid][0] > target) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_beauty(mut items: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        items.sort_unstable_by_key(|item| item[0]);
        for i in 1..items.len() {
            // 原地计算 beauty 的前缀最大值
            items[i][1] = items[i][1].max(items[i - 1][1]);
        }

        queries.into_iter().map(|q| {
            let j = items.partition_point(|item| item[0] <= q);
            if j > 0 { items[j - 1][1] } else { 0 }
        }).collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log n)$，其中 $n$ 是 $\textit{items}$ 的长度，$m$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

### 写法二：去除无用信息

再来看这个排序后的 $\textit{items}$：

$$
[[1,2],[2,4],[3,2],[3,5],[5,6]]
$$

其中 $[3,2]$ 价格比 $[2,4]$ 高，美丽值又比 $[2,4]$ 低，那么 $[3,2]$ 就完全不如 $[2,4]$，可以直接去掉。（这类似单调栈/单调队列的思想）

去掉这种无用数据后，数组变成

$$
[[1,2],[2,4],[3,5],[5,6]]
$$

此时 $\textit{beauty}$ 就是严格递增的。

这样做的好处是 $\textit{items}$ 更短，计算二分的时间也更短。

如何原地计算？做法可参考 [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)。

```py [sol-Python3]
class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort(key=lambda item: item[0])
        k = 0
        for i in range(1, len(items)):
            if items[i][1] > items[k][1]:  # 有用
                k += 1
                items[k] = items[i]

        for i, q in enumerate(queries):
            j = bisect_right(items, q, 0, k + 1, key=lambda item: item[0])
            queries[i] = items[j - 1][1] if j else 0
        return queries
```

```java [sol-Java]
class Solution {
    public int[] maximumBeauty(int[][] items, int[] queries) {
        Arrays.sort(items, (a, b) -> a[0] - b[0]);
        int k = 0;
        for (int i = 1; i < items.length; i++) {
            if (items[i][1] > items[k][1]) { // 有用
                items[++k] = items[i];
            }
        }

        for (int i = 0; i < queries.length; i++) {
            int j = upperBound(items, k + 1, queries[i]);
            queries[i] = j > 0 ? items[j - 1][1] : 0;
        }
        return queries;
    }

    private int upperBound(int[][] items, int right, int target) {
        int left = -1; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // items[left][0] <= target
            // items[right][0] > target
            int mid = (left + right) >>> 1;
            if (items[mid][0] > target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumBeauty(vector<vector<int>>& items, vector<int>& queries) {
        ranges::sort(items, {}, [](auto& item) { return item[0]; });
        int k = 0;
        for (int i = 1; i < items.size(); i++) {
            if (items[i][1] > items[k][1]) { // 有用
                items[++k] = items[i];
            }
        }

        for (int& q : queries) {
            int j = upper_bound(items.begin(), items.begin() + (k + 1), q, [](int value, auto& item) {
                return value < item[0];
            }) - items.begin();
            q = j ? items[j - 1][1] : 0;
        }
        return queries;
    }
};
```

```go [sol-Go]
func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] })
	k := 0
	for _, item := range items[1:] {
		if item[1] > items[k][1] { // 有用
			k++
			items[k] = item
		}
	}

	for i, q := range queries {
		j := sort.Search(k+1, func(i int) bool { return items[i][0] > q })
		if j > 0 {
			queries[i] = items[j-1][1]
		} else {
			queries[i] = 0
		}
	}
	return queries
}
```

```js [sol-JavaScript]
var maximumBeauty = function(items, queries) {
    items.sort((a, b) => a[0] - b[0]);
    let k = 0;
    for (let i = 1; i < items.length; i++) {
        if (items[i][1] > items[k][1]) { // 有用
            items[++k] = items[i];
        }
    }

    for (let i = 0; i < queries.length; i++) {
        const j = upperBound(items, k + 1, queries[i]);
        queries[i] = j > 0 ? items[j - 1][1] : 0;
    }
    return queries;
};

function upperBound(items, right, target) {
    let left = -1;
    while (left + 1 < right) {
        const mid = Math.floor((left + right) / 2);
        if (items[mid][0] > target) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_beauty(mut items: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        items.sort_unstable_by_key(|item| item[0]);
        items.dedup_by(|b, a| b[1] <= a[1]); // 去掉无用数据

        queries.into_iter().map(|q| {
            let j = items.partition_point(|item| item[0] <= q);
            if j > 0 { items[j - 1][1] } else { 0 }
        }).collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log n)$，其中 $n$ 是 $\textit{items}$ 的长度，$m$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

更多相似题目，见下面数据结构题单中的「**专题：离线算法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
