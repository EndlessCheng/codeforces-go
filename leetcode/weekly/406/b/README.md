如何在遍历链表的同时，删除链表节点？请看[【基础算法精讲 08】](https://www.bilibili.com/video/BV1VP4y1Q71e/)。

对于本题，由于直接判断节点值是否在 $\textit{nums}$ 中，需要遍历 $\textit{nums}$，时间复杂度为 $\mathcal{O}(n)$。把 $\textit{nums}$ 中的元素保存一个哈希集合中，然后判断节点值是否在哈希集合中，这样可以做到 $\mathcal{O}(1)$。

具体做法：

1. 把 $\textit{nums}$ 中的元素保存到一个哈希集合中。
2. 由于头节点可能会被删除，在头节点前面插入一个哨兵节点 $\textit{dummy}$，以简化代码逻辑。
3. 初始化 $\textit{cur} = \textit{dummy}$。
4. 遍历链表，如果 $\textit{cur}$ 的下一个节点的值在哈希集合中，则需要删除，更新 $\textit{cur}.\textit{next}$ 为 $\textit{cur}.\textit{next}.\textit{next}$；否则不删除，更新 $\textit{cur}$ 为 $\textit{cur}.\textit{next}$。
5. 循环结束后，返回 $\textit{dummy}.\textit{next}$。

> 注：$\textit{dummy}$ 和 $\textit{cur}$ 是同一个节点的引用，修改 $\textit{cur}.\textit{next}$ 也会修改 $\textit{dummy}.\textit{next}$。

[本题视频讲解](https://www.bilibili.com/video/BV1LZ421u7Ut/?t=4m)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        st = set(nums)
        cur = dummy = ListNode(next=head)
        while cur.next:
            nxt = cur.next
            if nxt.val in st:
                cur.next = nxt.next  # 从链表中删除 nxt 节点
            else:
                cur = nxt  # 不删除 nxt，继续向后遍历链表
        return dummy.next
```

```java [sol-Java]
class Solution {
    public ListNode modifiedList(int[] nums, ListNode head) {
        Set<Integer> set = new HashSet<>(nums.length, 1); // 预分配空间
        for (int x : nums) {
            set.add(x);
        }

        ListNode dummy = new ListNode(0, head);
        ListNode cur = dummy;
        while (cur.next != null) {
            ListNode nxt = cur.next;
            if (set.contains(nxt.val)) {
                cur.next = nxt.next; // 从链表中删除 nxt 节点
            } else {
                cur = nxt; // 不删除 nxt，继续向后遍历链表
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
            ListNode* nxt = cur->next;
            if (st.contains(nxt->val)) {
                cur->next = nxt->next; // 从链表中删除 nxt 节点
            } else {
                cur = nxt; // 不删除 nxt，继续向后遍历链表
            }
        }
        return dummy.next;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

struct ListNode* modifiedList(int* nums, int numsSize, struct ListNode* head) {
    int mx = 0;
    for (int i = 0; i < numsSize; i++) {
        mx = MAX(mx, nums[i]);
    }

    bool* has = calloc(mx + 1, sizeof(bool));
    for (int i = 0; i < numsSize; i++) {
        has[nums[i]] = true;
    }

    struct ListNode dummy = {0, head};
    struct ListNode* cur = &dummy;
    while (cur->next) {
        struct ListNode* nxt = cur->next;
        if (nxt->val <= mx && has[nxt->val]) {
            cur->next = nxt->next; // 从链表中删除 nxt 节点
            free(nxt);
        } else {
            cur = nxt; // 不删除 nxt，继续向后遍历链表
        }
    }

    free(has);
    return dummy.next;
}
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
		nxt := cur.Next
		if has[nxt.Val] {
			cur.Next = nxt.Next // 从链表中删除 nxt 节点
		} else {
			cur = nxt // 不删除 nxt，继续向后遍历链表
		}
	}
	return dummy.Next
}
```

```js [sol-JavaScript]
var modifiedList = function(nums, head) {
    const set = new Set(nums);
    const dummy = new ListNode(0, head);
    let cur = dummy;
    while (cur.next) {
        const nxt = cur.next;
        if (set.has(nxt.val)) {
            cur.next = nxt.next; // 从链表中删除 nxt 节点
        } else {
            cur = nxt; // 不删除 nxt，继续向后遍历链表
        }
    }
    return dummy.next;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn modified_list(nums: Vec<i32>, head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let set = nums.into_iter().collect::<HashSet<_>>();
        let mut dummy = Box::new(ListNode { val: 0, next: head });
        let mut cur = &mut dummy;
        while let Some(ref mut nxt) = cur.next {
            if set.contains(&nxt.val) {
                cur.next = nxt.next.take(); // 从链表中删除 nxt 节点
            } else {
                cur = cur.next.as_mut()?; // 不删除 nxt，继续向后遍历链表
            }
        }
        dummy.next
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m$ 是链表的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面链表题单的「**§1.2 删除节点**」。

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
