#ifndef BOOL_HPP
#define BOOL_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeBool.hpp>

namespace gazebo::g
{

class Bool : public BasicObject<TypeBool, bool>
{
};

}  // namespace gazebo::g

#endif /* BOOL_HPP */
