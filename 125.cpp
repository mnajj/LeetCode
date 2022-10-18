class Solution {
  public:
    bool isPalindrome(string s) {
      int l, r;
      string alnum = "";
      for (char c: s) {
        if (isalnum(c)) {
          alnum += c;
        }
      }
      l = 0;
      r = alnum.size() - 1;
      while (l < r) {
        if (tolower(alnum[l]) != tolower(alnum[r])) {
          return false;
        }
        ++l;
        --r;
      }
      return true;
    }
};

// Inplace Solution
class Solution {
  public:
    bool isPalindrome(string s) {
      int l, r;
      int n = s.size();
      l = 0;
      r = n - 1;
      while (l < r) {
        while (l < n && !isalnum(s[l])) {
          ++l;
        }
        while (r >= 0 && !isalnum(s[r])) {
          --r;
        }
        // If all string chars is not alphanum
        if (l <= n && r < 0) {
          break;
        }
        if (tolower(s[l]) != tolower(s[r])) {
          return false;
        }
        ++l;
        --r;
      }
      return true;
    }
};
