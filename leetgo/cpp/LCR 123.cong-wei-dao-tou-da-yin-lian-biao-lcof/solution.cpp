// Created by k at 2024/05/16 16:29
// leetgo: dev
// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    vector<int> reverseBookList(ListNode *head)
    {
        recur(head);
        return ans;
    }
    void recur(ListNode *head)
    {
        if (head == nullptr) {
            return;
        }
        recur(head->next);
        ans.push_back(head->val);
    }
    vector<int> ans;
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    ListNode *head;
    LeetCodeIO::scan(cin, head);

    Solution *obj = new Solution();
    auto res      = obj->reverseBookList(head);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
