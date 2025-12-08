#[cfg(test)]
mod tests {
    use crate::day3::solve;
    #[test]
    fn test_simple_pt1() {
        assert_eq!(solve::run(2, String::from("src/day3/input0.txt")), 357);
    }

    #[test]
    fn test_simple_pt2() {
        assert_eq!(solve::run(12, String::from("src/day3/input0.txt")), 3121910778619);
    }

    #[test]
    fn test_pt1() {
        assert_eq!(solve::run(2, String::from("src/day3/input.txt")), 17301);
    }

    #[test]
    fn test_pt2() {
        assert_eq!(solve::run(12, String::from("src/day3/input.txt")), 172162399742349);
    }
}