class Solution {
  public:
    bool isPalindrome(int x) {
      if (x < 0) return false;
      if (x < 10) return true;
      string str = to_string(x);
      int p1 = 0, p2 = str.size() - 1;
      while (p1 < p2) {
        if (str[p1] != str[p2]) {
          return false;
        }
        p1++;
        p2--;
      }
      return true;
    }
};


class Solution {
  public:
    bool isPalindrome(int x) {
      int digit = 0;
      int temp = x;
      while (x > 0) {
        int rem = x % 10;
        digit = digit * 10 + rem;
        x /= 10;
      }
      if (digit == temp) return true;
      return false;
    }
};
