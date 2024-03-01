impl Solution {
  pub fn divide_array(nums: Vec<i32>) -> bool {
    let mut counter = [0; 501];
    nums.iter().for_each(|&x| counter[x as usize] += 1);
    counter.iter().filter(|x| **x > 0).all(|&x| x % 2 == 0)
  }
}
