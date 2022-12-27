class Solution
{
	public:
		int finalValueAfterOperations(vector<string> &operations) {
			int x = operations.size();
			for (string s: operations) {
                if (s == "--X" || s == "X--") {
                    x -= 2;
                }
			}
            return x;
        }
    };
