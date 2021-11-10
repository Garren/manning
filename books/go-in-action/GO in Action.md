# GO in Action

Pointers

> “ Pointer variables are great for sharing variables between functions. They allow functions to access and change the state of a variable that was declared within the scope of a different function and possibly a different goroutine.”
>
> Excerpt From: William Kennedy. “Go in Action.” Apple Books. 



Passing arguments to an anonymous function rather than accessing local variables as closures.

> “Go supports closures and you’re seeing this in action. In fact, the searchTerm and results variables are also being accessed by the anonymous function via closures. Thanks to closures, the function can access those variables directly without the need to pass them in as parameters. The anonymous function isn’t given a copy of these variables; it has direct access to the same variables declared in the scope of the outer function. This is the reason why we don’t use closures for the matcher and feed variables.
>
> Listing 2.24. search/search.go: lines 29–32
> 29     // Launch a goroutine for each feed to find the results.
> 30     for _, feed := range feeds {
> 31         // Retrieve a matcher for the search.
> 32         matcher, exists := matchers[feed.Type]
>
> The values of the feed and matcher variables are changing with each iteration of the loop, as you can see on lines 30 and 32. If we used closures for these variables, as the values of these variables changed in the outer function, those changes would be reflected in the anonymous function. All the goroutines would be sharing the same variables as the outer function thanks to closures. Unless we passed these values in as function parameters, most likely the last one in the feeds slice.”
>
> Excerpt From: William Kennedy. “Go in Action.” Apple Books. 



Pointer receivers

> “It’s best practice to declare methods using pointer receivers, since many of the methods you implement need to manipulate the state of the value being used to make the method call. In the case of the defaultMatcher type, we want to use a value receiver because creating values of type defaultMatcher result in values of zero allocation. Using a pointer makes no sense since there’s no state to be manipulated.”
>
> Excerpt From: William Kennedy. “Go in Action.” Apple Books. 



## Packaging and tooling

### Packages

Go code is organized as packages. Packages are groups of files contained in a directory heirarchy that describes or indicates the relationship between files and functionality. 

The http std lib packages looks like this:

net/http/
    cgi/
    cookiejar/
        testdata/
    fcgi/
    httptest/
    httputil/
    pprof/
    testdata/

Packages can be imported individually (i.e., "net/http/cgi".)

All go files declare the package they belong to in the first line of code (excluding whitespace and comments.) Packages are contained in a single directory, they cannot be split across directories, nor can a single directory contain multiple packages.

Packages are named using the directory that contains the package.

The main package specifies the entry point for a program. A main package must have a main() entry point.

### Imports

Packages are imported using an imoprt statement. Import statements can be grouped into blocks.

```go
import "fmt"

import (
	"fmt"
  "string"
)
```

#### Remote Imports

Go has support for referencing source code and packages from distributed version control systems and code sharing sites. 

```go
import "github.com/spf13/viper"
```

Issuing a call to `go get` or `go build` will fetch the remote import, and fetch any related dependencies, then install into `$GOPATH`.

#### Named Imports

Go supports "named imports" allowing a user to specify an alternate package name. Named imports are useful in resolving package name collisions.

````go
import (
	"fmt"
  myfmt "mylib/fmt"
)
````

In the example above, two packages are imported that have the same name. The second one uses a named import "myfmt" that will be used instead of "fmt".

The Go compiler will throw an error when package is imported but not used. However, sometimes it's useful to import a package that you don't reference identifiers or code from directly. In these cases you can use Go's "blank" identifier as an import name.

````go
import (
	_ "mylib/fmt"
)
````

#### Import "bootstrapping" via init

Packages can provide an `init()` function that can be used to initialize variables, prepare the package for use, or perform some kind of bootstrapping. Package `init()` functions are discovered by the compiler and scheduled to run before the call to `main()` is executed.

Below is an example of an `init()` function used by a database driver:

````go
package postgres

import (
	"database/sql"
)

func init() {
  sql.Register("postgres", new(PostgresDriver))
}
....
````

Code that uses this package probably doesn't reference the driver functionality directly, and without a named import using the blank identifier, the compiler would throw an error.

````go
package main

import (
	"database/sql"

  // use a anonymous named import to allow the compiler to accept the 
  // package import even though we're not using it explicitly. Doing this
  // will also fire off the package's init(), which will register the driver
  // with the SQL package and allow us to use it below.
  _ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

func main() {
  sql.Open("postgres", "mydb")
}
````



## Arrays, Slices, and Maps

### Arrays

Go arrays are fixed length contiguous collections of elements having the same type. The type can be a built-in type or an aggregate created via structs. 

#### Declaring and Initializing

The easiest way to create an array is using an array literal. E.g.,

````go
// an array literal with explicit size
array := [5]int{1,2,3,4,5}

// an array literal using implicit size
array := [...]int{6,7,8,9,10}
````

Arrays also allow for index-specific initialization.

````go
array := [5]int{1:10, 2:20} // [0,10,20,0,0]
````

#### Working with arrays

Arrays are index accessible in the usual manner.

````go
array[2] = 35
````

Arrays can contain pointers.

````go
// an array of five pointers to int. the first two elements
// are initialized to heap allocated ints (zero initialized)
array := [5]*int{0: new(int), 1: new(int)}

// set the two elements through the array element's pointer
*array[0] = 10
*array[1] = 20
````

Arrays are value types in Go, meaning they can be assigned to variables and copied.

````go
var array1 [5]string
array2 := [5]string{"zero", "one", "two", "three", "four"}
array1 = array2 // create a copy
````

An array is defined by its type and length. Assignment and copying requires that the corresponding array variables have the same type and length.

Copying an array of pointers is a shallow operation. The result will be a second array that points to the same objects.

````go
var array1 = [3]*string
array2 := [3]*string{new(string), new(string), new(string)}
array1 = array2
// array1 contains copies of array2's pointers, the pointees are not copied
````



#### Multidimensional arrays

Multidimensional arrays are not directly supported in Go, but it's trivial to create arrays of arrays.

````go
var array = [4][2]int

// literal syntax is supported
array := [4][2]int{{0,1},{1,2},{2,3},{3,4}}

// index-initialization is supported
array := [4][2]int{0: {0,1},1: {1,2}}

// index-initialization is supported in both the outer and inner array
array := [4][2]int{0: {0: 33, 1: 44}, 1: {0: 11, 1: 222}}
````

Accessing multi-dimensional values isn't a surprise.

````go
var array [2][2]int

array[0][0] = 10
array[1][1] = 100
````

Copying multi-dimensional arrays is subject to the same type restrictions of single-dim arrays - namely, type and length must match.

#### Passing arrays as arguments

Arrays are value-types, and Go passes arguments using pass-by-value. As a result, passing an array into a function will create a copy of the array. This behavior has performance implications. It's better to pass a pointer to an array into a function rather than the array itself.

````go
var array [1e6]int // a 1 million int array (approximately 8mb)

func foo(ar *[1e6]int) {
  ...
}

foo(&array);
````



### Slices

Slices are vectors - they're growable, dynamic, arrays. Slices are effectively wrappers around an array. They consist of three parts: a pointer to an array, a number representing the length of the underlying array, and another representing the capacity of the array.

#### Creation and initialization

Slices are created and initialized using the built-in `make()` function.

````go
slice := make([]string, 5) // a slice of five strings (capacity == length == 5)

// we can also specify the length and capacity separately
slice := make([]string, 3, 5) // a slice of three strings (capacity = 5, length = 3)
// the capacity must be equal to, or greater than, the specified length
````

Slices can also be created using the same literal notation as arrays, just without specifying the size.

````go
slice := []string{"one", "two", "three"} // len = cap = 3
````

It's also possible to set the length and capacity of a slice by initializing its last element using index-initialization:

````go
slice := []string{99: ""} // len == cap == 100
````

Declaring a slice without initializing it will result in a nil slice. A nil slice has a length and capacity set to zero.

````go
var slice = []int
````

It's also possible to create empty slices. An empty slice differs from a nil slice in that it has an array pointer, but the array is zero sized.

````go
// create a empty slice using make()
slice := make([]int, 0)

// create a empty slice using literal notation
slice := []int{}
````



#### Working with slices

##### Slicing slices

Slices are index accessible the same as arrays. Slices also allow you to create a "slice of a slice", essentially creating a new slice from a selection of an existing slice. 

```go
slice := []int{0,1,2,3,4}

// create a slice by pulling three elements out of 'slice' starting at index 1
newSlice := slice[1:3] 
// newSlice = [1,2,3]
```

New slices will have their length and capacity calculated as follows:

For slice[i:j] with an underlying array of capacity k

Length: j - i

Capacity: k - i

So `newSlice` above will have a length of 3 - 1 = 2, and a capacity of 5 - 1 = 4.

Slices shadow the array they're created from - they're not copies. Changing one will change the other.

A slice can only be indexed up to the length of the slice, not the capacity. Attempting to access an index beyond the slice's length will result in a runtime error.

A slice can be "grown" by appending new values to it. Growing is accomplished using the `append()` routine. Appending to a slice will always update the length of the slice, but the capacity will only be updated if the underlying array needed to re-allocated.

Modifying a slice-of-a-slice by appending will have an affect on the parent slice. 

```go
slice := []int{10,20,30,40,50}
newSlice := slice[1:3]
newSlice = append(newSlice, 60)
```

![image-20200407220954987](/Users/adamgarren/Projects/manning/go-in-action/image-20200407220954987.png)

wut.

`append()` will result in a slice with a newly allocated array when the sliced array's capacity is too small.

```go
slice := []int{1,2,3,4} // len == cap == 4
newSlice := append(slice, 5) // parent slice's cap too small. 
```

When capacity needs to be increased, Go will double it so long as the array contains less than 1000 elements. At and beyond 1000 elements and Go will grow it by 25%. This behavior may change.

A third index can be used when slicing that will allow for finer control over the capacity of the new slice. 

```go
source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
slice := source[2:3:4] // {Plum, Banana, Grape}
// starting at index 2, take 3-2 elements (1), and set the cap to 2 ((k = 4) - 2 = 2)
```

The length and capacity of the new slice are calculated as follows:

For `slice[i:j:k]` or `[2:3:4]`

Length: j - i or 3 - 2 = 1

Capacity: k - i or 4 - 2 = 2

`i` represents the starting index of the slice

`j` represents the index to stop at (inclusive)

`k` indirectly represents the capacity we want the slice to have. This seems retarded to me, but we want the capacity of the new slice to be 2, so what value minus the starting index will give us 2? 4.

4 - 2 = 2.

If you set the length and the capacity of the new slice to be the same, then modifications to the new slice will not be reflected in parent slice.

```go
source := []string{"one", "two", "three", "four", "five"}
slice := source[2:3:3]
slice = append(slice, "Kiwi")
```

Create a slice starting at index 2 (exclusive), stopping at index 3 (inclusive), set the capacity to the length (1) by determining the number that, when 2 is subtracted, will give 1. 3. This creates a length 1, capacity 1 slice that will get reallocated on modification (copy on write?)

The `append()` routine is variadic, allowing for easy contatenation of slices. `...` will "explode" a collection.

```go
slice3 := append(slice1, slice2...)
```

##### Iterating over slices

The `range()` routine is used for iterating over collections including slices.

```go
slice := []int{0,1,2,3,4}
for index, value := range slice {
  fmt.Printf("index: %d, value %d\n", index, value)
}
```

`range()` makes a copy of the value, not a reference, and it always starts at the beginning.

A more conventional for loop can also be used.

```go
for index := 2; index < len(slice); index++ {
  fmt.Printf("Index %d, Value %d\n", index, slice[index]);
}
```

Note: `len()` and `cap()` return the length and capacity (respectively) for a given array.

`append()` also works for multi-dimensional slices. 

```go
	slice := [][]int{{10}, {100, 200}}

	fmt.Printf("%v\n", slice)

	slice[0] = append(slice[0], 20)

	fmt.Printf("%v\n", slice)


// [[10] [100 200]]
// [[10 20] [100 200]]
```

##### Passing slices as arguments

Slices are small and cheap to copy by value, unlike arrays. Passing a slice involves creating a shallow copy of the three element slice internal (pointer to array, int cap, int len), so it doesn't necessitate passing a pointer. Only the slice is copied, not the underlying array.

```go
slice := make([]int, 1e6) // 1 million element slice

func foo(slice []int) []int {
//  ...
  return slice
}

slice = foo(slice)
```



### Maps

#### Internals

Maps, like arrays and slices, are collections. Unlike arrays and slices, maps are unordered. There is no guarantee that iterating over a map will result in the same ordering.

Maps are implemented as hash-table, and a hash-table contains a collection of buckets. A map's hash function generates an even distribution of key/value pairs over its bucket collection. A generated hash key consists of low-order bits (LOBs) and high-order bits (HOBs), the LOBs are used to select the appropriate bucket.

Buckets are comprised of an array of HOBs that distinguish individual entries, and a collection containing both the keys and the values.

#### Creating and initializing

Maps can be created using the same `make()` routine used for arrays, or using a literal notation.

```go
dict := make(map[string]int) // a map of string to int using make()
dict := map[string]string{"zero", "one", "two", "three"} // literal
```

Map keys can be built-in types or structs (as long as the value can be used in an == expression.) Slices, functions, and structs containing slices *cannot* be used as keys.

Map values can be pretty much anything.

```go
dice := map[int][]string{} // a map of ints to string slices
```

Assigning values to maps is unremarkable.

```go
colors := map[string]string{} // an empty map of string to string
colors["Red"] = "#da1337"
```

Creating nil maps and assigning values to them blow up as one would expect.

```go
var colors = map[string]string // create a nil map (no initializer)
colors["Red"] = "#da1337" // Runtime error - panic
```

Testing a key for existance can be done by trying to get the value...

```go
value, exists := colors["Blue"]
if exists {
  fmt.Println(value)
}
```

 It's also possible to test for the "zero" value of a type so long as that value isn't a valid value in the map.

```go
value := colors["Blue"]
if value != "" {
  //...
}
```

A map index will always return a value for a key. If the key doesn't exist, then the value will be the zero value for the value type captured in the map.

#### Iterating over maps

Iterating over a map uses the `range()` routine just like arrays and slices. The difference is that it will return the key/value pair.

```go
colors := map[string]string{
  "Red": "#ff0000",
  "Blue": "#00ff00",
  "Green": "#0000ff",
}

for key, value := range colors {
  // ...
}
```

 

#### Map operations

Removing items from a map is accomplished using the `delete()` routine.

`delete(colors, "Coral")`

Getting the length of a map is done using the same `len()` routine as for slices and arrays.

Maps don't have a capacity and there are no restrictions on growth.

#### Passing maps to functions

There's no need to pass a map pointer into a function, passing the map will create a shallow copy of the map, but the hash-table is not copied. This means that modifying a map that was passed by value into function will allow the function to modify the map itself.    

Maps are designed to be cheap to copy just like slices.

## Go's Type System

Go is statically typed, so the compiler knows the type of every value that is created. Knowing the type allows the compiler to more accurately determine the amount of memory required to allocate for a value, as well as what operations are valid for a given type.

### User-defined types

Go provides a `struct` type for creating user-defined types. A struct can contain a collection of named, built-in, variables, or other structs. Go also allows for defining a type in terms of another type, similar to C's typedef, with the notable difference that these types are distinct and not interchangable.

#### Defining a user-defined struct type

Here we define a user as a struct containing a handful of built-in fields.

```go
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}
```

####  Declaring a user-defined struct type

A user-defined type can be declared using struct literal notation, inline, or "zero" initialized using the `var` keyword.

```go
// literal notation using named fields. Order doesn't matter.
lisa := user{
  name: "Lisa",
  email: "lisa@example.com",
  ext: 123,
  privilaged; true,
}

// inline notation doesn't use named fields. Order does matter.
linda := user{"Linda", "linda@example.com", 123, true}

// zero-inititialized sets every field to its zero value (empty string for strings
// false for bools, 0 for numbers, etc.)
var bill user
```

The `var` keyword should be reserved for zero-value initialization. It communicates the fact that the type and all of its fields are zero values. All other initialization should use the `:=` operator, since it provides double duty: declaration *and* initialization.

User-defined types can contain other user-defined types, and the same initialization rules apply.

```go
// define an aggregate UDT
type admin struct {
  person user
  level  string
}

// declare using struct literal initialization
fred := admin{
  person: user{
    name: "fred",
    email: "fred@example.com",
    ext: 345,
    privilaged: true,
  },
  level: "root"
}
```

#### Defining a 'typedef'

Go allows deriving a new type based upon another already existing type. In this case, the result is a distinct type that cannot be interchanged with its base type.

```go
type Duration int64

func main() {
  var dur Duration
  dur = int64(1000) // <--- Not valid.
}
```

Even though Duration is defined solely in terms of an int64 above, it cannot be initialized from an int64 type. Go's type system views int64 and Duration as completely separate types. Further, Go refuses to silently convert one type to another - implicit type coercion doesn't exist.

###  Methods

Method provide a way to combine operations for a given type. Defining a method for a type involves declaring a "receiver" of a type on its method. A receiver can be either a pointer or a value. The receiver provides a context, the same as "this" in Java or C++, in which the method can access instance-scoped fields.

```go
type user struct {
	name  string
	email string
}

// notify method for user type that uses a value as a receiver
func (u user) notify() { 
	//...
}

// changeEmail method for user that uses a pointer as a receiver
func (u *user) changeEmail(email string) {
  // ...
}

func main() {
	// create a value user type
	bill := user{"Bill", "bill@example.com"}
	bill.notify()

	// create pointer to a user instance
	lisa := &user{"Lisa", "lisa@example.com"}
  lisa.notify() // (*lisa).notify()

	// use a value type to call a method that accepts
	// a pointer. The compiler will take care of the
	// conversion to a pointer.
  bill.changeEmail("bill@corp.com") // (&bill).notify()
	bill.notify()

	lisa.changeEmail("lisa@micro.com")
	// Here too, the compiler will take care of the
	// details for us. lisa will be dereferenced
	// automatically so we can call methods with a
	// value reciever with a pointer without worry
	lisa.notify()
}
```

Go will adjust the  reciever to comply with the method's definition. If a method defined with a value type is called with a pointer, then Go will dereference the pointer for us. Similarly, if a pointer defined method is called with a value, Go will use the address of the value as the reciever instead.

When calling a value-defined method, a copy of the value is used. When calling a pointer defined method, the existing value is shared with the method.

##### Pointer or value recievers - when to use?

Does adding or removing something from a value of the type need to create a new value or mutate the existing one? If the answer is "create", then user value recievers. If the answer is "mutate", then user a pointer reciever.

#### Built-in types

Built-in types are simple, primative, and designed to be cheap to copy, modify, and pass as arguments. Built-in types including strings, numeric types, booleans, etc. should all be treated as values and copied rather than as pointers.

#### Reference Types

Maps, slices, arrays, interfaces, channels, and functions are all reference types. When declared, reference types consist of a locally scoped "header" with a pointer to the type's internal data structures. Think RAII. The header is designed to be cheap to copy and can be passed by value, copies will share the underlying internals.

#### Stucts and user-defined types

The nature of a user-defined type will determine whether or not its methods should specify pointer or value recievers. 

If a UDT represents a concept that shouldn't be modified, then it's methods should be defined with value recievers, will operate on a copy of the reciever, and will return new instances of the type if necessary.

E.g.,  Time structure in Go represents a particular point in time. An instance of Time should not change, but methods can return new instances of Time that have minutes added, etc.

Generally structs and UDTs will not exhibit a value-like nature, instead they'll be intended as mutatable instances. In these cases methods defined with a pointer reciever will get access to the types context, instance data, rather than a copy of it. Some types are unsafe to copy and must be passed as pointers. The File type, for example, protects against copying by wrapping a file struct as an embedded pointer.

```go
type File struct {
  *file
}

type file struct {
  fd      int
  name    string
  dirinfo *dirInfo
  nepipe   int32
}
```

If a function returns a pointer, then it's a good indicator that the type is nonprimative and should not be passed by value or copied.

```go
// Open() returns a pointer. The pointer tells us
// not to treat it as a value.
func Open(name string) (file *File, err error) {
	// ...
}

// Even if we don't change the state of a non-primative 
// type, we should still pass it as a pointer.
func (f *File) Chdir() error {
  if f == nil {
    return ErrInvalid
  }
  if e := syscol.Fchdir(f.fd); e != nil {
    return &PathError{"chdir", f.name, e}
  }
  return nil
}
```

The decision whether to use values or pointers as recievers should have less to do with mutation, and more to do with the nature of the type. 

### Interfaces

Interfaces are the primary mechanism Go provides for supporting polymorphism. Interfaces are types that declare only behavior, not implementation. Implementation is defined by methods within user-defined types. When a UDT implements all the methods described by an interface, then the type is said to implement the interface. 

Any user-defined type can implement any interface. UDTs that implement an interface are referred to as "concrete types", since they provide concrete behavior. When a method is called on an interface instance, then the equivalent method is called on the instances stored/backing type.

An interface instance is essentially a header that contains two pointers. The first points to an "iTable" that captures the instances stored/backing type and a "method set" The second is a pointer to the backing instance itself.

If an interface is assigned to a pointer, then the type captured by its iTable is a pointer.

```go
var n notifier    // a nil notifier interface
n = &user{"bill"} // a notifier interface with a pointer stored value

var n notifier
n = user{"bill"}  // a notifier interface instance with a value stored value
```

There are rules surrounding whether a values or pointers of a UDT satisfy an interface. 

```go
package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func main() {
	u := user{"Bill", "bill@example.com"}
	sendNotification(u)

	//$ go run listing36.go
	//# command-line-arguments
	//./listing36.go:21:18: cannot use u (type
	//	user) as type notifier in argument to
	// sendNotification:
	//        user does not implement notifier
	// (notify method has pointer receiver)

	// call fails because we've implemented
	// notify using a pointer receiver and we
	// called with a value.
}

func sendNotification(n notifier) {
	n.notify()
}
```

##### Method sets

Method sets define the set of methods associated with values or pointers of a particular type. Go's method set specification looks like the following:

| Values | Method Receivers |
| ------ | ---------------- |
| T      | (t T)            |
| *T     | (t T) or (t *T)  |

A value of T only has methods declared that have a value as a receiver. Pointers to T have methods declared for both values and pointeres.

However, looking at this table from the perspective of the Receivers changes things.

| Method Receivers | Values   |
| ---------------- | -------- |
| (t T)            | T and *T |
| (t *T)           | *T       |

If you implement an interface using a pointer, then only pointers of type T implement the interface. However, if you implement the interface using a value, then both pointers and values of T implement the interface.

In the example code above, calling `notify()` using a pointer would satisfy the restrictions of the method set since the interface was inplemented using a pointer receiver.

```go
sendNotification(&u)
```

This restriction is in place because it's not always possible to get the address of a value. E.g.,

```go
type duration int

func (d *duration) pretty() string {
  return fmt.Sprintf("Duration: %d", *d) // deref d pointer
}

func main() {
  duration(42).pretty() // <-- can't take an address of a literal value
}
```

It's for this reason that method sets for a value only include methods that are implemented with value receivers. However, value receivers accept both values and pointers.

#### Interfaces as a method of polymophism

In this example both user and admin are concrete types implementing the notifier interface.

```go
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name, a.email)
}

func main() {
	u := user{"Bill", "bill@example.com"}
	sendNotification(&u)

	a := admin{"Lisa", "lisa@example.com"}
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
```

```shell
$ go run listing37.go
Sending user email to Bill<bill@example.com>
Sending admin email to Lisa<lisa@example.com>

```

#### Type embedding and method promotion

Go support a convenient feature related to composition of types. Types can contain embedded types, and the methods of those embedded types can be "promoted" to the scope of the containing type.

Here we embed a user in an admin type. The `notify()` method of the user is accessible through the `user` reference, but the method is also promoted to the `admin` scope. Method promotion of embedded types allows us to access the `notify()` method directly on the `admin` instance as if it had been declared there in the first place.

```go
type user struct {
	name  string
	email string
}

type admin struct {
	user  // "embedded" type - no instance name
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "Jon Smith",
			email: "Jon@example.com",
		},
		level: "root",
	}

	ad.user.notify()
	ad.notify()
}

```

```shell
$ go run listing50.go
Sending user email to Jon Smith<Jon@example.com>
Sending user email to Jon Smith<Jon@example.com>

```

Method promotion for embedded types also means that "outer" types implement interfaces of embedded "inner" types.

```go
package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user  // "embedded" type - no instance name
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "Jon Smith",
			email: "Jon@example.com",
		},
		level: "root",
	}

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
```

```shell
$ go run listing51.go
Sending user email to Jon Smith<Jon@example.com>

```

In the example above, the user type implements the notification interface, and since it's embedded in the admin type, its `notify()` method is promoted to the outer admin type. The call to `sendNotification()` recognizes the promoted method as satisfying the notification interface.

If the admin type provides its own `notify()` implementation, then it will override the promoted method from the user type. The user's method can still be referenced, but it has to be done through the user "inner" instance.

```go
	ad.user.notify()
	ad.notify()
```

### Exporting Identifiers

