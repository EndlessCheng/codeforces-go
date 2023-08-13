下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

## 方法一

看成是 $\textit{head}$ 与 $\textit{head}$ 这两个链表相加。

直接调用 [445. 两数相加 II](https://leetcode.cn/problems/add-two-numbers-ii/solution/fan-zhuan-lian-biao-liang-shu-xiang-jia-okw6q/) 的代码即可。

```py [sol-Python3]
# https://space.bilibili.com/206214
class Solution:
    # 206. 反转链表
    # 视频讲解 https://www.bilibili.com/video/BV1sd4y1x7KN/
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
        cur = dummy = ListNode()  # 哨兵节点
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

```go [sol-Go]
// 206. 反转链表
// 视频讲解 https://www.bilibili.com/video/BV1sd4y1x7KN/
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
	return dummy.Next
}

func doubleIt(head *ListNode) *ListNode {
	head = reverseList(head)
	res := double(head) // 反转后，就变成【2. 两数相加】了
	return reverseList(res)
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
        if (head.val > 4)
            head = new ListNode(0, head);
        for (var cur = head; cur != null; cur = cur.next) {
            cur.val = cur.val * 2 % 10;
            if (cur.next != null && cur.next.val > 4)
                cur.val++;
        }
        return head;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    ListNode *doubleIt(ListNode *head) {
        if (head->val > 4)
            head = new ListNode(0, head);
        for (auto cur = head; cur; cur = cur->next) {
            cur->val = cur->val * 2 % 10;
            if (cur->next && cur->next->val > 4)
                cur->val++;
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$\mathcal{O}(1)$。
