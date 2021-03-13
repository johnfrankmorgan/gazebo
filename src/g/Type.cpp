#include <gazebo/Objects.hpp>
#include <sstream>

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

RefPtr<Bool> Type::g_bool(RefPtr<Object>) const
{
    return Object::create<Bool>(true);
}

RefPtr<String> Type::g_repr(RefPtr<Object> self) const
{
    std::stringstream ss;

    ss << name() << '(' << self.get() << ')';

    return Object::create<String>(ss.str());
}

RefPtr<String> Type::g_str(RefPtr<Object> self) const
{
    return g_repr(self);
}

}  // namespace gazebo::g
