use std::{
    fs::File,
    io::{BufRead, BufReader},
};

struct Position {
    x: usize,
    y: usize,
}

pub fn part1(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    let mut matrix = vec![];

    for line in reader.lines() {
        let line = line.unwrap();

        let mut row = vec![];

        for i in 0..line.len() {
            row.push(line[i..i + 1].to_string());
        }

        matrix.push(row);
    }

    for i in 0..matrix.len() {
        for j in 0..matrix[i].len() {
            if matrix[i][j] == "." {
                continue;
            }

            let mut rolls = 0;

            // NW
            if i != 0 && j != 0 && matrix[i - 1][j - 1] == "@" {
                rolls += 1;
            }

            // N
            if i != 0 && matrix[i - 1][j] == "@" {
                rolls += 1;
            }

            // NE
            if i != 0 && j + 1 < matrix[i].len() && matrix[i - 1][j + 1] == "@" {
                rolls += 1;
            }

            // E
            if j + 1 < matrix[i].len() && matrix[i][j + 1] == "@" {
                rolls += 1;
            }

            // SE
            if i + 1 < matrix.len() && j + 1 < matrix[i].len() && matrix[i + 1][j + 1] == "@" {
                rolls += 1;
            }

            // S
            if i + 1 < matrix.len() && matrix[i + 1][j] == "@" {
                rolls += 1;
            }

            // SW
            if i + 1 < matrix.len() && j != 0 && matrix[i + 1][j - 1] == "@" {
                rolls += 1;
            }

            // W
            if j != 0 && matrix[i][j - 1] == "@" {
                rolls += 1;
            }

            if rolls < 4 {
                result += 1;
            }

            // println!("{},{},{}", i, j, rolls)
        }
    }

    // for i in 0..matrix.len() {
    //     println!("{:?}", matrix[i]);
    // }

    result
}

pub fn part2(file: &File) -> i64 {
    let reader = BufReader::new(file);

    let mut result = 0;

    let mut matrix = vec![];

    for line in reader.lines() {
        let line = line.unwrap();

        let mut row = Vec::new();

        for i in 0..line.len() {
            row.push(line[i..i + 1].to_string());
        }

        matrix.push(row);
    }

    let mut removed = true;

    while removed {
        removed = false;

        let mut positions: Vec<Position> = Vec::new();

        for i in 0..matrix.len() {
            for j in 0..matrix[i].len() {
                if matrix[i][j] == "." {
                    continue;
                }

                let mut rolls = 0;

                // NW
                if i != 0 && j != 0 && matrix[i - 1][j - 1] == "@" {
                    rolls += 1;
                }

                // N
                if i != 0 && matrix[i - 1][j] == "@" {
                    rolls += 1;
                }

                // NE
                if i != 0 && j + 1 < matrix[i].len() && matrix[i - 1][j + 1] == "@" {
                    rolls += 1;
                }

                // E
                if j + 1 < matrix[i].len() && matrix[i][j + 1] == "@" {
                    rolls += 1;
                }

                // SE
                if i + 1 < matrix.len() && j + 1 < matrix[i].len() && matrix[i + 1][j + 1] == "@" {
                    rolls += 1;
                }

                // S
                if i + 1 < matrix.len() && matrix[i + 1][j] == "@" {
                    rolls += 1;
                }

                // SW
                if i + 1 < matrix.len() && j != 0 && matrix[i + 1][j - 1] == "@" {
                    rolls += 1;
                }

                // W
                if j != 0 && matrix[i][j - 1] == "@" {
                    rolls += 1;
                }

                if rolls < 4 {
                    result += 1;
                    removed = true;
                    positions.push(Position { x: i, y: j });
                }

                // println!("{},{},{}", i, j, rolls)
            }
        }

        for pos in &positions {
            matrix[pos.x][pos.y] = String::from(".");
        }
    }

    // for i in 0..matrix.len() {
    //     println!("{:?}", matrix[i]);
    // }

    result
}
