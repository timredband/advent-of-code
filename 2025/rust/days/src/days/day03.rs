use std::{
    cmp,
    fs::File,
    io::{BufRead, BufReader},
};

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    for line in reader.lines() {
        let line = line.unwrap();

        let mut max = 0;

        for i in 0..line.len() - 1 {
            for j in i + 1..line.len() {
                let first = &line[i..i + 1];
                let second = &line[j..j + 1];
                let mut num = String::from(first);
                num.push_str(second);

                let current: i64 = num.parse().unwrap();
                max = cmp::max(max, current);
            }
        }

        result += max;
    }

    result
}

pub fn part2(file: &File) -> i64 {
    1
}
