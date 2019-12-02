#[test]
fn it_works() {
    let input = b"Welcome to 30 Days of Code!\n";
    let expected_output = "Hello, World.\nWelcome to 30 Days of Code!\n";

    let mut output = Vec::new();

    super::run(&input[..], &mut output);

    let output = String::from_utf8(output).expect("Not UTF-8");

    assert_eq!(expected_output, output);
}
