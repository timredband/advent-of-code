use std::{
    fs::File,
    io::{BufRead, BufReader},
};

use regex::Regex;

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;
    let mut matrix = vec![];

    for line in reader.lines() {
        let line = line.unwrap().trim().to_string();

        if line == "" {
            break;
        }

        let re = Regex::new(r"\s+").unwrap();
        let row: Vec<String> = re.split(&line).map(|s| s.to_string()).collect();

        matrix.push(row);
    }

    for j in 0..matrix[0].len() {
        let operation = &matrix[matrix.len() - 1][j];

        let mut problem_result = 1;

        for i in 0..matrix.len() - 1 {
            problem_result = match operation.as_str() {
                "+" => problem_result + matrix[i][j].as_str().parse::<i64>().unwrap(),
                "*" => problem_result * matrix[i][j].as_str().parse::<i64>().unwrap(),
                _ => panic!(),
            };
        }

        if operation == "+" {
            problem_result -= 1;
        }

        result += problem_result;
    }

    result
}

pub fn part2(file: &File) -> i64 {
    1
}
