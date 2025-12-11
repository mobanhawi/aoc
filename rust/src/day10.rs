pub mod solve;
mod tests;
use crate::util::read::file;
pub fn solve() {
    println!("\n----- DAY 10 -----");
    let input0 = file(String::from("src/day10/input0.txt")).expect("Unable to read file");
    let input = file(String::from("src/day10/input.txt")).expect("Unable to read file");
    println!("Result Simple: {}", solve::run(&input0));
    println!("Resul Full: {}", solve::run(&input));
}
