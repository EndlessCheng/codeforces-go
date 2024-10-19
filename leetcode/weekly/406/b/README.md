如何在遍历链表的同时，删除链表节点？请看[【基础算法精讲 08】](https://www.bilibili.com/video/BV1VP4y1Q71e/)。

对于本题，由于直接判断节点值是否在 $\textit{nums}$ 中，需要遍历 $\textit{nums}$，时间复杂度为 $\mathcal{O}(n)$。通过把 $\textit{nums}$ 中的元素加到一个哈希集合中，然后判断节点值是否在哈希集合中，这样可以做到每次判断时间复杂度为 $\mathcal{O}(1)$。

具体做法：

1. 把 $\textit{nums}$ 中的元素加到一个哈希集合中。
2. 由于头节点可能会被删除，在头节点前面插入一个哨兵节点 $\textit{dummy}$，以简化代码逻辑。
3. 初始化 $\textit{cur} = \textit{dummy}$。
4. 不断循环，直到 $\textit{cur}$ 没有下一个节点。
5. 如果 $\textit{cur}$ 的下一个节点的值在哈希集合中，则需要删除，更新 $\textit{cur}.\textit{next}$ 为 $\textit{cur}.\textit{next}.\textit{next}$；否则不删除，更新 $\textit{cur}$ 为 $\textit{cur}.\textit{next}$。
6. 循环结束后，返回 $\textit{dummy}.\textit{next}$。

⚠注意：$\textit{dummy}$ 和 $\textit{cur}$ 是同一个节点的引用，修改 $\textit{cur}.\textit{next}$ 也会修改 $\textit{dummy}.\textit{next}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1LZ421u7Ut/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        st = set(nums)
        cur = dummy = ListNode(next=head)
        while cur.next:
            if cur.next.val in st:
                cur.next = cur.next.next  # 删除
            else:
                cur = cur.next  # 向后移动
        return dummy.next
```

```java [sol-Java]
class Solution {
    public ListNode modifiedList(int[] nums, ListNode head) {
        Set<Integer> set = new HashSet<>(nums.length); // 预分配空间
        for (int x : nums) {
            set.add(x);
        }
        ListNode dummy = new ListNode(0, head);
        ListNode cur = dummy;
        while (cur.next != null) {
            if (set.contains(cur.next.val)) {
                cur.next = cur.next.next; // 删除
            } else {
                cur = cur.next; // 向后移动
            }
        }
        return dummy.next;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    ListNode* modifiedList(vector<int>& nums, ListNode* head) {
        unordered_set<int> st(nums.begin(), nums.end());
        ListNode dummy(0, head);
        ListNode* cur = &dummy;
        while (cur->next) {
            if (st.contains(cur->next->val)) {
                cur->next = cur->next->next; // 删除
                // 注意力扣会在 modifiedList 调用结束后回收所有节点，自己手动删除反而不行
            } else {
                cur = cur->next; // 向后移动
            }
        }
        return dummy.next;
    }
};
```

```go [sol-Go]
func modifiedList(nums []int, head *ListNode) *ListNode {
	has := make(map[int]bool, len(nums)) // 预分配空间
	for _, x := range nums {
		has[x] = true
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if has[cur.Next.Val] {
			cur.Next = cur.Next.Next // 删除
		} else {
			cur = cur.Next // 向后移动
		}
	}
	return dummy.Next
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是链表的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

见[【基础算法精讲 08】](https://www.bilibili.com/video/BV1VP4y1Q71e/)视频简介。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
