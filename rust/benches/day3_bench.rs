use criterion::{Criterion, black_box, criterion_group, criterion_main};
use once_cell::sync::Lazy;
use rust::day3::solve;
use rust::util::read::file;
static INPUT: Lazy<String> = Lazy::new(|| file(String::from("src/day3/input.txt")).unwrap());

/// day3 pt1                time:   [347.16 µs 347.82 µs 348.55 µs]
fn benchmark_pt1(c: &mut Criterion) {
    c.bench_function("day3 pt1", |b| {
        b.iter(|| solve::run(black_box(2), black_box(&INPUT.clone())))
    });
}

/// day3 pt2                time:   [334.70 µs 335.69 µs 336.81 µs]
fn benchmark_pt2(c: &mut Criterion) {
    c.bench_function("day3 pt2", |b| {
        b.iter(|| solve::run(black_box(12), black_box(&INPUT.clone())))
    });
}

criterion_group!(benches, benchmark_pt1, benchmark_pt2);
criterion_main!(benches);
