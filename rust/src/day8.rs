pub mod solve;
mod tests;
use crate::util::read::file;
pub fn solve() {
    let input0 = file(String::from("src/day3/input0.txt")).expect("Unable to read file");
    let input = file(String::from("src/day3/input.txt")).expect("Unable to read file");
    println!("ResultP1 Simple: {}", solve::run_pt1(10, &input0));
    println!("ResultP1 Full: {}", solve::run_pt1(10, &input));
}
