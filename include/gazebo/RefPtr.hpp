#ifndef REFPTR_HPP
#define REFPTR_HPP

#include <memory>

template <class T>
using RefPtr = std::shared_ptr<T>;

template <class T, class... Args>
inline RefPtr<T> ref(Args... args)
{
    return std::make_shared<T>(std::forward<Args>(args)...);
}

#endif /* REFPTR_HPP */
