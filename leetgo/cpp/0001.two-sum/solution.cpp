// Created by k at 2024/05/10 18:09
// leetgo: dev
// https://leetcode.cn/problems/two-sum/

#include "LC_IO.h"
#include <bits/stdc++.h>
using namespace std;

// @lc code=begin

class Solution {
  public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        vector<int> ans;
        for (int i = 0; i < nums.size(); i++) {
            for (int j = i + 1; j < nums.size(); j++) {
                if (nums[i] + nums[j] == target) {
                    ans.push_back(i);
                    ans.push_back(j);
                }
            }
        }
        return ans;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<int> nums;
    LeetCodeIO::scan(cin, nums);
    int target;
    LeetCodeIO::scan(cin, target);

    Solution *obj = new Solution();
    auto res      = obj->twoSum(nums, target);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
