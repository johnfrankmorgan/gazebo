use std::fmt::Write;

use super::token;
use super::token::Token;

#[derive(Debug)]
pub struct Lexer {
    source: Vec<char>,
    buffer: String,
    position: usize,
}

fn is_digit(ch: char) -> bool {
    ch >= '0' && ch <= '9'
}

fn is_newline(ch: char) -> bool {
    ch == '\n'
}

fn is_whitespace(ch: char) -> bool {
    ch == ' ' || ch == '\t' || is_newline(ch)
}

fn is_identifier(ch: char) -> bool {
    if is_whitespace(ch) {
        return false;
    };

    match ch {
        '(' | ')' | ';' | '"' => false,
        _ => true,
    }
}

impl Lexer {
    pub fn new(source: String) -> Lexer {
        Lexer {
            source: source.trim().chars().collect(),
            buffer: String::new(),
            position: 0,
        }
    }

    fn finished(&self) -> bool {
        self.position >= self.source.len()
    }

    fn peek(&self) -> char {
        self.source[self.position]
    }

    fn next(&mut self) -> char {
        let ch = self.peek();
        self.buffer
            .write_char(ch)
            .expect("failed to write to buffer");
        self.position += 1;
        ch
    }

    fn token(&mut self, typ: token::Type) -> Token {
        let val = self.buffer.clone();
        self.buffer.clear();
        Token::new(typ, val)
    }

    fn lex_eol(&mut self, typ: token::Type) -> Token {
        while !self.finished() {
            if is_newline(self.next()) {
                return self.token(typ);
            }
        }

        unreachable!()
    }

    fn lex_whitespace(&mut self) -> Token {
        while !self.finished() {
            if !is_whitespace(self.peek()) {
                return self.token(token::Type::Whitespace);
            }

            self.next();
        }

        unreachable!()
    }

    fn lex_number(&mut self) -> Token {
        let mut float = false;

        while !self.finished() {
            let ch = self.peek();

            if ch == '.' && !float {
                self.next();
                float = true;
                continue;
            }

            if !is_digit(ch) {
                return self.token(token::Type::Number);
            }

            self.next();
        }

        unreachable!()
    }

    fn lex_string(&mut self) -> Token {
        while !self.finished() {
            if self.next() == '"' {
                return self.token(token::Type::String);
            }
        }

        unreachable!()
    }

    fn lex_identifier(&mut self) -> Token {
        while !self.finished() {
            if !is_identifier(self.peek()) {
                return self.token(token::Type::Identifier);
            }

            self.next();
        }

        unreachable!()
    }

    fn next_token(&mut self) -> Token {
        let ch = self.next();

        match ch {
            '(' => self.token(token::Type::ParenOpen),
            ')' => self.token(token::Type::ParenClose),
            ';' => self.lex_eol(token::Type::Comment),
            '"' => self.lex_string(),
            _ => {
                if is_whitespace(ch) {
                    return self.lex_whitespace();
                }

                if is_digit(ch) {
                    return self.lex_number();
                }

                if is_identifier(ch) {
                    return self.lex_identifier();
                }

                unreachable!()
            }
        }
    }

    pub fn lex(&mut self) -> Vec<Token> {
        let mut tokens = Vec::new();

        while !self.finished() {
            tokens.push(self.next_token());
        }

        tokens
    }
}
