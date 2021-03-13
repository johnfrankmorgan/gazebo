#include <gazebo.hpp>
#include <gazebo/objects.hpp>

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
    const RefPtr<Object> object = ref<String>("Test");
    const RefPtr<String> str    = Type::get<TypeString>()->cast<String>(object);

    REQUIRE(str->value() == "Test");
    REQUIRE(object.use_count() == str.use_count());
}
