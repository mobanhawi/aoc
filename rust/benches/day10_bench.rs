use aoc::day10::solve;
use aoc::util::read::file;
use criterion::{Criterion, black_box, criterion_group, criterion_main};
use once_cell::sync::Lazy;
static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day10/input.txt")).unwrap());

/// day10 pt2               time:   [8.0930 ms 8.1882 ms 8.3147 ms]
fn benchmark_pt1(c: &mut Criterion) {
    c.bench_function("day10 pt2", |b| {
        b.iter(|| solve::run(black_box(&INPUT.clone())))
    });
}
criterion_group!(benches, benchmark_pt1);
criterion_main!(benches);
