#ifndef REFPTR_HPP
#define REFPTR_HPP

#include <memory>

template <class T>
using RefPtr = std::shared_ptr<T>;

#endif /* REFPTR_HPP */
