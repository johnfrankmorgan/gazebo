#ifndef STRING_HPP
#define STRING_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeString.hpp>

namespace gazebo::g
{

class String : public BasicObject<TypeString, std::string>
{
  public:
    explicit String(char* value) : String(std::string(value))
    {
    }

    explicit String(std::string value) : BasicObject(value)
    {
    }

    inline size_t length() const
    {
        return value().length();
    }

    String replace(const String& find, const String& replace) const
    {
        size_t pos = value().find(find.value());

        if (pos == std::string::npos)
            return String("");

        String ret(value());

        ret.value().replace(pos, find.length(), replace.value());

        return ret;
    }

    String replace(const char* find, const char* replace) const
    {
        String f(find);
        String r(replace);
        return this->replace(f, r);
    }
};

}  // namespace gazebo::g

#endif /* STRING_HPP */
