mod compiler;

fn main() {
    let mut lexer = compiler::Lexer::new("(test 1 2 3)".to_string());
    println!("{:?}", lexer.lex())
}
