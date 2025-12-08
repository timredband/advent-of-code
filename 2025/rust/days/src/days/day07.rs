use std::{
    collections::HashSet,
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;
    let mut beams: HashSet<usize> = HashSet::new();

    for line in reader.lines() {
        let line = line.unwrap();

        // this is bad
        if let Some(index) = line.find("S") {
            beams.insert(index);
        }

        let row_beams: Vec<usize> = beams.clone().into_iter().collect();

        for beam in &row_beams {
            if &line[*beam..*beam + 1] == "^" {
                beams.remove(beam);

                if *beam != 0 {
                    beams.insert(beam - 1);
                }

                if *beam != line.len() - 1 {
                    beams.insert(beam + 1);
                }

                result += 1;
            }
        }
    }

    result
}

pub fn part2(file: &File) -> i64 {
    1
}
