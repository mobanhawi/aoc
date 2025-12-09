use std::collections::{BinaryHeap, HashMap, HashSet};
#[derive(PartialEq, Eq)]
struct Distance {
    dist: i64, // euclidean distance squared between p1 and p2
    p1: usize, // index of point 1
    p2: usize, // index of point 2
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
    points: HashSet<usize>, // indexes of points in the circuit
    connections: usize,     // number of connections in the circuit
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
    let dy = v1[1] - v2[1];
    let dz = v1[2] - v2[2];
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
    let mut c_key: usize = 0;
    // Circuits storage c_key to Circuit
    let mut circuits: HashMap<usize, Circuit> = HashMap::new();
    // Point index to circuit key
    let mut pt_to_circuit: HashMap<usize, usize> = HashMap::new();
    for _ in 0..n {
        if let Some(distance) = heap.pop() {
            let p1 = distance.p1;
            let p2 = distance.p2;
            let c1 = pt_to_circuit.get(&p1).cloned();
            let c2 = pt_to_circuit.get(&p2).cloned();
            // println!(
            //     "Distance between point {:?} and {:?}: {}",
            //     positions[p1], positions[p2], distance.dist
            // );

            match (c1, c2) {
                // merge circuit c2 into c1
                (Some(c1_key), Some(c2_key)) => {
                    if c1_key != c2_key {
                        let mut c2_points = circuits.get(&c2_key).unwrap().points.clone();
                        let c1_circuit = match circuits.get_mut(&c1_key) {
                            None => {
                                panic!(
                                    "Circuit {} not found keys:{:?}",
                                    c1_key,
                                    circuits.keys().collect::<Vec<&usize>>()
                                );
                            }
                            Some(_) => circuits.get_mut(&c1_key).unwrap(),
                        };
                        c1_circuit.points.extend(c2_points.iter());
                        c1_circuit.connections = c1_circuit.points.len();
                        circuits.remove(&c2_key);
                        // Update point to circuit mapping
                        for c in c2_points.drain() {
                            pt_to_circuit.insert(c, c1_key);
                        }
                    }
                }
                // Add point(p2) to existing circuit(c1)
                (Some(c1_key), None) => {
                    let c1_circuit = circuits.get_mut(&c1_key).unwrap();
                    c1_circuit.points.insert(p2);
                    pt_to_circuit.insert(p2, c1_key);
                    c1_circuit.connections += 1;
                }
                // Add point(p1) to existing circuit(c2)
                (None, Some(c2_key)) => {
                    let c2_circuit = circuits.get_mut(&c2_key).unwrap();
                    c2_circuit.points.insert(p1);
                    c2_circuit.connections += 1;
                    pt_to_circuit.insert(p1, c2_key);
                }
                // Create new circuit
                (None, None) => {
                    c_key += 1; // New circuit key
                    let mut ps: HashSet<usize> = HashSet::new();
                    ps.insert(p1);
                    ps.insert(p2);
                    let circuit = Circuit {
                        points: ps.clone(),
                        connections: 2,
                    };
                    circuits.insert(c_key, circuit);
                    pt_to_circuit.insert(p1, c_key);
                    pt_to_circuit.insert(p2, c_key);
                }
            }
        }
    }

    let mut c: Vec<&Circuit> = circuits.values().collect();
    c.sort_by(|a, b| b.cmp(a)); // Sorts in descend
    let mut result: u128 = 1;
    for i in 0..3 {
        result *= c[i].connections as u128;
    }
    result
}

pub fn run_pt2(_: u32, input: &String) -> u128 {
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
    let mut c_key: usize = 0;
    // Circuits storage c_key to Circuit
    let mut circuits: HashMap<usize, Circuit> = HashMap::new();
    // Point index to circuit key
    let mut pt_to_circuit: HashMap<usize, usize> = HashMap::new();
    let mut p1: usize;
    let mut p2: usize;
    loop {
        if let Some(distance) = heap.pop() {
            p1 = distance.p1;
            p2 = distance.p2;
            let c1 = pt_to_circuit.get(&p1).cloned();
            let c2 = pt_to_circuit.get(&p2).cloned();
            match (c1, c2) {
                // merge circuit c2 into c1
                (Some(c1_key), Some(c2_key)) => {
                    if c1_key != c2_key {
                        let mut c2_points = circuits.get(&c2_key).unwrap().points.clone();
                        let should_break = {
                            let c1_circuit = match circuits.get_mut(&c1_key) {
                                None => {
                                    panic!(
                                        "Circuit {} not found keys:{:?}",
                                        c1_key,
                                        circuits.keys().collect::<Vec<&usize>>()
                                    );
                                }
                                Some(_) => circuits.get_mut(&c1_key).unwrap(),
                            };
                            c1_circuit.points.extend(c2_points.iter());
                            c1_circuit.connections = c1_circuit.points.len();
                            c1_circuit.connections >= positions.len()
                        };
                        circuits.remove(&c2_key);
                        // Update point to circuit mapping
                        for c in c2_points.drain() {
                            pt_to_circuit.insert(c, c1_key);
                        }
                        if should_break {
                            break;
                        }
                    }
                }
                // Add point(p2) to existing circuit(c1)
                (Some(c1_key), None) => {
                    let c1_circuit = circuits.get_mut(&c1_key).unwrap();
                    c1_circuit.points.insert(p2);
                    pt_to_circuit.insert(p2, c1_key);
                    c1_circuit.connections += 1;
                    if c1_circuit.connections >= positions.len() {
                        break;
                    }
                }
                // Add point(p1) to existing circuit(c2)
                (None, Some(c2_key)) => {
                    let c2_circuit = circuits.get_mut(&c2_key).unwrap();
                    c2_circuit.points.insert(p1);
                    c2_circuit.connections += 1;
                    pt_to_circuit.insert(p1, c2_key);
                    if c2_circuit.connections >= positions.len() {
                        break;
                    }
                }
                // Create new circuit
                (None, None) => {
                    c_key += 1; // New circuit key
                    let mut ps: HashSet<usize> = HashSet::new();
                    ps.insert(p1);
                    ps.insert(p2);
                    let circuit = Circuit {
                        points: ps.clone(),
                        connections: 2,
                    };
                    circuits.insert(c_key, circuit);
                    pt_to_circuit.insert(p1, c_key);
                    pt_to_circuit.insert(p2, c_key);
                }
            }
        }
    }

    positions[p1][0].abs() as u128 * positions[p2][0].abs() as u128
}
