#ifndef TYPEBOOL_HPP
#define TYPEBOOL_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

class TypeBool : public Type
{
  public:
    virtual const char* name() const
    {
        return "Bool";
    }

    // protocols

    virtual RefPtr<Bool>   g_bool(RefPtr<Object>) const;
    virtual RefPtr<String> g_str(RefPtr<Object>) const;
};

}  // namespace gazebo::g

#endif /* TYPEBOOL_HPP */
