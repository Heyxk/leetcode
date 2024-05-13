// Created by k at 2024/05/13 21:34
// leetgo: dev
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    vector<int> decorateRecord(TreeNode *root)
    {
        vector<int> ans;
        if (root == nullptr) {
            return ans;
        }
        vector<TreeNode *> queue = {root};
        while (!queue.empty()) {
            TreeNode *node = queue.front();
            ans.push_back(node->val);
            if (node->left != nullptr) {
                queue.push_back(node->left);
            }
            if (node->right != nullptr) {
                queue.push_back(node->right);
            }
            queue.erase(queue.cbegin());
        }
        return ans;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    TreeNode *root;
    LeetCodeIO::scan(cin, root);

    Solution *obj = new Solution();
    auto res      = obj->decorateRecord(root);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
