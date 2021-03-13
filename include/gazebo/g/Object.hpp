#ifndef OBJECT_HPP
#define OBJECT_HPP

#include <gazebo.hpp>

namespace gazebo::g
{

class Object
{
  public:
    virtual ~Object()
    {
    }

    virtual RefPtr<Type> type() const = 0;

    template <class T, class... Args>
    static RefPtr<T> create(Args... args)
    {
        return std::make_shared<T>(std::forward<Args>(args)...);
    }
};

}  // namespace gazebo::g

#endif /* OBJECT_HPP */
