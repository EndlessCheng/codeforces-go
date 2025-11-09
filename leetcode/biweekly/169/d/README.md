## 转化

类似 [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)，把 $\textit{nums}$ 中的 $\textit{target}$ 视作 $1$，其余元素视作 $-1$，得到一个新数组 $a$，问题变成：

- 计算 $a$ 中元素和严格大于 $0$ 的子数组个数。

设 $a$ 的前缀和数组为 $s$。关于 $s$ 的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

问题变成：

- 计算有多少对 $(i,j)$ 满足 $0\le i  < j\le n$ 且 $s[j] - s[i] > 0$。

枚举 $j$，问题变成：

- 计算 $s$ 的下标区间 $[0,j-1]$ 中有多少个 $s[i] < s[j]$。

## 方法一：枚举右，有序集合维护左

枚举 $j$ 的同时，用有序集合维护左边的 $s[i]$。

在有序集合中查找第一个 $\ge s[j]$ 的数的位置，即为 $s[i] < s[j]$ 的个数。

[本题视频讲解](https://www.bilibili.com/video/BV19bkQBkEhG/?t=24m30s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        sl = SortedList([0])  # 为什么加个 0？见 525 题我的题解
        ans = s = 0
        for x in nums:
            s += 1 if x == target else -1
            ans += sl.bisect_left(s)
            sl.add(s)
        return ans
```

```cpp [sol-C++]
#include <ext/pb_ds/assoc_container.hpp>

using namespace __gnu_pbds;
// 使用 pair<key, index> 支持重复 key
using ordered_set = tree<pair<int, int>, null_type, less<>, rb_tree_tag, tree_order_statistics_node_update>;

class Solution {
public:
    long long countMajoritySubarrays(vector<int>& nums, int target) {
        int idx = 0; // 插入时自增，用来保证 st 中的元素互不相同
        ordered_set* st = new ordered_set();
        st->insert({0, ++idx});
        long long ans = 0;
        int s = 0;
        for (int x : nums) {
            s += x == target ? 1 : -1;
            // order_of_key(key) 计算 st 中的严格小于 key 的元素个数
            ans += st->order_of_key({s, 0});
            st->insert({s, ++idx});
        }
        delete st;
        return ans;
    }
};
```

```go [sol-Go]
// 我的 Treap 模板之一，相当于 std::multiset + 名次树
type nodeMS[K comparable] struct {
	son      [2]*nodeMS[K]
	priority uint
	key      K
	keyCnt   int
	subSize  int
}

func (o *nodeMS[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeMS[K]) maintain() {
	o.subSize = o.keyCnt + o.son[0].size() + o.son[1].size()
}

func (o *nodeMS[K]) rotate(d int) *nodeMS[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapMS[K comparable] struct {
	rd         uint
	root       *nodeMS[K]
	comparator func(a, b K) int
}

func (t *treapMS[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapMS[K]) size() int   { return t.root.size() }
func (t *treapMS[K]) empty() bool { return t.size() == 0 }

func (t *treapMS[K]) _put(o *nodeMS[K], key K) *nodeMS[K] {
	if o == nil {
		o = &nodeMS[K]{priority: t.fastRand(), key: key, keyCnt: 1}
	} else {
		c := t.comparator(key, o.key)
		if c == 0 {
			o.keyCnt++
		} else {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treapMS[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapMS[K]) _delete(o *nodeMS[K], key K) *nodeMS[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.keyCnt > 1 {
			o.keyCnt--
		} else {
			if o.son[1] == nil {
				return o.son[0]
			}
			if o.son[0] == nil {
				return o.son[1]
			}
			d := 0
			if o.son[0].priority > o.son[1].priority {
				d = 1
			}
			o = o.rotate(d)
			o.son[d] = t._delete(o.son[d], key)
		}
	}
	o.maintain()
	return o
}

func (t *treapMS[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapMS[K]) min() *nodeMS[K] { return t.kth(0) }
func (t *treapMS[K]) max() *nodeMS[K] { return t.kth(t.size() - 1) }

func (t *treapMS[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapMS[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + o.keyCnt
			o = o.son[1]
		} else {
			kth += o.son[0].size() + o.keyCnt
			break
		}
	}
	return
}

func (t *treapMS[K]) kth(k int) (o *nodeMS[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + o.keyCnt
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func (t *treapMS[K]) prev(key K) *nodeMS[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapMS[K]) next(key K) *nodeMS[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapMS[K]) find(key K) *nodeMS[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newMultiset[K cmp.Ordered]() *treapMS[K] {
	return &treapMS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func newMultisetWith[K comparable](comp func(a, b K) int) *treapMS[K] {
	return &treapMS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}

func countMajoritySubarrays(nums []int, target int) (ans int64) {
	t := newMultiset[int]()
	t.put(0) // 为什么加个 0？见 525 题我的题解
	s := 0
	for _, x := range nums {
		if x == target {
			s++
		} else {
			s--
		}
		ans += int64(t.lowerBoundIndex(s))
		t.put(s)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二

本题还有一个特殊性质：每次循环，$s$ 的变化量只有 $1$。

定义 $f[j]$ 表示满足 $i<j$ 且 $s[i] < s[j]$ 的 $i$ 的个数（即方法一每次循环加到 $\textit{ans}$ 中的值）。

如果知道 $f[j-1]$，能不能 $\mathcal{O}(1)$ 地把 $f[j]$ 算出来？

如果 $s[j] = s[j-1] + 1$，我们可以在 $f[j-1]$ 的基础上，加上恰好等于 $s[j]-1$ 的 $s[i]$ 的个数，就得到了 $f[j]$。

这启发我们统计 $s[j]$ 的出现次数 $\textit{cnt}[s[j]]$。

- 如果 $s[j] = s[j-1] + 1$，我们可以在 $f[j-1]$ 的基础上，加上恰好等于 $s[j]-1$ 的 $s[i]$ 的个数 $\textit{cnt}[s[j]-1]$，得到 $f[j] = f[j-1] + \textit{cnt}[s[j]-1] = f[j-1] + \textit{cnt}[s[j-1]]$。
- 如果 $s[j] = s[j-1] - 1$，我们可以在 $f[j-1]$ 的基础上，减去恰好等于 $s[j]$ 的 $s[i]$ 的个数 $\textit{cnt}[s[j]]$，得到 $f[j] = f[j-1] - \textit{cnt}[s[j]]$。

答案为 $f$ 的元素和。

## 优化前

```py [sol-Python3]
class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        cnt = defaultdict(int)
        cnt[0] = 1  # 为什么加个 0？见 525 题我的题解
        ans = s = f = 0
        for x in nums:
            if x == target:
                f += cnt[s]
                s += 1
            else:
                s -= 1
                f -= cnt[s]
            ans += f
            cnt[s] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countMajoritySubarrays(int[] nums, int target) {
        Map<Integer, Integer> cnt = new HashMap<>();
        cnt.put(0, 1); // 为什么加个 0？见 525 题我的题解
        long ans = 0;
        int s = 0;
        int f = 0;
        for (int x : nums) {
            if (x == target) {
                f += cnt.getOrDefault(s, 0);
                s++;
            } else {
                s--;
                f -= cnt.getOrDefault(s, 0);
            }
            ans += f;
            cnt.merge(s, 1, Integer::sum); // cnt[s]++
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countMajoritySubarrays(vector<int>& nums, int target) {
        unordered_map<int, int> cnt = {{0, 1}}; // 为什么加个 0？见 525 题我的题解
        long long ans = 0;
        int s = 0, f = 0;
        for (int x : nums) {
            if (x == target) {
                f += cnt[s];
                s++;
            } else {
                s--;
                f -= cnt[s];
            }
            ans += f;
            cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countMajoritySubarrays(nums []int, target int) (ans int64) {
	cnt := map[int]int{0: 1} // 为什么加个 0？见 525 题我的题解
	s, f := 0, 0
	for _, x := range nums {
		if x == target {
			f += cnt[s]
			s++
		} else {
			s--
			f -= cnt[s]
		}
		ans += int64(f)
		cnt[s]++
	}
	return
}
```

## 优化

用数组代替哈希表。

为避免下标越界，可以把 $s$ 初始化成 $n$。

```py [sol-Python3]
class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        s = n = len(nums)
        cnt = [0] * (n * 2 + 1)
        cnt[s] = 1
        ans = f = 0
        for x in nums:
            if x == target:
                f += cnt[s]
                s += 1
            else:
                s -= 1
                f -= cnt[s]
            ans += f
            cnt[s] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countMajoritySubarrays(int[] nums, int target) {
        int n = nums.length;
        int[] cnt = new int[n * 2 + 1];
        cnt[n] = 1;
        long ans = 0;
        int s = n;
        int f = 0;
        for (int x : nums) {
            if (x == target) {
                f += cnt[s];
                s++;
            } else {
                s--;
                f -= cnt[s];
            }
            ans += f;
            cnt[s]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countMajoritySubarrays(vector<int>& nums, int target) {
        int n = nums.size();
        vector<int> cnt(n * 2 + 1);
        cnt[n] = 1;
        long long ans = 0;
        int s = n, f = 0;
        for (int x : nums) {
            if (x == target) {
                f += cnt[s];
                s++;
            } else {
                s--;
                f -= cnt[s];
            }
            ans += f;
            cnt[s]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countMajoritySubarrays(nums []int, target int) (ans int64) {
	n := len(nums)
	cnt := make([]int, n*2+1)
	cnt[n] = 1
	s, f := n, 0
	for _, x := range nums {
		if x == target {
			f += cnt[s]
			s++
		} else {
			s--
			f -= cnt[s]
		}
		ans += int64(f)
		cnt[s]++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

改成计算**存在**绝对众数（出现次数严格大于子数组长度的一半）的子数组个数。

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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
