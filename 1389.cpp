class Solution {
public:
    vector<int> createTargetArray(vector<int>& nums, vector<int>& index) {
        vector<int> targ;
        for (int i = 0; i < nums.size(); ++i) {
            targ.insert(targ.begin() + index[i], nums[i]);
        }
        return targ;
    }
};
