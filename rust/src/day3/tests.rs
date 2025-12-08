#[cfg(test)]
mod tests {
    use crate::day3::solve;
    use crate::util::read::file;
    use once_cell::sync::Lazy;

    static INPUT0: Lazy<String> = Lazy::new(|| file(String::from("src/day3/input0.txt")).unwrap());

    static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day3/input.txt")).unwrap());
    #[test]
    fn test_simple_pt1() {
        assert_eq!(solve::run(2, &INPUT0.clone()), 357);
    }

    #[test]
    fn test_simple_pt2() {
        assert_eq!(solve::run(12, &INPUT0.clone()), 3121910778619);
    }

    #[test]
    fn test_pt1() {
        assert_eq!(solve::run(2, &INPUT.clone()), 17301);
    }

    #[test]
    fn test_pt2() {
        assert_eq!(solve::run(12, &INPUT.clone()), 172162399742349);
    }
}
