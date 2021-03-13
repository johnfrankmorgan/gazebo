CPP = g++
CPP_FLAGS = -I./include -g -Wall -Wextra -Wpedantic -std=c++17

CPP_FLAGS += -DTESTING

ifeq ($(shell uname -s),Darwin)
	CPP_FLAGS += -I/Library/Developer/CommandLineTools/usr/include/c++/v1
endif

PROG = gazebo

CPP_HED = $(shell find include -type f -name '*.hpp' ! -name catch.hpp)
CPP_SRC = $(shell find src -type f -name '*.cpp')
CPP_OBJ = $(CPP_SRC:.cpp=.o)

CPP_TEST_SRC = $(shell find tests -type f -name '*.cpp')
CPP_TEST_OBJ = $(CPP_TEST_SRC:.cpp=.o)

$(PROG): $(CPP_OBJ) $(CPP_TEST_OBJ)
	$(CPP) $(CPP_FLAGS) $(CPP_OBJ) $(CPP_TEST_OBJ) -o $(PROG)

%.o: %.cpp
	$(CPP) $(CPP_FLAGS) -c $< -o $@

.PHONY: clean
clean:
	rm -f $(PROG) $(CPP_OBJ) $(CPP_TEST_OBJ)

.PHONY: compile-commands
compile-commands: clean
	bear -- make

.PHONY: format
format:
	clang-format -i $(CPP_HED) $(CPP_SRC) $(CPP_TEST_SRC)

.PHONY: test
test: $(PROG)
	./$(PROG) -s
