class Solution {
public:
		TreeNode *getTargetCopy(TreeNode *original, TreeNode *cloned, TreeNode *target) {
			if (cloned != NULL) {
				if (cloned->val == target->val) return cloned;
				auto r = getTargetCopy(original, cloned->right, target);
				if (r != NULL) return r;
				auto l = getTargetCopy(original, cloned->left, target);
				if (l != NULL) return l;
			}
			return NULL;

		}
};
