# [LCR 147. 最小栈][link] (Easy)

[link]: https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/

请你设计一个 **最小栈** 。它提供 `push` ， `pop` ， `top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类:

- `MinStack()` 初始化堆栈对象。
- `void push(int val)` 将元素val推入堆栈。
- `void pop()` 删除堆栈顶部的元素。
- `int top()` 获取堆栈顶部的元素。
- `int getMin()` 获取堆栈中的最小元素。

**示例 1:**

```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[2],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,2,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(2);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 2.
minStack.getMin();   --> 返回 -2.
```

**提示：**

- `-2³¹ <= val <= 2³¹ - 1`
- `pop`、 `top` 和 `getMin` 操作总是在 **非空栈** 上调用
- `push`、 `pop`、 `top` 和 `getMin` 最多被调用 `3 * 10⁴` 次

注意：本题与主站 155 题相同： [https://leetcode-cn.com/problems/min-stack/](https://leetcode-cn.com/
problems/min-stack/)
