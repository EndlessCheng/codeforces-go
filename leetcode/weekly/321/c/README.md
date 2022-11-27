[视频讲解](https://www.bilibili.com/video/BV1sD4y1e7pr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

# 方法一：递归

既然要倒着看最大值，那么用递归解决是最合适的，毕竟**递归本质就是在倒着遍历链表**。

```py [sol1-Python3]
class Solution:
    def removeNodes(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head.next is None: return head  # 输入保证链表不为空
        node = self.removeNodes(head.next)  # 返回的链表头一定是最大的
        if node.val > head.val: return node  # 删除 head
        head.next = node  # 不删除 head
        return head
```

```go [sol1-Go]
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil { // 输入保证链表不为空
		return head
	}
	node := removeNodes(head.Next)
	if node.Val > head.Val { // 返回的链表头一定是最大的
		return node // 删除 head
	}
	head.Next = node // 不删除 head
	return head
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$O(n)$，需要 $O(n)$ 的栈空间。

# 方法二：迭代：两次反转链表

通过 [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)，我们可以从反转后的链表头开始，像 [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/) 那样，删除比当前节点值小的元素。

最后再次反转链表，即为答案。

```py [sol2-Python3]
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        pre, cur = None, head
        while cur:
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        return pre

    def removeNodes(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head = self.reverseList(head)
        while cur.next:
            if cur.val > cur.next.val:
                cur.next = cur.next.next
            else:
                cur = cur.next
        return self.reverseList(head)
```

```go [sol2-Go]
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

func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	cur := head
	for cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return reverseList(head)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为链表的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
