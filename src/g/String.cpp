#include <gazebo/g/String.hpp>

namespace gazebo::g
{

String::~String()
{
}

size_t String::hash() const
{
    return 1;
}

}  // namespace gazebo::g
