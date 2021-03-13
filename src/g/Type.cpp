#include <gazebo/g/Type.hpp>
#include <gazebo/g/TypeType.hpp>

namespace gazebo::g
{

RefPtr<Type> Type::type() const
{
    return get<TypeType>();
}

size_t Type::hash(RefPtr<Object>) const
{
    G_DEBUG("Type::hash called");
    G_UNREACHED();
}

}  // namespace gazebo::g
