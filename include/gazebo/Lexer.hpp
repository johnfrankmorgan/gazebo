#ifndef LEXER_HPP
#define LEXER_HPP

#include <sstream>
#include <string>

#include <gazebo/Token.hpp>
#include <vector>

namespace gazebo
{

class Lexer
{
  private:
    std::string       m_source;
    std::stringstream m_buffer;
    size_t            m_position;

  public:
    Lexer(std::string source) : m_source(source), m_buffer(""), m_position(0)
    {
    }

    Lexer(const char* source) : Lexer(std::string(source))
    {
    }

  private:
    inline bool finished() const
    {
        return m_position >= m_source.length();
    }

    inline char peek() const
    {
        if (finished())
            return 0;

        return m_source[m_position];
    }

    inline char next()
    {
        if (finished())
            return 0;

        char ch = m_source[m_position++];

        m_buffer.put(ch);

        return ch;
    }

    inline bool check(char ch) const
    {
        if (finished())
            return false;

        return peek() == ch;
    }

    inline bool match(char ch)
    {
        if (check(ch)) {
            next();
            return true;
        }

        return false;
    }

    inline Token token(Token::Type type)
    {
        std::string value = m_buffer.str();
        m_buffer.str("");

        return Token(type, value, m_position - value.length());
    }

    Token identifier();
    Token number();
    Token string();

  public:
    void               line();
    Token              lex();
    std::vector<Token> lex_all();
};

}  // namespace gazebo

#endif /* LEXER_HPP */
