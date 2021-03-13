#include <gazebo/Objects.hpp>

namespace gazebo::g
{

RefPtr<Bool> TypeNumber::g_bool(RefPtr<Object> self) const
{
    guard(self);

    return Object::create<Bool>(cast<Number>(self)->value() != 0.0);
}

RefPtr<String> TypeNumber::g_str(RefPtr<Object> self) const
{
    guard(self);

    return Object::create<String>(std::to_string(cast<Number>(self)->value()));
}

}  // namespace gazebo::g
