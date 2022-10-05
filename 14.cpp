class Solution {
public:
  string longestCommonPrefix(vector<string> &strs) {
    string longPre = strs[0];
    for (auto &str: strs) {
      for (int j = longPre.size() - 1; j >= 0; --j) {
        if (longPre[j] != str[j]) {
          longPre = longPre.substr(0, j);
        }
      }
    }
    return longPre;
  }
};
