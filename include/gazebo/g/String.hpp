#ifndef STRING_HPP
#define STRING_HPP

#include <gazebo/g/Object.hpp>

namespace gazebo::g
{

class String {
  public:
    virtual ~String();
    virtual size_t hash() const;
};

}  // namespace gazebo::g

#endif /* STRING_HPP */
