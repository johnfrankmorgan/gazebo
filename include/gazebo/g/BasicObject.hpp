#ifndef BASICOBJECT_HPP
#define BASICOBJECT_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

template <class GType, class T>
class BasicObject : public Object
{
  protected:
    T m_value;

  public:
    BasicObject(T value) : m_value(value)
    {
    }

    virtual RefPtr<Type> type() const
    {
        return Type::get<GType>();
    }

    inline const T& value() const
    {
        return m_value;
    }

    inline T& value()
    {
        return m_value;
    }
};

}  // namespace gazebo::g

#endif /* BASICOBJECT_HPP */
