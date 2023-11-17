为方便处理，先把 $\textit{nums}_1$ 和 $\textit{nums}_2$ 绑在一起（记作 pair 数组 $a$），按照 $\textit{nums}_1$ 的值**从大到小**排序，然后按照 $x_i$ **从大到小**的顺序回答询问。

对于示例 1 来说，排序后 $a=[(4, 2), (3, 4), (2, 5), (1, 9)]$，然后按照询问下标 $0,2,1$ 的顺序回答询问，即依次回答询问 $[4,1],[2,5],[1,3]$。

由于 $\textit{nums}_1[j]$ 和 $x_i$ 都是有序的，我们可以把重心放在 $\textit{nums}_2[j]$ 与 $y_i$ 上。

整体思路是，从 $j=0$ 开始遍历 pair 数组，并逐一回答询问 $(x_i,y_i)$。在这个过程中，把满足 $\textit{nums}_1[j]\ge x_i$ 的 $\textit{nums}_2[j]$ 记录到一个数据结构中，然后在这个数据结构中查找 $\ge y_i$ 的 $\textit{nums}_2[j]$，并想办法求出 $\textit{nums}_1[j] + \textit{nums}_2[j]$ 的最大值。

具体来说，按照 $\textit{nums}_2[j]$ 与之前遍历过的 $\textit{nums}_2[j']$ 的大小关系，分类讨论：

- 如果 $\textit{nums}_2[j]$ 比之前遍历过的 $\textit{nums}_2[j']$ 要**小**，由于 $\textit{nums}_1[j]$ 已经从大到小排序，所以 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 也比之前遍历过的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 要小。所以在回答询问时，最大值不可能是 $\textit{nums}_1[j]+\textit{nums}_2[j]$，所以无需考虑这样的 $\textit{nums}_2[j]$。（这种单调性启发我们用**单调栈**来维护。）
- 如果**相等**，同理，无需考虑。
- 如果**大于**，就可以把 $\textit{nums}_2[j]$ 入栈（同时把 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 也入栈）。在入栈前，去掉一些无效数据：如果 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 不低于栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$，那么可以弹出栈顶。因为更大的 $\textit{nums}_2[j]$ 更能满足 $\ge y_i$ 的要求，所以栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 在后续的询问中，永远不会作为最大值。
- 代码实现时，可以直接比较 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 与栈顶的值，这是因为如果这一条件成立，由于 $\textit{nums}_1[j]$ 是从大到小处理的，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 能比栈顶的大，说明 $\textit{nums}_2[j]$ 必然不低于栈顶的 $\textit{nums}_2[j']$。

这样我们会得到一个从栈底到栈顶，$\textit{nums}_2[j]$ 递增，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 递减的单调栈。在单调栈中二分 $\ge y_i$ 的最小的 $\textit{nums}_2[j]$，对应的 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 就是最大的。

```py [sol-Python3]
class Solution:
    def maximumSumQueries(self, nums1: List[int], nums2: List[int], queries: List[List[int]]) -> List[int]:
        ans = [-1] * len(queries)
        a = sorted(((a, b) for a, b in zip(nums1, nums2)), key=lambda p: -p[0])
        j = 0
        st = []
        for i, (x, y) in sorted(enumerate(queries), key=lambda p: -p[1][0]):
            while j < len(a) and a[j][0] >= x:  # 下面只需关心 ay (a[j][1])
                ax, ay = a[j]
                while st and st[-1][1] <= ax + ay:  # ay >= st[-1][0]
                    st.pop()
                if not st or st[-1][0] < ay:
                    st.append((ay, ax + ay))
                j += 1
            p = bisect_left(st, (y,))
            if p < len(st):
                ans[i] = st[p][1]
        return ans
```

```java [sol-Java]
class Solution {
    public int[] maximumSumQueries(int[] nums1, int[] nums2, int[][] queries) {
        int n = nums1.length;
        int[][] a = new int[n][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = nums1[i];
            a[i][1] = nums2[i];
        }
        Arrays.sort(a, (x, y) -> y[0] - x[0]);

        Integer[] qid = new Integer[queries.length];
        for (int i = 0; i < queries.length; i++) {
            qid[i] = i;
        }
        Arrays.sort(qid, (i, j) -> queries[j][0] - queries[i][0]);

        int[] ans = new int[queries.length];
        List<int[]> st = new ArrayList<>();
        int j = 0;
        for (int i : qid) {
            int x = queries[i][0], y = queries[i][1];
            for (; j < n && a[j][0] >= x; j++) { // 下面只需关心 a[j][1]
                while (!st.isEmpty() && st.get(st.size() - 1)[1] <= a[j][0] + a[j][1]) { // a[j][1] >= st.get(st.size()-1)[0]
                    st.remove(st.size() - 1);
                }
                if (st.isEmpty() || st.get(st.size() - 1)[0] < a[j][1]) {
                    st.add(new int[]{a[j][1], a[j][0] + a[j][1]});
                }
            }
            int p = lowerBound(st, y);
            ans[i] = p < st.size() ? st.get(p)[1] : -1;
        }
        return ans;
    }

    // 开区间写法，原理请看 b23.tv/AhwfbS2
    private int lowerBound(List<int[]> st, int target) {
        int left = -1, right = st.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            int mid = (left + right) >>> 1;
            if (st.get(mid)[0] >= target) {
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
    vector<int> maximumSumQueries(vector<int> &nums1, vector<int> &nums2, vector<vector<int>> &queries) {
        vector<pair<int, int>> a(nums1.size());
        for (int i = 0; i < nums1.size(); i++) {
            a[i] = {nums1[i], nums2[i]};
        }
        sort(a.begin(), a.end(),
             [](auto &a, auto &b) { return a.first > b.first; });

        vector<int> qid(queries.size());
        iota(qid.begin(), qid.end(), 0);
        sort(qid.begin(), qid.end(),
             [&](int i, int j) { return queries[i][0] > queries[j][0]; });

        vector<int> ans(queries.size());
        vector<pair<int, int>> st;
        int j = 0;
        for (int i: qid) {
            int x = queries[i][0], y = queries[i][1];
            for (; j < a.size() && a[j].first >= x; j++) { // 下面只需关心 a[j].second
                while (!st.empty() && st.back().second <= a[j].first + a[j].second) { // a[j].second >= st.back().first
                    st.pop_back();
                }
                if (st.empty() || st.back().first < a[j].second) {
                    st.emplace_back(a[j].second, a[j].first + a[j].second);
                }
            }
            auto it = lower_bound(st.begin(), st.end(), y,
                          [](const auto &p, int val) { return p.first < val; });
            ans[i] = it != st.end() ? it->second : -1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumSumQueries(nums1, nums2 []int, queries [][]int) []int {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	slices.SortFunc(a, func(a, b pair) int { return b.x - a.x })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}
	slices.SortFunc(qid, func(i, j int) int { return queries[j][0] - queries[i][0] })

	ans := make([]int, len(queries))
	type data struct{ y, s int }
	st := []data{}
	j := 0
	for _, i := range qid {
		x, y := queries[i][0], queries[i][1]
		for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
			for len(st) > 0 && st[len(st)-1].s <= a[j].x+a[j].y { // a[j].y >= st[len(st)-1].y
				st = st[:len(st)-1]
			}
			if len(st) == 0 || st[len(st)-1].y < a[j].y {
				st = append(st, data{a[j].y, a[j].x + a[j].y})
			}
		}
		p := sort.Search(len(st), func(i int) bool { return st[i].y >= y })
		if p < len(st) {
			ans[i] = st[p].s
		} else {
			ans[i] = -1
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var maximumSumQueries = function (nums1, nums2, queries) {
    const a = _.zip(nums1, nums2).sort((a, b) => b[0] - a[0])
    const qid = [...queries.keys()].sort((i, j) => queries[j][0] - queries[i][0]);

    const ans = Array(queries.length);
    const st = [];
    let j = 0;
    for (const i of qid) {
        const [x, y] = queries[i];
        for (; j < a.length && a[j][0] >= x; j++) { // 下面只需关心 a[j][1]
            while (st.length && st[st.length - 1][1] <= a[j][0] + a[j][1]) { // a[j][1] >= st[st.length-1][0]
                st.pop();
            }
            if (!st.length || st[st.length - 1][0] < a[j][1]) {
                st.push([a[j][1], a[j][0] + a[j][1]]);
            }
        }
        const p = lowerBound(st, y);
        ans[i] = p < st.length ? st[p][1] : -1;
    }
    return ans;
};

// 开区间写法，原理请看 b23.tv/AhwfbS2
var lowerBound = function (st, target) {
    let left = -1, right = st.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        const mid = left + ((right - left) >> 1);
        if (st[mid][0] >= target) {
            right = mid; // 范围缩小到 (left, mid)
        } else {
            left = mid; // 范围缩小到 (mid, right)
        }
    }
    return right; // 或者 left+1
}
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_sum_queries(nums1: Vec<i32>, nums2: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut a: Vec<(i32, i32)> = nums1.into_iter().zip(nums2.into_iter()).collect();
        a.sort_by(|x, y| y.0.cmp(&x.0));

        let mut qid: Vec<usize> = (0..queries.len()).collect();
        qid.sort_by(|&i, &j| queries[j][0].cmp(&queries[i][0]));

        let mut ans = vec![-1; queries.len()];
        let mut st: Vec<(i32, i32)> = Vec::new();
        let mut j = 0;
        for &i in &qid {
            let x = queries[i][0];
            let y = queries[i][1];
            while j < a.len() && a[j].0 >= x { // 下面只需关心 a[j].1
                while !st.is_empty() && st.last().unwrap().1 <= a[j].0 + a[j].1 { // a[j].1 >= st.last().unwrap().0
                    st.pop();
                }
                if st.is_empty() || st.last().unwrap().0 < a[j].1 {
                    st.push((a[j].1, a[j].0 + a[j].1));
                }
                j += 1;
            }
            let p = st.partition_point(|&p| p.0 < y);
            if p < st.len() {
                ans[i] = st[p].1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + q\log q + q\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + q)$。

[往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
