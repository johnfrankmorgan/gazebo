#ifndef TYPE_HPP
#define TYPE_HPP

#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>

namespace gazebo::g
{

class Type : public Object
{
  public:
    virtual const char* name() const = 0;

    virtual RefPtr<Type> type() const;

    virtual size_t hash(RefPtr<Object>) const;

    // protocol methods

    virtual RefPtr<Bool>   g_bool(RefPtr<Object>) const;
    virtual RefPtr<String> g_repr(RefPtr<Object>) const;
    virtual RefPtr<String> g_str(RefPtr<Object>) const;

    template <class T>
    inline bool is() const
    {
        return is(get<T>().get());
    }

    virtual inline bool is(Type* type) const
    {
        G_ASSERT(type);

        return this == type;
    }

    virtual inline bool is(RefPtr<Type> type) const
    {
        return is(type.get());
    }

    virtual inline void guard(RefPtr<Object> object) const
    {
        G_ASSERT(is(object->type()));
    }

    template <class T>
    static RefPtr<T> get()
    {
        static RefPtr<T> type_object = std::make_shared<T>();
        return type_object;
    }

    template <class T>
    static RefPtr<T> cast(RefPtr<Object> object)
    {
        return *(RefPtr<T>*)&object;
    }
};

}  // namespace gazebo::g

#endif /* TYPE_HPP */
