#include <gazebo/Objects.hpp>

namespace gazebo::g
{

RefPtr<Bool> TypeBool::g_bool(RefPtr<Object> self) const
{
    guard(self);

    return Object::create<Bool>(cast<Bool>(self)->value());
}

RefPtr<String> TypeBool::g_str(RefPtr<Object> self) const
{
    guard(self);

    if (cast<Bool>(self)->value())
        return Object::create<String>("true");

    return Object::create<String>("false");
}

}  // namespace gazebo::g
