class Solution {
public:
  int removeDuplicates(vector<int> &nums) {
    int p1 = 0, p2 = 1;
    while (p2 < nums.size()) {
      if (nums[p1] == nums[p2]) {
        nums.erase(nums.begin() + p2);
        continue;
      }
      p1 = p2;
      p2++;
    }
    return nums.size();
  }
};
