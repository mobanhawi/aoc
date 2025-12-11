pub fn run(input: &String) -> u128 {
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

        println!("Switches: {:?}, Jolts: {:?}", switches, jolts);
    }
    0
}
