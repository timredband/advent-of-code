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
    let reader = BufReader::new(file);

    let mut result = 0;
    let mut operation_indexes = Vec::new();
    let mut operations = Vec::new();
    let mut lines = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap().to_string();

        if !line.contains("+") {
            lines.push(line);
            continue;
        }

        for (i, c) in line.char_indices() {
            let operation = char::to_string(&c);
            if operation != " " {
                operation_indexes.push(i);
                operations.push(operation);
            }
        }
    }

    let line_length = lines[0].len();

    for i in 0..operation_indexes.len() {
        let operation_index = operation_indexes[i];

        let end = if i == operation_indexes.len() - 1 {
            line_length
        } else {
            operation_indexes[i + 1] - 1
        };

        let mut problem_result = 1;

        for j in operation_index..end {
            let mut num_as_string = "".to_string();

            for line in &lines {
                if &line[j..j + 1] != " " {
                    num_as_string.push_str(&line[j..j + 1]);
                }
            }

            let num: i64 = num_as_string.parse().unwrap();

            problem_result = match operations[i].as_str() {
                "+" => problem_result + num,
                "*" => problem_result * num,
                _ => panic!(),
            };
        }

        if operations[i] == "+" {
            problem_result -= 1;
        }

        result += problem_result;
    }

    result
}
