#include <catch.hpp>
#include <vector>

#include <gazebo/Lexer.hpp>

using namespace gazebo;

TEST_CASE("Lexer.cpp")
{
    Lexer lexer("()[]{}= == ! != > >= < <= + - * / 123.33. 123 ");
    auto  tokens = lexer.lex_all();

    const std::vector<Token::Type> expected = {
        Token::Type::PAREN_OPEN,    Token::Type::PAREN_CLOSE, Token::Type::BRACKET_OPEN,
        Token::Type::BRACKET_CLOSE, Token::Type::BRACE_OPEN,  Token::Type::BRACE_CLOSE,
        Token::Type::EQUAL,         Token::Type::EQUAL_EQUAL, Token::Type::BANG,
        Token::Type::BANG_EQUAL,    Token::Type::GREATER,     Token::Type::GREATER_EQUAL,
        Token::Type::LESS,          Token::Type::LESS_EQUAL,  Token::Type::PLUS,
        Token::Type::MINUS,         Token::Type::STAR,        Token::Type::SLASH,
        Token::Type::NUMBER,        Token::Type::DOT,         Token::Type::NUMBER,
        Token::Type::END,
    };

    for (size_t i = 0; i < tokens.size(); ++i) {
        REQUIRE(tokens[i].is(expected[i]));
    }
}
