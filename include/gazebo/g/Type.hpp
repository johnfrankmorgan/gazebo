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

    template <class T>
    inline bool is() const
    {
        return is(get<T>().get());
    }

    inline bool is(Type* type) const
    {
        G_ASSERT(type);

        return this == type;
    }

    template <class T>
    RefPtr<T> cast(RefPtr<Object> object)
    {
        G_ASSERT(object->type()->is(this));

        return *(RefPtr<T>*)&object;
    }

    template <class T>
    static RefPtr<T> get()
    {
        static RefPtr<T> type_object = ref<T>();
        return type_object;
    }
};

}  // namespace gazebo::g

#endif /* TYPE_HPP */
