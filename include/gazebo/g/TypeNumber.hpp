#ifndef TYPENUMBER_HPP
#define TYPENUMBER_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

class TypeNumber : public Type
{
  public:
    virtual const char* name() const
    {
        return "Number";
    }
};

}  // namespace gazebo::g

#endif /* TYPENUMBER_HPP */
