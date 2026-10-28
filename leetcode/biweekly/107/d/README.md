**核心思路**：把数据按照时间升序排序，问题变成长为 $x+1$ 的**滑动窗口**。

为了不打乱 $\textit{queries}$ 的顺序，可以创建一个下标数组，把下标数组按照 $\textit{queries}[i]$ 的大小排序。

然后就可以滑动窗口了。在从上一个询问到当前询问 $q$ 的过程中：

- $\textit{time}\le q$ 的新增请求进入窗口。
- $\textit{time} < q-x$ 的请求离开窗口。

同时维护每个服务器收到了多少个请求 $\textit{cnt}$，以及没有收到请求的服务器数目 $\textit{notReceives}$（初始值为 $n$）：

- 服务器 $i$ 收到的请求进入窗口，把 $\textit{cnt}[i]$ 增加一，如果加一前 $\textit{cnt}[i] = 0$，那么把 $\textit{notReceives}$ 减少一。
- 服务器 $i$ 收到的请求离开窗口，把 $\textit{cnt}[i]$ 减少一，如果减一后 $\textit{cnt}[i] = 0$，那么把 $\textit{notReceives}$ 增加一。

```py [sol-Python3]
class Solution:
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        logs.sort(key=lambda p: p[1])  # 按照 time 排序
        idx = sorted(range(len(queries)), key=lambda i: queries[i])
        cnt = [0] * (n + 1)  # 每个服务器收到了多少个请求
        not_receives = n  # 没有收到请求的服务器数目
        left = right = 0

        for qid in idx:
            q = queries[qid]
            # 进入窗口
            while right < len(logs) and logs[right][1] <= q:
                i = logs[right][0]
                if cnt[i] == 0:
                    not_receives -= 1
                cnt[i] += 1
                right += 1

            # 离开窗口
            while left < len(logs) and logs[left][1] < q - x:
                i = logs[left][0]
                cnt[i] -= 1
                if cnt[i] == 0:
                    not_receives += 1
                left += 1

            queries[qid] = not_receives

        return queries
```

```java [sol-Java]
class Solution {
    public int[] countServers(int n, int[][] logs, int x, int[] queries) {
        Arrays.sort(logs, (a, b) -> a[1] - b[1]); // 按照 time 排序
        Integer[] idx = new Integer[queries.length];
        Arrays.setAll(idx, i -> i);
        Arrays.sort(idx, (i, j) -> queries[i] - queries[j]);
        int[] cnt = new int[n + 1]; // 每个服务器收到了多少个请求
        int notReceives = n; // 没有收到请求的服务器数目
        int left = 0;
        int right = 0;

        for (int qid : idx) {
            int q = queries[qid];
            // 进入窗口
            while (right < logs.length && logs[right][1] <= q) {
                int i = logs[right][0];
                if (cnt[i] == 0) {
                    notReceives--;
                }
                cnt[i]++;
                right++;
            }

            // 离开窗口
            while (left < logs.length && logs[left][1] < q - x) {
                int i = logs[left][0];
                cnt[i]--;
                if (cnt[i] == 0) {
                    notReceives++;
                }
                left++;
            }

            queries[qid] = notReceives;
        }

        return queries;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> countServers(int n, vector<vector<int>>& logs, int x, vector<int>& queries) {
        ranges::sort(logs, {}, [](auto& p) { return p[1]; }); // 按照 time 排序
        vector<int> idx(queries.size());
        ranges::iota(idx, 0);
        ranges::sort(idx, {}, [&](int i) { return queries[i]; });
        vector<int> cnt(n + 1); // 每个服务器收到了多少个请求
        int not_receives = n; // 没有收到请求的服务器数目
        int left = 0, right = 0;

        for (int qid : idx) {
            int& q = queries[qid];
            // 进入窗口
            while (right < logs.size() && logs[right][1] <= q) {
                int i = logs[right][0];
                if (cnt[i] == 0) {
                    not_receives--;
                }
                cnt[i]++;
                right++;
            }

            // 离开窗口
            while (left < logs.size() && logs[left][1] < q - x) {
                int i = logs[left][0];
                cnt[i]--;
                if (cnt[i] == 0) {
                    not_receives++;
                }
                left++;
            }

            q = not_receives;
        }

        return queries;
    }
};
```

