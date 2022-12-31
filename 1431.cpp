class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        vector<bool> vec(candies.size(), false);
        int max = 0;
        for (int x : candies) {
            if (max < x) max = x;
        }
        for (int i = 0; i < candies.size(); ++i) {
            if (candies[i] + extraCandies >= max) vec[i] = true;
        }
        return vec;        
    }
};
