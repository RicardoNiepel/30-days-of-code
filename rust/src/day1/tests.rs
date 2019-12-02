#[test]
fn it_works() {
    let input = b"12\n4.0\nis the best place to learn and practice coding!";
    let expected_output = "16\n8.0\nHackerRank is the best place to learn and practice coding!";

    let mut output = Vec::new();

    super::run(&input[..], &mut output);

    let output = String::from_utf8(output).expect("Not UTF-8");

    assert_eq!(expected_output, output);
}
