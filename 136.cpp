class Solution {
public:
		int singleNumber(vector<int> &nums) {
			sort(nums.begin(), nums.end());
			int s = 0, f = 1;
			while (f < nums.size()) {
				if (nums[s] == nums[f]) {
					s = ++f;
				}
				++f;
			}
			return nums[s];
		}
};
