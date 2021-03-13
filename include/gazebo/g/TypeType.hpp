#ifndef TYPETYPE_HPP
#define TYPETYPE_HPP

#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

class TypeType : public Type
{
  public:
    virtual const char* name() const
    {
        return "Type";
    }
};

}  // namespace gazebo::g

#endif /* TYPETYPE_HPP */
