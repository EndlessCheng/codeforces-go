从最左边开始遍历，第一段的最小值必须是 $1$，最大值必须等于我们遍历过的数的个数，后面依此类推。例如 $\textit{nums}=[3,2,1,6,4,5]$，第一段为 $[3,2,1]$，第二段为 $[6,4,5]$。

每一段排好序后，$\textit{nums}$ 就是有序数组了，所以每一段可以独立计算。用 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/) 找到每一段。

对于每一段，交换次数为这一段的逆序对的个数，[证明](https://leetcode.cn/problems/minimum-adjacent-swaps-to-partition-array/solutions/3998484/ji-suan-ni-xu-dui-de-ge-shu-pythonjavacg-o1ig/)。

长为 $1$ 的段无需排序，其中元素无需解锁。其余元素都需要解锁。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 模板来源 https://leetcode.cn/discuss/post/3583665/
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res += t[i]
            i &= i - 1
        return res

    # 计算区间和 a[l] + ... + a[r]
    # 1 <= l <= r <= n
    # 时间复杂度 O(log n)
    def query(self, l: int, r: int) -> int:
        if l > r:
            return 0
        return self.pre(r) - self.pre(l - 1)

class Solution:
    def minOperations(self, nums: list[int], locked: list[int], k: int) -> int:
        n = len(nums)
        t = FenwickTree(n)
        ans = i = 0
        while i < n:
            if nums[i] == i + 1:  # nums[i] 已在正确位置上
                i += 1
                continue
            start = i
            mn, mx = inf, 0
            # 这一段的最小值必须是 start+1，且值域大小必须等于段长度
            while i < n and (mn != start + 1 or mx - mn + 1 != i - start):
                if locked[i]:
                    if k == 0:  # 无法解锁
                        return -1
                    k -= 1
                    ans += 1
                x = nums[i]
                ans += t.query(x, mx)  # x 与其左侧元素的逆序对个数
                t.update(x, 1)
                mn = min(mn, x)
                mx = max(mx, x)
                i += 1
        return ans
```

```java [sol-Java]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
class FenwickTree {
    private final int[] tree;

    public FenwickTree(int n) {
        tree = new int[n + 1]; // 使用下标 1 到 n
    }

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    public int query(int l, int r) {
        if (l > r) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    public int minOperations(int[] nums, int[] locked, int k) {
        int n = nums.length;
        FenwickTree t = new FenwickTree(n);
        int ans = 0;
        for (int i = 0; i < n; ) {
            if (nums[i] == i + 1) { // nums[i] 已在正确位置上
                i++;
                continue;
            }
            int start = i;
            int mn = n + 1, mx = 0;
            // 这一段的最小值必须是 start+1，且值域大小必须等于段长度
            for (; i < n && (mn != start + 1 || mx - mn + 1 != i - start); i++) {
                if (locked[i] > 0) {
                    if (k == 0) { // 无法解锁
                        return -1;
                    }
                    k--;
                    ans++;
                }
                int x = nums[i];
                ans += t.query(x, mx); // x 与其左侧元素的逆序对个数
                t.update(x, 1);
                mn = Math.min(mn, x);
                mx = Math.max(mx, x);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res{};
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    T query(int l, int r) const {
        if (l > r) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    int minOperations(vector<int>& nums, vector<int>& locked, int k) {
        int n = nums.size();
        FenwickTree<int> t(n);
        int ans = 0;
        for (int i = 0; i < n;) {
            if (nums[i] == i + 1) { // nums[i] 已在正确位置上
                i++;
                continue;
            }
            int start = i;
            int mn = n + 1, mx = 0;
            // 这一段的最小值必须是 start+1，且值域大小必须等于段长度
            for (; i < n && (mn != start + 1 || mx - mn + 1 != i - start); i++) {
                if (locked[i]) {
                    if (k == 0) { // 无法解锁
                        return -1;
                    }
                    k--;
                    ans++;
                }
                int x = nums[i];
                ans += t.query(x, mx); // x 与其左侧元素的逆序对个数
                t.update(x, 1);
                mn = min(mn, x);
                mx = max(mx, x);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
type fenwick []int

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
    for ; i < len(f); i += i & -i {
        f[i] += val
    }
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res += f[i]
    }
    return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
    if l > r {
        return 0
    }
    return f.pre(r) - f.pre(l-1)
}

func minOperations(nums, locked []int, k int) (ans int) {
	n := len(nums)
	t := make(fenwick, n+1)
	for i := 0; i < n; {
		if nums[i] == i+1 { // nums[i] 已在正确位置上
			i++
			continue
		}
		start := i
		mn, mx := n+1, 0
		// 这一段的最小值必须是 start+1，且值域大小必须等于段长度
		for ; i < n && (mn != start+1 || mx-mn+1 != i-start); i++ {
			if locked[i] > 0 {
				if k == 0 { // 无法解锁
					return -1
				}
				k--
				ans++
			}
			x := nums[i]
			ans += t.query(x, mx) // x 与其左侧元素的逆序对个数
			t.update(x, 1)
			mn = min(mn, x)
			mx = max(mx, x)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 滑动窗口与双指针题单的「**六、分组循环**」。
2. 数据结构题单的「**§8.2 逆序对**」。

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
