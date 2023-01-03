class Solution {
public:
    int balancedStringSplit(string s) {
        int l(0), r(0), sum(0);
        for (char c : s) {
            if(c == 'L') ++l;
            else ++r;
            if (l != 0 && l == r) {
                l = r = 0;
                ++sum;
            }
        }
        return sum;
    }
};