```c [sol-C]
int cmp_logs(const void* p, const void* q) {
    int* a = *(int**)p;
    int* b = *(int**)q;
    return a[1] - b[1]; // 按照 time 排序
}

int* _queries;

int cmp_queries(const void* p, const void* q) {
    return _queries[*(int*)p] - _queries[*(int*)q];
}

int* countServers(int n, int** logs, int logsSize, int* logsColSize, int x, int* queries, int queriesSize, int* returnSize) {
    qsort(logs, logsSize, sizeof(int*), cmp_logs);

    int* idx = malloc(queriesSize * sizeof(int));
    for (int i = 0; i < queriesSize; i++) {
        idx[i] = i;
    }
    _queries = queries;
    qsort(idx, queriesSize, sizeof(int), cmp_queries);

    int* cnt = calloc(n + 1, sizeof(int)); // 每个服务器收到了多少个请求
    int not_receives = n; // 没有收到请求的服务器数目
    int left = 0, right = 0;

    for (int k = 0; k < queriesSize; k++) {
        int qid = idx[k];
        int q = queries[qid];

        // 进入窗口
        while (right < logsSize && logs[right][1] <= q) {
            int i = logs[right][0];
            if (cnt[i] == 0) {
                not_receives--;
            }
            cnt[i]++;
            right++;
        }

        // 离开窗口
        while (left < logsSize && logs[left][1] < q - x) {
            int i = logs[left][0];
            cnt[i]--;
            if (cnt[i] == 0) {
                not_receives++;
            }
            left++;
        }

        queries[qid] = not_receives;
    }

    free(idx);
    free(cnt);
    *returnSize = queriesSize;
    return queries;
}
```

```go [sol-Go]
func countServers(n int, logs [][]int, x int, queries []int) []int {
	slices.SortFunc(logs, func(a, b []int) int { return a[1] - b[1] }) // 按照 time 排序
	idx := make([]int, len(queries))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return queries[i] - queries[j] })
	cnt := make([]int, n+1) // 每个服务器收到了多少个请求
	notReceives := n // 没有收到请求的服务器数目
	left, right := 0, 0

	for _, qid := range idx {
		q := queries[qid]
		// 进入窗口
		for right < len(logs) && logs[right][1] <= q {
			i := logs[right][0]
			if cnt[i] == 0 {
				notReceives--
			}
			cnt[i]++
			right++
		}

		// 离开窗口
		for left < len(logs) && logs[left][1] < q-x {
			i := logs[left][0]
			cnt[i]--
			if cnt[i] == 0 {
				notReceives++
			}
			left++
		}

		queries[qid] = notReceives
	}

	return queries
}
```

```js [sol-JavaScript]
var countServers = function(n, logs, x, queries) {
    logs.sort((a, b) => a[1] - b[1]); // 按照 time 排序
    const idx = Array.from({ length: queries.length }, (_, i) => i);
    idx.sort((i, j) => queries[i] - queries[j]);
    const cnt = Array(n + 1).fill(0); // 每个服务器收到了多少个请求
    let notReceives = n; // 没有收到请求的服务器数目
    let left = 0, right = 0;

    for (const qid of idx) {
        const q = queries[qid];
        // 进入窗口
        while (right < logs.length && logs[right][1] <= q) {
            const i = logs[right][0];
            if (cnt[i] === 0) {
                notReceives--;
            }
            cnt[i]++;
            right++;
        }

        // 离开窗口
        while (left < logs.length && logs[left][1] < q - x) {
            const i = logs[left][0];
            cnt[i]--;
            if (cnt[i] === 0) {
                notReceives++;
            }
            left++;
        }

        queries[qid] = notReceives;
    }

    return queries;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_servers(n: i32, mut logs: Vec<Vec<i32>>, x: i32, mut queries: Vec<i32>) -> Vec<i32> {
        logs.sort_unstable_by_key(|p| p[1]); // 按照 time 排序
        let mut idx = (0..queries.len()).collect::<Vec<_>>();
        idx.sort_unstable_by_key(|&i| queries[i]);
        let mut cnt = vec![0; n as usize + 1]; // 每个服务器收到了多少个请求
        let mut not_receives = n; // 没有收到请求的服务器数目
        let mut left = 0;
        let mut right = 0;

        for qid in idx {
            let q = queries[qid];
            // 进入窗口
            while right < logs.len() && logs[right][1] <= q {
                let i = logs[right][0] as usize;
                if cnt[i] == 0 {
                    not_receives -= 1;
                }
                cnt[i] += 1;
                right += 1;
            }

            // 离开窗口
            while left < logs.len() && logs[left][1] < q - x {
                let i = logs[left][0] as usize;
                cnt[i] -= 1;
                if cnt[i] == 0 {
                    not_receives += 1;
                }
                left += 1;
            }

            queries[qid] = not_receives;
        }

        queries
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m+ q\log q + n)$，其中 $m$ 是 $\textit{logs}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。瓶颈在排序上。创建大小为 $\textit{cnt}$ 数组需要 $\mathcal{O}(n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n+q)$。忽略排序的栈开销。

## 专题训练

1. 滑动窗口与双指针题单的「**五、三指针**」。
2. 数据结构题单的「**专题：离线算法**」

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
