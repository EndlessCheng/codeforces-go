请看 [视频讲解](https://www.bilibili.com/video/BV1wh4y1Q7XW/) 第二题。

## 方法一

看成是 $\textit{head}$ 与 $\textit{head}$ 这两个链表相加。

直接调用 [445. 两数相加 II](https://leetcode.cn/problems/add-two-numbers-ii/solution/fan-zhuan-lian-biao-liang-shu-xiang-jia-okw6q/) 的代码即可。

```py [sol-Python3]
class Solution:
    # 206. 反转链表
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        pre = None
        cur = head
        while cur:
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        return pre

    # 2. 两数相加：自己和自己相加
    # 题解 https://leetcode.cn/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
    def double(self, l1: Optional[ListNode]) -> Optional[ListNode]:
        cur = dummy = ListNode()  # 哨兵节点，作为新链表的头节点的前一个节点
        carry = 0  # 进位
        while l1:  # 有一个不是空节点，或者还有进位，就继续迭代
            carry += l1.val * 2  # 节点值和进位加在一起
            cur.next = ListNode(carry % 10)  # 每个节点保存一个数位
            carry //= 10  # 新的进位
            cur = cur.next  # 下一个节点
            l1 = l1.next  # 下一个节点
        if carry:
            cur.next = ListNode(carry)
        return dummy.next  # 哨兵节点的下一个节点就是头节点

    def doubleIt(self, head: Optional[ListNode]) -> Optional[ListNode]:
        head = self.reverseList(head)
        res = self.double(head)  # 反转后，就变成【2. 两数相加】了
        return self.reverseList(res)
```

```java [sol-Java]
class Solution {
    public ListNode doubleIt(ListNode head) {
        head = reverseList(head);
        ListNode res = double2(head); // 反转后，就变成【2. 两数相加】了
        return reverseList(res);
    }

    // 206. 反转链表
    private ListNode reverseList(ListNode head) {
        ListNode pre = null;
        ListNode cur = head;
        while (cur != null) {
            ListNode nxt = cur.next;
            cur.next = pre;
            pre = cur;
            cur = nxt;
        }
        return pre;
    }

    // 2. 两数相加：自己和自己相加
    // 题解 https://leetcode.cn/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
    private ListNode double2(ListNode l1) {
        ListNode dummy = new ListNode(); // 哨兵节点，作为新链表的头节点的前一个节点
        ListNode cur = dummy;
        int carry = 0; // 进位
        while (l1 != null) {
            carry += l1.val * 2; // 节点值和进位加在一起
            cur.next = new ListNode(carry % 10); // 每个节点保存一个数位
            carry /= 10; // 新的进位
            cur = cur.next; // 下一个节点
            l1 = l1.next; // 下一个节点
        }
        if (carry != 0) {
            cur.next = new ListNode(carry);
        }
        return dummy.next; // 哨兵节点的下一个节点就是头节点
    }
}
```

```cpp [sol-C++]
class Solution {
    // 206. 反转链表
    ListNode* reverseList(ListNode* head) {
        ListNode* pre = nullptr, *cur = head;
        while (cur) {
            auto nxt = cur->next;
            cur->next = pre;
            pre = cur;
            cur = nxt;
        }
        return pre;
    }

    // 2. 两数相加：自己和自己相加
    // 题解 https://leetcode.cn/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
    ListNode* double_(ListNode* l1) {
        ListNode dummy; // 哨兵节点，作为新链表的头节点的前一个节点
        auto cur = &dummy;
        int carry = 0; // 进位
        while (l1) {
            carry += l1->val * 2; // 节点值和进位加在一起
            cur->next = new ListNode(carry % 10); // 每个节点保存一个数位
            carry /= 10; // 新的进位
            cur = cur->next; // 下一个节点
            l1 = l1->next; // 下一个节点
        }
        if (carry) {
            cur->next = new ListNode(carry);
        }
        return dummy.next;
    }

public:
    ListNode* doubleIt(ListNode* head) {
        head = reverseList(head);
        auto res = double_(head);
        return reverseList(res);
    }
};
```

```go [sol-Go]
// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 2. 两数相加：自己和自己相加
// 题解 https://leetcode.cn/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
func double(l1 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点，作为新链表的头节点的前一个节点
	cur := dummy
	carry := 0 // 进位
	for l1 != nil {
		carry += l1.Val * 2                   // 节点值和进位加在一起
		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry /= 10                           // 新的进位
		cur = cur.Next                        // 下一个节点
		l1 = l1.Next                          // 下一个节点
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next // 哨兵节点的下一个节点就是头节点
}

func doubleIt(head *ListNode) *ListNode {
	head = reverseList(head)
	res := double(head) // 反转后，就变成【2. 两数相加】了
	return reverseList(res)
}
```

```js [sol-JavaScript]
// 206. 反转链表
var reverseList = function(head) {
    let pre = null;
    let cur = head;
    while (cur) {
        let nxt = cur.next;
        cur.next = pre;
        pre = cur;
        cur = nxt;
    }
    return pre;
};

// 2. 两数相加：自己和自己相加
// 题解 https://leetcode.cn/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
var double = function(l1) {
    let dummy = new ListNode(); // 哨兵节点，作为新链表的头节点的前一个节点
    let cur = dummy;
    let carry = 0; // 进位
    while (l1) {
        carry += l1.val * 2; // 节点值和进位加在一起
        cur.next = new ListNode(carry % 10); // 每个节点保存一个数位
        carry = Math.floor(carry / 10); // 新的进位
        cur = cur.next; // 下一个节点
        l1 = l1.next; // 下一个节点
    }
    if (carry) {
        cur.next = new ListNode(carry);
    }
    return dummy.next; // 哨兵节点的下一个节点就是头节点
};

var doubleIt = function(head) {
    head = reverseList(head);
    const res = double(head); // 反转后，就变成【2. 两数相加】了
    return reverseList(res);
}
```

```rust [sol-Rust]
impl Solution {
    // 206. 反转链表
    fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut pre = None;
        let mut cur = head;
        while let Some(mut node) = cur {
            let nxt = node.next.take();
            node.next = pre;
            pre = Some(node);
            cur = nxt;
        }
        pre
    }

    // 2. 两数相加：自己和自己相加
    // 题解 https://leetcode-cn.com/problems/add-two-numbers/solution/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
    fn double(l1: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = Some(Box::new(ListNode::new(0))); // 哨兵节点，作为新链表的头节点的前一个节点
        let mut cur = &mut dummy;
        let mut carry = 0; // 进位
        let mut l1 = l1;
        while let Some(mut node) = l1 {
            carry += node.val * 2; // 节点值和进位加在一起
            cur.as_mut()?.next = Some(Box::new(ListNode::new(carry % 10)));// 每个节点保存一个数位
            carry /= 10; // 新的进位
            cur = &mut cur.as_mut()?.next; // 下一个节点
            l1 = node.next.take(); // 下一个节点
        }
        if carry != 0 {
            cur.as_mut()?.next = Some(Box::new(ListNode::new(carry)));
        }
        dummy?.next // 哨兵节点的下一个节点就是头节点
    }

    pub fn double_it(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let head = Self::reverse_list(head);
        let res = Self::double(head); // 反转后，就变成【2. 两数相加】了
        Self::reverse_list(res)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 方法二

如果不考虑进位，就是每个节点的值乘以 $2$。

什么时候会受到进位的影响呢？只有下一个节点大于 $4$ 的时候，才会因为进位多加一。

特别地，如果链表头的值大于 $4$，那么需要在前面插入一个新的节点。

```py [sol-Python3]
class Solution:
    def doubleIt(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head.val > 4:
            head = ListNode(0, head)
        cur = head
        while cur:
            cur.val = cur.val * 2 % 10
            if cur.next and cur.next.val > 4:
                cur.val += 1
            cur = cur.next
        return head
```

```java [sol-Java]
class Solution {
    public ListNode doubleIt(ListNode head) {
        if (head.val > 4) {
            head = new ListNode(0, head);
        }
        for (ListNode cur = head; cur != null; cur = cur.next) {
            cur.val = cur.val * 2 % 10;
            if (cur.next != null && cur.next.val > 4) {
                cur.val++;
            }
        }
        return head;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    ListNode* doubleIt(ListNode* head) {
        if (head->val > 4) {
            head = new ListNode(0, head);
        }
        for (auto cur = head; cur; cur = cur->next) {
            cur->val = cur->val * 2 % 10;
            if (cur->next && cur->next->val > 4) {
                cur->val++;
            }
        }
        return head;
    }
};
```

```go [sol-Go]
func doubleIt(head *ListNode) *ListNode {
	if head.Val > 4 {
		head = &ListNode{0, head}
	}
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = cur.Val * 2 % 10
		if cur.Next != nil && cur.Next.Val > 4 {
			cur.Val++
		}
	}
	return head
}
```

```js [sol-JavaScript]
var doubleIt = function(head) {
    if (head.val > 4) {
        head = new ListNode(0, head);
    }
    for (let cur = head; cur; cur = cur.next) {
        cur.val = cur.val * 2 % 10;
        if (cur.next && cur.next.val > 4) {
            cur.val++;
        }
    }
    return head;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn double_it(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.as_ref()?.val > 4 {
            head = Some(Box::new(ListNode { val: 0, next: head }));
        }
        let mut cur = head.as_mut();
        while let Some(node) = cur {
            node.val = node.val * 2 % 10;
            if let Some(next) = node.next.as_mut() {
                if next.val > 4 {
                    node.val += 1;
                }
            }
            cur = node.next.as_mut();
        }
        head
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
