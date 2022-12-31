class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        int arr[101] = {0};
        int sum = 0;
        for (int x : nums) ++arr[x];
        for (int x : arr) sum += x * (x - 1) / 2;
        return sum;
    }
};
