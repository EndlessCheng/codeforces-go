下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

把反转看成是后续往字符串的头部添加字符。

这可以用双端队列实现。

```py [sol-Python3]
class Solution:
    def finalString(self, s: str) -> str:
        q = deque()
        tail = True
        for c in s:
            if c == 'i':
                tail = not tail  # 修改添加方向
            elif tail:  # 加尾部
                q.append(c)
            else:  # 加头部
                q.appendleft(c)
        return ''.join(q if tail else reversed(q))
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
