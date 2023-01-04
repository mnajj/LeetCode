/*
* N*log(N)
*/
class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
        int a[501] = {0};
        vector<int> vec(nums.size(), 0);
        for (int i = 0; i < nums.size(); ++i) {
            vec[i] = nums[i];
        }
        std::sort(vec.begin(), vec.end());
        for (int i = 1; i < nums.size(); ++i) {
            if (vec[i] != vec[i - 1]) {
                a[vec[i]] = i;
            }
        }
        for (int i = 0; i < nums.size(); ++i) {
            vec[i] = a[nums[i]];
        }
        return vec;
    }
};

/*
* N
*/
class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
        int c[101] = {0};
        vector<int> res(nums.size(), 0);
        for (int i = 0; i < nums.size(); i++) {
            ++c[nums[i]];
        }
        for (int i = 1; i <= 100; i++) {
            c[i] += c[i - 1];    
        }
        for (int i = 0; i < nums.size(); i++) {
            if (!nums[i]) continue;
            res[i] = c[nums[i] - 1];
        }
        return res;                       
    }
};
