#ifndef BOOL_HPP
#define BOOL_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeBool.hpp>

namespace gazebo::g
{

class Bool : public BasicObject<TypeBool, bool>
{
  public:
    explicit Bool(bool value) : BasicObject(value)
    {
    }
};

}  // namespace gazebo::g

#endif /* BOOL_HPP */
