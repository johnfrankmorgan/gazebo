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
    bool is() const
    {
        return this == get<T>().get();
    }

    template <class T>
    static RefPtr<T> get()
    {
        static RefPtr<T> type_object = ref<T>();
        return type_object;
    }

    template <class T, class V>
    static RefPtr<V> cast(RefPtr<Object> object)
    {
        G_ASSERT(object->type()->is<T>());
        return *(RefPtr<V>*)&object;
    }
};

}  // namespace gazebo::g

#endif /* TYPE_HPP */
