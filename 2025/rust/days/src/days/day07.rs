use std::{
    collections::{HashMap, HashSet},
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
    let reader = BufReader::new(file);

    let mut beams: HashMap<usize, usize> = HashMap::new();

    for line in reader.lines() {
        let line = line.unwrap();

        // this is bad
        if let Some(index) = line.find("S") {
            beams.insert(index, 1);
        }

        let row_beams: Vec<usize> = beams.clone().into_keys().collect();

        for beam in &row_beams {
            if &line[*beam..*beam + 1] == "^" {
                let timelines = beams[&beam];
                beams.remove(beam);

                if let Some(existing) = beams.get(&(beam - 1)) {
                    beams.insert(beam - 1, existing + timelines);
                } else {
                    beams.insert(beam - 1, timelines);
                }

                if let Some(existing) = beams.get(&(beam + 1)) {
                    beams.insert(beam + 1, existing + timelines);
                } else {
                    beams.insert(beam + 1, timelines);
                }
            }
        }
    }

    let result = beams.into_values().fold(0, |acc, value| acc + value);

    result.try_into().unwrap()
}
