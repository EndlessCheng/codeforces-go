问题本质上是给你前缀和，还原回原数组。

由于前缀和的差分是元素组，因此直接两两异或即可。

```
class Solution:
    def findArray(self, pref: List[int]) -> List[int]:
        return [pref[0]] + [x ^ y for x, y in pairwise(pref)]
```


---

欢迎关注我的B站频道：[灵茶山艾府](https://space.bilibili.com/206214)，定期更新算法讲解视频哦~
