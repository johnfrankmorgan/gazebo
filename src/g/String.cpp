#include <gazebo/g/String.hpp>
#include <gazebo/g/Type.hpp>
#include <gazebo/g/TypeString.hpp>

namespace gazebo::g
{

String::String(std::string value) : m_value(value)
{
}

RefPtr<Type> String::type() const
{
    return Type::get<TypeString>();
}

}  // namespace gazebo::g
