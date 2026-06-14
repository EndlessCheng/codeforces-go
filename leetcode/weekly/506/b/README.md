## 题意解读

有两种符合要求的子数组：

1. 子数组只包含一种元素。例如 $[5,5,5,5]$。
2. 子数组**恰好**有两种**不同**的出现次数，且这两种出现次数，一个必须是另一个的两倍。例如 $[20, 10, 20, 30, 30]$，元素 $10,20,30$ 的出现次数分别为 $1,2,2$，去重后为 $1,2$，这有两个，且满足一个是另一个的两倍。

## 思路

由于 $n\le 1000$，我们可以枚举子数组的左右端点。

外层循环枚举左端点，内层循环枚举右端点，不断向右扩大子数组长度。

在扩大子数组的过程中，维护子数组的**元素出现次数** $\textit{cnt}$，以及**元素出现次数的出现次数** $\textit{cc}$。后者用来表达「多重集」，如上面例子中的 $\{1,2,2\}$。 

1. 子数组只包含一种元素，等价于 $\textit{cnt}$ 的大小为 $1$。
2. 子数组恰好有两种不同的出现次数，等价于 $\textit{cc}$ 的大小为 $2$。此外，这两种出现次数，一个必须是另一个的两倍。

[本题视频讲解](https://www.bilibili.com/video/BV1ptJw6hENZ/)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def getLength(self, nums: list[int]) -> int:
        ans = 0
        n = len(nums)

        for i in range(n):
            cnt = defaultdict(int)  # 出现次数
            cc = defaultdict(int)  # 出现次数的出现次数

            for j in range(i, n):
                x = nums[j]
                c = cnt[x]
                if c > 0:
                    cc[c] -= 1
                    if cc[c] == 0:
                        del cc[c]  # 保证我们可以正确计算 cc 的大小
                cnt[x] += 1
                cc[cnt[x]] += 1

                if len(cnt) == 1:  # 子数组只包含一种元素
                    ans = max(ans, j - i + 1)
                elif len(cc) == 2:  # 子数组有两种不同的出现次数
                    c1, c2 = sorted(cc)
                    if c1 * 2 == c2:
                        ans = max(ans, j - i + 1)

        return ans
```

```java [sol-Java]
class Solution {
    public int getLength(int[] nums) {
        int n = nums.length;
        int ans = 0;

        for (int i = 0; i < n; i++) {
            HashMap<Integer, Integer> cnt = new HashMap<>(); // 出现次数
            HashMap<Integer, Integer> cc = new HashMap<>(); // 出现次数的出现次数

            for (int j = i; j < n; j++) {
                int x = nums[j];
                Integer c = cnt.get(x);
                if (c != null) {
                    if (cc.merge(c, -1, Integer::sum) == 0) { // --cc[c] == 0
                        cc.remove(c); // 保证我们可以正确计算 cc 的大小
                    }
                }
                c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
                cc.merge(c, 1, Integer::sum); // ++cc[c]

                if (cnt.size() == 1) { // 子数组只包含一种元素
                    ans = Math.max(ans, j - i + 1);
                } else if (cc.size() == 2) { // 子数组有两种不同的出现次数
                    Iterator<Integer> it = cc.keySet().iterator();
                    int c1 = it.next();
                    int c2 = it.next();
                    if (Math.min(c1, c2) * 2 == Math.max(c1, c2)) {
                        ans = Math.max(ans, j - i + 1);
                    }
                }
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLength(vector<int>& nums) {
        int n = nums.size();
        int ans = 0;

        for (int i = 0; i < n; i++) {
            unordered_map<int, int> cnt; // 出现次数
            unordered_map<int, int> cc; // 出现次数的出现次数

            for (int j = i; j < n; j++) {
                int x = nums[j];
                int c = cnt[x];
                if (c > 0 && --cc[c] == 0) {
                    cc.erase(c); // 保证我们可以正确计算 cc 的大小
                }
                cnt[x]++;
                cc[cnt[x]]++;

                if (cnt.size() == 1) { // 子数组只包含一种元素
                    ans = max(ans, j - i + 1);
                } else if (cc.size() == 2) { // 子数组有两种不同的出现次数
                    int c1 = cc.begin()->first;
                    int c2 = next(cc.begin())->first;
                    if (min(c1, c2) * 2 == max(c1, c2)) {
                        ans = max(ans, j - i + 1);
                    }
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func getLength(nums []int) (ans int) {
	for i := range nums {
		cnt := map[int]int{} // 出现次数
		cc := map[int]int{}  // 出现次数的出现次数
		for j := i; j < len(nums); j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					delete(cc, c) // 保证我们可以正确计算 cc 的大小
				}
			}
			cnt[x]++
			cc[cnt[x]]++

			if len(cnt) == 1 { // 子数组只包含一种元素
				ans = max(ans, j-i+1)
			} else if len(cc) == 2 { // 子数组有两种不同的出现次数
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}
```

## 优化

1. 特判：如果 $\textit{nums}$ 所有元素互不相同，返回 $1$。
2. 统计 $\textit{nums}$ 中连续相同元素的最长长度，作为答案的初始值。做法同 [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/)。相关题目见双指针题单的「**六、分组循环**」。
3. 有了优化 2，后续代码无需处理子数组只包含一种元素的情况。
4. 最优性剪枝：对于以 $i$ 为左端点的子数组，其长度至多为 $n-i$。所以如果发现 $n-i\le \textit{ans}$，说明答案不会再增大，跳出循环。

```py [sol-Python3]
class Solution:
    def getLength(self, nums: list[int]) -> int:
        n = len(nums)
        # 所有元素互不相同
        if len(set(nums)) == n:
            return 1

        # 子数组只有一种元素
        ans = same = 0
        for i, x in enumerate(nums):
            same += 1
            if i == n - 1 or x != nums[i + 1]:  # 到达连续相同段的末尾
                ans = max(ans, same)
                same = 0

        for i in range(n):
            # 最优性剪枝：答案不会变大
            if n - i <= ans:
                break

            cnt = defaultdict(int)  # 出现次数
            cc = defaultdict(int)  # 出现次数的出现次数

            for j in range(i, n):
                x = nums[j]
                c = cnt[x]
                if c > 0:
                    cc[c] -= 1
                    if cc[c] == 0:
                        del cc[c]  # 保证我们可以正确计算 cc 的大小
                cnt[x] += 1
                cc[cnt[x]] += 1

                if len(cc) == 2:  # 子数组有两种不同的出现次数
                    c1, c2 = sorted(cc)
                    if c1 * 2 == c2:
                        ans = max(ans, j - i + 1)

        return ans
```

```java [sol-Java]
class Solution {
    public int getLength(int[] nums) {
        int n = nums.length;
        // 所有元素互不相同
        HashSet<Integer> st = new HashSet<>();
        for (int x : nums) {
            st.add(x);
        }
        if (st.size() == n) {
            return 1;
        }

        int ans = 0;
        // 子数组只有一种元素
        int same = 0;
        for (int i = 0; i < n; i++) {
            same++;
            if (i == n - 1 || nums[i] != nums[i + 1]) { // 到达连续相同段的末尾
                ans = Math.max(ans, same);
                same = 0;
            }
        }

        // 最优性剪枝：答案不会变大
        for (int i = 0; i < n - ans; i++) {
            HashMap<Integer, Integer> cnt = new HashMap<>(); // 出现次数
            HashMap<Integer, Integer> cc = new HashMap<>(); // 出现次数的出现次数

            for (int j = i; j < n; j++) {
                int x = nums[j];
                Integer c = cnt.get(x);
                if (c != null) {
                    if (cc.merge(c, -1, Integer::sum) == 0) { // --cc[c] == 0
                        cc.remove(c); // 保证我们可以正确计算 cc 的大小
                    }
                }
                c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
                cc.merge(c, 1, Integer::sum); // ++cc[c]

                if (cc.size() == 2) { // 子数组有两种不同的出现次数
                    Iterator<Integer> it = cc.keySet().iterator();
                    int c1 = it.next();
                    int c2 = it.next();
                    if (Math.min(c1, c2) * 2 == Math.max(c1, c2)) {
                        ans = Math.max(ans, j - i + 1);
                    }
                }
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int getLength(vector<int>& nums) {
        int n = nums.size();
        // 所有元素互不相同
        unordered_set<int> st(nums.begin(), nums.end());
        if (st.size() == n) {
            return 1;
        }

        // 子数组只有一种元素
        int ans = 0, same = 0;
        for (int i = 0; i < n; i++) {
            same++;
            if (i == n - 1 || nums[i] != nums[i + 1]) { // 到达连续相同段的末尾
                ans = max(ans, same);
                same = 0;
            }
        }

        // 最优性剪枝：答案不会变大
        for (int i = 0; i < n - ans; i++) {
            unordered_map<int, int> cnt; // 出现次数
            unordered_map<int, int> cc; // 出现次数的出现次数

            for (int j = i; j < n; j++) {
                int x = nums[j];
                int c = cnt[x];
                if (c > 0 && --cc[c] == 0) {
                    cc.erase(c); // 保证我们可以正确计算 cc 的大小
                }
                cnt[x]++;
                cc[cnt[x]]++;

                if (cc.size() == 2) { // 子数组有两种不同的出现次数
                    int c1 = cc.begin()->first;
                    int c2 = next(cc.begin())->first;
                    if (min(c1, c2) * 2 == max(c1, c2)) {
                        ans = max(ans, j - i + 1);
                    }
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func getLength(nums []int) (ans int) {
	n := len(nums)
	set := map[int]struct{}{}
	// 子数组只有一种元素
	same := 0
	for i, x := range nums {
		set[x] = struct{}{}
		same++
		if i == n-1 || x != nums[i+1] { // 到达连续相同段的末尾
			ans = max(ans, same)
			same = 0
		}
	}

	// 所有元素互不相同
	if len(set) == n {
		return 1
	}

	// 最优性剪枝：答案不会变大
	for i := 0; i < n-ans; i++ {
		cnt := map[int]int{} // 出现次数
		cc := map[int]int{}  // 出现次数的出现次数
		for j := i; j < n; j++ {
			x := nums[j]
			if c := cnt[x]; c > 0 {
				cc[c]--
				if cc[c] == 0 {
					delete(cc, c) // 保证我们可以正确计算 cc 的大小
				}
			}
			cnt[x]++
			cc[cnt[x]]++

			if len(cc) == 2 { // 子数组的元素出现次数恰好有两种
				c := slices.Sorted(maps.Keys(cc))
				if c[0]*2 == c[1] {
					ans = max(ans, j-i+1)
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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
