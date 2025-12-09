use std::collections::{BinaryHeap, HashMap, HashSet};
#[derive(PartialEq, Eq)]
struct Distance {
    dist: i64,
    p1: usize,
    p2: usize,
}
// Implement Ord and PartialOrd for Distance to make it work with BinaryHeap
// see: https://stackoverflow.com/questions/76249349/rust-binary-heap-or-priority-queue-on-struct-with-2-ways-to-order
impl Ord for Distance {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.dist.cmp(&self.dist) // Reversed comparison = min-heap
        // self.dist.cmp(&other.dist)  // Normal comparison = max-heap
    }
}
impl PartialOrd for Distance {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(PartialEq, Eq)]
struct Circuit {
    points: HashSet<Vec<i64>>,
    connections: usize,
}

impl Ord for Circuit {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.connections.cmp(&other.connections) // Normal comparison = max-heap
    }
}
impl PartialOrd for Circuit {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

fn euclidean_distance(v1: &[i64], v2: &[i64]) -> i64 {
    let dx = v1[0] - v2[0];
    let dy = v1[0] - v2[0];
    let dz = v1[0] - v2[0];
    dx * dx + dy * dy + dz * dz
}

pub fn run_pt1(n: u32, input: &String) -> u128 {
    let positions: Vec<Vec<i64>> = input
        .lines()
        .map(|line| line.split(',').map(|s| s.parse::<i64>().unwrap()).collect())
        .collect();
    // Build a min-heap of distances
    let mut heap: BinaryHeap<Distance> = BinaryHeap::new();

    for i in 0..positions.len() {
        for j in (i + 1)..positions.len() {
            let dist = euclidean_distance(&positions[i], &positions[j]);
            let distance = Distance { dist, p1: i, p2: j };
            heap.push(distance);
        }
    }
    let mut connection: usize = 0;
    let mut circuits: HashMap<usize, Circuit> = HashMap::new();
    for _ in 0..n {
        if let Some(distance) = heap.pop() {
            let p1 = distance.p1;
            let p2 = distance.p2;
            let mut c1: Option<usize> = None;
            let mut c2: Option<usize> = None;
            for (key, circuit) in &circuits {
                if circuit.points.contains(&positions[p1]) {
                    c1 = Some(*key);
                }
                if circuit.points.contains(&positions[p2]) {
                    c2 = Some(*key);
                }
            }
            match (c1, c2) {
                // merge circuits
                (Some(c1_key), Some(c2_key)) => {
                    if c1_key != c2_key {
                        let mut c2_points = circuits.get(&c2_key).unwrap().points.clone();
                        let c2_connections = circuits.get(&c2_key).unwrap().connections;
                        let c1_circuit = circuits.get_mut(&c1_key).unwrap();
                        c1_circuit.points.extend(c2_points.drain());
                        c1_circuit.connections += c2_connections + 1;
                        circuits.remove(&c2_key);
                    } else {
                        circuits.get_mut(&c1_key).unwrap().connections += 1;
                    }
                }
                // Add point to existing circuit
                (Some(c1_key), None) => {
                    let c1_circuit = circuits.get_mut(&c1_key).unwrap();
                    c1_circuit.points.insert(positions[p2].clone());
                    c1_circuit.connections += 1;
                }
                // Add point to existing circuit
                (None, Some(c2_key)) => {
                    let c2_circuit = circuits.get_mut(&c2_key).unwrap();
                    c2_circuit.points.insert(positions[p1].clone());
                    c2_circuit.connections += 1;
                }
                // Create new circuit
                (None, None) => {
                    let mut points: HashSet<Vec<i64>> = HashSet::new();
                    points.insert(positions[p1].clone());
                    points.insert(positions[p2].clone());
                    let circuit = Circuit {
                        points,
                        connections: 1,
                    };
                    circuits.insert(connection, circuit);
                }
            }
        }
    }

    let mut c: Vec<&Circuit> = circuits.values().collect();
    c.sort_by(|a, b| b.cmp(a)); // Sorts in descend
    let mut result: u128 = 1;
    for i in 0..3 {
        println!("Circuit {}: {} connections", i + 1, c[i].connections);
        result *= c[i].connections as u128;
    }
    result
}
