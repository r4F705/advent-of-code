use std::fs::File;
use std::path::Path;
use std::str::FromStr;
use std::io::{self, BufRead, BufReader};

#[derive(Debug, PartialEq)]
enum RPS {
    // elf options
    A = -1, // rock
    B = -2, // paper
    C = -3, // scissors

    // human options
    X = 1, // rock
    Y = 2, // paper
    Z = 3, // scissors
}

enum Outcomes {
    Win = 6,
    Draw = 3,
    Lose = 0, 
}

impl FromStr for RPS {

    type Err = ();

    fn from_str(input: &str) -> Result<RPS, Self::Err> {
        match input {
            "A"  => Ok(RPS::A),
            "B"  => Ok(RPS::B),
            "C"  => Ok(RPS::C),
            "X" => Ok(RPS::X),
            "Y" => Ok(RPS::Y),
            "Z" => Ok(RPS::Z),
            _      => Err(()),
        }
    }
}

fn play_x(elf_choice: RPS) -> i32 {
    let score: i32 = RPS::X as i32; 
    match elf_choice {
        RPS::A => score + Outcomes::Draw as i32,
        RPS::B => score + Outcomes::Lose as i32,
        RPS::C => score + Outcomes::Win as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}

fn play_y(elf_choice: RPS) -> i32 {
    let score: i32 = RPS::Y as i32;
    match elf_choice {
        RPS::A => score + Outcomes::Win as i32,
        RPS::B => score + Outcomes::Draw as i32,
        RPS::C => score + Outcomes::Lose as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}

fn play_z(elf_choice: RPS) -> i32 {
    let score: i32 = RPS::Z as i32;
    match elf_choice {
        RPS::A => score + Outcomes::Lose as i32,
        RPS::B => score + Outcomes::Win as i32,
        RPS::C => score + Outcomes::Draw as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}

fn lose_agaist(elf_choice: RPS) -> i32 {
    let score: i32 = Outcomes::Lose as i32;
    match elf_choice {
        RPS::A => score + RPS::Z as i32,
        RPS::B => score + RPS::X as i32,
        RPS::C => score + RPS::Y as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}

fn draw_agaist(elf_choice: RPS) -> i32 {
    let score: i32 = Outcomes::Draw as i32;
    match elf_choice {
        RPS::A => score + RPS::X as i32,
        RPS::B => score + RPS::Y as i32,
        RPS::C => score + RPS::Z as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}

fn win_agaist(elf_choice: RPS) -> i32 {
    let score: i32 = Outcomes::Win as i32; 
    match elf_choice {
        RPS::A => score + RPS::Y as i32,
        RPS::B => score + RPS::Z as i32,
        RPS::C => score + RPS::X as i32,
        RPS::X | RPS::Y | RPS::Z => -1, 
    }
}


fn play_round(elf_choice: RPS, human_choice: RPS) -> i32 {
    match human_choice {
        RPS::X => play_x(elf_choice),
        RPS::Y => play_y(elf_choice),
        RPS::Z => play_z(elf_choice),
        RPS::A | RPS::B | RPS::C => -1, 
    }
}

fn cook_round(elf_choice: RPS, outcome: RPS) -> i32 {
    match outcome {
        RPS::X => lose_agaist(elf_choice),
        RPS::Y => draw_agaist(elf_choice),
        RPS::Z => win_agaist(elf_choice),
        RPS::A | RPS::B | RPS::C => -1, 
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn main() {
    let mut score_p1: i32 = 0;
    let mut score_p2: i32 = 0;

    // File hosts must exist in current path before this produces output
    if let Ok(lines) = read_lines("input.txt") {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(l) = line {
                let round_data: Vec<&str> = l.split(" ").collect(); 

                let round_score_p1 = play_round(
                    RPS::from_str(round_data[0]).unwrap(),
                    RPS::from_str(round_data[1]).unwrap()
                );

               let round_score_p2 = cook_round(
                   RPS::from_str(round_data[0]).unwrap(),
                   RPS::from_str(round_data[1]).unwrap()
               );


                println!("Round P1: {:} | Round P2 {:}", round_score_p2, round_score_p2);
                score_p1 += round_score_p1;
                score_p2 += round_score_p2;
            }
        }
    }

    println!("Final score part one: {:}", score_p1);
    println!("Final score part two: {:}", score_p2);
}
