#ifndef GAZEBO_H
#define GAZEBO_H

#include <assert.h>
#include <stdio.h>
#include <string>
#include <utility>

#include <gazebo/RefPtr.hpp>

namespace gazebo::g
{

class Type;
class Object;
class Bool;
class Number;
class String;

}  // namespace gazebo::g

#define G_DEBUG(...)                                                               \
    do {                                                                           \
        fprintf(stderr, "%15.15s:%-4d %10.10s :: ", __FILE__, __LINE__, __func__); \
        fprintf(stderr, __VA_ARGS__);                                              \
        fprintf(stderr, "\n");                                                     \
    } while (0)

#define G_ASSERT(condition) assert(condition)
#define G_UNREACHED() G_ASSERT(false)

#endif /* GAZEBO_H */
