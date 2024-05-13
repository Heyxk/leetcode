// Created by k at 2024/05/13 14:35
// leetgo: dev
// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    ListNode *trainningPlan(ListNode *head)
    {
        ListNode *pre = nullptr;
        for (; head != nullptr;) {
            ListNode *next = head->next;
            head->next     = pre;
            pre            = head;
            head           = next;
        }
        return pre;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    ListNode *head;
    LeetCodeIO::scan(cin, head);

    Solution *obj = new Solution();
    auto res      = obj->trainningPlan(head);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
