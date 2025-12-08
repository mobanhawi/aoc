pub mod solve;
mod tests;

pub fn solve() {
    println!("ResultP1 Simple: {}", solve::run(2, String::from("src/day3/input.txt")));
    println!("ResultP1: {}", solve::run(12, String::from("src/day3/input.txt")));
    println!("ResultP2 Simple: {}", solve::run(2, String::from("src/day3/input0.txt")));
    println!("ResultP2: {}", solve::run(12, String::from("src/day3/input0.txt")));
}