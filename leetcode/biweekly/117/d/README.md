下午两点[【b站@灵茶山艾府】](https://b23.tv/JMcHRRp)直播讲题，欢迎关注！

---

## 提示 1

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，$\textit{values}[i][j]$ 越小的数，应该越早购买。

## 提示 2

$\textit{values}$ 的最后一列一定包含最小的数，因为其余列的数都不会比最后一列的这 $m$ 个数小。

次小的数呢？

## 提示 3

考虑每行当前最后一个数，这 $m$ 个数中一定包含次小的数，因为和上面一样，其余数字都不会比这 $m$ 个数小。

## 提示 4

根据数学归纳法，我们可以按照 $\textit{values}[i][j]$ 从小到大的顺序取到所有元素。

所以，把所有数合并在一起排序即可。

也可以用最小堆实现，那样空间更小。

#### 写法一：排序

```py [sol-Python3]
class Solution:
    def maxSpending(self, values: List[List[int]]) -> int:
        a = sorted(x for row in values for x in row)
        return sum(x * i for i, x in enumerate(a, 1))
```

```java [sol-Java]
class Solution {
    public long maxSpending(int[][] values) {
        int m = values.length, n = values[0].length;
        int[] a = new int[m * n];
        int i = 0;
        for (int[] row : values) {
            System.arraycopy(row, 0, a, i, n);
            i += n;
        }
        Arrays.sort(a);
        long ans = 0;
        for (i = 0; i < a.length; i++) {
            ans += (long) a[i] * (i + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSpending(vector<vector<int>> &values) {
        int m = values.size(), n = values[0].size();
        vector<int> a;
        a.reserve(m * n);
        for (auto &row: values) {
            a.insert(a.end(), row.begin(), row.end());
        }
        sort(a.begin(), a.end());
        long long ans = 0;
        for (int i = 0; i < a.size(); i++) {
            ans += (long long) a[i] * (i + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	a := make([]int, 0, m*n)
	for _, row := range values {
		a = append(a, row...)
	}
	slices.Sort(a)
	ans := 0
	for i, x := range a {
		ans += x * (i + 1)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$，其中 $m$ 和 $n$ 分别为 $\textit{values}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

#### 写法二：最小堆

```py [sol-Python3]
class Solution:
    def maxSpending(self, values: List[List[int]]) -> int:
        h = [(a[-1], i) for i, a in enumerate(values)]
        heapify(h)
        ans = 0
        for d in range(1, len(values) * len(values[0]) + 1):
            v, i = heappop(h)
            ans += v * d
            values[i].pop()
            if values[i]:
                heappush(h, (values[i][-1], i))
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSpending(int[][] values) {
        int m = values.length, n = values[0].length;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> values[a[0]][a[1]] - values[b[0]][b[1]]);
        for (int i = 0; i < m; i++) {
            pq.offer(new int[]{i, n - 1});
        }
        long ans = 0;
        for (int d = 1; d <= m * n; d++) {
            int[] p = pq.poll();
            int i = p[0], j = p[1];
            ans += (long) values[i][j] * d;
            if (j > 0) {
                pq.offer(new int[]{i, j - 1});
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSpending(vector<vector<int>> &values) {
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<>> pq;
        int m = values.size(), n = values[0].size();
        for (int i = 0; i < m; i++) {
            pq.emplace(values[i].back(), i);
        }
        long long ans = 0;
        for (int d = 1; d <= m * n; d++) {
            auto [v, i] = pq.top();
            pq.pop();
            ans += (long long) v * d;
            values[i].pop_back();
            if (!values[i].empty()) {
                pq.push({values[i].back(), i});
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	id := make([]int, m)
	for i := range id {
		id[i] = i
	}
	h := &hp{id, values}
	heap.Init(h)
	ans := 0
	for d := 1; d <= m*n; d++ {
		a := values[id[0]]
		ans += a[len(a)-1] * d
		if len(a) > 1 {
			values[id[0]] = a[:len(a)-1]
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return int64(ans)
}

type hp struct {
	sort.IntSlice
	values [][]int
}
func (h hp) Less(i, j int) bool {
	a, b := h.values[h.IntSlice[i]], h.values[h.IntSlice[j]]
	return a[len(a)-1] < b[len(b)-1]
}
func (hp) Push(any) {}
func (h *hp) Pop() (_ any) { a := h.IntSlice; h.IntSlice = a[:len(a)-1]; return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log m)$，其中 $m$ 和 $n$ 分别为 $\textit{values}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。
