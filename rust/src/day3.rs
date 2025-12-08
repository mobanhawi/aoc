pub mod solve;

pub fn solve() {
    solve::run(2, String::from("src/day3/input.txt"));
    solve::run(12, String::from("src/day3/input.txt"));
    solve::run(2, String::from("src/day3/input0.txt"));
    solve::run(12, String::from("src/day3/input0.txt"));
}