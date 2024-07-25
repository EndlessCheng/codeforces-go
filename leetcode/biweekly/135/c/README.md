## 方法一：枚举 X

### 提示 1

想一想，什么情况下答案是 $0$？什么情况下答案是 $1$？

如果答案是 $0$，意味着所有 $|\textit{nums}[i]-\textit{nums}[n-1-i]|$ 都等于同一个数 $X$。

如果答案是 $1$，意味着有 $n/2-1$ 个 $|\textit{nums}[i]-\textit{nums}[n-1-i]|$ 都等于同一个数 $X$。我们只需要修改那对不相等的，设这两个数分别为 $p=\textit{nums}[i],\ q=\textit{nums}[n-1-i]$。

不妨设 $p\le q$，分类讨论：

- 如果修改 $p$，那么把 $p$ 改成 $0$ 可以让差值尽量大，此时差值为 $q$。
- 如果修改 $q$，那么把 $q$ 改成 $k$ 可以让差值尽量大，此时差值为 $k-p$。
- 如果 $\max(q, k-p)\ge X$，改其中一个数就行。
- 如果 $\max(q, k-p) < X$，$p$ 和 $q$ 两个数都要改。

注意题目保证 $n$ 是偶数。

### 提示 2

枚举 $X=0,1,2,\cdots,k$，计算至少要修改多少个数。

设 $p=\textit{nums}[i],\ q=\textit{nums}[n-1-i]$，且 $p\le q$（如果 $p>q$ 则交换 $p$ 和 $q$）。

统计 $q-p$ 的出现次数，记录到一个数组 $\textit{cnt}$ 中。

统计 $\max(q, k-p)$ 的出现次数，记录到另一个数组 $\textit{cnt}_2$ 中。

讨论哪些数对无需修改，哪些数对要改一个数，哪些数对要改两个数：

1. 有 $\textit{cnt}[X]$ 对 $(p,q)$ 无需修改。
2. 有 $n/2-\textit{cnt}[X]$ 对 $(p,q)$ 至少要改一个数。
3. 在 2 的基础上，有额外的 $\textit{cnt}_2[0] + \textit{cnt}_2[1] + \cdots + \textit{cnt}_2[X-1]$ 对 $(p,q)$ **还要再改一个数**（根据提示 1）。这可以在枚举 $X$ 的同时，维护一个变量 $\textit{sum}_2$ 表示这些 $\textit{cnt}_2[i]$ 的和。

综上所述，至少要修改

$$
\dfrac{n}{2} - \textit{cnt}[X] + \textit{sum}_2
$$

