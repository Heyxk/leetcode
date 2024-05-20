// Created by k at 2024/05/20 16:45
// leetgo: dev
// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    bool isSubStructure(TreeNode *A, TreeNode *B)
    {
        if (A == nullptr || B == nullptr) {
            return false;
        }
        return recur(A, B) || isSubStructure(A->left, B)
               || isSubStructure(A->right, B);
    }
    bool recur(TreeNode *a, TreeNode *b)
    {
        if (b == nullptr) {
            return true;
        }
        if (a == nullptr || a->val != b->val) {
            return false;
        }
        return recur(a->left, b->left) && recur(a->right, b->right);
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    TreeNode *A;
    LeetCodeIO::scan(cin, A);
    TreeNode *B;
    LeetCodeIO::scan(cin, B);

    Solution *obj = new Solution();
    auto res      = obj->isSubStructure(A, B);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
