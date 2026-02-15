## 前置题目/知识点

1. [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)
2. [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)
3. [421. 数组中两个数的最大异或值](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/)

## 思路

「最大值与最小值之间的差值不超过 $k$」这个约束和 1438 题是一样的，用滑动窗口和单调队列解决。

子数组异或和可以用前缀和快速计算。设 $\textit{nums}$ 的**前缀异或和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

题目要求子数组的最大异或和，我们可以枚举子数组的右端点下标 $r$，问题变成：

- 计算子数组 $[i,r]$ 的最大异或和，其中 $i$ 在闭区间 $[l,r]$ 中，其中 $l$ 是滑动窗口的左端点。

这等价于：

- 计算 $s[r+1]$ 与 $i$ 在 $[l,r]$ 中的一个 $s[i]$ 的最大异或和。

这可以用 0-1 字典树解决，见 421 题。

[本题视频讲解](https://www.bilibili.com/video/BV1VgZ4BCETj/?t=26m12s)，欢迎点赞关注~

```py [sol-Python3]
WIDTH = 15  # nums[i] 二进制长度的最大值


class Node:
    __slots__ = 'son', 'leaf'

    def __init__(self):
        self.son = [None] * 2
        self.leaf = 0  # 子树叶子个数


class Trie:
    def __init__(self):
        self.root = Node()

    def put(self, val: int) -> Node:
        cur = self.root
        for i in range(WIDTH - 1, -1, -1):
            bit = val >> i & 1
            if cur.son[bit] is None:
                cur.son[bit] = Node()
            cur = cur.son[bit]
            cur.leaf += 1

    def delete(self, val: int) -> Node:
        cur = self.root
        for i in range(WIDTH - 1, -1, -1):
            cur = cur.son[val >> i & 1]
            cur.leaf -= 1  # 如果减成 0 了，说明子树是空的，可以理解成 cur is None

    def max_xor(self, val: int) -> int:
        cur = self.root
        ans = 0
        for i in range(WIDTH - 1, -1, -1):
            bit = val >> i & 1
            if cur.son[bit ^ 1] and cur.son[bit ^ 1].leaf:
                ans |= 1 << i
                bit ^= 1
            cur = cur.son[bit]
        return ans


class Solution:
    def maxXor(self, nums: list[int], k: int) -> int:
        pre = list(accumulate(nums, xor, initial=0))

        t = Trie()
        min_q = deque()
        max_q = deque()
        ans = left = 0
        for right, x in enumerate(nums):
            # 1. 入
            t.put(pre[right])

            while min_q and x <= nums[min_q[-1]]:
                min_q.pop()
            min_q.append(right)

            while max_q and x >= nums[max_q[-1]]:
                max_q.pop()
            max_q.append(right)

            # 2. 出
            while nums[max_q[0]] - nums[min_q[0]] > k:
                t.delete(pre[left])
                left += 1
                if min_q[0] < left:
                    min_q.popleft()
                if max_q[0] < left:
                    max_q.popleft()

            # 3. 更新答案
            ans = max(ans, t.max_xor(pre[right + 1]))
        return ans
```

```java [sol-Java]
class Trie {
    private static final int WIDTH = 15; // nums[i] 二进制长度的最大值

    private static class Node {
        Node[] son = new Node[2];
        int leaf; // 子树叶子个数
    }

    private final Node root = new Node();

    public void put(int val) {
        Node cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (cur.son[bit] == null) {
                cur.son[bit] = new Node();
            }
            cur = cur.son[bit];
            cur.leaf++;
        }
    }

    public void del(int val) {
        Node cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            cur = cur.son[bit];
            cur.leaf--; // 如果减成 0 了，说明子树是空的，可以理解成 cur is None
        }
    }

    public int maxXor(int val) {
        Node cur = root;
        int ans = 0;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (cur.son[bit ^ 1] != null && cur.son[bit ^ 1].leaf > 0) {
                ans |= 1 << i;
                bit ^= 1;
            }
            cur = cur.son[bit];
        }
        return ans;
    }
}

class Solution {
    public int maxXor(int[] nums, int k) {
        int n = nums.length;
        int[] pre = new int[n + 1];
        for (int i = 0; i < n; i++) {
            pre[i + 1] = pre[i] ^ nums[i];
        }

        Trie t = new Trie();
        ArrayDeque<Integer> minQ = new ArrayDeque<>();
        ArrayDeque<Integer> maxQ = new ArrayDeque<>();
        int ans = 0;
        int left = 0;

        for (int right = 0; right < n; right++) {
            // 1. 入
            t.put(pre[right]);

            int x = nums[right];
            while (!minQ.isEmpty() && x <= nums[minQ.peekLast()]) {
                minQ.pollLast();
            }
            minQ.addLast(right);

            while (!maxQ.isEmpty() && x >= nums[maxQ.peekLast()]) {
                maxQ.pollLast();
            }
            maxQ.addLast(right);

            // 2. 出
            while (nums[maxQ.peekFirst()] - nums[minQ.peekFirst()] > k) {
                t.del(pre[left]);
                left++;
                if (minQ.peekFirst() < left) {
                    minQ.pollFirst();
                }
                if (maxQ.peekFirst() < left) {
                    maxQ.pollFirst();
                }
            }

            // 3. 更新答案
            ans = Math.max(ans, t.maxXor(pre[right + 1]));
        }

        return ans;
    }
}
```

```cpp [sol-C++]
constexpr int WIDTH = 15; // nums[i] 二进制长度的最大值

struct Node {
    Node* son[2]{};
    int leaf = 0; // 子树叶子个数
};

class Trie {
    // 注：我没有写析构函数，想写的同学可以自己补上
    Node* root = new Node();

public:
    void put(int val) {
        Node* cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (!cur->son[bit]) {
                cur->son[bit] = new Node();
            }
            cur = cur->son[bit];
            cur->leaf++;
        }
    }

    void del(int val) {
        Node* cur = root;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            cur = cur->son[bit];
            cur->leaf--; // 如果减成 0 了，说明子树是空的，可以理解成 cur is None
        }
    }

    int max_xor(int val) {
        Node* cur = root;
        int ans = 0;
        for (int i = WIDTH - 1; i >= 0; i--) {
            int bit = val >> i & 1;
            if (cur->son[bit ^ 1] && cur->son[bit ^ 1]->leaf) {
                ans |= 1 << i;
                bit ^= 1;
            }
            cur = cur->son[bit];
        }
        return ans;
    }
};

class Solution {
public:
    int maxXor(vector<int>& nums, int k) {
        int n = nums.size();
        vector<int> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] ^ nums[i];
        }

        Trie t;
        deque<int> min_q, max_q;
        int ans = 0, left = 0;
        for (int right = 0; right < n; right++) {
            // 1. 入
            t.put(sum[right]);

            int x = nums[right];
            while (!min_q.empty() && x <= nums[min_q.back()]) {
                min_q.pop_back();
            }
            min_q.push_back(right);

            while (!max_q.empty() && x >= nums[max_q.back()]) {
                max_q.pop_back();
            }
            max_q.push_back(right);

            // 2. 出
            while (nums[max_q.front()] - nums[min_q.front()] > k) {
                t.del(sum[left]);
                left++;
                if (min_q.front() < left) {
                    min_q.pop_front();
                }
                if (max_q.front() < left) {
                    max_q.pop_front();
                }
            }

            // 3. 更新答案
            ans = max(ans, t.max_xor(sum[right + 1]));
        }
        return ans;
    }
};
```

```go [sol-Go]
const width = 15 // nums[i] 二进制长度的最大值

type node struct {
	son  [2]*node
	leaf int // 子树叶子个数
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t *trie) put(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit] == nil {
			cur.son[bit] = &node{}
		}
		cur = cur.son[bit]
		cur.leaf++
	}
}

func (t *trie) del(val int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		cur = cur.son[val>>i&1]
		cur.leaf-- // 如果减成 0 了，说明子树是空的，可以理解成 cur == nil
	}
}

func (t *trie) maxXor(val int) (ans int) {
	cur := t.root
	for i := width - 1; i >= 0; i-- {
		bit := val >> i & 1
		if cur.son[bit^1] != nil && cur.son[bit^1].leaf > 0 {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.son[bit]
	}
	return
}

func maxXor(nums []int, k int) (ans int) {
	sum := make([]int, len(nums)+1)
	for i, x := range nums {
		sum[i+1] = sum[i] ^ x
	}

	t := newTrie()
	var minQ, maxQ []int
	left := 0
	for right, x := range nums {
		// 1. 入
		t.put(sum[right])

		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, right)

		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, right)

		// 2. 出
		for nums[maxQ[0]]-nums[minQ[0]] > k {
			t.del(sum[left])
			left++
			if minQ[0] < left {
				minQ = minQ[1:]
			}
			if maxQ[0] < left {
				maxQ = maxQ[1:]
			}
		}

		// 3. 更新答案
		ans = max(ans, t.maxXor(sum[right+1]))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 专题训练

见下面数据结构题单的「**§4.4 单调队列**」和「**§6.4 0-1 字典树**」。

其中和本题相似的题目是 [2935. 找出强数对的最大异或值 II](https://leetcode.cn/problems/maximum-strong-pair-xor-ii/)。

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
