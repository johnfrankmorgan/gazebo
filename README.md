# Gazebo :circus_tent:

An **awesome** project that helps you create other awesome (and cool) projects!

## Building

You must have `go` installed to build Gazebo.

```bash
go build -o gazebo main.go
```

## Running

Once built, you can use `./gazebo` to run a file or start a REPL.

```bash
# To start a REPL
./gazebo

# To run a file
./gazebo filename.gaz

# To enable (lots) of (useless) debug output
./gazebo -d ...
```

## Syntax

```
name = "Frank";

greet = fun (name) {
    out.println("Hello,", name + "!");
};

greet(name);

fun fibonacci(n) {
    if n < 2 {
        return n;
    }

    return fibonacci(n - 1) + fibonacci(n - 2);
};

load time;

N = 22;
start = time.now();
out.printf("fibonacci(%d) = %d\n", N, fibonacci(N));
out.println("Took:", time.since(start));
```

## Types / Methods / Modules

The built in types are:
* `g.Nil`
* `g.Bool`
* `g.String`
* `g.Number`
* `g.List`
* `g.Map`
* `g.Reader`
* `g.Writer`
* `vm.Fun`

```
nil         // g.Nil
true, false // g.Bool
"string"    // g.String
123.4       // g.Number
[]          // g.List
{}          // g.Map
in          // g.Reader (wrapping stdin)
out         // g.Writer (wrapping stdout)
fun {}      // vm.Fun
```

The built in modules are:
* `http`
* `inspect`
* `os`
* `testing`
* `time`

Objects can have arbitrary attributes set on them:
```
list = [1, 2, 3];
list.name = "Whatever";
list.other = -list;
```

Methods can be called on Gazebo objects.

To see the available methods for a Gazebo object (include modules), you can use the `inspect` module:

```
load inspect;

methods = inspect.methods(inspect);

// Filter out protocol methods
methods.filter(fun (method) {
    return !inspect.protocol(method);
}).each(fun (method) {
    out.println(method.name());
});


// Strings

str = "Test";
str.upper();             // TEST
str.lower();             // test
str.find("e");           // 1
str.replace("es", "ES"); // TESt
str.empty();             // false
"".empty();              // true
str.from(1);             // "est"
str.until(3);            // "Tes";
str.numeric();           // false
"123.4".numeric();       // true
"123.4".num();           // 123.4
str = -str;              // tseT


// Lists

list = [];
list.append("", nil, 21, true, out.println, inspect);
list.get(1); // nil
list.len();  // 6

list.map(fun (value, index) {});
list.filter(fun (value, index) { return value == "test"; });
list.each(fun (value, index) { out.println(value); });

[false, false, true].any(); // true
[false, false, true].all(); // false

-[1, 2, 3]; // [3, 2, 1]


// Maps

map = {
    name: "Frank",
    address: [
        "somewhere",
        "over",
        "the",
        "rainbow",
    ],
};

map.keys();          // ["name", "address"]
map.values();        // ["Frank", ["somewhere", "over", ...]]
map.has("age");      // false
map.get("age");      // nil
map.get("age", 101); // 101
map.pop("name");     // Frank
map.has("name");     // false
map.set("age", 123);
map.keys();          // ["address", "age"]

map.each(fun (key, value) {
    out.println(key, "->", value);
});
```

## Reading / Writing Files

File operations are performed using the `os` module:

```
load os;

FILEPATH = os.path.join("./output", "file.txt");

if !os.exists(os.path.dirname(FILEPATH)) {
    os.mkdir(os.path.dirname(FILEPATH));
}

file = os.open(FILEPATH, "w");
file.println("Hello!");
file.println("test");
file.close();

file = os.open(FILEPATH, "r");
lines = [];

while true {
    line = file.readln();

    if !line {
        break;
    }

    lines.append(line);
}

file.close();
out.println(lines);
```
