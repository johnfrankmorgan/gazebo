#ifndef TYPESTRING_HPP
#define TYPESTRING_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

class TypeString : public Type
{
  public:
    virtual const char* name() const
    {
        return "String";
    }
};

}  // namespace gazebo::g

#endif /* TYPESTRING_HPP */
