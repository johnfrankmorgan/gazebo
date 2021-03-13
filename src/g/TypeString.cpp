#include <gazebo/Objects.hpp>

namespace gazebo::g
{

RefPtr<Bool> TypeString::g_bool(RefPtr<Object> self) const
{
    guard(self);

    return Object::create<Bool>(cast<String>(self)->length() > 0);
}

RefPtr<String> TypeString::g_str(RefPtr<Object> self) const
{
    guard(self);

    return Object::create<String>(cast<String>(self)->value());
}

}  // namespace gazebo::g
