## 思路

设合法子数组有 $k$ 个不同值，那么合法子数组的长度必须恰好等于

$$
1 + 2 +\cdots + k = \dfrac{k(k+1)}{2}
$$

设 $n$ 是 $\textit{nums}$ 的长度。由于 $\dfrac{k(k+1)}{2} \le n$，所以 $k = \mathcal{O}(\sqrt n)$。

我们可以暴力枚举子数组左端点 $L$ 和子数组长度 $\dfrac{k(k+1)}{2}$，其中 $k=1,2,3,\ldots$。

假设现在枚举到了子数组 $[L,R]$，其中 $R-L+1 = \dfrac{k(k+1)}{2}$。定义：

- $p_i$ 为子数组第 $i$ 个不同值在 $[L,n)$ 中的第 $i$ 次出现位置。如果不存在这样的位置，不仅说明 $[L,R]$ 的第 $i$ 个不同值的出现次数小于 $i$，还说明其余右端点大于 $R$ 的子数组 $[L,R']$ 也不满足要求（第 $i$ 个不同值的出现次数小于 $i$），所以跳出枚举 $k$ 的循环。
- $q_i$ 为子数组第 $i$ 个不同值在 $[L,n)$ 中的第 $i+1$ 次出现位置。如果不存在这样的位置，令 $q_i=n$。

**关键点**：如果 $p_i\le R < q_i$，则说明第 $i$ 个不同值在子数组中恰好出现 $i$ 次。

对于 $i=1,2,\ldots,k$，联立这 $k$ 个不等式（求 $k$ 个区间的交集），得

$$
\max_{i=1}^{k} p_i \le R < \min_{i=1}^{k} q_i
$$

这意味着，我们需要维护 $p$ 的**前缀最大值**，以及 $q$ 的**前缀最小值**。由于我们是从小到大枚举 $k$，所以只需用两个变量 $\textit{mx}$ 和 $\textit{mn}$ 分别维护前缀最大值和前缀最小值，上式变为

$$
\textit{mx} \le R < \textit{mn}
$$

这样就能 $\mathcal{O}(1)$ 判断一个子数组是否合法。

## 细节

为了快速求出 $p_i$ 和 $q_i$，我们可以预处理：

- $\textit{pos}[x]$：元素 $x$ 在 $\textit{nums}$ 中的所有位置组成的列表。
- $\textit{idxInPos}[i]$：$\textit{nums}[i]$ 在 $\textit{pos}[\textit{nums}[i]]$ 中的下标。

设子数组的第 $i$ 个不同值为 $x$，其在 $\textit{nums}$ 中的下标为 $j$，那么 $p_i = \textit{pos}[x][\textit{idxInPos}[j] + i - 1]$，$q_i = \textit{pos}[x][\textit{idxInPos}[j] + i]$。

为了维护子数组中的不同值列表 $\textit{firsts}$，我们可以倒着遍历 $\textit{nums}$，把遍历到的 $\textit{nums}[i]$ 移到 $\textit{firsts}$ 的开头。

[本题视频讲解](https://www.bilibili.com/video/BV1KLHG6TEWJ/)，欢迎点赞关注~

## 写法一：哈希表

```py [sol-Python3]
class Solution:
    def countStaircase(self, nums: list[int]) -> int:
        n = len(nums)
        pos = defaultdict(list)
        idx_in_pos = [0] * n  # nums[i] 在 pos[nums[i]] 中的下标
        for i, x in enumerate(nums):
            idx_in_pos[i] = len(pos[x])
            pos[x].append(i)

        firsts = []
        first_pos = {}
        ans = 0
        for i in range(n - 1, -1, -1):
            x = nums[i]
            first_pos[x] = i

            # 把 firsts 中的 x 移到最前面
            if x in firsts:
                firsts.remove(x)
            firsts.insert(0, x)

            m = len(firsts)
            if m * (m + 1) // 2 > n:  # firsts 不需要这么多元素
                firsts.pop()

            mx, mn = 0, n
            r = i  # 子数组右开端点
            for cnt, x in enumerate(firsts, 1):
                r += cnt
                # [i,r) 必须包含恰好 cnt 个 x
                ps = pos[x]
                k = idx_in_pos[first_pos[x]] + cnt - 1
                if k >= len(ps):  # [i,n) 中的 x 的个数小于 cnt
                    break

                # 更新 mx 和 mn 后，必须满足 mx < r <= mn
                if k < len(ps) - 1:
                    mn = min(mn, ps[k + 1])
                if r > mn:  # 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break

                mx = max(mx, ps[k])
                if mx < r:
                    ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    public int countStaircase(int[] nums) {
        int n = nums.length;
        Map<Integer, List<Integer>> pos = new HashMap<>();
        int[] idxInPos = new int[n];
        for (int i = 0; i < n; i++) {
            List<Integer> ps = pos.computeIfAbsent(nums[i], _ -> new ArrayList<>());
            idxInPos[i] = ps.size();
            ps.add(i);
        }

        List<Integer> firsts = new ArrayList<>();
        Map<Integer, Integer> firstPos = new HashMap<>();
        int ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            firstPos.put(x, i);

            // 把 firsts 中的 x 移到最前面
            firsts.remove(Integer.valueOf(x));
            firsts.addFirst(x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.removeLast();
            }

            int mx = 0;
            int mn = n;
            int r = i; // 子数组右开端点
            for (int cnt = 1; cnt <= firsts.size(); cnt++) {
                x = firsts.get(cnt - 1);
                r += cnt;
                // [i,r) 必须包含恰好 cnt 个 x
                List<Integer> ps = pos.get(x);
                int k = idxInPos[firstPos.get(x)] + cnt - 1;
                if (k >= ps.size()) { // [i,n) 中的 x 的个数小于 cnt
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                if (k < ps.size() - 1) {
                    mn = Math.min(mn, ps.get(k + 1));
                }
                if (r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }

                mx = Math.max(mx, ps.get(k));
                if (mx < r) {
                    ans++;
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
    int countStaircase(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> pos;
        vector<int> idx_in_pos(n);
        for (int i = 0; i < n; i++) {
            auto& ps = pos[nums[i]];
            idx_in_pos[i] = ps.size();
            ps.push_back(i);
        }

        vector<int> firsts;
        unordered_map<int, int> first_pos;
        int ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            first_pos[x] = i;

            // 把 firsts 中的 x 移到最前面
            auto it = ranges::find(firsts, x);
            if (it != firsts.end()) {
                firsts.erase(it);
            }
            firsts.insert(firsts.begin(), x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.pop_back();
            }

            int mx = 0, mn = n;
            int r = i; // 子数组右开端点
            for (int cnt = 1; cnt <= firsts.size(); cnt++) {
                x = firsts[cnt - 1];
                r += cnt;
                // [i,r) 必须包含恰好 cnt 个 x
                auto& ps = pos[x];
                int k = idx_in_pos[first_pos[x]] + cnt - 1;
                if (k >= ps.size()) { // [i,n) 中的 x 的个数小于 cnt
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                if (k + 1 < ps.size()) {
                    mn = min(mn, ps[k + 1]);
                }
                if (r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }

                mx = max(mx, ps[k]);
                if (mx < r) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countStaircase(nums []int) (ans int) {
	n := len(nums)
	pos := map[int][]int{}
	idxInPos := make([]int, n) // nums[i] 在 pos[nums[i]] 中的下标
	for i, x := range nums {
		idxInPos[i] = len(pos[x])
		pos[x] = append(pos[x], i)
	}

	firsts := []int{}
	firstPos := map[int]int{}
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		firstPos[x] = i

		// 把 firsts 中的 x 移到最前面
		j := slices.Index(firsts, x)
		if j != 0 { // 优化：j = 0 时无需移动
			if j >= 0 {
				firsts = append(firsts[:j], firsts[j+1:]...)
			}
			firsts = slices.Insert(firsts, 0, x)
		}

		m := len(firsts)
		if m*(m+1)/2 > n { // firsts 不需要这么多元素
			firsts = firsts[:len(firsts)-1]
		}

		mx, mn := 0, n
		r := i // 子数组右开端点
		for cnt, x := range firsts {
			r += cnt + 1
			// [i,r) 必须包含恰好 cnt+1 个 x
			ps := pos[x]
			k := idxInPos[firstPos[x]] + cnt
			if k >= len(ps) { // [i,n) 中的 x 的个数小于 cnt+1
				break
			}

			// 更新 mx 和 mn 后，必须满足 mx < r <= mn
			if k < len(ps)-1 {
				mn = min(mn, ps[k+1])
			}
			if r > mn { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
				break
			}

			mx = max(mx, ps[k])
			if mx < r {
				ans++
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：数组

> 试了一下 Python 反而变慢了，所以只写其他语言的代码。

```java [sol-Java]
class Solution {
    public int countStaircase(int[] nums) {
        int n = nums.length;
        int maxNum = 0;
        for (int x : nums) {
            maxNum = Math.max(maxNum, x);
        }

        List<Integer>[] pos = new ArrayList[maxNum + 1];
        Arrays.setAll(pos, _ -> new ArrayList<>());
        int[] idxInPos = new int[n];
        for (int i = 0; i < n; i++) {
            List<Integer> ps = pos[nums[i]];
            idxInPos[i] = ps.size();
            ps.add(i);
        }

        List<Integer> firsts = new ArrayList<>();
        int[] firstPos = new int[maxNum + 1];
        int ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            firstPos[x] = i;

            // 把 firsts 中的 x 移到最前面
            firsts.remove(Integer.valueOf(x));
            firsts.addFirst(x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.removeLast();
            }

            int mx = 0;
            int mn = n;
            int r = i; // 子数组右开端点
            for (int cnt = 1; cnt <= firsts.size(); cnt++) {
                x = firsts.get(cnt - 1);
                r += cnt;
                // [i,r) 必须包含恰好 cnt 个 x
                List<Integer> ps = pos[x];
                int k = idxInPos[firstPos[x]] + cnt - 1;
                if (k >= ps.size()) { // [i,n) 中的 x 的个数小于 cnt
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                if (k < ps.size() - 1) {
                    mn = Math.min(mn, ps.get(k + 1));
                }
                if (r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }

                mx = Math.max(mx, ps.get(k));
                if (mx < r) {
                    ans++;
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
    int countStaircase(vector<int>& nums) {
        int n = nums.size();
        int max_num = ranges::max(nums);

        vector<vector<int>> pos(max_num + 1);
        vector<int> idx_in_pos(n);
        for (int i = 0; i < n; i++) {
            auto& ps = pos[nums[i]];
            idx_in_pos[i] = ps.size();
            ps.push_back(i);
        }

        vector<int> firsts;
        vector<int> first_pos(max_num + 1);
        int ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            first_pos[x] = i;

            // 把 firsts 中的 x 移到最前面
            auto it = ranges::find(firsts, x);
            if (it != firsts.end()) {
                firsts.erase(it);
            }
            firsts.insert(firsts.begin(), x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.pop_back();
            }

            int mx = 0, mn = n;
            int r = i; // 子数组右开端点
            for (int cnt = 1; cnt <= firsts.size(); cnt++) {
                x = firsts[cnt - 1];
                r += cnt;
                // [i,r) 必须包含恰好 cnt 个 x
                auto& ps = pos[x];
                int k = idx_in_pos[first_pos[x]] + cnt - 1;
                if (k >= ps.size()) { // [i,n) 中的 x 的个数小于 cnt
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                if (k + 1 < ps.size()) {
                    mn = min(mn, ps[k + 1]);
                }
                if (r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }

                mx = max(mx, ps[k]);
                if (mx < r) {
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countStaircase(nums []int) (ans int) {
	n := len(nums)
	maxNum := slices.Max(nums)

	pos := make([][]int, maxNum+1)
	idxInPos := make([]int, n) // nums[i] 在 pos[nums[i]] 中的下标
	for i, x := range nums {
		idxInPos[i] = len(pos[x])
		pos[x] = append(pos[x], i)
	}

	firsts := []int{}
	firstPos := make([]int, maxNum+1)
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		firstPos[x] = i

		// 把 firsts 中的 x 移到最前面
		j := slices.Index(firsts, x)
		if j != 0 { // 优化：j = 0 时无需移动
			if j >= 0 {
				firsts = append(firsts[:j], firsts[j+1:]...)
			}
			firsts = slices.Insert(firsts, 0, x)
		}

		m := len(firsts)
		if m*(m+1)/2 > n { // firsts 不需要这么多元素
			firsts = firsts[:len(firsts)-1]
		}

		mx, mn := 0, n
		r := i // 子数组右开端点
		for cnt, x := range firsts {
			r += cnt + 1
			// [i,r) 必须包含恰好 cnt+1 个 x
			ps := pos[x]
			k := idxInPos[firstPos[x]] + cnt
			if k >= len(ps) { // [i,n) 中的 x 的个数小于 cnt+1
				break
			}

			// 更新 mx 和 mn 后，必须满足 mx < r <= mn
			if k < len(ps)-1 {
				mn = min(mn, ps[k+1])
			}
			if r > mn { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
				break
			}

			mx = max(mx, ps[k])
			if mx < r {
				ans++
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n + U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + U)$。

## 相似题目

[3234. 统计 1 显著的字符串的数量](https://leetcode.cn/problems/count-the-number-of-substrings-with-dominant-ones/)

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
