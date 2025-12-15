#[cfg(test)]
mod tests {
    use crate::day10::solve;
    use crate::util::read::file;
    use once_cell::sync::Lazy;

    static INPUT0: Lazy<String> = Lazy::new(|| file(String::from("src/day10/input0.txt")).unwrap());
    static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day10/input.txt")).unwrap());
    #[test]
    fn test_run_simple() {
        assert_eq!(solve::run(&INPUT0.clone()), 33);
    }

    #[test]
    fn test_run() {
        assert_eq!(solve::run(&INPUT.clone()), 17576);
    }
}
