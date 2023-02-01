use std::fs::File;
use std::io::{self, BufRead, BufReader};
use std::path::Path;

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {

    let max_calories: i32; 
    let mut current_calories: i32 = 0;
    let mut top_three_sum: i32 = 0;
    let mut calorie_sums: Vec<i32> = vec![];

    // File hosts must exist in current path before this produces output
    if let Ok(lines) = read_lines("../input.txt") {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(l) = line {
                
                if l == "" {
                    calorie_sums.push(current_calories);
                    current_calories = 0;
                } else {
                    current_calories += l.parse::<i32>().unwrap();
                }
            }
        }
    }

    calorie_sums.sort_by(|a, b| b.cmp(a));
    max_calories = calorie_sums[0];
    
    for i in 0..3 {
        top_three_sum += calorie_sums[i];
    }

    println!("The elf with the most calories has: {max_calories} ");
    println!("The sums of calories of the top three are: {top_three_sum}")

}
