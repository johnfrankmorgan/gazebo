#ifndef OBJECT_HPP
#define OBJECT_HPP

#include <stdlib.h>

namespace gazebo::g
{

class Object {
  public:
    virtual ~Object()           = 0;
    virtual size_t hash() const = 0;
};

}  // namespace gazebo::g

#endif /* OBJECT_HPP */
