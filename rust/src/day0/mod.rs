use std::io::{self, BufRead, Write};

#[cfg(test)]
mod tests;

fn run<R, W>(mut reader: R, mut writer: W)
where
    R: BufRead,
    W: Write,
{
    let mut input = String::new();
    match reader.read_line(&mut input) {
        Ok(_n) => {
            write!(&mut writer, "Hello, World.\n").expect("Unable to write");
            write!(&mut writer, "{}", input).expect("Unable to write");
        }
        Err(error) => write!(&mut writer, "Error: {}", error).expect("Unable to write"),
    }
}

#[allow(dead_code)]
fn main() {
    let stdio = io::stdin();
    let input = stdio.lock();

    let output = io::stdout();

    run(input, output);
}
