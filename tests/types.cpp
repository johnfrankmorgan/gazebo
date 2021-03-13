#include <gazebo.hpp>
#include <gazebo/g/Object.hpp>
#include <gazebo/g/String.hpp>
#include <gazebo/g/Type.hpp>
#include <gazebo/g/TypeString.hpp>
#include <gazebo/g/TypeType.hpp>

#include <catch.hpp>

using namespace gazebo::g;

TEST_CASE("g::Type::is")
{
    const auto type_string = Type::get<TypeString>();

    REQUIRE(type_string->is<TypeString>());
    REQUIRE(type_string->type()->is<TypeType>());
}

TEST_CASE("g::String")
{
    const RefPtr<Object> object = Object::create<String>("Test");
    const RefPtr<String> str    = Type::cast<TypeString, String>(object);

    REQUIRE(str->value() == "Test");
    REQUIRE(object.use_count() == str.use_count());
}
