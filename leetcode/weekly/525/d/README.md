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

倒着遍历 $\textit{nums}$，同时维护元素 $x$ 在 $\textit{nums}$ 中的位置组成的列表，即可快速得到 $p_i$ 和 $q_i$。

为了维护子数组中的不同值列表 $\textit{firsts}$，把遍历到的 $\textit{nums}[i]$ 移到 $\textit{firsts}$ 的开头。

[本题视频讲解](https://www.bilibili.com/video/BV1KLHG6TEWJ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countStaircase(self, nums: list[int]) -> int:
        n = len(nums)
        pos = defaultdict(list)
        firsts = []
        ans = 0

        for i in range(n - 1, -1, -1):
            x = nums[i]
            pos[x].append(i)  # pos[x][-k] 即为右侧第 k 个 x 的下标

            # 把 firsts 中的 x 移到最前面
            if x in firsts:
                firsts.remove(x)
            firsts.insert(0, x)

            m = len(firsts)
            if m * (m + 1) // 2 > n:  # firsts 不需要这么多元素
                firsts.pop()

            mx, mn = 0, n
            r = i  # 子数组右开端点
            for k, x in enumerate(firsts, 1):
                r += k
                # [i,r) 必须包含恰好 k 个 x
                ps = pos[x]
                if k > len(ps):  # [i,n) 中的 x 的个数小于 k
                    break

                # 更新 mx 和 mn 后，必须满足 mx < r <= mn
                mx = max(mx, ps[-k])
                if k < len(ps):
                    mn = min(mn, ps[-1 - k])
                if mx >= mn or r > mn:  # 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break
                if mx < r:
                    ans += 1

        return ans
```

```java [sol-Java]
class Solution {
    public int countStaircase(int[] nums) {
        int n = nums.length;
        Map<Integer, List<Integer>> pos = new HashMap<>();
        List<Integer> firsts = new ArrayList<>();
        int ans = 0;

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            pos.computeIfAbsent(x, _ -> new ArrayList<>()).add(i);
            // pos[x][pos[x].size()-k] 即为右侧第 k 个 x 的下标

            // 把 firsts 中的 x 移到最前面
            firsts.remove(Integer.valueOf(x));
            firsts.addFirst(x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.removeLast();
                m--;
            }

            int mx = 0;
            int mn = n;
            int r = i; // 子数组右开端点
            for (int k = 1; k <= m; k++) {
                x = firsts.get(k - 1);
                r += k;
                // [i,r) 必须包含恰好 k 个 x
                List<Integer> ps = pos.get(x);
                int sz = ps.size();
                if (k > sz) { // [i,n) 中的 x 的个数小于 k
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                mx = Math.max(mx, ps.get(sz - k));
                if (k < sz) {
                    mn = Math.min(mn, ps.get(sz - k - 1));
                }
                if (mx >= mn || r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }
                if (mx < r) {
                    ans++;
                }
            }
        }

        return ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int countStaircase(int[] nums) {
        int n = nums.length;
        int maxNum = 0;
        for (int x : nums) {
            maxNum = Math.max(maxNum, x);
        }

        List<Integer>[] pos = new ArrayList[maxNum + 1];
        List<Integer> firsts = new ArrayList<>();
        int ans = 0;

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            if (pos[x] == null) {
                pos[x] = new ArrayList<>();
            }
            pos[x].add(i); // pos[x][pos[x].size()-k] 即为右侧第 k 个 x 的下标

            // 把 firsts 中的 x 移到最前面
            firsts.remove(Integer.valueOf(x));
            firsts.addFirst(x);

            int m = firsts.size();
            if (m * (m + 1) / 2 > n) { // firsts 不需要这么多元素
                firsts.removeLast();
                m--;
            }

            int mx = 0;
            int mn = n;
            int r = i; // 子数组右开端点
            for (int k = 1; k <= m; k++) {
                x = firsts.get(k - 1);
                r += k;
                // [i,r) 必须包含恰好 k 个 x
                List<Integer> ps = pos[x];
                int sz = ps.size();
                if (k > sz) { // [i,n) 中的 x 的个数小于 k
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                mx = Math.max(mx, ps.get(sz - k));
                if (k < sz) {
                    mn = Math.min(mn, ps.get(sz - k - 1));
                }
                if (mx >= mn || r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }
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
        vector<int> firsts;
        int ans = 0;

        for (int i = n - 1; i >= 0; i--) {
            int x = nums[i];
            pos[x].push_back(i); // pos[x][pos[x].size()-k] 即为右侧第 k 个 x 的下标

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
            for (int k = 1; k <= firsts.size(); k++) {
                x = firsts[k - 1];
                r += k;
                // [i,r) 必须包含恰好 k 个 x
                auto& ps = pos[x];
                int sz = ps.size();
                if (k > sz) { // [i,n) 中的 x 的个数小于 k
                    break;
                }

                // 更新 mx 和 mn 后，必须满足 mx < r <= mn
                mx = max(mx, ps[sz - k]);
                if (k < sz) {
                    mn = min(mn, ps[sz - k - 1]);
                }
                if (mx >= mn || r > mn) { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
                    break;
                }
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
	//pos := map[int][]int{}
	pos := make([][]int, slices.Max(nums)+1) // slice 比 map 快
	firsts := []int{}

	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		pos[x] = append(pos[x], i) // pos[x][len(pos[x])-k] 即为右侧第 k 个 x 的下标

		// 把 firsts 中的 x 移到最前面
		j := slices.Index(firsts, x)
		if j >= 0 {
			firsts = append(firsts[:j], firsts[j+1:]...)
		}
		firsts = slices.Insert(firsts, 0, x)

		m := len(firsts)
		if m*(m+1)/2 > n { // firsts 不需要这么多元素
			firsts = firsts[:len(firsts)-1]
		}

		mx, mn := 0, n
		r := i // 子数组右开端点
		for k, x := range firsts {
			k++
			r += k
			// [i,r) 必须包含恰好 k 个 x
			ps := pos[x]
			sz := len(ps)
			if k > sz { // [i,n) 中的 x 的个数小于 k
				break
			}

			// 更新 mx 和 mn 后，必须满足 mx < r <= mn
			mx = max(mx, ps[sz-k])
			if k < len(ps) {
				mn = min(mn, ps[sz-k-1])
			}
			if mx >= mn || r > mn { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
				break
			}
			if mx < r {
				ans++
			}
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：本题只需判断子数组是否恰好有 $k$ 个 $x$，无需二分。如果要精确计算子数组中的 $x$ 的个数，需要二分，见 [2080. 区间内查询数字的频率](https://leetcode.cn/problems/range-frequency-queries/)。 

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
