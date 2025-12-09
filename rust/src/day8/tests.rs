#[cfg(test)]
mod tests {
    use crate::day8::solve;
    use crate::util::read::file;
    use once_cell::sync::Lazy;

    static INPUT0: Lazy<String> = Lazy::new(|| file(String::from("src/day8/input0.txt")).unwrap());
    static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day8/input.txt")).unwrap());
    #[test]
    fn test_run_pt1_simple() {
        assert_eq!(solve::run_pt1(10, &INPUT0.clone()), 40);
    }
    //
    // #[test]
    // // fn test_simple_pt2() {
    // //     assert_eq!(solve::run(12, &INPUT0.clone()), 3121910778619);
    // // }
    //
    #[test]
    fn test_pt1() {
        assert_eq!(solve::run_pt1(1000, &INPUT.clone()), 131580);
    }
    //
    // #[test]
    // // fn test_pt2() {
    // //     assert_eq!(solve::run(12, &INPUT.clone()), 172162399742349);
    // // }
}
