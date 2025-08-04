## 莫队算法的原理

暴力的想法是，对于每个询问，遍历下标在 $[l,r]$ 中的元素，统计元素的出现次数，维护元素出现次数的最大值，及其对应元素的最小值。

考虑这样两个下标区间：$[1,999]$ 和 $[2,1000]$。这两个区间有大量公共元素，难道都要完整地遍历一遍吗？

借鉴 [定长滑动窗口](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/) 或者 [不定长滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 的思想，我们只需额外增删不在交集中的元素。

莫队算法，本质是通过巧妙地改变回答询问的顺序，使访问的元素个数（区间左右端点的总移动次数）由 $\mathcal{O}(nq)$ 降至 $\mathcal{O}(n\sqrt q)$，其中 $q$ 是 $\textit{queries}$ 的长度。

把数组分成若干块，每块大小为 $B$，分成 $k=\dfrac{n}{B}$ 块。最后一块的大小 $\le B$。

莫队算法的核心思想是：**把询问按照左端点所在块分组，左端点在同一个块的询问分到同一组**。

对于每个块，把右端点排序，使得**右端点在数组中一直向右移动**（对于本题来说），而**左端点只在块内「抖动」**。如此一来，两个相邻询问区间的交集就能尽量大，就能减少增删元素的次数。

对于每个块，右端点的平均总移动次数约为 $\dfrac{n}{2}$。对于所有询问，右端点的总移动次数约为 $\dfrac{nk}{2} = \dfrac{n^2}{2B}$。

对于每个询问，左端点由于只在块内移动，移动次数不超过 $2B$，左端点的总移动次数不超过 $2qB$。注意这里乘了 $2$，是因为本题有回滚操作，每个询问左端点都要来回移动一遍。

相加得

$$
\frac{n^2}{2B} + 2qB
$$

由基本不等式可知，当

$$
B=\frac{n}{\sqrt q}
$$

时取到最小值

$$
n\sqrt q
$$

为保证 $B$ 为正整数，实际取

$$
B = \left\lceil\frac{n}{\sqrt q}\right\rceil
$$

如此一来，对于 $\textit{nums}$ 元素的总访问次数，就从暴力算法的 $\mathcal{O}(nq)$，降低至 $\mathcal{O}(n\sqrt q)$，足以通过本题。

这便是莫队算法，仅仅改变回答询问的顺序，就能加快算法的效率。

## 细节

添加元素 $x$ 是好维护的：

- 用一个哈希表 $\textit{cnt}$ 维护元素的出现次数。
- 把 $\textit{cnt}[x]$ 加一。
- 如果 $c=\textit{cnt}[x]$ 大于出现次数的最大值 $\textit{maxCnt}$，那么更新 $\textit{maxCnt}=c$，同时更新对应元素的最小值 $\textit{minVal}=x$。
- 如果 $c = \textit{maxCnt}$，只更新 $\textit{minVal}$ 为 $\min(\textit{minVal},x)$。

然而，删除元素是不好维护的，如果我们刚好把 $\textit{minVal}$ 的出现次数减少后，谁作为新的 $\textit{minVal}$ 呢？这可以用有序集合/懒删除堆维护。

但其实，并不需要这些数据结构。

如果保证只有「添加元素」，没有「删除元素」，不就搞定了吗？

具体来说，对于每个询问，如果区间长度 $\le B$，直接暴力算。如果区间长度 $> B$，那么左右端点必然不在同一个块中，我们可以把这些询问用莫队分块，然后回答每个块中的询问：

- 右端点 $r$ 从块的右端点加一处开始，**向右移动**。
- 左端点 $l$ 从块的右端点开始，**向左移动**。
- 这样只有添加元素。但是，在回答完一个询问后，要把 $l$ **回滚**到块的右端点处，从而保证回答每个询问时，$l$ 都是从块的右端点开始的。
- 怎么回滚？
   - 对于元素出现次数的最大值 $\textit{maxCnt}$ 及其对应元素的最小值 $\textit{minVal}$，用两个临时变量记录。回答完询问后，直接回滚为临时变量的值。
   - 对于元素出现次数，直接减少即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1p3h3zYEbc/)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def subarrayMajority(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n, m = len(nums), len(queries)

        cnt = defaultdict(int)
        max_cnt = min_val = 0

        # 添加元素 x
        def add(x: int) -> None:
            nonlocal max_cnt, min_val
            cnt[x] += 1
            c = cnt[x]
            if c > max_cnt:
                max_cnt, min_val = c, x
            elif c == max_cnt:
                min_val = min(min_val, x)

        ans = [-1] * m
        block_size = ceil(n / sqrt(m))

        qs = []  # (bid, ql, qr, threshold, qid) 其中 bid 是块号，qid 是询问的编号
        for i, (l, r, threshold) in enumerate(queries):
            r += 1  # 左闭右开

            # 大区间离线（保证 l 和 r 不在同一个块中）
            if r - l > block_size:
                qs.append((l // block_size, l, r, threshold, i))
                continue

            # 小区间暴力
            for x in nums[l: r]:
                add(x)
            if max_cnt >= threshold:
                ans[i] = min_val

            # 重置数据
            cnt.clear()
            max_cnt = 0

        qs.sort(key=lambda q: (q[0], q[2]))

        for i, (bid, ql, qr, threshold, qid) in enumerate(qs):
            l0 = (bid + 1) * block_size
            if i == 0 or bid > qs[i - 1][0]:  # 遍历到一个新的块
                r = l0  # 右端点移动的起点
                # 重置数据
                cnt.clear()
                max_cnt = 0

            # 右端点从 r 移动到 qr（qr 不计入）
            while r < qr:
                add(nums[r])
                r += 1

            tmp_max_cnt, tmp_min_val = max_cnt, min_val

            # 左端点从 l0 移动到 ql（l0 不计入）
            for x in nums[ql: l0]:
                add(x)
            if max_cnt >= threshold:
                ans[qid] = min_val

            # 回滚
            max_cnt, min_val = tmp_max_cnt, tmp_min_val
            for x in nums[ql: l0]:
                cnt[x] -= 1

        return ans
```

```java [sol-Java]
class Solution {
    public int[] subarrayMajority(int[] nums, int[][] queries) {
        int n = nums.length;
        int m = queries.length;
        int[] ans = new int[m];
        int blockSize = (int) Math.ceil(n / Math.sqrt(m));

        record Query(int bid, int l, int r, int threshold, int qid) { // [l,r) 左闭右开
        }

        List<Query> qs = new ArrayList<>();
        for (int i = 0; i < m; i++) {
            int[] q = queries[i];
            int l = q[0];
            int r = q[1] + 1; // 左闭右开
            int threshold = q[2];

            // 大区间离线（保证 l 和 r 不在同一个块中）
            if (r - l > blockSize) {
                qs.add(new Query(l / blockSize, l, r, threshold, i));
                continue;
            }

            // 小区间暴力
            for (int j = l; j < r; j++) {
                add(nums[j]);
            }
            ans[i] = maxCnt >= threshold ? minVal : -1;

            // 重置数据
            cnt.clear();
            maxCnt = 0;
        }

        qs.sort((a, b) -> a.bid != b.bid ? a.bid - b.bid : a.r - b.r);

        int r = 0;
        for (int i = 0; i < qs.size(); i++) {
            Query q = qs.get(i);
            int l0 = (q.bid + 1) * blockSize;
            if (i == 0 || q.bid > qs.get(i - 1).bid) { // 遍历到一个新的块
                r = l0; // 右端点移动的起点
                // 重置数据
                cnt.clear();
                maxCnt = 0;
            }

            // 右端点从 r 移动到 q.r（q.r 不计入）
            for (; r < q.r; r++) {
                add(nums[r]);
            }

            int tmpMaxCnt = maxCnt;
            int tmpMinVal = minVal;

            // 左端点从 l0 移动到 q.l（l0 不计入）
            for (int j = q.l; j < l0; j++) {
                add(nums[j]);
            }
            ans[q.qid] = maxCnt >= q.threshold ? minVal : -1;

            // 回滚
            maxCnt = tmpMaxCnt;
            minVal = tmpMinVal;
            for (int j = q.l; j < l0; j++) {
                cnt.merge(nums[j], -1, Integer::sum);
            }
        }

        return ans;
    }

    private final Map<Integer, Integer> cnt = new HashMap<>();
    private int maxCnt = 0;
    private int minVal = 0;

    // 添加元素 x
    private void add(int x) {
        int c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
        if (c > maxCnt) {
            maxCnt = c;
            minVal = x;
        } else if (c == maxCnt) {
            minVal = Math.min(minVal, x);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> subarrayMajority(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size(), m = queries.size();

        unordered_map<int, int> cnt;
        int max_cnt = 0, min_val = 0;

        // 添加元素 x
        auto add = [&](int x) {
            int c = ++cnt[x];
            if (c > max_cnt) {
                max_cnt = c;
                min_val = x;
            } else if (c == max_cnt) {
                min_val = min(min_val, x);
            }
        };

        vector<int> ans(m, -1);
        int block_size = ceil(n / sqrt(m));

        struct Query {
            int bid, l, r, threshold, qid; // [l,r) 左闭右开
        };

        vector<Query> qs;
        for (int i = 0; i < m; i++) {
            auto& q = queries[i];
            int l = q[0], r = q[1] + 1, threshold = q[2]; // 左闭右开

            // 大区间离线（保证 l 和 r 不在同一个块中）
            if (r - l > block_size) {
                qs.emplace_back(l / block_size, l, r, threshold, i);
                continue;
            }

            // 小区间暴力
            for (int j = l; j < r; j++) {
                add(nums[j]);
            }
            if (max_cnt >= threshold) {
                ans[i] = min_val;
            }

            // 重置数据
            cnt.clear();
            max_cnt = 0;
        }

        ranges::sort(qs, {}, [](auto& q) { return pair(q.bid, q.r); });

        int r;
        for (int i = 0; i < qs.size(); i++) {
            auto& q = qs[i];
            int l0 = (q.bid + 1) * block_size;
            if (i == 0 || q.bid > qs[i - 1].bid) { // 遍历到一个新的块
                r = l0; // 右端点移动的起点
                // 重置数据
                cnt.clear();
                max_cnt = 0;
            }

            // 右端点从 r 移动到 q.r（q.r 不计入）
            for (; r < q.r; r++) {
                add(nums[r]);
            }

            int tmp_max_cnt = max_cnt, tmp_min_val = min_val;

            // 左端点从 l0 移动到 q.l（l0 不计入）
            for (int j = q.l; j < l0; j++) {
                add(nums[j]);
            }
            if (max_cnt >= q.threshold) {
                ans[q.qid] = min_val;
            }

            // 回滚
            max_cnt = tmp_max_cnt;
            min_val = tmp_min_val;
            for (int j = q.l; j < l0; j++) {
                cnt[nums[j]]--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func subarrayMajority(nums []int, queries [][]int) []int {
	n, m := len(nums), len(queries)

	cnt := map[int]int{}
	maxCnt, minVal := 0, 0
	// 添加元素 x
	add := func(x int) {
		cnt[x]++
		c := cnt[x]
		if c > maxCnt {
			maxCnt, minVal = c, x
		} else if c == maxCnt {
			minVal = min(minVal, x)
		}
	}

	ans := make([]int, m)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
	type query struct{ bid, l, r, threshold, qid int } // [l,r) 左闭右开
	qs := []query{}
	for i, q := range queries {
		l, r, threshold := q[0], q[1]+1, q[2] // 左闭右开

		// 大区间离线（保证 l 和 r 不在同一个块中）
		if r-l > blockSize {
			qs = append(qs, query{l / blockSize, l, r, threshold, i})
			continue
		}

		// 小区间暴力
		for _, x := range nums[l:r] {
			add(x)
		}
		if maxCnt >= threshold {
			ans[i] = minVal
		} else {
			ans[i] = -1
		}

		// 重置数据
		clear(cnt)
		maxCnt = 0
	}

	slices.SortFunc(qs, func(a, b query) int { return cmp.Or(a.bid-b.bid, a.r-b.r) })

	var r int
	for i, q := range qs {
		l0 := (q.bid + 1) * blockSize
		if i == 0 || q.bid > qs[i-1].bid { // 遍历到一个新的块
			r = l0 // 右端点移动的起点
			// 重置数据
			clear(cnt)
			maxCnt = 0
		}

		// 右端点从 r 移动到 q.r（q.r 不计入）
		for ; r < q.r; r++ {
			add(nums[r])
		}

		tmpMaxCnt, tmpMinVal := maxCnt, minVal

		// 左端点从 l0 移动到 q.l（l0 不计入）
		for _, x := range nums[q.l:l0] {
			add(x)
		}
		if maxCnt >= q.threshold {
			ans[q.qid] = minVal
		} else {
			ans[q.qid] = -1
		}

		// 回滚
		maxCnt, minVal = tmpMaxCnt, tmpMinVal
		for _, x := range nums[q.l:l0] {
			cnt[x]--
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q\log q + n\sqrt q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 优化

把 $\textit{nums}$ 离散化，这样可以用数组代替哈希表，效率更高。

```py [sol-Python3]
class Solution:
    def subarrayMajority(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n, m = len(nums), len(queries)

        # 离散化
        a = sorted(set(nums))
        index_to_value = [bisect_left(a, x) for x in nums]

        cnt = [0] * (len(a) + 1)
        max_cnt = min_val = 0

        def add(i: int) -> None:
            nonlocal max_cnt, min_val
            v = index_to_value[i]
            cnt[v] += 1
            c = cnt[v]
            x = nums[i]
            if c > max_cnt:
                max_cnt, min_val = c, x
            elif c == max_cnt:
                min_val = min(min_val, x)

        ans = [-1] * m
        block_size = ceil(n / sqrt(m))

        qs = []  # (bid, ql, qr, threshold, qid) 其中 bid 是块号，qid 是询问的编号
        for i, (l, r, threshold) in enumerate(queries):
            r += 1  # 左闭右开

            # 大区间离线（保证 l 和 r 不在同一个块中）
            if r - l > block_size:
                qs.append((l // block_size, l, r, threshold, i))
                continue

            # 小区间暴力
            for j in range(l, r):
                add(j)
            if max_cnt >= threshold:
                ans[i] = min_val

            # 重置数据
            for v in index_to_value[l: r]:
                cnt[v] -= 1
            max_cnt = 0

        qs.sort(key=lambda q: (q[0], q[2]))

        for i, (bid, ql, qr, threshold, qid) in enumerate(qs):
            l0 = (bid + 1) * block_size
            if i == 0 or bid > qs[i - 1][0]:  # 遍历到一个新的块
                r = l0  # 右端点移动的起点
                # 重置数据
                cnt = [0] * (len(a) + 1)
                max_cnt = 0

            # 右端点从 r 移动到 qr（qr 不计入）
            while r < qr:
                add(r)
                r += 1

            tmp_max_cnt, tmp_min_val = max_cnt, min_val

            # 左端点从 l0 移动到 ql（l0 不计入）
            for j in range(ql, l0):
                add(j)
            if max_cnt >= threshold:
                ans[qid] = min_val

            # 回滚
            max_cnt, min_val = tmp_max_cnt, tmp_min_val
            for v in index_to_value[ql: l0]:
                cnt[v] -= 1

        return ans
```

```java [sol-Java]
class Solution {
    public int[] subarrayMajority(int[] nums, int[][] queries) {
        int n = nums.length;
        int m = queries.length;
        this.nums = nums;
        cnt = new int[n + 1];

        // 离散化
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        indexToValue = new int[n];
        for (int i = 0; i < n; i++) {
            indexToValue[i] = Arrays.binarySearch(sorted, nums[i]);
        }

        int[] ans = new int[m];
        int blockSize = (int) Math.ceil(n / Math.sqrt(m));

        record Query(int bid, int l, int r, int threshold, int qid) { // [l,r) 左闭右开
        }

        List<Query> qs = new ArrayList<>();
        for (int i = 0; i < m; i++) {
            int[] q = queries[i];
            int l = q[0];
            int r = q[1] + 1; // 左闭右开
            int threshold = q[2];
            
            // 大区间离线（保证 l 和 r 不在同一个块中）
            if (r - l > blockSize) {
                qs.add(new Query(l / blockSize, l, r, threshold, i));
                continue;
            }

            // 小区间暴力
            for (int j = l; j < r; j++) {
                add(j);
            }
            ans[i] = maxCnt >= threshold ? minVal : -1;

            // 重置数据
            for (int j = l; j < r; j++) {
                cnt[indexToValue[j]]--;
            }
            maxCnt = 0;
        }

        qs.sort((a, b) -> a.bid != b.bid ? a.bid - b.bid : a.r - b.r);

        int r = 0;
        for (int i = 0; i < qs.size(); i++) {
            Query q = qs.get(i);
            int l0 = (q.bid + 1) * blockSize;
            if (i == 0 || q.bid > qs.get(i - 1).bid) { // 遍历到一个新的块
                r = l0; // 右端点移动的起点
                // 重置数据
                Arrays.fill(cnt, 0);
                maxCnt = 0;
            }

            // 右端点从 r 移动到 q.r（q.r 不计入）
            for (; r < q.r; r++) {
                add(r);
            }

            int tmpMaxCnt = maxCnt;
            int tmpMinVal = minVal;

            // 左端点从 l0 移动到 q.l（l0 不计入）
            for (int j = q.l; j < l0; j++) {
                add(j);
            }
            ans[q.qid] = maxCnt >= q.threshold ? minVal : -1;

            // 回滚
            maxCnt = tmpMaxCnt;
            minVal = tmpMinVal;
            for (int j = q.l; j < l0; j++) {
                cnt[indexToValue[j]]--;
            }
        }

        return ans;
    }

    private int[] nums;
    private int[] indexToValue;
    private int[] cnt;
    private int maxCnt = 0;
    private int minVal = 0;

    // 添加元素 x
    private void add(int i) {
        int v = indexToValue[i];
        int c = ++cnt[v];
        int x = nums[i];
        if (c > maxCnt) {
            maxCnt = c;
            minVal = x;
        } else if (c == maxCnt) {
            minVal = Math.min(minVal, x);
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> subarrayMajority(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size(), m = queries.size();

        // 离散化
        auto a = nums;
        ranges::sort(a);
        a.erase(ranges::unique(a).begin(), a.end());
        vector<int> index_to_value(n);
        for (int i = 0; i < n; i++) {
            index_to_value[i] = ranges::lower_bound(a, nums[i]) - a.begin();
        }

        vector<int> cnt(a.size() + 1);
        int max_cnt = 0, min_val = 0;

        auto add = [&](int i) {
            int v = index_to_value[i];
            int c = ++cnt[v];
            int x = nums[i];
            if (c > max_cnt) {
                max_cnt = c;
                min_val = x;
            } else if (c == max_cnt) {
                min_val = min(min_val, x);
            }
        };

        vector<int> ans(m, -1);
        int block_size = ceil(n / sqrt(m));

        struct Query {
            int bid, l, r, threshold, qid; // [l,r) 左闭右开
        };

        vector<Query> qs;
        for (int i = 0; i < m; i++) {
            auto& q = queries[i];
            int l = q[0], r = q[1] + 1, threshold = q[2]; // 左闭右开

            // 大区间离线（保证 l 和 r 不在同一个块中）
            if (r - l > block_size) {
                qs.emplace_back(l / block_size, l, r, threshold, i);
                continue;
            }

            // 小区间暴力
            for (int j = l; j < r; j++) {
                add(j);
            }
            if (max_cnt >= threshold) {
                ans[i] = min_val;
            }

            // 重置数据
            for (int j = l; j < r; j++) {
                cnt[index_to_value[j]]--;
            }
            max_cnt = 0;
        }

        ranges::sort(qs, {}, [](auto& q) { return pair(q.bid, q.r); });

        int r;
        for (int i = 0; i < qs.size(); i++) {
            auto& q = qs[i];
            int l0 = (q.bid + 1) * block_size;
            if (i == 0 || q.bid > qs[i - 1].bid) { // 遍历到一个新的块
                r = l0; // 右端点移动的起点
                // 重置数据
                ranges::fill(cnt, 0);
                max_cnt = 0;
            }

            // 右端点从 r 移动到 q.r（q.r 不计入）
            for (; r < q.r; r++) {
                add(r);
            }

            int tmp_max_cnt = max_cnt, tmp_min_val = min_val;

            // 左端点从 l0 移动到 q.l（l0 不计入）
            for (int j = q.l; j < l0; j++) {
                add(j);
            }
            if (max_cnt >= q.threshold) {
                ans[q.qid] = min_val;
            }

            // 回滚
            max_cnt = tmp_max_cnt;
            min_val = tmp_min_val;
            for (int j = q.l; j < l0; j++) {
                cnt[index_to_value[j]]--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func subarrayMajority(nums []int, queries [][]int) []int {
	n, m := len(nums), len(queries)

	a := slices.Clone(nums)
	slices.Sort(a)
	a = slices.Compact(a)
	indexToValue := make([]int, n)
	for i, x := range nums {
		indexToValue[i] = sort.SearchInts(a, x)
	}

	cnt := make([]int, len(a)+1)
	maxCnt, minVal := 0, 0
	add := func(i int) {
		v := indexToValue[i]
		cnt[v]++
		c := cnt[v]
		x := nums[i]
		if c > maxCnt {
			maxCnt, minVal = c, x
		} else if c == maxCnt {
			minVal = min(minVal, x)
		}
	}

	ans := make([]int, m)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(m))))
	type query struct{ bid, l, r, threshold, qid int } // [l,r) 左闭右开
	qs := []query{}
	for i, q := range queries {
		l, r, threshold := q[0], q[1]+1, q[2] // 左闭右开

		// 大区间离线（保证 l 和 r 不在同一个块中）
		if r-l > blockSize {
			qs = append(qs, query{l / blockSize, l, r, threshold, i})
			continue
		}

		// 小区间暴力
		for j := l; j < r; j++ {
			add(j)
		}
		if maxCnt >= threshold {
			ans[i] = minVal
		} else {
			ans[i] = -1
		}

		// 重置数据
		for _, v := range indexToValue[l:r] {
			cnt[v]--
		}
		maxCnt = 0
	}

	slices.SortFunc(qs, func(a, b query) int { return cmp.Or(a.bid-b.bid, a.r-b.r) })

	var r int
	for i, q := range qs {
		l0 := (q.bid + 1) * blockSize
		if i == 0 || q.bid > qs[i-1].bid { // 遍历到一个新的块
			r = l0 // 右端点移动的起点
			// 重置数据
			clear(cnt)
			maxCnt = 0
		}

		// 右端点从 r 移动到 q.r（q.r 不计入）
		for ; r < q.r; r++ {
			add(r)
		}

		tmpMaxCnt, tmpMinVal := maxCnt, minVal

		// 左端点从 l0 移动到 q.l（l0 不计入）
		for l := q.l; l < l0; l++ {
			add(l)
		}
		if maxCnt >= q.threshold {
			ans[q.qid] = minVal
		} else {
			ans[q.qid] = -1
		}

		// 回滚
		maxCnt, minVal = tmpMaxCnt, tmpMinVal
		for _, v := range indexToValue[q.l:l0] {
			cnt[v]--
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + q\log q + n\sqrt q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

推荐做做 [P5906 【模板】回滚莫队&不删除莫队](https://www.luogu.com.cn/problem/P5906) 加深对回滚莫队的理解。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
