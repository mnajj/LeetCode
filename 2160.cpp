class Solution {
public:
    int minimumSum(int num) {
        string s = to_string(num);
        sort(s.begin(), s.end());
        swap(s[2], s[1]);
        return stoi(s.substr(0, 2)) + stoi(s.substr(2, 2));
    }
};
