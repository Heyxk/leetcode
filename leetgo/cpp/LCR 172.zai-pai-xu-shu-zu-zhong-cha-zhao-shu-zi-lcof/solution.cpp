// Created by k at 2024/05/15 16:12
// leetgo: dev
// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    int countTarget(vector<int> &scores, int target)
    {
        return helper(scores, target) - helper(scores, target - 1);
    }
    int helper(vector<int> &scores, int target)
    {
        int l = 0;
        int r = scores.size() - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (scores[mid] <= target) {
                l = mid + 1;
            }
            else {
                r = mid - 1;
            }
        }
        return l;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<int> scores;
    LeetCodeIO::scan(cin, scores);
    int target;
    LeetCodeIO::scan(cin, target);

    Solution *obj = new Solution();
    auto res      = obj->countTarget(scores, target);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}
