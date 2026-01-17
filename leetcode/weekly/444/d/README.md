为了快速模拟题目的操作，我们需要维护三种信息：

1. 把相邻元素和 $s$，以及相邻元素中的左边元素的下标 $i$，组成一个 pair $(s,i)$。我们需要添加 pair、删除 pair 以及查询这些 pair 的最小值（双关键字比较），这可以用**有序集合**，或者**懒删除堆**。
2. 维护剩余下标。我们需要查询每个下标 $i$ 左侧最近剩余下标，以及右侧最近剩余下标。这可以用**有序集合**，或者**两个并查集**，或者**双向链表**。
3. 在相邻元素中，满足「左边元素大于右边元素」的个数，记作 $\textit{dec}$。

不断模拟操作，直到 $\textit{dec} = 0$。

题目说「用它们的和替换这对元素」，设操作的这对元素的下标为 $i$ 和 $\textit{nxt}$，操作相当于把 $\textit{nums}[i]$ 增加 $\textit{nums}[\textit{nxt}]$，然后删除 $\textit{nums}[\textit{nxt}]$。

在这个过程中，$\textit{dec}$ 如何变化？

设操作的这对元素的下标为 $i$ 和 $\textit{nxt}$，$i$ 左侧最近剩余下标为 $\textit{pre}$，$\textit{nxt}$ 右侧最近剩余下标为 $\textit{nxt}_2$。

操作会影响 $\textit{nums}[i]$ 和 $\textit{nums}[\textit{nxt}]$，也会影响周边相邻元素的大小关系。所以**每次操作，我们需要重新考察 $4$ 个元素值的大小关系**，下标从左到右为 $\textit{pre},i,\textit{nxt},\textit{nxt}_2$。

1. 删除 $\textit{nums}[\textit{nxt}]$。如果删除前 $\textit{nums}[i] > \textit{nums}[\textit{nxt}]$，把 $\textit{dec}$ 减一。
2. 如果删除前 $\textit{nums}[\textit{pre}] > \textit{nums}[i]$，把 $\textit{dec}$ 减一。如果删除后 $\textit{nums}[\textit{pre}] > s$，把 $\textit{dec}$ 加一。这里 $s$ 表示操作的这对元素之和，也就是新的 $\textit{nums}[i]$ 的值。
3. 如果删除前 $\textit{nums}[\textit{nxt}] > \textit{nums}[\textit{nxt}_2]$，把 $\textit{dec}$ 减一。删除后 $i$ 和 $\textit{nxt}_2$ 相邻，如果删除后 $s > \textit{nums}[\textit{nxt}_2]$，把 $\textit{dec}$ 加一。

上述过程中，同时维护（添加删除）新旧相邻元素和以及下标。

