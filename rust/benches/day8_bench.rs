use aoc::day8::solve;
use aoc::util::read::file;
use criterion::{Criterion, black_box, criterion_group, criterion_main};
use once_cell::sync::Lazy;
static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day8/input.txt")).unwrap());

/// day8 pt1                time:   [8.6685 ms 8.7126 ms 8.7594 ms]
fn benchmark_pt1(c: &mut Criterion) {
    c.bench_function("day8 pt1", |b| {
        b.iter(|| solve::run_pt1(black_box(1000), black_box(&INPUT.clone())))
    });
}
/// day8 pt2                time:   [11.677 ms 11.948 ms 12.241 ms]
fn benchmark_pt2(c: &mut Criterion) {
    c.bench_function("day8 pt2", |b| {
        b.iter(|| solve::run_pt2(black_box(1000), black_box(&INPUT.clone())))
    });
}
criterion_group!(benches, benchmark_pt1, benchmark_pt2);
criterion_main!(benches);
