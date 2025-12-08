pub mod solve;
mod tests;
use crate::util::read::file;
pub fn solve() {
    let input0 = file(String::from("src/day3/input0.txt")).expect("Unable to read file");
    let input = file(String::from("src/day3/input.txt")).expect("Unable to read file");
    println!("ResultP1 Simple: {}", solve::run(2, &input));
    println!("ResultP1: {}", solve::run(12, &input));
    println!("ResultP2 Simple: {}", solve::run(2, &input0));
    println!("ResultP2: {}", solve::run(12, &input0));
}
