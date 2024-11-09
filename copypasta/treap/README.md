目前有四个版本：

- [对标 C++ 的 set](./set/set.go)
- [对标 C++ 的 multiset](./multiset/multiset.go)
- [对标 C++ 的 map](./map/map.go)
- [把 treap 当作有序数组（允许重复元素），同时动态维护这个有序数组的前缀和](./prefixsum/prefixsum.go)

具体使用方法，见对应目录下的 test.go 文件。

关于遍历 treap 的逻辑，见 [bst.go](../bst.go)。

TIPS：某些题目通过插入 $-\infty$ 和 $\infty$ 哨兵，可以减少边界判断。
