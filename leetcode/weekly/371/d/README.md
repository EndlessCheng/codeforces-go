请看 [视频讲解](https://www.bilibili.com/video/BV1MG411X7zR/) 第四题。

## 方法一：0-1 trie + 滑动窗口

由于答案和 $\textit{nums}$ 的元素顺序无关，先排序。

排序后设 $x\le y$，那么 $|x-y|\le \min(x,y)$ 可以化简为

$$
2x \ge y
$$

这意味着对于每个 $y=\textit{nums}[i]$，我们需要选择 $y$ 及其左边的满足 $2x\ge y$ 的 $x$，与 $y$ 异或，求最大异或和。这可以用 0-1 trie 实现。

由于 $y$ 越大，能选到的最小 $x$ 也越大，所以需要用**滑动窗口**。每次滑出元素时，把它从 0-1 trie 中删掉。

```py [sol-Python3]
class Node:
    __slots__ = 'children', 'cnt'

    def __init__(self):
        self.children = [None, None]
        self.cnt = 0  # 子树大小

class Trie:
    HIGH_BIT = 19

    def __init__(self):
        self.root = Node()

    # 添加 val
    def insert(self, val: int) -> None:
        cur = self.root
        for i in range(Trie.HIGH_BIT, -1, -1):
            bit = (val >> i) & 1
            if cur.children[bit] is None:
                cur.children[bit] = Node()
            cur = cur.children[bit]
            cur.cnt += 1  # 维护子树大小
        return cur

    # 删除 val，但不删除节点
    # 要求 val 必须在 trie 中
    def remove(self, val: int) -> None:
        cur = self.root
        for i in range(Trie.HIGH_BIT, -1, -1):
            cur = cur.children[(val >> i) & 1]
            cur.cnt -= 1  # 维护子树大小
        return cur

    # 返回 val 与 trie 中一个元素的最大异或和
    # 要求 trie 中至少有一个元素
    def max_xor(self, val: int) -> int:
        cur = self.root
        ans = 0
        for i in range(Trie.HIGH_BIT, -1, -1):
            bit = (val >> i) & 1
            # 如果 cur.children[bit^1].cnt == 0，视作空节点
            if cur.children[bit ^ 1] and cur.children[bit ^ 1].cnt:
                ans |= 1 << i
                bit ^= 1
            cur = cur.children[bit]
        return ans

class Solution:
    def maximumStrongPairXor(self, nums: List[int]) -> int:
        nums.sort()
        t = Trie()
        ans = left = 0
        for y in nums:
            t.insert(y)
            while nums[left] * 2 < y:
                t.remove(nums[left])
                left += 1
            ans = max(ans, t.max_xor(y))
        return ans
```

```java [sol-Java]
class Node {
    Node[] children = new Node[2];
    int cnt; // 子树大小
}

class Trie {
    private static final int HIGH_BIT = 19;
    private Node root = new Node();

    // 添加 val
    public void insert(int val) {
        Node cur = root;
        for (int i = HIGH_BIT; i >= 0; i--) {
            int bit = (val >> i) & 1;
            if (cur.children[bit] == null) {
                cur.children[bit] = new Node();
            }
            cur = cur.children[bit];
            cur.cnt++; // 维护子树大小
        }
    }

    // 删除 val，但不删除节点
    // 要求 val 必须在 trie 中
    public void remove(int val) {
        Node cur = root;
        for (int i = HIGH_BIT; i >= 0; i--) {
            cur = cur.children[(val >> i) & 1];
            cur.cnt--; // 维护子树大小
        }
    }

    // 返回 val 与 trie 中一个元素的最大异或和
    // 要求 trie 不能为空
    public int maxXor(int val) {
        Node cur = root;
        int ans = 0;
        for (int i = HIGH_BIT; i >= 0; i--) {
            int bit = (val >> i) & 1;
            // 如果 cur.children[bit^1].cnt == 0，视作空节点
            if (cur.children[bit ^ 1] != null && cur.children[bit ^ 1].cnt > 0) {
                ans |= 1 << i;
                bit ^= 1;
            }
            cur = cur.children[bit];
        }
        return ans;
    }
}

class Solution {
    public int maximumStrongPairXor(int[] nums) {
        Arrays.sort(nums);
        Trie t = new Trie();
        int ans = 0, left = 0;
        for (int y : nums) {
            t.insert(y);
            while (nums[left] * 2 < y) {
                t.remove(nums[left++]);
            }
            ans = Math.max(ans, t.maxXor(y));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Node {
public:
    array<Node*, 2> children{};
    int cnt = 0; // 子树大小
};

class Trie {
    static const int HIGH_BIT = 19;
public:
    Node *root = new Node();

    // 添加 val
    void insert(int val) {
        Node *cur = root;
        for (int i = HIGH_BIT; i >= 0; i--) {
            int bit = (val >> i) & 1;
            if (cur->children[bit] == nullptr) {
                cur->children[bit] = new Node();
            }
            cur = cur->children[bit];
            cur->cnt++; // 维护子树大小
        }
    }

    // 删除 val，但不删除节点
    // 要求 val 必须在 trie 中
    void remove(int val) {
        Node *cur = root;
        for (int i = HIGH_BIT; i >= 0; i--) {
            cur = cur->children[(val >> i) & 1];
            cur->cnt--; // 维护子树大小
        }
    }

    // 返回 val 与 trie 中一个元素的最大异或和
    // 要求 trie 不能为空
    int max_xor(int val) {
        Node *cur = root;
        int ans = 0;
        for (int i = HIGH_BIT; i >= 0; i--) {
            int bit = (val >> i) & 1;
            // 如果 cur.children[bit^1].cnt == 0，视作空节点
            if (cur->children[bit ^ 1] && cur->children[bit ^ 1]->cnt) {
                ans |= 1 << i;
                bit ^= 1;
            }
            cur = cur->children[bit];
        }
        return ans;
    }
};

class Solution {
public:
    int maximumStrongPairXor(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        Trie t{};
        int ans = 0, left = 0;
        for (int y: nums) {
            t.insert(y);
            while (nums[left] * 2 < y) {
                t.remove(nums[left++]);
            }
            ans = max(ans, t.max_xor(y));
        }
        return ans;
    }
};
```

```go [sol-Go]
const highBit = 19

type node struct {
	children [2]*node
	cnt      int // 子树大小
}

type trie struct{ root *node }

// 添加 val
func (t trie) insert(val int) *node {
	cur := t.root
	for i := highBit; i >= 0; i-- {
		bit := val >> i & 1
		if cur.children[bit] == nil {
			cur.children[bit] = &node{}
		}
		cur = cur.children[bit]
		cur.cnt++ // 维护子树大小
	}
	return cur
}

// 删除 val，但不删除节点
// 要求 val 必须在 t 中
func (t trie) remove(val int) *node {
	cur := t.root
	for i := highBit; i >= 0; i-- {
		cur = cur.children[val>>i&1]
		cur.cnt-- // 维护子树大小
	}
	return cur
}

// 返回 val 与 t 中一个元素的最大异或和
// 要求 t 不能为空
func (t trie) maxXor(val int) (ans int) {
	cur := t.root
	for i := highBit; i >= 0; i-- {
		bit := val >> i & 1
		// 如果 cur.children[bit^1].cnt == 0，视作空节点
		if cur.children[bit^1] != nil && cur.children[bit^1].cnt > 0 {
			ans |= 1 << i
			bit ^= 1
		}
		cur = cur.children[bit]
	}
	return
}

func maximumStrongPairXor(nums []int) (ans int) {
	slices.Sort(nums)
	t := trie{&node{}}
	left := 0
	for _, y := range nums {
		t.insert(y)
		for nums[left]*2 < y {
			t.remove(nums[left])
			left++
		}
		ans = max(ans, t.maxXor(y))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，本题 $U=2^{20}-1$，也就是说 $\textit{nums}[i]$ 二进制长度不会超过 $20$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 方法二：哈希表

原理请看我的这篇题解：[【图解】简洁高效，一图秒懂！](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solution/tu-jie-jian-ji-gao-xiao-yi-tu-miao-dong-1427d/)

把 hashset 改成 hashmap，一边遍历数组，一边记录每个 key 对应的最大的 $\textit{nums}[i]$。

由于数组已经排好序，所以每个 key 对应的 $x=\textit{nums}[i]$ 一定是当前最大的，只要 $2x\ge y$，就说明这个比特位可以是 $1$。

```py [sol-Python3]
class Solution:
    def maximumStrongPairXor(self, nums: List[int]) -> int:
        nums.sort()
        ans = mask = 0
        high_bit = nums[-1].bit_length() - 1
        for i in range(high_bit, -1, -1):  # 从最高位开始枚举
            mask |= 1 << i
            new_ans = ans | (1 << i)  # 这个比特位可以是 1 吗？
            d = {}
            for y in nums:
                mask_y = y & mask  # 低于 i 的比特位置为 0
                if new_ans ^ mask_y in d and d[new_ans ^ mask_y] * 2 >= y:
                    ans = new_ans  # 这个比特位可以是 1
                    break
                d[mask_y] = y
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumStrongPairXor(int[] nums) {
        Arrays.sort(nums);
        int highBit = 31 - Integer.numberOfLeadingZeros(nums[nums.length - 1]);

        int ans = 0, mask = 0;
        Map<Integer, Integer> mp = new HashMap<>();
        for (int i = highBit; i >= 0; i--) { // 从最高位开始枚举
            mp.clear();
            mask |= 1 << i;
            int newAns = ans | (1 << i); // 这个比特位可以是 1 吗？
            for (int y : nums) {
                int maskY = y & mask; // 低于 i 的比特位置为 0
                if (mp.containsKey(newAns ^ maskY) && mp.get(newAns ^ maskY) * 2 >= y) {
                    ans = newAns; // 这个比特位可以是 1
                    break;
                }
                mp.put(maskY, y);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumStrongPairXor(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int high_bit = 31 - __builtin_clz(nums.back());

        int ans = 0, mask = 0;
        unordered_map<int, int> mp;
        for (int i = high_bit; i >= 0; i--) { // 从最高位开始枚举
            mp.clear();
            mask |= 1 << i;
            int new_ans = ans | (1 << i); // 这个比特位可以是 1 吗？
            for (int y: nums) {
                int mask_y = y & mask; // 低于 i 的比特位置为 0
                auto it = mp.find(new_ans ^ mask_y);
                if (it != mp.end() && it->second * 2 >= y) {
                    ans = new_ans; // 这个比特位可以是 1
                    break;
                }
                mp[mask_y] = y;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumStrongPairXor(nums []int) (ans int) {
	slices.Sort(nums)
	highBit := bits.Len(uint(nums[len(nums)-1])) - 1
	mp := map[int]int{}
	mask := 0
	for i := highBit; i >= 0; i-- { // 从最高位开始枚举
		clear(mp)
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		for _, y := range nums {
			maskY := y & mask // 低于 i 的比特位置为 0
			if x, ok := mp[newAns^maskY]; ok && x*2 >= y {
				ans = newAns // 这个比特位可以是 1
				break
			}
			mp[maskY] = y
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，本题 $U=2^{20}-1$，也就是说 $\textit{nums}[i]$ 二进制长度不会超过 $20$。
- 空间复杂度：$\mathcal{O}(n)$。

## 练习：0-1 trie（右边分数为题目难度）

- [1707. 与数组中元素的最大异或值](https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/) 2359
- [1803. 统计异或值在范围内的数对有多少](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/) 2479
- [1938. 查询最大基因差](https://leetcode.cn/problems/maximum-genetic-difference-query/) 2503
- [2479. 两个不重叠子树的最大异或值](https://leetcode.cn/problems/maximum-xor-of-two-non-overlapping-subtrees/)（会员题）

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
