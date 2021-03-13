#ifndef STRING_HPP
#define STRING_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeString.hpp>

namespace gazebo::g
{

class String : public BasicObject<TypeString, std::string>
{
  public:
    explicit String(const char* value) : String(std::string(value))
    {
    }

    explicit String(std::string value) : BasicObject(value)
    {
    }

    inline size_t length() const
    {
        return value().length();
    }
};

}  // namespace gazebo::g

#endif /* STRING_HPP */
