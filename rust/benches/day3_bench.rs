use criterion::{black_box, criterion_group, criterion_main, Criterion};
use rust::day3::solve;

fn benchmark_pt1(c: &mut Criterion) {
    c.bench_function("day3 pt1", |b| {
        b.iter(|| solve::run(black_box(2), black_box(String::from("src/day3/input.txt"))))
    });
}

fn benchmark_pt2(c: &mut Criterion) {
    c.bench_function("day3 pt2", |b| {
        b.iter(|| solve::run(black_box(12), black_box(String::from("src/day3/input.txt"))))
    });
}

criterion_group!(benches, benchmark_pt1, benchmark_pt2);
criterion_main!(benches);
