遍历链表，在当前节点 $\textit{cur}$ 后面插入 $\textit{gcd}$ 节点，同时 $\textit{gcd}$ 节点指向 $\textit{cur}$ 的下一个节点。

插入后，$\textit{cur}$ 更新为 $\textit{cur}.\textit{next}.\textit{next}$，也就是 $\textit{cur}$ 原来的下一个节点，开始下一轮循环。

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

```java [sol-Java]
public class Solution {
    public ListNode insertGreatestCommonDivisors(ListNode head) {
        for (ListNode cur = head; cur.next != null; cur = cur.next.next) {
            cur.next = new ListNode(gcd(cur.val, cur.next.val), cur.next);
        }
        return head;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int temp = a;
            a = b % a;
            b = temp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    ListNode *insertGreatestCommonDivisors(ListNode *head) {
        for (auto cur = head; cur->next; cur = cur->next->next) {
            cur->next = new ListNode(gcd(cur->val, cur->next->val), cur->next);
        }
        return head;
    }
};
```

```go [sol-Go]
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
    for cur := head; cur.Next != nil; cur = cur.Next.Next {
        cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
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

```js [sol-JavaScript]
var insertGreatestCommonDivisors = function(head) {
    for (let cur = head; cur.next; cur = cur.next.next) {
        cur.next = new ListNode(gcd(cur.val, cur.next.val), cur.next);
    }
    return head;
};

function gcd(a, b) {
    while (a !== 0) {
        [a, b] = [b % a, a];
    }
    return b;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn insert_greatest_common_divisors(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut cur = &mut head;
        while cur.as_ref().unwrap().next.is_some() {
            let x = cur.as_mut().unwrap();
            let next = x.next.take();
            x.next = Some(Box::new(ListNode {
                val: Self::gcd(x.val, next.as_ref().unwrap().val),
                next,
            }));
            cur = &mut cur.as_mut().unwrap().next.as_mut().unwrap().next;
        }
        head
    }

    fn gcd(mut a: i32, mut b: i32) -> i32 {
        while a != 0 {
            (a, b) = (b % a, a);
        }
        b
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为链表长度，$U$ 为节点值的最大值。每次计算 `gcd` 需要 $\mathcal{O}(\log U)$ 的时间。
- 空间复杂度：$\mathcal{O}(1)$。返回值的空间不计入。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
