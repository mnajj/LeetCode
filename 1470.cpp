class Solution {
public:
    vector<int> shuffle(vector<int>& nums, int n) {
        vector<int> vec(2 * n, 0);
        for (int i = 0, j = 0; i < n; ++i, j += 2) {
            vec[j] = nums[i];
            vec[j + 1] = nums[i + n];
        }
        return vec;
    }
};
