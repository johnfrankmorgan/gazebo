#ifndef TYPESTRING_HPP
#define TYPESTRING_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/Type.hpp>

namespace gazebo::g
{

class TypeString : public Type
{
  public:
    virtual const char* name() const
    {
        return "String";
    }

    // protocols

    virtual RefPtr<Bool>   g_bool(RefPtr<Object>) const;
    virtual RefPtr<String> g_str(RefPtr<Object>) const;
};

}  // namespace gazebo::g

#endif /* TYPESTRING_HPP */
