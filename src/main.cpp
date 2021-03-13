#ifdef TESTING

#define CATCH_CONFIG_MAIN
#include <catch.hpp>

#else

#include <gazebo/Lexer.hpp>
#include <iostream>

using namespace gazebo;

int main()
{
    std::string source = "// test\n";

    Lexer lexer(source);
    Token tk;

    while (!(tk = lexer.lex()).is(Token::Type::END))
        std::cout << tk.string() << "\n";

    return 0;
}

#endif