个数，用它更新答案的最小值。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JE4m1d7br/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minChanges(self, nums: List[int], k: int) -> int:
        cnt = [0] * (k + 1)
        cnt2 = [0] * (k + 1)
        n = len(nums)
        for i in range(n // 2):
            p, q = nums[i], nums[n - 1 - i]
            if p > q:  # 保证 p <= q
                p, q = q, p
            cnt[q - p] += 1
            cnt2[max(q, k - p)] += 1

        ans = n
        sum2 = 0  # 统计有多少对 (p,q) 都要改
        for c, c2 in zip(cnt, cnt2):
            # 其他 n/2-c 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
            ans = min(ans, n // 2 - c + sum2)
            # 对于后面的更大的 x，当前的这 c2 对 (p,q) 都要改
            sum2 += c2
        return ans
```

```java [sol-Java]
class Solution {
    public int minChanges(int[] nums, int k) {
        int[] cnt = new int[k + 1];
        int[] cnt2 = new int[k + 1];
        int n = nums.length;
        for (int i = 0; i < n / 2; i++) {
            int p = nums[i];
            int q = nums[n - 1 - i];
            if (p > q) { // 保证 p <= q
                int tmp = p;
                p = q;
                q = tmp;
            }
            cnt[q - p]++;
            cnt2[Math.max(q, k - p)]++;
        }

        int ans = n;
        int sum2 = 0; // 统计有多少对 (p,q) 都要改
        for (int x = 0; x <= k; x++) {
            // 其他 n/2-cnt[x] 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
            ans = Math.min(ans, n / 2 - cnt[x] + sum2);
            // 对于后面的更大的 x，当前的这 cnt2[x] 对 (p,q) 都要改
            sum2 += cnt2[x];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minChanges(vector<int>& nums, int k) {
        vector<int> cnt(k + 1), cnt2(k + 1);
        int n = nums.size();
        for (int i = 0; i < n / 2; i++) {
            int p = nums[i], q = nums[n - 1 - i];
            if (p > q) { // 保证 p <= q
                swap(p, q);
            }
            cnt[q - p]++;
            cnt2[max(q, k - p)]++;
        }

        int ans = n;
        int sum2 = 0; // 统计有多少对 (p,q) 都要改
        for (int x = 0; x <= k; x++) {
            // 其他 n/2-cnt[x] 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
            ans = min(ans, n / 2 - cnt[x] + sum2);
            // 对于后面的更大的 x，当前的这 cnt2[x] 对 (p,q) 都要改
            sum2 += cnt2[x];
        }
        return ans;
    }
};
```

```go [sol-Go]
func minChanges(nums []int, k int) int {
	cnt := make([]int, k+1)
	cnt2 := make([]int, k+1)
	n := len(nums)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		cnt[q-p]++
		cnt2[max(q, k-p)]++
	}

	ans := n
	sum2 := 0 // 统计有多少对 (p,q) 都要改
	for x, c := range cnt {
		// 其他 n/2-c 对 (p,q) 至少要改一个数，在此基础上，有额外的 sum2 对 (p,q) 还要再改一个数
		ans = min(ans, n/2-c+sum2)
		// 对于后面的更大的 x，当前的这 cnt2[x] 对 (p,q) 都要改
		sum2 += cnt2[x]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 方法二：差分数组

方法一相当于计算出了一个 $\textit{minModify}$ 数组，其中 $\textit{minModify}[X]$ 表示把所有数对 $(p,q)$ 的差值都改成 $X$ 所需的最少操作次数。

如何计算 $\textit{minModify}$ 数组，我们还有一种方法。

枚举数对 $(p,q)$ 的同时，想一想，把这两个数的差值改成 $X=0$ 要操作几次？改成 $X=1$ 要操作几次？改成 $X=2$ 要操作几次？…… 改成 $X=k$ 要操作几次？

设 $x=q-p,\ \textit{mx}=\max(q, k-p)$。讨论把差值改成多少：

- 改成 $[0,x-1]$ 中的数。这意味着 $p$ 和 $q$ 的距离变小，两者靠近即可，只需改其中一个数。
- 改成 $x$。无需修改。
- 改成 $[x+1,\textit{mx}]$ 中的数。根据方法一中的分析，只需改其中一个数。
- 改成 $[\textit{mx}+1,k]$ 中的数。根据方法一中的分析，两个数都要改。

把这些操作次数加到 $\textit{minModify}$ 数组中。这可以用**差分数组**实现，具体见 [差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

遍历这 $n/2$ 个数对，都按上述方法更新 $\textit{minModify}$ 数组，最终我们得到的 $\textit{minModify}$ 数组就和方法一一样了。答案为 $\min(\textit{minModify})$。

注意差分数组需要开 $k+2$ 大小，因为 $\textit{mx}+1$ 可以等于 $k+1$。

```py [sol-Python3]
class Solution:
    def minChanges(self, nums: List[int], k: int) -> int:
        n = len(nums)
        d = [0] * (k + 2)
        for i in range(n // 2):
            p, q = nums[i], nums[n - 1 - i]
            if p > q:  # 保证 p <= q
                p, q = q, p
            x = q - p
            mx = max(q, k - p)
            # [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
            d[0] += 1
            d[x] -= 1
            # [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
            d[x + 1] += 1
            d[mx + 1] -= 1
            # [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
            d[mx + 1] += 2
        return min(accumulate(d))
```

```java [sol-Java]
class Solution {
    public int minChanges(int[] nums, int k) {
        int n = nums.length;
        int[] d = new int[k + 2];
        for (int i = 0; i < n / 2; i++) {
            int p = nums[i];
            int q = nums[n - 1 - i];
            if (p > q) { // 保证 p <= q
                int tmp = p;
                p = q;
                q = tmp;
            }
            int x = q - p;
            int mx = Math.max(q, k - p);
            // [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
            d[0]++;
            d[x]--;
            // [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
            d[x + 1]++;
            d[mx + 1]--;
            // [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
            d[mx + 1] += 2;
        }

        int ans = n;
        int minModify = 0;
        for (int v : d) {
            minModify += v;
            ans = Math.min(ans, minModify);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minChanges(vector<int>& nums, int k) {
        int n = nums.size();
        vector<int> d(k + 2);
        for (int i = 0; i < n / 2; i++) {
            int p = nums[i], q = nums[n - 1 - i];
            if (p > q) { // 保证 p <= q
                swap(p, q);
            }
            int x = q - p;
            int mx = max(q, k - p);
            // [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
            d[0]++;
            d[x]--;
            // [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
            d[x + 1]++;
            d[mx + 1]--;
            // [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
            d[mx + 1] += 2;
        }
        partial_sum(d.begin(), d.end(), d.begin()); // 计算前缀和，得到 minModify 数组
        return ranges::min(d);
    }
};
```

```go [sol-Go]
func minChanges(nums []int, k int) int {
	n := len(nums)
	d := make([]int, k+2)
	for i := 0; i < n/2; i++ {
		p, q := nums[i], nums[n-1-i]
		if p > q { // 保证 p <= q
			p, q = q, p
		}
		x := q - p
		mx := max(q, k-p)
		// [0, x-1] 全部 +1：把 q-p 改成小于 x 的，只需要改 p 或 q 中的一个数
		d[0]++
		d[x]--
		// [x+1, mx] 全部 +1：把 q-p 改成大于 x 小于等于 mx 的，也只需要改 p 或 q 中的一个数
		d[x+1]++
		d[mx+1]--
		// [mx+1, k] 全部 +2：把 q-p 改成大于 mx 的，p 和 q 都需要改
		d[mx+1] += 2
	}

	ans := n
	minModify := 0
	for _, v := range d {
		minModify += v
		ans = min(ans, minModify)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+k)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 相似题目

- [1674. 使数组互补的最少操作次数](https://leetcode.cn/problems/minimum-moves-to-make-array-complementary/)

更多差分题目见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
