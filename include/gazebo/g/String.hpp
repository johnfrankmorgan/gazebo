#ifndef STRING_HPP
#define STRING_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeString.hpp>

namespace gazebo::g
{

class String : public BasicObject<TypeString, std::string>
{
  public:
    String(const char* value) : BasicObject(value)
    {
    }
};

}  // namespace gazebo::g

#endif /* STRING_HPP */
