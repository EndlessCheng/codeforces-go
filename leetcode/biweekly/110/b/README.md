下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

遍历链表，在当前节点 $\textit{cur}$ 后面插入 $\textit{gcd}$ 节点，$\textit{gcd}$ 节点指向 $\textit{cur}$ 的下一个节点。

插入后，$\textit{cur}$ 更新为 $\textit{cur}.\textit{next}.\textit{next}$。

循环直到 $\textit{cur}$ 没有下一个节点为止。

```py [sol-Python3]
class Solution:
    def insertGreatestCommonDivisors(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        while cur.next:
            cur.next = ListNode(gcd(cur.val, cur.next.val), cur.next)
            cur = cur.next.next
        return head
```

```go [sol-Go]
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
	cur := head
	for cur.Next != nil {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
		cur = cur.Next.Next
	}
	return head
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为链表长度，$U$ 为节点值的最大值。
- 空间复杂度：$\mathcal{O}(1)$。返回值的空间不计入，仅用到若干额外变量。
