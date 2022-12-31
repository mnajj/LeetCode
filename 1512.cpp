class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        int arr[101] = {0};
        int sum = 0;
        for (int x : nums) ++arr[x];
        for (int i = 0; i < nums.size(); ++i) {
            sum += arr[i] * (arr[i] - 1) / 2;
        }
        return sum;
    }
};
