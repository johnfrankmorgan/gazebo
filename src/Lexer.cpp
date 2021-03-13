#include <gazebo.hpp>
#include <gazebo/Lexer.hpp>
#include <vector>

namespace gazebo
{

void Lexer::line()
{
    while (next() != '\n')
        ;
}

Token Lexer::identifier()
{
    return token(Token::Type::IDENTIFIER);
}

Token Lexer::number()
{
    bool isfloat = false;

    while (isdigit(peek())) {
        next();

        if (!isfloat && match('.')) {
            isfloat = true;
            G_ASSERT(isdigit(peek()));
        }
    }

    return token(Token::Type::NUMBER);
}

Token Lexer::string()
{
    return token(Token::Type::STRING);
}

Token Lexer::lex()
{
    char ch = next();

    if (!ch)
        return token(Token::Type::END);

    if (isspace(ch)) {
        while (isspace(peek()))
            next();

        return token(Token::Type::WHITESPACE);
    }

    if (ch == '#') {
        line();
        return token(Token::Type::COMMENT);
    }

    switch (ch) {
    case '(':
        return token(Token::Type::PAREN_OPEN);

    case ')':
        return token(Token::Type::PAREN_CLOSE);

    case '[':
        return token(Token::Type::BRACKET_OPEN);

    case ']':
        return token(Token::Type::BRACKET_CLOSE);

    case '{':
        return token(Token::Type::BRACE_OPEN);

    case '}':
        return token(Token::Type::BRACE_CLOSE);

    case ':':
        return token(Token::Type::COLON);

    case ';':
        return token(Token::Type::SEMICOLON);

    case '.':
        return token(Token::Type::DOT);

    case ',':
        return token(Token::Type::COMMA);

    case '=':
        return match('=') ? token(Token::Type::EQUAL_EQUAL) : token(Token::Type::EQUAL);

    case '!':
        return match('=') ? token(Token::Type::BANG_EQUAL) : token(Token::Type::BANG);

    case '<':
        return match('=') ? token(Token::Type::LESS_EQUAL) : token(Token::Type::LESS);

    case '>':
        return match('=') ? token(Token::Type::GREATER_EQUAL) : token(Token::Type::GREATER);

    case '+':
        return token(Token::Type::PLUS);

    case '-':
        return token(Token::Type::MINUS);

    case '*':
        return token(Token::Type::STAR);

    case '/':
        if (check('/')) {
            line();
            return token(Token::Type::COMMENT);
        }

        return token(Token::Type::SLASH);
    }

    if (isdigit(ch))
        return number();

    G_UNREACHED();
}

std::vector<Token> Lexer::lex_all()
{
    std::vector<Token> tokens;

    Token token;

    do {
        token = lex();

        if (!token.is(Token::Type::WHITESPACE))
            tokens.push_back(token);

    } while (!token.is(Token::Type::END));

    return tokens;
}

}  // namespace gazebo
