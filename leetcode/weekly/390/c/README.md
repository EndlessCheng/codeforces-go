请看 [视频讲解](https://www.bilibili.com/video/BV1wr421h7xY/) 第三题。

## 方法一：哈希表 + 有序集合

做法类似前几天的每日一题 [2671. 频率跟踪器](https://leetcode.cn/problems/frequency-tracker/)：

- 用哈希表 $\textit{cnt}$ 记录 $x=\textit{nums}[i]$ 的出现次数 $\textit{cnt}[x]$（用 $\textit{freq}$ 更新出现次数）。
- 用有序集合记录 $\textit{cnt}[x]$ 的出现次数，从而可以 $\mathcal{O}(\log n)$ 知道最大的 $\textit{cnt}[x]$ 是多少。

```py [sol-Python3]
from sortedcontainers import SortedList

class Solution:
    def mostFrequentIDs(self, nums: List[int], freq: List[int]) -> List[int]:
        cnt = Counter()
        sl = SortedList()
        ans = []
        for x, f in zip(nums, freq):
            sl.discard(cnt[x])  # 多个 cnt[x] 只会移除一个
            cnt[x] += f
            sl.add(cnt[x])
            ans.append(sl[-1])
        return ans
```

```java [sol-Java]
class Solution {
    public long[] mostFrequentIDs(int[] nums, int[] freq) {
        Map<Integer, Long> cnt = new HashMap<>();
        TreeMap<Long, Integer> m = new TreeMap<>();
        int n = nums.length;
        long[] ans = new long[n];
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (cnt.containsKey(x) && m.containsKey(cnt.get(x)) && m.merge(cnt.get(x), -1, Integer::sum) == 0) { // --m[cnt[x]] == 0
                m.remove(cnt.get(x));
            }
            long c = cnt.merge(x, (long) freq[i], Long::sum); // cnt[x] += freq[i]
            m.merge(c, 1, Integer::sum); // ++m[cnt[x]]
            ans[i] = m.lastKey();
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> mostFrequentIDs(vector<int> &nums, vector<int> &freq) {
        unordered_map<int, long long> cnt;
        multiset<long long> m;
        int n = nums.size();
        vector<long long> ans(n);
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            auto it = m.find(cnt[x]);
            if (it != m.end()) {
                m.erase(it);
            }
            cnt[x] += freq[i];
            m.insert(cnt[x]);
            ans[i] = *m.rbegin();
        }
        return ans;
    }
};
```

```go [sol-Go]
// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha
func mostFrequentIDs(nums, freq []int) []int64 {
	cnt := map[int]int{}
	t := redblacktree.New[int, int]()
	ans := make([]int64, len(nums))
	for i, x := range nums {
		// 减少一次 cnt[x] 的出现次数
		node := t.GetNode(cnt[x])
		if node != nil {
			node.Value--
			if node.Value == 0 {
				t.Remove(node.Key)
			}
		}

		cnt[x] += freq[i]

		// 增加一次 cnt[x] 的出现次数
		node = t.GetNode(cnt[x])
		if node == nil {
			t.Put(cnt[x], 1)
		} else {
			node.Value++
		}
		ans[i] = int64(t.Right().Key)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：哈希表 + 懒删除堆

也可以不用有序集合，而是用一个**最大堆**保存数对 $(\textit{cnt}[x], x)$。

在堆中查询 $\textit{cnt}[x]$ 的最大值时，如果堆顶保存的数据并不是目前实际的 $\textit{cnt}[x]$，那么就弹出堆顶。

```py [sol-Python3]
class Solution:
    def mostFrequentIDs(self, nums: List[int], freq: List[int]) -> List[int]:
        ans = []
        cnt = Counter()
        h = []
        for x, f in zip(nums, freq):
            cnt[x] += f
            heappush(h, (-cnt[x], x))  # 取负号变成最大堆
            while -h[0][0] != cnt[h[0][1]]:  # 堆顶保存的数据已经发生变化
                heappop(h)  # 删除
            ans.append(-h[0][0])
        return ans
```

```java [sol-Java]
class Solution {
    public long[] mostFrequentIDs(int[] nums, int[] freq) {
        int n = nums.length;
        long[] ans = new long[n];
        Map<Integer, Long> cnt = new HashMap<>();
        PriorityQueue<Pair<Long, Integer>> pq = new PriorityQueue<>((a, b) -> Long.compare(b.getKey(), a.getKey()));
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            long c = cnt.merge(x, (long) freq[i], Long::sum);
            pq.add(new Pair<>(c, x));
            while (!pq.peek().getKey().equals(cnt.get(pq.peek().getValue()))) { // 堆顶保存的数据已经发生变化
                pq.poll(); // 删除
            }
            ans[i] = pq.peek().getKey();
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> mostFrequentIDs(vector<int> &nums, vector<int> &freq) {
        int n = nums.size();
        vector<long long> ans(n);
        unordered_map<int, long long> cnt;
        priority_queue<pair<long long, int>> pq;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            cnt[x] += freq[i];
            pq.emplace(cnt[x], x);
            while (pq.top().first != cnt[pq.top().second]) { // 堆顶保存的数据已经发生变化
                pq.pop(); // 删除
            }
            ans[i] = pq.top().first;
        }
        return ans;
    }
};
```

```go [sol-Go]
func mostFrequentIDs(nums, freq []int) []int64 {
	ans := make([]int64, len(nums))
	cnt := make(map[int]int)
	h := hp{}
	heap.Init(&h)
	for i, x := range nums {
		cnt[x] += freq[i]
		heap.Push(&h, pair{cnt[x], x})
		for h[0].c != cnt[h[0].x] { // 堆顶保存的数据已经发生变化
			heap.Pop(&h) // 删除
		}
		ans[i] = int64(h[0].c)
	}
	return ans
}

type pair struct{ c, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].c > h[j].c } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。由于至多入堆 $\mathcal{O}(n)$ 次，所以出堆也至多 $\mathcal{O}(n)$ 次，二重循环的时间复杂度为 $\mathcal{O}(n\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。