[本题视频讲解](https://www.bilibili.com/video/BV1ezRvYiE27/?t=41m12s)，欢迎点赞关注~

## 写法一：两个有序集合

```py [sol-Python3]
class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        sl = SortedList()  # (相邻元素和，左边那个数的下标)
        idx = SortedList(range(len(nums)))  # 剩余下标
        dec = 0  # 递减的相邻对的个数

        for i, (x, y) in enumerate(pairwise(nums)):
            if x > y:
                dec += 1
            sl.add((x + y, i))

        ans = 0
        while dec > 0:
            ans += 1

            s, i = sl.pop(0)  # 删除相邻元素和最小的一对
            k = idx.bisect_left(i)

            # (当前元素，下一个数)
            nxt = idx[k + 1]
            if nums[i] > nums[nxt]:  # 旧数据
                dec -= 1

            # (前一个数，当前元素)
            if k > 0:
                pre = idx[k - 1]
                if nums[pre] > nums[i]:  # 旧数据
                    dec -= 1
                if nums[pre] > s:  # 新数据
                    dec += 1
                sl.remove((nums[pre] + nums[i], pre))
                sl.add((nums[pre] + s, pre))

            # (下一个数，下下一个数)
            if k + 2 < len(idx):
                nxt2 = idx[k + 2]
                if nums[nxt] > nums[nxt2]:  # 旧数据
                    dec -= 1
                if s > nums[nxt2]:  # 新数据（当前元素，下下一个数）
                    dec += 1
                sl.remove((nums[nxt] + nums[nxt2], nxt))
                sl.add((s + nums[nxt2], i))

            nums[i] = s  # 把 nums[nxt] 加到 nums[i] 中
            idx.remove(nxt)  # 删除 nxt

        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(long s, int i) {
    }

    public int minimumPairRemoval(int[] nums) {
        int n = nums.length;
        // (相邻元素和，左边那个数的下标)
        TreeSet<Pair> pairs = new TreeSet<>((a, b) -> a.s != b.s ? Long.compare(a.s, b.s) : a.i - b.i);
        int dec = 0; // 递减的相邻对的个数
        for (int i = 0; i < n - 1; i++) {
            int x = nums[i];
            int y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            pairs.add(new Pair(x + y, i));
        }

        // 剩余下标
        TreeSet<Integer> idx = new TreeSet<>();
        for (int i = 0; i < n; i++) {
            idx.add(i);
        }

        long[] a = new long[n];
        for (int i = 0; i < n; i++) {
            a[i] = nums[i];
        }

        int ans = 0;
        while (dec > 0) {
            ans++;

            // 删除相邻元素和最小的一对
            Pair p = pairs.pollFirst();
            long s = p.s;
            int i = p.i;

            // (当前元素，下一个数)
            int nxt = idx.higher(i);
            if (a[i] > a[nxt]) { // 旧数据
                dec--;
            }

            // (前一个数，当前元素)
            Integer pre = idx.lower(i);
            if (pre != null) {
                if (a[pre] > a[i]) { // 旧数据
                    dec--;
                }
                if (a[pre] > s) { // 新数据
                    dec++;
                }
                pairs.remove(new Pair(a[pre] + a[i], pre));
                pairs.add(new Pair(a[pre] + s, pre));
            }

            // (下一个数，下下一个数)
            Integer nxt2 = idx.higher(nxt);
            if (nxt2 != null) {
                if (a[nxt] > a[nxt2]) { // 旧数据
                    dec--;
                }
                if (s > a[nxt2]) { // 新数据（当前元素，下下一个数）
                    dec++;
                }
                pairs.remove(new Pair(a[nxt] + a[nxt2], nxt));
                pairs.add(new Pair(s + a[nxt2], i));
            }

            a[i] = s; // 把 a[nxt] 加到 a[i] 中
            idx.remove(nxt); // 删除 nxt
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumPairRemoval(vector<int>& nums) {
        int n = nums.size();
        set<pair<long long, int>> pairs; // (相邻元素和，左边那个数的下标)
        int dec = 0; // 递减的相邻对的个数
        for (int i = 0; i + 1 < n; i++) {
            int x = nums[i], y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            pairs.emplace(x + y, i);
        }

        set<int> idx; // 剩余下标
        for (int i = 0; i < n; i++) {
            idx.insert(i);
        }

        vector<long long> a(nums.begin(), nums.end());
        int ans = 0;
        while (dec > 0) {
            ans++;

            // 删除相邻元素和最小的一对
            auto [s, i] = *pairs.begin();
            pairs.erase(pairs.begin());

            auto it = idx.lower_bound(i);

            // (当前元素，下一个数)
            auto nxt_it = next(it);
            int nxt = *nxt_it;
            dec -= a[i] > a[nxt]; // 旧数据

            // (前一个数，当前元素)
            if (it != idx.begin()) {
                int pre = *prev(it);
                dec -= a[pre] > a[i]; // 旧数据
                dec += a[pre] > s; // 新数据
                pairs.erase({a[pre] + a[i], pre});
                pairs.emplace(a[pre] + s, pre);
            }

            // (下一个数，下下一个数)
            auto nxt2_it = next(nxt_it);
            if (nxt2_it != idx.end()) {
                int nxt2 = *nxt2_it;
                dec -= a[nxt] > a[nxt2]; // 旧数据
                dec += s > a[nxt2]; // 新数据（当前元素，下下一个数）
                pairs.erase({a[nxt] + a[nxt2], nxt});
                pairs.emplace(s + a[nxt2], i);
            }

            a[i] = s; // 把 a[nxt] 加到 a[i] 中
            idx.erase(nxt); // 删除 nxt
        }
        return ans;
    }
};
```

```go [sol-Go]
// import "github.com/emirpasic/gods/v2/trees/redblacktree"
func minimumPairRemoval(nums []int) (ans int) {
	n := len(nums)
	type pair struct{ s, i int }
	// (相邻元素和，左边那个数的下标)
	pairs := redblacktree.NewWith[pair, struct{}](func(a, b pair) int { return cmp.Or(a.s-b.s, a.i-b.i) })
	dec := 0 // 递减的相邻对的个数
	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		pairs.Put(pair{x + y, i}, struct{}{})
	}

	// 剩余下标
	idx := redblacktree.New[int, struct{}]()
	for i := range n {
		idx.Put(i, struct{}{})
	}

	for dec > 0 {
		ans++

		it := pairs.Left()
		s := it.Key.s
		i := it.Key.i
		pairs.Remove(it.Key) // 删除相邻元素和最小的一对

		// (当前元素，下一个数)
		node, _ := idx.Ceiling(i + 1)
		nxt := node.Key
		if nums[i] > nums[nxt] { // 旧数据
			dec--
		}

		// (前一个数，当前元素)
		node, _ = idx.Floor(i - 1)
		if node != nil {
			pre := node.Key
			if nums[pre] > nums[i] { // 旧数据
				dec--
			}
			if nums[pre] > s { // 新数据
				dec++
			}
			pairs.Remove(pair{nums[pre] + nums[i], pre})
			pairs.Put(pair{nums[pre] + s, pre}, struct{}{})
		}

		// (下一个数，下下一个数)
		node, _ = idx.Ceiling(nxt + 1)
		if node != nil {
			nxt2 := node.Key
			if nums[nxt] > nums[nxt2] { // 旧数据
				dec--
			}
			if s > nums[nxt2] { // 新数据（当前元素，下下一个数）
				dec++
			}
			pairs.Remove(pair{nums[nxt] + nums[nxt2], nxt})
			pairs.Put(pair{s + nums[nxt2], i}, struct{}{})
		}

		nums[i] = s // 把 nums[nxt] 加到 nums[i] 中
		idx.Remove(nxt) // 删除 nxt
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：懒删除堆 + 两个数组模拟双向链表

用最小堆（懒删除堆）代替维护 pair 的有序集合。

用双向链表代替维护下标的有序集合。进一步地，可以用两个数组模拟双向链表的 $\textit{prev}$ 指针和 $\textit{next}$ 指针。

如果堆顶下标 $i$ 被删除，或者 $i$ 右边下标 $\textit{nxt}$ 被删除，或者堆顶元素和不等于 $\textit{nums}[i]+\textit{nums}[\textit{nxt}]$，则弹出堆顶。

```py [sol-Python3]
class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        n = len(nums)
        h = []  # (相邻元素和，左边那个数的下标)
        dec = 0  # 递减的相邻对的个数
        for i, (x, y) in enumerate(pairwise(nums)):
            if x > y:
                dec += 1
            h.append((x + y, i))
        heapify(h)
        lazy = defaultdict(int)

        # 每个下标的左右最近的未删除下标
        left = list(range(-1, n))  # 加一个哨兵，防止下标越界
        right = list(range(1, n + 1))

        ans = 0
        while dec:
            ans += 1

            while lazy[h[0]]:
                lazy[heappop(h)] -= 1
            s, i = heappop(h)  # 删除相邻元素和最小的一对

            # (当前元素，下一个数)
            nxt = right[i]
            if nums[i] > nums[nxt]:  # 旧数据
                dec -= 1

            # (前一个数，当前元素)
            pre = left[i]
            if pre >= 0:
                if nums[pre] > nums[i]:  # 旧数据
                    dec -= 1
                if nums[pre] > s:  # 新数据
                    dec += 1
                lazy[(nums[pre] + nums[i], pre)] += 1  # 懒删除
                heappush(h, (nums[pre] + s, pre))

            # (下一个数，下下一个数)
            nxt2 = right[nxt]
            if nxt2 < n:
                if nums[nxt] > nums[nxt2]:  # 旧数据
                    dec -= 1
                if s > nums[nxt2]:  # 新数据（当前元素，下下一个数）
                    dec += 1
                lazy[(nums[nxt] + nums[nxt2], nxt)] += 1  # 懒删除
                heappush(h, (s + nums[nxt2], i))

            nums[i] = s
            # 删除 nxt
            l, r = left[nxt], right[nxt]
            right[l] = r  # 模拟双向链表的删除操作
            left[r] = l

        return ans
```

```py [sol-Python3 不用 lazy]
class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        n = len(nums)
        h = []  # (相邻元素和，左边那个数的下标)
        dec = 0  # 递减的相邻对的个数
        for i, (x, y) in enumerate(pairwise(nums)):
            if x > y:
                dec += 1
            h.append((x + y, i))
        heapify(h)

        # 每个下标的左右最近的未删除下标
        left = list(range(-1, n))  # 加一个哨兵，防止下标越界
        right = list(range(1, n + 1))  # 注意最下面的代码，删除 nxt 的时候额外把 right[nxt] 置为 n

        ans = 0
        while dec:
            ans += 1

            # 如果堆顶数据与实际数据不符，说明堆顶数据是之前本应删除，但没有删除的数据（懒删除）
            while right[h[0][1]] >= n or h[0][0] != nums[h[0][1]] + nums[right[h[0][1]]]:
                heappop(h)
            s, i = heappop(h)  # 删除相邻元素和最小的一对

            # (当前元素，下一个数)
            nxt = right[i]
            if nums[i] > nums[nxt]:  # 旧数据
                dec -= 1

            # (前一个数，当前元素)
            pre = left[i]
            if pre >= 0:
                if nums[pre] > nums[i]:  # 旧数据
                    dec -= 1
                if nums[pre] > s:  # 新数据
                    dec += 1
                heappush(h, (nums[pre] + s, pre))

            # (下一个数，下下一个数)
            nxt2 = right[nxt]
            if nxt2 < n:
                if nums[nxt] > nums[nxt2]:  # 旧数据
                    dec -= 1
                if s > nums[nxt2]:  # 新数据（当前元素，下下一个数）
                    dec += 1
                heappush(h, (s + nums[nxt2], i))

            nums[i] = s
            # 删除 nxt
            l, r = left[nxt], right[nxt]
            right[l] = r  # 模拟双向链表的删除操作
            left[r] = l
            right[nxt] = n  # 表示删除 nxt

        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(long s, int i) {
    }

    public int minimumPairRemoval(int[] nums) {
        int n = nums.length;
        // (相邻元素和，左边那个数的下标)
        PriorityQueue<Pair> h = new PriorityQueue<>((a, b) -> a.s != b.s ? Long.compare(a.s, b.s) : a.i - b.i);
        int dec = 0; // 递减的相邻对的个数
        for (int i = 0; i < n - 1; i++) {
            int x = nums[i];
            int y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            h.offer(new Pair(x + y, i));
        }

        // 每个下标的左右最近的未删除下标
        int[] left = new int[n + 1];
        int[] right = new int[n + 1];
        for (int i = 0; i <= n; i++) {
            left[i] = i - 1;
            right[i] = i + 1;
        }

        long[] a = new long[n];
        for (int i = 0; i < n; i++) {
            a[i] = nums[i];
        }

        int ans = 0;
        while (dec > 0) {
            ans++;

            // 如果堆顶数据与实际数据不符，说明堆顶数据是之前本应删除，但没有删除的数据（懒删除）
            while (right[h.peek().i] >= n || h.peek().s != a[h.peek().i] + a[right[h.peek().i]]) {
                h.poll();
            }

            // 删除相邻元素和最小的一对
            Pair p = h.poll();
            long s = p.s;
            int i = p.i;

            // (当前元素，下一个数)
            int nxt = right[i];
            if (a[i] > a[nxt]) { // 旧数据
                dec--;
            }

            // (前一个数，当前元素)
            int pre = left[i];
            if (pre >= 0) {
                if (a[pre] > a[i]) { // 旧数据
                    dec--;
                }
                if (a[pre] > s) { // 新数据
                    dec++;
                }
                h.offer(new Pair(a[pre] + s, pre));
            }

            // (下一个数，下下一个数)
            int nxt2 = right[nxt];
            if (nxt2 < n) {
                if (a[nxt] > a[nxt2]) { // 旧数据
                    dec--;
                }
                if (s > a[nxt2]) { // 新数据（当前元素，下下一个数）
                    dec++;
                }
                h.add(new Pair(s + a[nxt2], i));
            }

            a[i] = s; // 把 a[nxt] 加到 a[i] 中
            // 删除 nxt
            int l = left[nxt];
            int r = right[nxt];
            right[l] = r; // 模拟双向链表的删除操作
            left[r] = l;
            right[nxt] = n; // 表示删除 nxt
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumPairRemoval(vector<int>& nums) {
        int n = nums.size();
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> pq; // (相邻元素和，左边那个数的下标)
        int dec = 0; // 递减的相邻对的个数
        for (int i = 0; i + 1 < n; i++) {
            int x = nums[i], y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            pq.emplace(x + y, i);
        }

        // 每个下标的左右最近的未删除下标
        vector<int> left(n + 1), right(n);
        ranges::iota(left, -1);
        ranges::iota(right, 1);

        vector<long long> a(nums.begin(), nums.end());
        int ans = 0;
        while (dec) {
            ans++;

            // 如果堆顶数据与实际数据不符，说明堆顶数据是之前本应删除，但没有删除的数据（懒删除）
            while (right[pq.top().second] >= n || pq.top().first != a[pq.top().second] + a[right[pq.top().second]]) {
                pq.pop();
            }
            auto [s, i] = pq.top();
            pq.pop(); // 删除相邻元素和最小的一对

            // (当前元素，下一个数)
            int nxt = right[i];
            dec -= a[i] > a[nxt]; // 旧数据

            // (前一个数，当前元素)
            int pre = left[i];
            if (pre >= 0) {
                dec -= a[pre] > a[i]; // 旧数据
                dec += a[pre] > s; // 新数据
                pq.emplace(a[pre] + s, pre);
            }

            // (下一个数，下下一个数)
            int nxt2 = right[nxt];
            if (nxt2 < n) {
                dec -= a[nxt] > a[nxt2]; // 旧数据
                dec += s > a[nxt2]; // 新数据（当前元素，下下一个数）
                pq.emplace(s + a[nxt2], i);
            }

            a[i] = s;
            // 删除 nxt
            int l = left[nxt], r = right[nxt];
            right[l] = r; // 模拟双向链表的删除操作
            left[r] = l;
            right[nxt] = n; // 表示删除 nxt
        }

        return ans;
    }
};
```

```go [sol-Go]
func minimumPairRemoval(nums []int) (ans int) {
	n := len(nums)
	h := make(hp, n-1)
	dec := 0 // 递减的相邻对的个数
	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		h[i] = pair{x + y, i}
	}
	heap.Init(&h)
	lazy := map[pair]int{}

	// 每个下标的左右最近的未删除下标
	left := make([]int, n+1) // 加一个哨兵，防止下标越界
	right := make([]int, n)
	for i := range n {
		left[i] = i - 1
		right[i] = i + 1
	}
	remove := func(i int) {
		l, r := left[i], right[i]
		right[l] = r // 模拟双向链表的删除操作
		left[r] = l
	}

	for dec > 0 {
		ans++

		for lazy[h[0]] > 0 {
			lazy[h[0]]--
			heap.Pop(&h)
		}
		p := heap.Pop(&h).(pair) // 删除相邻元素和最小的一对
		s := p.s
		i := p.i

		// (当前元素，下一个数)
		nxt := right[i]
		if nums[i] > nums[nxt] { // 旧数据
			dec--
		}

		// (前一个数，当前元素)
		pre := left[i]
		if pre >= 0 {
			if nums[pre] > nums[i] { // 旧数据
				dec--
			}
			if nums[pre] > s { // 新数据
				dec++
			}
			lazy[pair{nums[pre] + nums[i], pre}]++ // 懒删除
			heap.Push(&h, pair{nums[pre] + s, pre})
		}

		// (下一个数，下下一个数)
		nxt2 := right[nxt]
		if nxt2 < n {
			if nums[nxt] > nums[nxt2] { // 旧数据
				dec--
			}
			if s > nums[nxt2] { // 新数据（当前元素，下下一个数）
				dec++
			}
			lazy[pair{nums[nxt] + nums[nxt2], nxt}]++ // 懒删除
			heap.Push(&h, pair{s + nums[nxt2], i})
		}

		nums[i] = s
		remove(nxt)
	}
	return
}

type pair struct{ s, i int } // (相邻元素和，左边那个数的下标)
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.s < b.s || a.s == b.s && a.i < b.i }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

