class Solution {
public:
		TreeNode *traverse(vector<int> &nums, int l, int r) {
			if (l > r) {
				return NULL;
			}
			int m = (l + r) / 2;
			TreeNode *tmp = new TreeNode(nums[m]);
			tmp->left = traverse(nums, l, m - 1);
			tmp->right = traverse(nums, m + 1, r);
			return tmp;
		}

		TreeNode *sortedArrayToBST(vector<int> &nums) {
			return traverse(nums, 0, nums.size() - 1);
		}
};
