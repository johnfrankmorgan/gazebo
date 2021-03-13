#ifndef NUMBER_HPP
#define NUMBER_HPP

#include <gazebo/g/BasicObject.hpp>
#include <gazebo/g/TypeNumber.hpp>

namespace gazebo::g
{

class Number : public BasicObject<TypeNumber, long double>
{
};

}  // namespace gazebo::g

#endif /* NUMBER_HPP */