```go [sol-Go 不用 lazy]
func minimumPairRemoval(nums []int) (ans int) {
	n := len(nums)
	h := make(hp, n-1)
	dec := 0 // 递减的相邻对的个数
	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		h[i] = pair{x + y, i}
	}
	heap.Init(&h)

	// 每个下标的左右最近的未删除下标
	left := make([]int, n+1) // 加一个哨兵，防止下标越界
	right := make([]int, n)
	for i := range n {
		left[i] = i - 1
		right[i] = i + 1
	}
	remove := func(i int) {
		l, r := left[i], right[i]
		right[l] = r // 模拟双向链表的删除操作
		left[r] = l
		right[i] = n // 表示 i 已被删除
	}

	for dec > 0 {
		ans++

		// 如果堆顶数据与实际数据不符，说明堆顶数据是之前本应删除，但没有删除的数据（懒删除）
		for right[h[0].i] >= n || nums[h[0].i]+nums[right[h[0].i]] != h[0].s {
			heap.Pop(&h)
		}
		p := heap.Pop(&h).(pair) // 删除相邻元素和最小的一对
		s := p.s
		i := p.i

		// (当前元素，下一个数)
		nxt := right[i]
		if nums[i] > nums[nxt] { // 旧数据
			dec--
		}

		// (前一个数，当前元素)
		pre := left[i]
		if pre >= 0 {
			if nums[pre] > nums[i] { // 旧数据
				dec--
			}
			if nums[pre] > s { // 新数据
				dec++
			}
			heap.Push(&h, pair{nums[pre] + s, pre})
		}

		// (下一个数，下下一个数)
		nxt2 := right[nxt]
		if nxt2 < n {
			if nums[nxt] > nums[nxt2] { // 旧数据
				dec--
			}
			if s > nums[nxt2] { // 新数据（当前元素，下下一个数）
				dec++
			}
			heap.Push(&h, pair{s + nums[nxt2], i})
		}

		nums[i] = s
		remove(nxt)
	}
	return
}

type pair struct{ s, i int } // (相邻元素和，左边那个数的下标)
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.s < b.s || a.s == b.s && a.i < b.i }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
