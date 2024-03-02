impl Solution {
    pub fn number_game(nums: Vec<i32>) -> Vec<i32> {
        let mut clone = nums.clone();
        clone.sort();
        for i in (0..nums.len()).step_by(2) {
            if i + 1 < nums.len() {
                clone.swap(i, i + 1);
            }
        }
        clone
    }
}
