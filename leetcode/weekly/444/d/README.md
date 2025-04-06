为了快速模拟题目的操作，我们需要维护三种信息：

1. 把相邻元素和 $s$，以及这对相邻元素中的左边那个数的下标 $i$，组成一个 pair $(s,i)$。需要支持维护和查询这些 pair 的最小值。这可以用**有序集合**，或者**懒删除堆**。
2. 维护剩余下标，需要支持查询每个下标 $i$ 左侧最近剩余下标，以及右侧最近剩余下标。这可以用**有序集合**，或者**两个并查集**。
3. 在相邻元素中，左边大于右边的个数，记作 $\textit{dec}$。

不断模拟操作，直到 $\textit{dec} = 0$。

题目说「用它们的和替换这对元素」，设操作的这对元素的下标为 $i$ 和 $\textit{nxt}$，操作相当于把 $\textit{nums}[i]$ 增加 $\textit{nums}[\textit{nxt}]$，然后删除下标 $\textit{nxt}$。

在这个过程中，$\textit{dec}$ 如何变化？

设操作的这对元素的下标为 $i$ 和 $\textit{nxt}$，$i$ 左侧最近剩余下标为 $\textit{pre}$，$\textit{nxt}$ 右侧最近剩余下标为 $\textit{nxt}_2$。

也就是说，下标的顺序为 $\textit{pre},i,\textit{nxt},\textit{nxt}_2$。

一个一个来看：

1. 删除 $\textit{nxt}$。如果删除之前 $\textit{nums}[i] > \textit{nums}[\textit{nxt}]$，把 $\textit{dec}$ 减一。
2. 如果删除前 $\textit{nums}[\textit{pre}] > \textit{nums}[i]$，把 $\textit{dec}$ 减一。如果删除后 $\textit{nums}[\textit{pre}] > s$，把 $\textit{dec}$ 加一。这里 $s$ 表示操作的这对元素之和，也就是新的 $\textit{nums}[i]$ 的值。
3. 如果删除前 $\textit{nums}[\textit{nxt}] > \textit{nums}[\textit{nxt}_2]$，把 $\textit{dec}$ 减一。删除后 $i$ 和 $\textit{nxt}_2$ 相邻，如果删除后 $s > \textit{nums}[\textit{nxt}_2]$，把 $\textit{dec}$ 加一。

上述过程中，同时维护（添加删除）新旧相邻元素和以及下标。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ezRvYiE27/?t=41m12s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumPairRemoval(self, nums: List[int]) -> int:
        n = len(nums)
        sl = SortedList()  # (相邻元素和，左边那个数的下标)
        idx = SortedList(range(n))  # 剩余下标
        dec = 0  # 递减的相邻对的个数

        for i, (x, y) in enumerate(pairwise(nums)):
            if x > y:
                dec += 1
            sl.add((x + y, i))

        ans = 0
        while dec:
            ans += 1

            s, i = sl.pop(0)  # 删除相邻元素和最小的一对
            k = idx.bisect_left(i)

            # (当前元素，下一个数)
            nxt = idx[k + 1]
            if nums[i] > nums[nxt]:  # 旧数据
                dec -= 1

            # (前一个数，当前元素)
            if k:
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

            nums[i] = s
            idx.remove(nxt)

        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(long s, int i) {
    }

    public int minimumPairRemoval(int[] nums) {
        int n = nums.length;
        TreeSet<Pair> pairs = new TreeSet<>((a, b) -> a.s != b.s ? Long.compare(a.s, b.s) : a.i - b.i); // (相邻元素和，左边那个数的下标)
        TreeSet<Integer> idx = new TreeSet<>(); // 剩余下标
        int dec = 0; // 递减的相邻对的个数

        for (int i = 0; i < n - 1; i++) {
            int x = nums[i];
            int y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            pairs.add(new Pair(x + y, i));
        }
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

            int nxt = idx.higher(i); // 当前元素的下一个数
            if (a[i] > a[nxt]) { // 旧数据
                dec--;
            }

            Integer pre = idx.lower(i); // 当前元素的前一个数
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

            Integer nxt2 = idx.higher(nxt); // 下下一个数
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

            a[i] = s;
            idx.remove(nxt);
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
        set<int> idx; // 剩余下标
        int dec = 0; // 递减的相邻对的个数

        for (int i = 0; i < n - 1; i++) {
            int x = nums[i], y = nums[i + 1];
            if (x > y) {
                dec++;
            }
            pairs.emplace(x + y, i);
        }
        for (int i = 0; i < n; i++) {
            idx.insert(i);
        }

        vector<long long> a(nums.begin(), nums.end());
        int ans = 0;
        while (dec) {
            ans++;

            // 删除相邻元素和最小的一对
            auto [s, i] = *pairs.begin();
            pairs.erase(pairs.begin());

            auto it = idx.lower_bound(i);
            auto nxt_it = next(it);
            int nxt = *nxt_it;

            // (当前元素，下一个数)
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

            a[i] = s;
            idx.erase(nxt);
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
	// 剩余下标
	idx := redblacktree.New[int, struct{}]()
	// 递减的相邻对的个数
	dec := 0

	for i := range n - 1 {
		x, y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		pairs.Put(pair{x + y, i}, struct{}{})
	}
	for i := range n {
		idx.Put(i, struct{}{})
	}

	for dec > 0 {
		ans++

		it := pairs.Left()
		s := it.Key.s
		i := it.Key.i
		pairs.Remove(it.Key) // 删除相邻元素和最小的一对

		// 找到 i 的位置
		node, _ := idx.Ceiling(i + 1)
		nxt := node.Key

		// (当前元素，下一个数)
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

		nums[i] = s
		idx.Remove(nxt)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

其他做法稍后补充。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
