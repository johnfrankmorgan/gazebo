NAME = gazebo

BUILD_DIR = ./build
SOURCE_DIR = .

.PHONY: all
all: build

.PHONY: build
build: $(BUILD_DIR)/$(NAME)

.PHONY: run
run: build
	$(BUILD_DIR)/$(NAME)

$(BUILD_DIR)/$(NAME): generate $(shell find $(SOURCE_DIR) -name '*.go')
	go build -o $@ $(SOURCE_DIR)/cmd/$(NAME)

.PHONY:
generate: \
	$(SOURCE_DIR)/grammar/yy.go \
	$(SOURCE_DIR)/ast/expr/binary_string.go \
	$(SOURCE_DIR)/ast/expr/unary_string.go \
	$(SOURCE_DIR)/compile/opcode/op_string.go

$(SOURCE_DIR)/grammar/yy.go: $(SOURCE_DIR)/grammar/yy.go.y
	goyacc -v $(SOURCE_DIR)/grammar/yy.output -o $@ $<

$(SOURCE_DIR)/ast/expr/binary_string.go: $(SOURCE_DIR)/ast/expr/binary.go
	stringer -type=BinaryOp -output=$@ $<

$(SOURCE_DIR)/ast/expr/unary_string.go: $(SOURCE_DIR)/ast/expr/unary.go
	stringer -type=UnaryOp -output=$@ $<

$(SOURCE_DIR)/compile/opcode/op_string.go: $(SOURCE_DIR)/compile/opcode/op.go
	stringer -type=Op -output=$@ $<
