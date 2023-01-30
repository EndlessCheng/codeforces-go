本题 [视频讲解](https://www.bilibili.com/video/BV1GV4y1J7kc) 已出炉，欢迎点赞三连~

---

## 方法一：直接排序

对每个询问，按照题目要求排序，取第 $k$ 小的元素的下标。

#### 复杂度分析

- 时间复杂度：$O(qmn\log n)$，其中 $q$ 是数组 $\textit{queries}$ 的长度，$n$ 是数组 $\textit{nums}$ 的长度，$m$ 是每个 $\textit{nums}[i]$ 的长度。每次询问，都需要对一个长为 $n$ 的数组排序，排序共发生 $O(n\log n)$ 次比较，每次比较的耗时为 $O(m)$，故总的时间复杂度为 $O(qmn\log n)$。
- 空间复杂度：$O(n)$。返回值不计入空间复杂度。

```py [sol1-Python3]
class Solution:
    def smallestTrimmedNumbers(self, nums: List[str], queries: List[List[int]]) -> List[int]:
        return [sorted((s[-trim:], i) for i, s in enumerate(nums))[k - 1][1] for k, trim in queries]
```

```java [sol1-Java]
class Solution {
    public int[] smallestTrimmedNumbers(String[] nums, int[][] queries) {
        var ans = new int[queries.length];
        var m = nums[0].length();
        for (var p = 0; p < queries.length; p++) {
            var q = queries[p];
            var idx = new ArrayList<>(Arrays.asList(IntStream.range(0, nums.length).boxed().toArray(Integer[]::new)));
            Collections.sort(idx, (i, j) -> nums[i].substring(m - q[1]).compareTo(nums[j].substring(m - q[1]))); // 稳定排序
            ans[p] = idx.get(q[0] - 1);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> smallestTrimmedNumbers(vector<string> &nums, vector<vector<int>> &queries) {
        vector<int> ans(queries.size());
        int n = nums.size(), m = nums[0].length();
        int idx[n];
        for (int i = 0; i < queries.size(); ++i) {
            auto &q = queries[i];
            iota(idx, idx + n, 0);
            stable_sort(idx, idx + n, [&](int a, int b) {
                auto &s = nums[a], &t = nums[b];
                for (int j = m - q[1]; j < m; ++j)
                    if (s[j] != t[j]) return s[j] < t[j];
                return false;
            });
            ans[i] = idx[q[0] - 1];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, len(queries))
	type pair struct { s string; i int }
	ps := make([]pair, len(nums))
	for i, q := range queries {
		for j, s := range nums {
			ps[j] = pair{s[len(s)-q[1]:], j}
		}
		// 也可以用稳定排序，但是要慢一些 sort.SliceStable(ps, func(i, j int) bool { return ps[i].s < ps[j].s })
		sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.s < b.s || a.s == b.s && a.i < b.i })
		ans[i] = ps[q[0]-1].i
	}
	return ans
}
```

## 方法二：离线 + 增量排序

方法一排序时，每次都需要比较整个字符串。若采用增量排序，每次排序只需要比较单个字符。

具体来说，先将询问按照 $\textit{trim}$ 从小到大排序，并按照这一顺序回答询问。随着 $\textit{tirm}$ 的不断增加，我们可以在排好序的字符串数组的基础上，向每个字符串的前面添加一个对应的字符，由于字符串数组已经是有序的，我们只需要比较这个新增的字符的大小，即可比较整个字符串的大小。

#### 复杂度分析

- 时间复杂度：$O(q\log q+mn\log n)$ 或 $O(q\log q+mn)$。对询问排序需要 $O(q\log q)$，后面至多排序 $m$ 次，每次排序的时间复杂度为 $O(n\log n)$，故总的时间复杂度为 $O(q\log q+mn\log n)$。如果用基数排序可以做到 $O(q\log q+mn)$。
- 空间复杂度：$O(q+n)$。

```py [sol2-Python3]
class Solution:
    def smallestTrimmedNumbers(self, nums: List[str], queries: List[List[int]]) -> List[int]:
        idx = list(range(len(nums)))
        ans, j = [0] * len(queries), 1
        for qi, (k, trim) in sorted(enumerate(queries), key=lambda q: q[1][1]):  # 按 trim 排序
            while j <= trim:
                idx.sort(key=lambda i: nums[i][-j])  # 只比较倒数第 j 个字符的大小
                j += 1
            ans[qi] = idx[k - 1]
        return ans
```

```java [sol2-Java]
class Solution {
    public int[] smallestTrimmedNumbers(String[] nums, int[][] queries) {
        var qid = IntStream.range(0, queries.length).boxed().toArray(Integer[]::new);
        Arrays.sort(qid, (i, j) -> queries[i][1] - queries[j][1]);

        var m = nums[0].length();
        var idx = new ArrayList<>(Arrays.asList(IntStream.range(0, nums.length).boxed().toArray(Integer[]::new)));

        var ans = new int[queries.length];
        var p = 1;
        for (var qi : qid) {
            var q = queries[qi];
            while (p <= q[1]) {
                final var pp = p++;
                Collections.sort(idx, (i, j) -> nums[i].charAt(m - pp) - nums[j].charAt(m - pp)); // 稳定排序
            }
            ans[qi] = idx.get(q[0] - 1);
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    vector<int> smallestTrimmedNumbers(vector<string> &nums, vector<vector<int>> &queries) {
        int nq = queries.size();
        int qid[nq];
        iota(qid, qid + nq, 0);
        sort(qid, qid + nq, [&](int a, int b) { return queries[a][1] < queries[b][1]; });

        int n = nums.size(), m = nums[0].length();
        int idx[n];
        iota(idx, idx + n, 0);

        vector<int> ans(nq);
        int p = 1;
        for (int qi : qid) {
            auto &q = queries[qi];
            for (; p <= q[1]; ++p)
                stable_sort(idx, idx + n, [&](int a, int b) { return nums[a][m - p] < nums[b][m - p]; });
            ans[qi] = idx[q[0] - 1];
        }
        return ans;
    }
};
```

```go [sol2-Go]
func smallestTrimmedNumbers(nums []string, queries [][]int) (ans []int) {
	for i, q := range queries {
		q[0] |= i << 32 // 把询问的下标整合到 k 里面，相比 append 到 q 里面可以避免扩容
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][1] < queries[j][1] }) // 按 trim 排序

	m := len(nums[0])
	type pair struct { s string; i int }
	ps := make([]pair, len(nums))
	for i, s := range nums {
		ps[i] = pair{s, i}
	}

	ans = make([]int, len(queries))
	p := 1
	for _, q := range queries {
		for ; p <= q[1]; p++ {
			sort.SliceStable(ps, func(i, j int) bool { return ps[i].s[m-p] < ps[j].s[m-p] }) // 只比较第 m-p 个字符的大小
		}
		ans[q[0]>>32] = ps[q[0]&math.MaxUint32-1].i
	}
	return
}
```
