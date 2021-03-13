#ifndef STRING_HPP
#define STRING_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>

namespace gazebo::g
{

class String : public Object
{
  private:
    std::string m_value;

  public:
    String(std::string value);
    virtual RefPtr<Type> type() const;

    inline const std::string& value() const
    {
        return m_value;
    }
};

}  // namespace gazebo::g

#endif /* STRING_HPP */
