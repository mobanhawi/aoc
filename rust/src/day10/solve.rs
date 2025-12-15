pub fn run(input: &String) -> u64 {
    use good_lp::{Expression, Solution, SolverModel, microlp, variable, variables};
    let mut result: u64 = 0;
    // Parse each line of input to extract switch configurations and jolt values
    for line in input.lines() {
        // Split line at first space to separate pattern from switch/jolt data
        // Example: "[.##.]|(3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
        let (_, switch_jolt_str) = line.split_once(' ').unwrap();

        // Split at last space to separate switches from jolts
        // Example: "(3) (1,3) (2) (2,3) (0,2) (0,1)|{3,5,4,7}"
        let (switches_str, jolts_str) = switch_jolt_str.rsplit_once(' ').unwrap();

        // Parse switch configurations: "(3) (1,3) (2) (2,3) (0,2) (0,1)"
        let switches: Vec<Vec<usize>> = switches_str
            // Split by space to get individual switch groups: "(3)", "(1,3)", etc.
            .split(' ')
            // Process each switch group
            .map(|s| {
                // Remove parentheses and split by comma to get individual switch indices
                s.trim_matches(['(', ')'])
                    .split(',')
                    // Parse each index string to usize
                    .map(|s| s.parse::<usize>().unwrap())
                    // Collect parsed indices into a vector
                    .collect::<Vec<usize>>()
            })
            .collect();

        // Parse jolt values: "{3,5,4,7}"
        let jolts: Vec<usize> = jolts_str
            // Remove curly braces from the jolt list
            .trim_matches(['{', '}'])
            // Split by comma to get individual jolt values
            .split(',')
            // Parse each jolt value string to integer
            .map(|s| s.parse::<usize>().unwrap())
            .collect();

        // println!(
        //     "Switches: {:?}, Jolts: {:?}",
        //     switches, jolts
        // );
        let n = switches.len();
        // use good_lp to set up and solve the optimization problem
        let mut vars = variables!();
        // this creates n variables, one for each switch
        let presses: Vec<variable::Variable> = (0..n)
            .map(|_| vars.add(variable().min(0).integer()))
            .collect();
        // we want to minimize the total number of presses
        // need to define the solver for the problem first
        // otherwise the problem constraints cannot be added
        let mut problem = vars
            .minimise(presses.iter().sum::<Expression>())
            .using(microlp);
        // add constraints for each jolt
        for (jolt_index, &jolt) in jolts.iter().enumerate() {
            let mut expr = Expression::default();
            // each switch that affects this jolt
            for (switch_index, switch) in switches.iter().enumerate() {
                if switch.contains(&jolt_index) {
                    // presses[switch_index]
                    expr.add_mul(1, presses[switch_index]);
                }
            }
            // add the constraint that the sum of presses for this jolt is odd
            // this requires the problem solver to be defined first
            problem.add_constraint(expr.eq(jolt as f64));
        }
        let sol = problem.solve().unwrap();
        // get the total number of presses
        // print the solution
        for (_i, press_var) in presses.iter().enumerate() {
            let value = sol.value(*press_var);
            // println!("Switch {} pressed {} times", i, value);
            result += value.round() as u64;
        }
    }
    result
}
