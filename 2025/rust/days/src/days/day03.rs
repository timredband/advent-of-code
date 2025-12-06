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
    let reader = BufReader::new(file);

    let mut result = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        let length = line.len();
        let mut matrix = vec![vec![0; length]; 12];

        for j in (0..length).rev() {
            for i in 0..matrix.len() {
                if ((length as i64) - (j as i64) - (i as i64) - 1) < 0 {
                    matrix[i][j] = -1;
                }
            }
        }

        for j in (0..length).rev() {
            for i in 0..matrix.len() {
                let mut current = line[j..j + 1].to_string();
                let mut previous: i64 = -1;

                if i != 0 && j + 1 < length && matrix[i - 1][j + 1] != -1 {
                    current.push_str(&matrix[i - 1][j + 1].to_string());
                }

                if j + 1 < length {
                    previous = matrix[i][j + 1];
                }

                matrix[i][j] = cmp::max(current.parse().unwrap(), previous);
            }
        }

        result += matrix[11][0];
    }

    result
}
