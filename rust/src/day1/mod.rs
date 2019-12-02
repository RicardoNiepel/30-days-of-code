use std::io::{self, BufRead, Write};

#[cfg(test)]
mod tests;

pub fn run<R, W>(reader: R, mut writer: W)
where
    R: BufRead,
    W: Write,
{
    let mut i = 4u64;
    let mut d = 4.0f64;
    let string = "HackerRank ";

    let mut lines_iter = reader.lines().map(|l| l.unwrap());

    let mut input = lines_iter.next().unwrap();
    i = i + input.parse::<u64>().unwrap();

    input = lines_iter.next().unwrap();
    d = d + input.parse::<f64>().unwrap();

    input = lines_iter.next().unwrap();
    let output_string = &format!("{}{}", string, &input).to_string();

    write!(&mut writer, "{}\n", i).expect("Unable to write");
    write!(&mut writer, "{:.1}\n", d).expect("Unable to write");
    write!(&mut writer, "{}", output_string).expect("Unable to write");
}

#[allow(dead_code)]
fn main() {
    let stdio = io::stdin();
    let input = stdio.lock();

    let output = io::stdout();

    run(input, output);
}
