class Solution {
public:
    vector<int> runningSum(vector<int>& nums) {
        vector<int> rs(nums.size(), 0);
        rs[0] = nums[0];
        for (int i = 1; i < nums.size(); ++i) {
            rs[i] = rs[i - 1] + nums [i];
        }
        return rs;
    }
};
