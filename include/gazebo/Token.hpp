#ifndef TOKEN_HPP
#define TOKEN_HPP

#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

#include <gazebo/g/String.hpp>

#define _EACH_TOKEN_TYPE(CB) \
    CB(INVALID)              \
    CB(END)                  \
    CB(COMMENT)              \
    CB(WHITESPACE)           \
    CB(PAREN_OPEN)           \
    CB(PAREN_CLOSE)          \
    CB(BRACKET_OPEN)         \
    CB(BRACKET_CLOSE)        \
    CB(BRACE_OPEN)           \
    CB(BRACE_CLOSE)          \
    CB(COLON)                \
    CB(SEMICOLON)            \
    CB(DOT)                  \
    CB(COMMA)                \
    CB(EQUAL)                \
    CB(EQUAL_EQUAL)          \
    CB(BANG)                 \
    CB(BANG_EQUAL)           \
    CB(LESS)                 \
    CB(LESS_EQUAL)           \
    CB(GREATER)              \
    CB(GREATER_EQUAL)        \
    CB(PLUS)                 \
    CB(MINUS)                \
    CB(STAR)                 \
    CB(SLASH)                \
    CB(IDENTIFIER)           \
    CB(NUMBER)               \
    CB(STRING)

#define _TYPE(name) name,
#define _TYPE_STRING(name) #name,

namespace gazebo
{

const std::vector<const char*> _token_type_names = {_EACH_TOKEN_TYPE(_TYPE_STRING)};

class Token
{
  public:
    enum class Type { _EACH_TOKEN_TYPE(_TYPE) };

  private:
    Type        m_type;
    std::string m_value;
    size_t      m_position;

  public:
    Token() : m_type(Type::INVALID), m_value(""), m_position(0)
    {
    }

    Token(Type type, std::string& value, size_t position)
        : m_type(type), m_value(value), m_position(position)
    {
    }

    inline Type type() const
    {
        return m_type;
    }

    inline const std::string& value() const
    {
        return m_value;
    }

    inline size_t position() const
    {
        return m_position;
    }

    inline bool is(Type type) const
    {
        return this->type() == type;
    }

    std::string string() const
    {
        std::stringstream ss;

        if ((size_t)type() < 0 || (size_t)type() >= _token_type_names.size())
            ss << "UNKNOWN";
        else
            ss << _token_type_names[(size_t)type()];

        g::String lexeme = g::String(value()).replace("\n", "\\n");

        ss << "{ " << position() << ", " << std::quoted(lexeme.value()) << " }";

        return ss.str();
    }
};

}  // namespace gazebo

#endif /* TOKEN_HPP */
