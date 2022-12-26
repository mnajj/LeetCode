class Solution {
public:
    vector<int> getConcatenation(vector<int>& nums) {
        int n = nums.size();
        vector<int> vec(n * 2, 0);
        for (int i = 0; i < n; ++i) {
            vec[i] = vec[i + n] = nums[i];
        }
        return vec;
    }
};
