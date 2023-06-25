为方便回答询问，可以把 $\textit{logs}$ 的时间和询问都从小到大排序。对于询问，为了不打乱顺序，可以创建一个下标数组对其排序。

由于询问的窗口大小是固定的，所以可以用**滑动窗口**（双指针）来算，维护窗口内的各个服务器收到了多少次请求 $\textit{cnt}$，以及没有收到请求的服务器数目 $\textit{outOfRange}$。

具体见[【双周赛 107】](https://www.bilibili.com/video/BV1am4y1a7Zi/)第四题讲解，欢迎点赞投币！

```py [sol-Python3]
class Solution:
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        logs.sort(key=lambda p: p[1])  # 按照 time 排序
        ans = [0] * len(queries)
        cnt = [0] * (n + 1)
        out_of_range = n
        left = right = 0
        for qi, q in sorted(enumerate(queries), key=lambda p: p[1]):
            while right < len(logs) and logs[right][1] <= q:  # 进入窗口
                i = logs[right][0]
                if cnt[i] == 0: out_of_range -= 1
                cnt[i] += 1
                right += 1
            while left < len(logs) and logs[left][1] < q - x:  # 离开窗口
                i = logs[left][0]
                cnt[i] -= 1
                if cnt[i] == 0: out_of_range += 1
                left += 1
            ans[qi] = out_of_range
        return ans
```

```java [sol-Java]
class Solution {
    public int[] countServers(int n, int[][] logs, int x, int[] queries) {
        int nq = queries.length;
        var id = new Integer[nq];
        for (int i = 0; i < nq; i++) id[i] = i;
        Arrays.sort(id, (i, j) -> queries[i] - queries[j]);
        Arrays.sort(logs, (a, b) -> a[1] - b[1]); // 按照 time 排序

        int[] ans = new int[nq], cnt = new int[n + 1];
        int outOfRange = n, left = 0, right = 0;
        for (int i : id) {
            while (right < logs.length && logs[right][1] <= queries[i]) // 进入窗口
                if (cnt[logs[right++][0]]++ == 0)
                    outOfRange--;
            while (left < logs.length && logs[left][1] < queries[i] - x) // 离开窗口
                if (--cnt[logs[left++][0]] == 0)
                    outOfRange++;
            ans[i] = outOfRange;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> countServers(int n, vector<vector<int>> &logs, int x, vector<int> &queries) {
        int nq = queries.size(), id[nq], cnt[n + 1];
        memset(cnt, 0, sizeof(cnt));
        iota(id, id + nq, 0);
        sort(id, id + nq, [&](int i, int j) {
            return queries[i] < queries[j];
        });
        sort(logs.begin(), logs.end(), [](const auto &a, const auto &b) {
            return a[1] < b[1]; // 按照 time 排序
        });

        vector<int> ans(nq);
        int out_of_range = n, left = 0, right = 0;
        for (int i: id) {
            while (right < logs.size() && logs[right][1] <= queries[i]) // 进入窗口
                if (cnt[logs[right++][0]]++ == 0)
                    out_of_range--;
            while (left < logs.size() && logs[left][1] < queries[i] - x) // 离开窗口
                if (--cnt[logs[left++][0]] == 0)
                    out_of_range++;
            ans[i] = out_of_range;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countServers(n int, logs [][]int, x int, queries []int) []int {
	type pair struct{ q, i int }
	qs := make([]pair, len(queries))
	for i, q := range queries {
		qs[i] = pair{q, i}
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].q < qs[j].q })
	sort.Slice(logs, func(i, j int) bool { return logs[i][1] < logs[j][1] }) // 按照 time 排序

	ans := make([]int, len(queries))
	cnt := make([]int, n+1)
	outOfRange, left, right := n, 0, 0
	for _, p := range qs {
		for ; right < len(logs) && logs[right][1] <= p.q; right++ { // 进入窗口
			i := logs[right][0]
			if cnt[i] == 0 {
				outOfRange--
			}
			cnt[i]++
		}
		for ; left < len(logs) && logs[left][1] < p.q-x; left++ { // 离开窗口
			i := logs[left][0]
			cnt[i]--
			if cnt[i] == 0 {
				outOfRange++
			}
		}
		ans[p.i] = outOfRange
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m+ q\log q)$，其中 $m$ 为 $\textit{logs}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。用数组统计出现次数，复杂度要加 $n$。
- 空间复杂度：$\mathcal{O}(n+q)$。

#### 思考题

如果询问的区间长度不固定，要怎么做？也就是输入的询问是 $[\textit{left}_i, \textit{right}_i]$。
