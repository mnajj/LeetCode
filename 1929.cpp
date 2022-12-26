class Solution {
public:
    vector<int> getConcatenation(vector<int>& nums) {
        int n = nums.size();
        vector<int> vec(n * 2, 0);
        int l(0), r(n);
        while (l < n) {
            vec[l] = vec[r] = nums[l];
            ++l;
            ++r;
        }
        return vec;
    }
};
